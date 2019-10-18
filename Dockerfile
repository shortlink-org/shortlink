FROM golang:1.13-alpine as builder

ARG CI_COMMIT_TAG

# Build project
WORKDIR /go/src/github/batazor/shortlink
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
  go build \
  -a \
  -mod vendor \
  -ldflags "-X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo -o app ./cmd/shortlink

FROM alpine:latest

RUN addgroup -S 997 && adduser -S -g 997 997
USER 997

WORKDIR /app/
COPY --from=builder /go/src/github/batazor/shortlink/app .
CMD ["./app"]
