# postgres-operator

![Version: 1.6.0](https://img.shields.io/badge/Version-1.6.0-informational?style=flat-square) ![AppVersion: 1.6.0](https://img.shields.io/badge/AppVersion-1.6.0-informational?style=flat-square)

Postgres Operator creates and manages PostgreSQL clusters running in Kubernetes

**Homepage:** <https://github.com/zalando/postgres-operator>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Zalando | opensource@zalando.de |  |

## Source Code

* <https://github.com/zalando/postgres-operator>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| configAwsOrGcp.aws_region | string | `"eu-central-1"` |  |
| configAwsOrGcp.enable_ebs_gp3_migration | string | `"false"` |  |
| configConnectionPooler.connection_pooler_default_cpu_limit | string | `"1"` |  |
| configConnectionPooler.connection_pooler_default_cpu_request | string | `"500m"` |  |
| configConnectionPooler.connection_pooler_default_memory_limit | string | `"100Mi"` |  |
| configConnectionPooler.connection_pooler_default_memory_request | string | `"100Mi"` |  |
| configConnectionPooler.connection_pooler_image | string | `"registry.opensource.zalan.do/acid/pgbouncer:master-9"` |  |
| configConnectionPooler.connection_pooler_max_db_connections | string | `"60"` |  |
| configConnectionPooler.connection_pooler_mode | string | `"transaction"` |  |
| configConnectionPooler.connection_pooler_number_of_instances | string | `"2"` |  |
| configConnectionPooler.connection_pooler_schema | string | `"pooler"` |  |
| configConnectionPooler.connection_pooler_user | string | `"pooler"` |  |
| configDebug.debug_logging | string | `"true"` |  |
| configDebug.enable_database_access | string | `"true"` |  |
| configGeneral.docker_image | string | `"registry.opensource.zalan.do/acid/spilo-13:2.0-p2"` |  |
| configGeneral.enable_crd_validation | string | `"true"` |  |
| configGeneral.enable_lazy_spilo_upgrade | string | `"false"` |  |
| configGeneral.enable_pgversion_env_var | string | `"true"` |  |
| configGeneral.enable_shm_volume | string | `"true"` |  |
| configGeneral.enable_spilo_wal_path_compat | string | `"false"` |  |
| configGeneral.etcd_host | string | `""` |  |
| configGeneral.max_instances | string | `"-1"` |  |
| configGeneral.min_instances | string | `"-1"` |  |
| configGeneral.repair_period | string | `"5m"` |  |
| configGeneral.resync_period | string | `"30m"` |  |
| configGeneral.workers | string | `"8"` |  |
| configKubernetes.cluster_domain | string | `"cluster.local"` |  |
| configKubernetes.cluster_labels | string | `"application:spilo"` |  |
| configKubernetes.cluster_name_label | string | `"cluster-name"` |  |
| configKubernetes.enable_init_containers | string | `"true"` |  |
| configKubernetes.enable_pod_antiaffinity | string | `"false"` |  |
| configKubernetes.enable_pod_disruption_budget | string | `"true"` |  |
| configKubernetes.enable_sidecars | string | `"true"` |  |
| configKubernetes.pdb_name_format | string | `"postgres-{cluster}-pdb"` |  |
| configKubernetes.pod_antiaffinity_topology_key | string | `"kubernetes.io/hostname"` |  |
| configKubernetes.pod_management_policy | string | `"ordered_ready"` |  |
| configKubernetes.pod_role_label | string | `"spilo-role"` |  |
| configKubernetes.pod_terminate_grace_period | string | `"5m"` |  |
| configKubernetes.secret_name_template | string | `"{username}.{cluster}.credentials.{tprkind}.{tprgroup}"` |  |
| configKubernetes.spilo_privileged | string | `"false"` |  |
| configKubernetes.storage_resize_mode | string | `"pvc"` |  |
| configKubernetes.watched_namespace | string | `"*"` |  |
| configLoadBalancer.db_hosted_zone | string | `"db.example.com"` |  |
| configLoadBalancer.enable_master_load_balancer | string | `"false"` |  |
| configLoadBalancer.enable_replica_load_balancer | string | `"false"` |  |
| configLoadBalancer.external_traffic_policy | string | `"Cluster"` |  |
| configLoadBalancer.master_dns_name_format | string | `"{cluster}.{team}.{hostedzone}"` |  |
| configLoadBalancer.replica_dns_name_format | string | `"{cluster}-repl.{team}.{hostedzone}"` |  |
| configLoggingRestApi.api_port | string | `"8080"` |  |
| configLoggingRestApi.cluster_history_entries | string | `"1000"` |  |
| configLoggingRestApi.ring_log_lines | string | `"100"` |  |
| configLogicalBackup.logical_backup_docker_image | string | `"registry.opensource.zalan.do/acid/logical-backup:v1.6.0"` |  |
| configLogicalBackup.logical_backup_job_prefix | string | `"logical-backup-"` |  |
| configLogicalBackup.logical_backup_provider | string | `"s3"` |  |
| configLogicalBackup.logical_backup_s3_access_key_id | string | `""` |  |
| configLogicalBackup.logical_backup_s3_bucket | string | `"my-bucket-url"` |  |
| configLogicalBackup.logical_backup_s3_endpoint | string | `""` |  |
| configLogicalBackup.logical_backup_s3_region | string | `""` |  |
| configLogicalBackup.logical_backup_s3_secret_access_key | string | `""` |  |
| configLogicalBackup.logical_backup_s3_sse | string | `"AES256"` |  |
| configLogicalBackup.logical_backup_schedule | string | `"30 00 * * *"` |  |
| configPostgresPodResources.default_cpu_limit | string | `"1"` |  |
| configPostgresPodResources.default_cpu_request | string | `"100m"` |  |
| configPostgresPodResources.default_memory_limit | string | `"500Mi"` |  |
| configPostgresPodResources.default_memory_request | string | `"100Mi"` |  |
| configPostgresPodResources.min_cpu_limit | string | `"250m"` |  |
| configPostgresPodResources.min_memory_limit | string | `"250Mi"` |  |
| configTarget | string | `"ConfigMap"` |  |
| configTeamsApi.enable_postgres_team_crd | string | `"false"` |  |
| configTeamsApi.enable_teams_api | string | `"false"` |  |
| configTimeouts.pod_deletion_wait_timeout | string | `"10m"` |  |
| configTimeouts.pod_label_wait_timeout | string | `"10m"` |  |
| configTimeouts.ready_wait_interval | string | `"3s"` |  |
| configTimeouts.ready_wait_timeout | string | `"30s"` |  |
| configTimeouts.resource_check_interval | string | `"3s"` |  |
| configTimeouts.resource_check_timeout | string | `"10m"` |  |
| configUsers.replication_username | string | `"standby"` |  |
| configUsers.super_username | string | `"postgres"` |  |
| controllerID.create | bool | `false` |  |
| controllerID.name | string | `nil` |  |
| crd.create | bool | `true` |  |
| enableJsonLogging | bool | `false` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.registry | string | `"registry.opensource.zalan.do"` |  |
| image.repository | string | `"acid/postgres-operator"` |  |
| image.tag | string | `"v1.7.0"` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podLabels | object | `{}` |  |
| podPriorityClassName | string | `""` |  |
| podServiceAccount.name | string | `"postgres-pod"` |  |
| priorityClassName | string | `""` |  |
| rbac.create | bool | `true` |  |
| resources.limits.cpu | string | `"500m"` |  |
| resources.limits.memory | string | `"500Mi"` |  |
| resources.requests.cpu | string | `"100m"` |  |
| resources.requests.memory | string | `"250Mi"` |  |
| securityContext.allowPrivilegeEscalation | bool | `false` |  |
| securityContext.readOnlyRootFilesystem | bool | `true` |  |
| securityContext.runAsNonRoot | bool | `true` |  |
| securityContext.runAsUser | int | `1000` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `nil` |  |
| tolerations | list | `[]` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
