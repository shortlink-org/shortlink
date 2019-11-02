# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}
CI_COMMIT_TAG := 0.1.0

DOCKER_LOGIN := "batazor"

# export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME
export CI_REGISTRY_IMAGE
export CI_COMMIT_TAG

# Regular Makefile part for buildpypi itself
help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo '  dep			install dependencies for this project'
	@echo '  generate		code generation'
	@echo '  golint		linter for golang'
	@echo '  test			run all test'
	@echo '  run  			run thisproject in docker-compose'
	@echo '  down			down docker-compose'
	@echo '  docker-login		docker login'
	@echo '  docker-build		docker build'
	@echo '  docker-push		docker push'

.: generate

dep:
	@echo "install protoc"
	@sudo ./ops/scripts/install-protobuf.sh
	@echo "install protoc addons"
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u moul.io/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

generate:
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

golint:
	@golangci-lint run

test:
	@echo "run test"
	@sh ./ops/scripts/coverage.sh

run:
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/database/redis.yaml \
         -f ops/docker-compose/gataway/traefik.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         up -d

down:
	@docker-compose down --remove-orphans

docker-login:
	@echo docker login as ${DOCKER_LOGIN}
	@docker login -u ${DOCKER_LOGIN} -p ${DOCKER_PASS}

docker-build:
	@echo docker build image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG} .

docker-push:
	@echo docker push image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
