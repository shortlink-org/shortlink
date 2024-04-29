# DOCKER TASKS =========================================================================================================
# This is the default. It can be overridden in the main Makefile after
# including docker.mk

# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

DOCKER_USERNAME := "batazor"
DOCKER_BUILDKIT := 1
# disable becouse more images don't have signature
DOCKER_CONTENT_TRUST := 0
BUILDX_GIT_LABELS := 1
BUILDX_EXPERIMENTAL := 1
SOURCE_DATE_EPOCH := $(git log -1 --pretty=%ct)
export CURRENT_UID=$(id -u):$(id -g)

# DOCKER TASKS =========================================================================================================
docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

### Runnings ===========================================================================================================

dev: ## Run for development mode
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f ops/docker-compose/gateway/traefik/traefik.yaml \
		-f ops/docker-compose/application/auth/kratos/kratos.yaml \
		-f ops/docker-compose/database/postgres/postgres.yaml \
		up -d --remove-orphans --build

watch: ## Run for development mode with watch
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f ops/docker-compose/database/redis/redis.yaml \
		-f ops/docker-compose/application/api/api.yaml \
		watch

down: confirm ## Down docker compose
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f ops/docker-compose/tooling/services/docker-registry/docker-registry.yaml \
		-f ops/docker-compose/tooling/services/feature-toggle/feature-toggle.yaml \
		-f ops/docker-compose/tooling/saas/airflow/airflow.yaml \
		-f ops/docker-compose/tooling/saas/nifi/nifi.yaml \
		-f ops/docker-compose/tooling/saas/gitlab/gitlab.yaml \
		-f ops/docker-compose/tooling/saas/novu/novu.yaml \
		-f ops/docker-compose/tooling/saas/localstack/localstack.yaml \
		-f ops/docker-compose/tooling/saas/temporal/temporal.yaml \
		-f ops/docker-compose/tooling/observability/grafana/grafana-loki.yaml \
		-f ops/docker-compose/tooling/observability/grafana/grafana-tempo.yaml \
		-f ops/docker-compose/tooling/observability/grafana/grafana-pyroscope.yaml \
		-f ops/docker-compose/tooling/observability/grafana/grafana-oncall.yaml \
		-f ops/docker-compose/tooling/observability/grafana/grafana-beyla.yaml \
		-f ops/docker-compose/tooling/observability/prometheus/prometheus.yaml \
		-f ops/docker-compose/tooling/observability/fluent-bit/fluent-bit.yaml \
		-f ops/docker-compose/tooling/minio/minio.yaml \
		-f ops/docker-compose/gateway/caddy/caddy.yaml \
		-f ops/docker-compose/gateway/nginx/nginx.yaml \
		-f ops/docker-compose/gateway/traefik/traefik.yaml \
		-f ops/docker-compose/application/auth/keycloak/keycloak.yaml \
		-f ops/docker-compose/application/auth/kratos/kratos.yaml \
		-f ops/docker-compose/application/auth/hydra/hydra.yaml \
		-f ops/docker-compose/application/auth/keto/keto.yaml \
		-f ops/docker-compose/application/auth/spicedb/spicedb.yaml \
		-f ops/docker-compose/application/api/api.yaml \
		-f ops/docker-compose/application/metadata/metadata.yaml \
		-f ops/docker-compose/application/logger/logger.yaml \
		-f ops/docker-compose/application/support/support.yaml \
		-f ops/docker-compose/application/ui-next/ui-next.yaml \
		-f ops/docker-compose/database/aerospike/aerospike.yaml \
		-f ops/docker-compose/database/cassandra/cassandra.yaml \
		-f ops/docker-compose/database/clickhouse/clickhouse.yaml \
		-f ops/docker-compose/database/couchbase/couchbase.yaml \
		-f ops/docker-compose/database/mysql/mysql.yaml \
		-f ops/docker-compose/database/trino/trino.yaml \
		-f ops/docker-compose/database/cockroachdb/cockroachdb.yaml \
		-f ops/docker-compose/database/dgraph/dgraph.yaml \
		-f ops/docker-compose/database/dragonfly/dragonfly.yaml \
		-f ops/docker-compose/database/edgedb/edgedb.yaml \
		-f ops/docker-compose/database/elasticsearch/elasticsearch.yaml \
		-f ops/docker-compose/database/etcd/etcd.yaml \
		-f ops/docker-compose/database/foundation/foundation.yaml \
		-f ops/docker-compose/database/ignite/ignite.yaml \
		-f ops/docker-compose/database/mongo/mongo.yaml \
		-f ops/docker-compose/database/neo4j/neo4j.yaml \
		-f ops/docker-compose/database/surrealdb/surrealdb.yaml \
		-f ops/docker-compose/database/postgres/postgres.yaml \
		-f ops/docker-compose/database/postgres/pgbouncer.yaml \
		-f ops/docker-compose/database/postgres/patroni.yaml \
		-f ops/docker-compose/database/postgres/backup.yaml \
		-f ops/docker-compose/database/redis/redis.yaml \
		-f ops/docker-compose/database/rethinkdb/rethinkdb.yaml \
		-f ops/docker-compose/database/scylla/scylla.yaml \
		-f ops/docker-compose/database/tarantool/tarantool.yaml \
		-f ops/docker-compose/database/tidb/tidb.yaml \
		-f ops/docker-compose/database/victoria-metrics/victoria-metrics.yaml \
		-f ops/docker-compose/mq/rabbitmq/rabbitmq.yaml \
		-f ops/docker-compose/mq/kafka/kafka-connector-elasticsearch.yaml \
		-f ops/docker-compose/mq/kafka/kafka.yaml \
		-f ops/docker-compose/mq/kafka/kafka-schema-registry.yaml \
		-f ops/docker-compose/mq/kafka/kafka-connect.yaml \
		-f ops/docker-compose/mq/kafka/kafka-connector-postgres.yaml \
		-f ops/docker-compose/mq/kafka/kafka-connector-elasticsearch.yaml \
		-f ops/docker-compose/mq/zookeeper/zookeeper.yaml \
		-f ops/docker-compose/mq/pulsar/pulsar.yaml \
		-f ops/docker-compose/mq/nats/nats.yaml \
	down --remove-orphans
	@docker network prune -f

### Helpers ============================================================================================================

docker_ip: ## View docker ip and container name
	@docker ps -q | xargs docker inspect --format "{{range .NetworkSettings.Networks}}{{print .IPAddress}} {{end}}{{.Name}}"

