# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

# Where to push the docker image.
CI_REGISTRY_IMAGE ?= batazor/${PROJECT_NAME}

# This version-strategy uses a manual value to set the version string
CI_COMMIT_TAG ?= latest
#
# This version-strategy uses git tags to set the version string
# CI_COMMIT_TAG ?= $(shell git describe --tags --always --dirty)

DOCKER_USERNAME := "batazor"

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

###
### These variables should not need tweaking.
###

# Used internally.  Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# APPLICATION ==========================================================================================================
.PHONY: dep generate golint gosec test run run-dep down clean
dep: ## Install dependencies for this project
	@echo "install protoc"
	@sudo ./ops/scripts/install-protobuf.sh
	@echo "install protoc addons"
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u moul.io/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go get -u github.com/securego/gosec/cmd/gosec

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

	@echo "Generate from .go code"
	@go generate internal/store/postgres/postgres.go

	@echo "Generate go static"
	@pkger -o cmd/shortlink

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
         -f ops/docker-compose/mq/kafka.yaml \
         -f ops/docker-compose/application/shortlink.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         up -d --force-recreate

run-dep: ## Run only dep for this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/mq/kafka.yaml \
         -f ops/docker-compose/application/shortlink.yaml \
         -f ops/docker-compose/database/postgres.yaml \
         -f ops/docker-compose/gateway/traefik.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         up -d

down: ## Down docker-compose
	@docker-compose down --remove-orphans

clean: ## Clean artifacts
	@docker rmi -f shortlink_shortlink

# DOCKER TASKS =========================================================================================================
.PHONY: docker docker-login docker-build docker-push
docker: docker-login docker-build docker-push ## docker login > build > push

docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

# Build the container
docker-build: ## Build the container for preferred architecture
	@echo docker build image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG} .

docker-push: ## Publish the container
	@echo docker push image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}

# KUBERNETES ===========================================================================================================
.PHONY: helm-lint helm-deploy helm-clean
helm-lint: ## Check Helm chart
	@helm lint ops/Helm/${PROJECT_NAME}

helm-deploy: ## Deploy Helm chart to default kube-context and default namespace
	@echo helm install/update ${PROJECT_NAME}
	@helm upgrade ${PROJECT_NAME} ops/Helm/shortlink \
		--install \
		--force \
		--wait

helm-clean: ## Clean artifact from K8S
	@helm del --purge ${PROJECT_NAME}
