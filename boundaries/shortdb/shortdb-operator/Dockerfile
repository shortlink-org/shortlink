# Build the manager binary
FROM --platform=$BUILDPLATFORM golang:1.23-alpine@sha256:ac67716dd016429be8d4c2c53a248d7bcdf06d34127d3dc451bda6aa5a87bc06 AS builder

ENV GOCACHE=/root/.cache/go-build

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# will cache go packages while downloading packages
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Field Alignment
RUN go run golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment -fix ./internal/...; exit 0

# Build project
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot@sha256:26f9b99f2463f55f20db19feb4d96eb88b056e0f1be7016bb9296a464a89d772
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
