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
	@go get -u github.com/jteeuwen/go-bindata/...
	@go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
  # CFSSL: Cloudflare's PKI and TLS toolkit
	@go get -u github.com/cloudflare/cfssl/cmd/cfssl
	@go get -u github.com/cloudflare/cfssl/cmd/cfssljson

	# install wire
	@go get -u github.com/google/wire/cmd/wire

export CURRENT_UID=$(id -u):$(id -g)

do: ## Run for specific job
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/application/api.yaml \
         -f ops/docker-compose/application/metadata.yaml \
         -f ops/docker-compose/database/postgres.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/prometheus.yaml \
         up -d --remove-orphans

run: ## Run this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         -f ops/docker-compose/tooling/fluent-bit.yaml \
         -f ops/docker-compose/gateway/traefik.yaml \
         -f ops/docker-compose/application/api.yaml \
         -f ops/docker-compose/application/metadata.yaml \
         -f ops/docker-compose/application/logger.yaml \
         -f ops/docker-compose/application/ui-next.yaml \
         -f ops/docker-compose/database/mongo.yaml \
         -f ops/docker-compose/tooling/prometheus.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         -f ops/docker-compose/tooling/grafana.yaml \
         -f ops/docker-compose/tooling/loki.yaml \
         -f ops/docker-compose/mq/rabbitmq.yaml \
         up -d --remove-orphans

run-dep: ## Run only dep for this project in docker-compose
	@docker-compose \
         -f docker-compose.yaml \
         -f ops/docker-compose/mq/kafka.yaml \
         -f ops/docker-compose/application/api.yaml \
         -f ops/docker-compose/database/postgres.yaml \
         -f ops/docker-compose/gateway/traefik.yaml \
         -f ops/docker-compose/tooling/opentracing.yaml \
         -f ops/docker-compose/tooling/coredns.yaml \
         up -d --remove-orphans

down: ## Down docker-compose
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/coredns.yaml \
		-f ops/docker-compose/tooling/fluent-bit.yaml \
		-f ops/docker-compose/gateway/traefik.yaml \
		-f ops/docker-compose/application/api.yaml \
 		-f ops/docker-compose/application/metadata.yaml \
		-f ops/docker-compose/application/logger.yaml \
		-f ops/docker-compose/application/ui-next.yaml \
		-f ops/docker-compose/database/mongo.yaml \
		-f ops/docker-compose/tooling/prometheus.yaml \
		-f ops/docker-compose/tooling/opentracing.yaml \
		-f ops/docker-compose/mq/rabbitmq.yaml \
		down --remove-orphans
	@docker network prune -f

logger: ## Run logger infra
		@docker-compose \
				-f docker-compose.yaml \
        -f ops/docker-compose/application/api.yaml \
        -f ops/docker-compose/tooling/coredns.yaml \
				-f ops/docker-compose/tooling/grafana.yaml \
				-f ops/docker-compose/tooling/loki.yaml \
				-f ops/docker-compose/tooling/prometheus.yaml \
				-f ops/docker-compose/tooling/opentracing.yaml \
				up -d --remove-orphans
