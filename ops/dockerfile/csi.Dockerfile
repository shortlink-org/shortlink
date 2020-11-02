FROM golang:1.15-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/src/github/batazor/shortlink
COPY . .

# Load dependencies
RUN ! -d "/go/src/github/batazor/shortlink/vendor" && go mod vendor

# Build project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -ldflags "-X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./cmd/k8s/csi

FROM alpine:latest

# 9090: metrics
EXPOSE 9090

# Install dependencies
RUN \
    apk update && \
    apk add curl util-linux && \
    rm -rf /var/cache/apk/*

WORKDIR /app/
COPY --from=builder /go/src/github/batazor/shortlink/app .
ENTRYPOINT ["./app"]
