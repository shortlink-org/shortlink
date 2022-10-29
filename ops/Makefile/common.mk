# APPLICATION TASKS ====================================================================================================
dep: ## Install dependencies for this project
	# install protoc addons
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/srikrsna/protoc-gen-gotag@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/securego/gosec/cmd/gosec@latest
	@go install moul.io/protoc-gen-gotemplate@latest
	@go install github.com/cloudflare/cfssl/cmd/...@latest

	# for NodeJS
	@npm install -g grpc-tools grpc_tools_node_protoc_ts

	# install wire
	@go install github.com/google/wire/cmd/wire@latest

	#i18n
	@go install golang.org/x/text/cmd/gotext@latest

export CURRENT_UID=$(id -u):$(id -g)

up: ## Run for specific job
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns.yaml \
		-f ops/docker-compose/mq/rabbitmq.yaml \
		-f ops/docker-compose/tooling/observability/prometheus.yaml \
		-f ops/docker-compose/tooling/observability/grafana.yaml \
		-f ops/docker-compose/tooling/observability/grafana-loki.yaml \
		-f ops/docker-compose/tooling/observability/grafana-tempo.yaml \
		up -d --remove-orphans

run: ## Run this project in docker compose
	@docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns.yaml \
		-f ops/docker-compose/tooling/observability/fluent-bit.yaml \
		-f ops/docker-compose/gateway/traefik.yaml \
		-f ops/docker-compose/application/auth.yaml \
		-f ops/docker-compose/application/api.yaml \
		-f ops/docker-compose/application/metadata.yaml \
		-f ops/docker-compose/application/logger.yaml \
		-f ops/docker-compose/application/ui-next.yaml \
		-f ops/docker-compose/database/mongo.yaml \
		-f ops/docker-compose/tooling/observability/prometheus.yaml \
		-f ops/docker-compose/tooling/observability/grafana.yaml \
		-f ops/docker-compose/tooling/observability/grafana-loki.yaml \
		-f ops/docker-compose/tooling/observability/grafana-tempo.yaml \
		-f ops/docker-compose/mq/rabbitmq.yaml \
		up -d --remove-orphans

down: ## Down docker compose
	@docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns.yaml \
		-f ops/docker-compose/tooling/saas/airflow/airflow.yaml \
		-f ops/docker-compose/tooling/observability/grafana.yaml \
		-f ops/docker-compose/tooling/observability/grafana-tempo.yaml \
		-f ops/docker-compose/tooling/observability/prometheus.yaml \
		-f ops/docker-compose/tooling/observability/fluent-bit.yaml \
		-f ops/docker-compose/gateway/traefik.yaml \
		-f ops/docker-compose/application/auth.yaml \
		-f ops/docker-compose/application/auth.yaml \
		-f ops/docker-compose/application/api.yaml \
		-f ops/docker-compose/application/metadata.yaml \
		-f ops/docker-compose/application/logger.yaml \
		-f ops/docker-compose/application/ui-next.yaml \
		-f ops/docker-compose/database/mongo.yaml \
		-f ops/docker-compose/database/redis.yaml \
		-f ops/docker-compose/database/patroni.yaml \
		-f ops/docker-compose/database/postgres.yaml \
		-f ops/docker-compose/database/elasticsearch.yaml \
		-f ops/docker-compose/mq/rabbitmq.yaml \
		-f ops/docker-compose/mq/kafka.yaml \
		-f ops/docker-compose/mq/kafka-schema-registry.yaml \
		-f ops/docker-compose/mq/kafka-connect.yaml \
		-f ops/docker-compose/mq/kafka-connector-postgres.yaml \
		-f ops/docker-compose/mq/kafka-connector-elasticsearch.yaml \
	down --remove-orphans
	@docker network prune -f
