# DEV ========================================================================
dev: ### Run the development environment
	@echo "Running the development environment..."
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
    		-f $(ROOT_DIR)/docker-compose.yaml \
    		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
			-f $(ROOT_DIR)/ops/docker-compose/tooling/observability/grafana/grafana-alloy.yaml \
    		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
    		up -d --remove-orphans --build

down: confirm ## Down docker compose
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(ROOT_DIR)/docker-compose.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/observability/grafana/grafana-alloy.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
	down --remove-orphans
	@docker network prune -f
