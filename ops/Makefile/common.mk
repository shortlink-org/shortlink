# APPLICATION TASKS ====================================================================================================
dep: ## Install dependencies for this project
	# install protoc
	@sudo ./ops/scripts/install-protobuf.sh
	@sudo rm -rf bin

	# install protoc addons
	@go get -u github.com/golang/protobuf/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u github.com/batazor/protoc-gen-gotemplate
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go get -u github.com/securego/gosec/cmd/gosec
	@go get -u moul.io/protoc-gen-gotemplate

	# install wire
	@go get -u github.com/google/wire/cmd/wire

run: ## Run this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/application/shortlink.yaml \
         -f ops/docker-compose/application/logger.yaml \
         -f ops/docker-compose/application/ui-next.yaml \
         -f ops/docker-compose/database/mongo.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/fluentd.yaml \
         -f ops/docker-compose/tooling/prometheus.yaml \
         -f ops/docker-compose/mq/rabbitmq.yaml \
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
