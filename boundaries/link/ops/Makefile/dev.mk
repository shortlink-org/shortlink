# DEV ========================================================================
dev: ### Run the development environment
	@echo "Running the development environment..."
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(ROOT_DIR)/docker-compose.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/observability/grafana/grafana-alloy.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/postgres/postgres.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/minio/minio.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/application/auth/kratos/kratos.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/application/auth/spicedb/spicedb.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/gateway/traefik/traefik.yaml \
    up -d --remove-orphans --build

down: confirm ## Down docker compose
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(ROOT_DIR)/docker-compose.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/observability/grafana/grafana-alloy.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/postgres/postgres.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/minio/minio.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/application/auth/kratos/kratos.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/application/auth/spicedb/spicedb.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/gateway/traefik/traefik.yaml \
	down --remove-orphans
	@docker network prune -f
