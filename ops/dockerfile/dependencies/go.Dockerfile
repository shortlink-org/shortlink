FROM golang:1.15-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/src/github/batazor/shortlink
COPY . .

# Load dependencies
RUN ! -d "/go/src/github/batazor/shortlink/vendor" && go mod vendor
