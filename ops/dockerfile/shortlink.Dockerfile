FROM golang:1.14-alpine as builder

ARG CI_COMMIT_TAG

# Build project
WORKDIR /go/src/github/batazor/shortlink
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -ldflags "-X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./cmd/shortlink

FROM alpine:latest

# 7070: API
# 9090: metrics
EXPOSE 7070 9090

# Install dependencies
RUN \
    apk update && \
    apk add curl && \
    rm -rf /var/cache/apk/*

RUN addgroup -S shortlink && adduser -S -g shortlink shortlink
USER shortlink

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /app/
COPY --from=builder /go/src/github/batazor/shortlink/app .
CMD ["./app"]
