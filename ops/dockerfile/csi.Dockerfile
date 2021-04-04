FROM golang:1.16-alpine as builder

ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /go/github.com/batazor/shortlink

# Load dependencies
COPY go.mod go.sum ./
RUN go mod download

# COPY the source code as the last step
COPY . .

# Build project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -gcflags="${SKAFFOLD_GO_GCFLAGS}" \
  -ldflags "-s -w -X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./cmd/k8s/csi

FROM alpine:3.13

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=all

# 9090: metrics
EXPOSE 9090

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl util-linux

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/batazor/shortlink/app /app
