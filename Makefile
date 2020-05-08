# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

CI_COMMIT_TAG := latest

DOCKER_USERNAME := "batazor"

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

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
	@golangci-lint run ./...

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
         -f ops/docker-compose/database/postgres.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/prometheus.yaml \
         -f ops/docker-compose/tooling/loki.yaml \
         -f ops/docker-compose/tooling/fluentd.yaml \
         up -d --remove-orphans

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
	-@docker network rm simple shortlink_default
	@docker rmi -f shortlink_shortlink

# DOCKER TASKS =========================================================================================================
CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}

docker: docker-login docker-build docker-push ## docker login > build > push

docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

docker-build: ## Build the container
	@echo docker build image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG} -f ops/dockerfile/shortlink.Dockerfile .

	@echo docker build image ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG} -f ops/dockerfile/logger.Dockerfile .

	@echo docker build image ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG} -f ops/dockerfile/ui-nuxt.Dockerfile .

docker-push: ## Publish the container
	@echo docker push image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}

	@echo docker push image ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}

	@echo docker push image ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}

# KUBERNETES ===========================================================================================================
PATH_TO_SHORTLINK_CHART := ops/Helm/shortlink-ui
PATH_TO_COMMON_CHART := ops/Helm/common

helm-init: ## helm init
	@add custom repo for helm
	@helm repo add jaegertracing https://jaegertracing.github.io/helm-charts

helm-lint: ## Check Helm chart
	@helm lint ${PATH_TO_SHORTLINK_CHART}
	@helm lint ${PATH_TO_COMMON_CHART}

helm-deploy: ## Deploy Helm chart to default kube-context and default namespace
	@echo helm install/update ${PROJECT_NAME}
	@helm upgrade ${PROJECT_NAME} ${PATH_TO_SHORTLINK_CHART} \
		--install \
		--force \
		--wait

helm-clean: ## Clean artifact from K8S
	@helm del ${PROJECT_NAME}

helm-common: ## run common service for
	@make helm-init
	@echo helm install/update common service
	@helm upgrade common ${PATH_TO_COMMON_CHART} \
		--install \
		--force \
		--wait

# MINIKUBE =============================================================================================================
minikube-init: ## run minikube for dev mode
	@minikube start \
		--cpus 4 \
		--memory "16384" \
		--extra-config=apiserver.enable-admission-plugins=PodSecurityPolicy\
		--extra-config=apiserver.enable-admission-plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"
	@minikube addons enable ingress
	@eval $(minikube docker-env) # Set docker env

minikube-update: ## update image to last version
	@eval $(minikube docker-env) # Set docker env
	@make docker-build           # Build docker images on remote host (minikube)
	@make helm-deploy            # Deploy shortlink HELM-chart

minikube-down: ## minikube delete
	@minikube delete

# ISTIO ================================================================================================================
istio-run: ## Run istio
	@istioctl manifest apply --set profile=demo


# GITLAB ===============================================================================================================
gitlab-minikube: ## Install GitLab to minikube
	@helm repo add gitlab https://charts.gitlab.io/
	@helm repo update
	@helm upgrade -n gitlab --install gitlab gitlab/gitlab \
      --namespace=gitlab \
      --create-namespace=true \
	  -f ops/docker-compose/tooling/gitlab/helm-value.yaml

# UI ===================================================================================================================
PATH_TO_UI_NUXT := pkg/ui/nuxt

nuxt_generate: ## Deploy nuxt UI
	@npm --prefix ${PATH_TO_UI_NUXT} install
	@npm --prefix ${PATH_TO_UI_NUXT} run generate
