# syntax=docker/dockerfile:1.7

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

ENV GOOS=wasip1
ENV GOARCH=wasm

WORKDIR /go/github.com/shortlink-org/shortlink

# Load dependencies
COPY go.mod go.sum ./

# will cache go packages while downloading packages
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# COPY the source code AS the last step
COPY . .

# Build project
RUN go build -o main.wasm ./boundaries/platform/istio-extension/shortlink_go/main.go

FROM scratch

COPY --from=builder /go/github.com/shortlink-org/shortlink/main.wasm ./plugin.wasm
