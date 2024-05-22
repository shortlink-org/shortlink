# syntax=docker/dockerfile:1.7

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ARG TARGETOS
ARG TARGETARCH

ENV GOEXPERIMENT=rangefunc,newinliner,arenas,cgocheck2

WORKDIR /go/github.com/shortlink-org/shortlink

# Load dependencies
COPY go.mod go.sum ./

# will cache go packages while downloading packages
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# COPY the source code AS the last step
COPY . .

# Field Alignment
RUN go run golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment -fix ./internal/...; exit 0

# Build project
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build \
  -a \
  -gcflags="${SKAFFOLD_GO_GCFLAGS}" \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./boundaries/platform/csi/cmd

FROM --platform=$TARGETPLATFORM alpine:3.20

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-csi"
LABEL org.opencontainers.image.description="shortlink-csi"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# Define GOTRACEBACK to mark this container AS using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=all

# 9090: metrics
EXPOSE 9090

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl util-linux tini

ENTRYPOINT ["/sbin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/shortlink-org/shortlink/app /app
