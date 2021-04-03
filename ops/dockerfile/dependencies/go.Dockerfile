FROM golang:1.16-alpine as builder

ARG CI_COMMIT_TAG
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /go/github.com/batazor/shortlink

# Load dependencies
COPY go.mod .
COPY go.sum .
RUN go mod vendor
