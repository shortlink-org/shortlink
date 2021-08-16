# syntax=docker/dockerfile:1.3

FROM alpine:latest

# Install dependencies
RUN apk add --update curl

# 7070: API
# 9090: Prometheus metrics
EXPOSE 7070 9090

WORKDIR /app/
CMD ["./api"]
COPY api .
