# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}
CI_COMMIT_TAG := latest

DOCKER_USERNAME := "batazor"

PATH_TO_UI_NUXT := pkg/ui/nuxt

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# APPLICATION ==========================================================================================================
dep: ## Install dependencies for this project
	@echo "install protoc"
	@sudo ./ops/scripts/install-protobuf.sh

	@echo "install protoc addons"
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u github.com/batazor/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go get -u github.com/securego/gosec/cmd/gosec

	@echo "install wire"
	@go get -u github.com/google/wire/cmd/wire

generate: ## Code generation
	@echo "proto generation link entity"
	@protoc -I/usr/local/include -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=plugins=grpc:. \
	pkg/link/link.proto

	@protoc -I/usr/local/include -I. \
    	--gotemplate_out=all=true,template_dir=internal/store/query/template:internal/store/query \
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
	@go generate internal/di/wire.go

	@echo "Generate go static"
	@pkger -o cmd/shortlink

	@make fmt

golint: ## Linter for golang
	@golangci-lint run

.PHONY: fmt
fmt: ## Format source using gofmt
	@echo Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@gosec ./...

test: ## Run all test
	@echo "run test"
	@sh ./ops/scripts/coverage.sh

bench: ## Run benchmark tests
	go test -bench ./...

run: ## Run this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/grafana.yaml \
         -f ops/docker-compose/tooling/loki.yaml \
         -f ops/docker-compose/tooling/fluentd.yaml \
         -f ops/docker-compose/tooling/prometheus.yaml \
         -f ops/docker-compose/database/mysql.yaml \
         up -d --force-recreate

run-dep: ## Run only dep for this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/mq/kafka.yaml \
         -f ops/docker-compose/application/shortlink.yaml \
         -f ops/docker-compose/database/postgres.yaml \
         -f ops/docker-compose/gateway/traefik.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         up -d

down: ## Down docker-compose
	@docker-compose down --remove-orphans
	@docker network rm simple shortlink_default

clean: ## Clean artifacts
	@docker network rm simple
	@docker rmi -f shortlink_shortlink

# DOCKER TASKS =========================================================================================================
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

# KUBERNETES ===========================================================================================================
helm-lint: ## Check Helm chart
	@helm lint ops/Helm/shortlink

helm-deploy: ## Deploy Helm chart to default kube-context and default namespace
	@echo helm install/update ${PROJECT_NAME}
	@helm upgrade ${PROJECT_NAME} ops/Helm/shortlink \
		--install \
		--force \
		--wait

helm-clean: ## Clean artifact from K8S
	@helm del --purge ${PROJECT_NAME}

# UI ===================================================================================================================
nuxt_generate: ## Deploy nuxt UI
	@npm --prefix ${PATH_TO_UI_NUXT} install
	@npm --prefix ${PATH_TO_UI_NUXT} run generate
