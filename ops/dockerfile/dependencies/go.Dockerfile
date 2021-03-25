FROM golang:1.16-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/github.com/batazor/shortlink
COPY . .

# Load dependencies
RUN go mod vendor
