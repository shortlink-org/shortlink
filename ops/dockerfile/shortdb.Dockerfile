# syntax=docker/dockerfile:1.4

FROM --platform=$BUILDPLATFORM golang:1.20-rc AS builder

ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ARG TARGETOS TARGETARCH

WORKDIR /go/github.com/shortlink-org/shortlink

# Load io_uring
RUN apt-get update && apt-get install -y liburing-dev

# Load dependencies
COPY go.mod go.sum ./
RUN go mod download

# COPY the source code as the last step
COPY . .

# Build project
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=1 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build \
  -a \
  -gcflags="${SKAFFOLD_GO_GCFLAGS}" \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./internal/pkg/shortdb/cli

FROM debian:11.6

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=all

# Load io_uring
RUN apt-get update && apt-get install -y liburing-dev curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/shortlink-org/shortlink/app /app
