# syntax=docker/dockerfile:1.7

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ARG PGO_PATH
ARG CMD_PATH
ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ARG TARGETOS
ARG TARGETARCH

ENV GOCACHE=/root/.cache/go-build
ENV GOEXPERIMENT=rangefunc,newinliner,arenas,cgocheck2
ENV PGO_PATH=auto

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
  --mount=type=cache,target=/go/pkg/mod \
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build \
  -a \
  -pgo=${PGO_PATH} \
  -gcflags="${SKAFFOLD_GO_GCFLAGS}" \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app $CMD_PATH

FROM --platform=$TARGETPLATFORM alpine:3.20

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-${CMD_PATH}"
LABEL org.opencontainers.image.description="shortlink-${CMD_PATH}"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"
LABEL org.opencontainers.image.revision=$CI_COMMIT_SHA

# Define GOTRACEBACK to mark this container AS using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=all
ENV USER golang

# 50051: gRPC
# 9090: metrics
EXPOSE 50051 9090

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl tini

RUN addgroup -S golang && adduser -S -g golang golang
USER golang

ENTRYPOINT ["/sbin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/shortlink-org/shortlink/app /app
