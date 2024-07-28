# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/ops/Makefile/docker.mk

### Runnings ===========================================================================================================

up: ## Run for development mode
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f ops/docker-compose/database/redis/redis.yaml \
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
		-f ops/docker-compose/tooling/registry/docker-registry/docker-registry.yaml \
		-f ops/docker-compose/tooling/registry/zot/zot.yaml \
		-f ops/docker-compose/tooling/services/coredns/coredns.yaml \
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
