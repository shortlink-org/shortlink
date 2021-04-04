FROM golang:1.16-alpine as builder

ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /go/github.com/batazor/shortlink

# Load dependencies
COPY go.mod go.sum ./
RUN go mod vendor

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
  -o app ./cmd/api

FROM alpine:3.13

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=all

# 7070: API
# 9090: metrics
EXPOSE 7070 9090

# Install dependencies
RUN \
    apk update && \
    apk add --no-cache curl ca-certificates git

RUN addgroup -S api && adduser -S -g api api
USER api

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
CMD ["./app"]
COPY --from=builder /go/github.com/batazor/shortlink/app /app
