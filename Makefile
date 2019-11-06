# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}
CI_COMMIT_TAG := 0.1.0

DOCKER_USERNAME := "batazor"

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dep: ## Install dependencies for this project
	@echo "install protoc"
	@sudo ./ops/scripts/install-protobuf.sh
	@echo "install protoc addons"
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u moul.io/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

generate: ## Code generation
	@echo "proto generation link entity"
	@protoc -I/usr/local/include -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=plugins=grpc:. \
	pkg/link/link.proto

	@echo "proto generation gRPC-web"
	@protoc -I/usr/local/include -I. \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--go_out=plugins=grpc:. \
	--swagger_out=logtostderr=true,allow_delete_body=true:. \
	--grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
	pkg/api/grpc-web/api.proto
	@mv pkg/api/grpc-web/api.swagger.json docs/api.swagger.json

	@echo "Generate go static"
	@go generate pkg/api/graphql/schema/schema.go

golint: ## Linter for golang
	@golangci-lint run

gosec: ## Golang security checker
	@gosec ./...

test: ## Run all test
	@echo "run test"
	@sh ./ops/scripts/coverage.sh

run: ## Run this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/database/redis.yaml \
         -f ops/docker-compose/gateway/traefik.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         up -d

down: ## Down docker-compose
	@docker-compose down --remove-orphans

# DOCKER TASKS
docker: docker-login docker-build docker-push ## docker login > build > push

docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

# Build the container
docker-build: ## Build the container
	@echo docker build image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG} .

docker-push: ## Publish the container
	@echo docker push image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
