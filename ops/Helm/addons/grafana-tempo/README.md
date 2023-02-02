# grafana

![Version: 0.4.0](https://img.shields.io/badge/Version-0.4.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.22.0 || >= v1.22.0-0`

| Repository | Name | Version |
|------------|------|---------|
| file://../../shortlink-common | shortlink-common | 0.2.4 |
| https://grafana.github.io/helm-charts | grafana | 6.48.2 |
| https://grafana.github.io/helm-charts | loki | 3.8.0 |
| https://grafana.github.io/helm-charts | phlare | 0.1.2 |
| https://grafana.github.io/helm-charts | promtail | 6.7.4 |
| https://grafana.github.io/helm-charts | tempo | 0.16.8 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| grafana."grafana.ini".server.root_url | string | `"https://shortlink.best/grafana"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".apiVersion | int | `1` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].disableDeletion | bool | `false` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].editable | bool | `true` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].folder | string | `""` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].name | string | `"default"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].options.path | string | `"/var/lib/grafana/dashboards/default"` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].orgId | int | `1` |  |
| grafana.dashboardProviders."dashboardproviders.yaml".providers[0].type | string | `"file"` |  |
| grafana.dashboards.default.cert-manager.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.cert-manager.gnetId | int | `11001` |  |
| grafana.dashboards.default.go-runtime.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.go-runtime.gnetId | int | `14061` |  |
| grafana.dashboards.default.ingress-nginx.url | string | `"https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/grafana/dashboards/nginx.json"` |  |
| grafana.dashboards.default.istio-control-plane.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-control-plane.gnetId | int | `7645` |  |
| grafana.dashboards.default.istio-galley.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-galley.gnetId | int | `7648` |  |
| grafana.dashboards.default.istio-mesh.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-mesh.gnetId | int | `14061` |  |
| grafana.dashboards.default.istio-mixer.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-mixer.gnetId | int | `7642` |  |
| grafana.dashboards.default.istio-service.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-service.gnetId | int | `7636` |  |
| grafana.dashboards.default.istio-wasm-extension.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-wasm-extension.gnetId | int | `13277` |  |
| grafana.dashboards.default.istio-workload.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.istio-workload.gnetId | int | `7630` |  |
| grafana.dashboards.default.kyverno.url | string | `"https://raw.githubusercontent.com/kyverno/grafana-dashboard/master/grafana/dashboard.json"` |  |
| grafana.dashboards.default.node-exporter.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.node-exporter.gnetId | int | `1860` |  |
| grafana.dashboards.default.rabbitmq-overview.datasource | string | `"Prometheus"` |  |
| grafana.dashboards.default.rabbitmq-overview.gnetId | int | `10991` |  |
| grafana.dashboards.default.rabbitmq-overview.revision | int | `11` |  |
| grafana.datasources."datasources.yaml".apiVersion | int | `1` |  |
| grafana.datasources."datasources.yaml".datasources[0].access | string | `"proxy"` |  |
| grafana.datasources."datasources.yaml".datasources[0].isDefault | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[0].jsonData.alertmanagerUid | string | `"alertmanager"` |  |
| grafana.datasources."datasources.yaml".datasources[0].jsonData.httpMethod | string | `"POST"` |  |
| grafana.datasources."datasources.yaml".datasources[0].jsonData.queryTimeout | string | `"30s"` |  |
| grafana.datasources."datasources.yaml".datasources[0].jsonData.timeInterval | string | `"10s"` |  |
| grafana.datasources."datasources.yaml".datasources[0].name | string | `"Prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].type | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].uid | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[0].url | string | `"http://prometheus-operated.prometheus-operator:9090/prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[1].access | string | `"proxy"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.alertmanagerUid | string | `"alertmanager"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[0].datasourceUid | string | `"tempo"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[0].matcherRegex | string | `"traceID\":\"(\\w+)"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[0].name | string | `"TraceID"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[0].url | string | `"$${__value.raw}"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[1].matcherRegex | string | `"traceID\":\"(\\w+)"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[1].name | string | `"TraceID"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.derivedFields[1].url | string | `"http://grafana-tempo:16686/trace/$${__value.raw}"` |  |
| grafana.datasources."datasources.yaml".datasources[1].jsonData.maxLines | int | `1000` |  |
| grafana.datasources."datasources.yaml".datasources[1].name | string | `"Loki"` |  |
| grafana.datasources."datasources.yaml".datasources[1].type | string | `"loki"` |  |
| grafana.datasources."datasources.yaml".datasources[1].uid | string | `"loki"` |  |
| grafana.datasources."datasources.yaml".datasources[1].url | string | `"http://grafana-grafana-loki:3100"` |  |
| grafana.datasources."datasources.yaml".datasources[2].editable | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.lokiSearch.datasourceUid | string | `"loki"` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.nodeGraph.enabled | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.search.hide | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.serviceMap.datasourceUid | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.spanBar.type | string | `"Duration"` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.tracesToLogs.datasourceUid | string | `"loki"` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.tracesToLogs.filterBySpanID | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.tracesToLogs.filterByTraceID | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[2].jsonData.tracesToLogs.mapTagNamesEnabled | bool | `true` |  |
| grafana.datasources."datasources.yaml".datasources[2].name | string | `"Tempo"` |  |
| grafana.datasources."datasources.yaml".datasources[2].type | string | `"tempo"` |  |
| grafana.datasources."datasources.yaml".datasources[2].uid | string | `"tempo"` |  |
| grafana.datasources."datasources.yaml".datasources[2].url | string | `"http://grafana-tempo:3100"` |  |
| grafana.datasources."datasources.yaml".datasources[3].editable | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[3].jsonData.implementation | string | `"prometheus"` |  |
| grafana.datasources."datasources.yaml".datasources[3].name | string | `"Alertmanager"` |  |
| grafana.datasources."datasources.yaml".datasources[3].readOnly | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[3].type | string | `"alertmanager"` |  |
| grafana.datasources."datasources.yaml".datasources[3].uid | string | `"alertmanager"` |  |
| grafana.datasources."datasources.yaml".datasources[3].url | string | `"http://prometheus-alertmanager.prometheus-operator:9093"` |  |
| grafana.datasources."datasources.yaml".datasources[4].access | string | `"proxy"` |  |
| grafana.datasources."datasources.yaml".datasources[4].basicAuth | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[4].editable | bool | `false` |  |
| grafana.datasources."datasources.yaml".datasources[4].name | string | `"Phlare"` |  |
| grafana.datasources."datasources.yaml".datasources[4].type | string | `"phlare"` |  |
| grafana.datasources."datasources.yaml".datasources[4].uid | string | `"phlare"` |  |
| grafana.datasources."datasources.yaml".datasources[4].url | string | `"http://phlare:4100"` |  |
| grafana.datasources."datasources.yaml".datasources[4].version | int | `1` |  |
| grafana.defaultDashboardsEnabled | bool | `true` |  |
| grafana.enabled | bool | `true` |  |
| grafana.imageRenderer.enabled | bool | `false` |  |
| grafana.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| grafana.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| grafana.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| grafana.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| grafana.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| grafana.ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| grafana.ingress.enabled | bool | `true` |  |
| grafana.ingress.hosts[0] | string | `"shortlink.best"` |  |
| grafana.ingress.path | string | `"/grafana/?(.*)"` |  |
| grafana.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| grafana.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| grafana.namespaceOverride | string | `""` |  |
| grafana.plugins | list | `[]` |  |
| grafana.serviceMonitor.enabled | bool | `true` |  |
| grafana.serviceMonitor.interval | string | `"1m"` |  |
| grafana.serviceMonitor.selfMonitor | bool | `true` |  |
| grafana.sidecar.dashboards.enabled | bool | `true` |  |
| grafana.sidecar.dashboards.searchNamespace | string | `"ALL"` |  |
| grafana.sidecar.datasources.enabled | bool | `true` |  |
| grafana.sidecar.datasources.searchNamespace | string | `"ALL"` |  |
| loki.enabled | bool | `true` |  |
| phlare.enabled | bool | `true` |  |
| promtail.config.clients[0].url | string | `"http://grafana-grafana-loki:3100/loki/api/v1/push"` |  |
| promtail.enabled | bool | `true` |  |
| promtail.extraScrapeConfigs[0].job_name | string | `"syslog"` |  |
| promtail.extraScrapeConfigs[0].syslog.labels.job | string | `"syslog"` |  |
| promtail.extraScrapeConfigs[0].syslog.listen_address | string | `"0.0.0.0:1514"` |  |
| promtail.extraScrapeConfigs[1].job_name | string | `"journal"` |  |
| promtail.extraScrapeConfigs[1].journal.labels.job | string | `"systemd-journal"` |  |
| promtail.extraScrapeConfigs[1].journal.max_age | string | `"12h"` |  |
| promtail.extraScrapeConfigs[1].journal.path | string | `"/var/log/journal"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[0].source_labels[0] | string | `"__syslog_message_hostname"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[0].target_label | string | `"hostname"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[1].source_labels[0] | string | `"__journal__systemd_unit"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[1].target_label | string | `"unit"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[2].source_labels[0] | string | `"__journal__hostname"` |  |
| promtail.extraScrapeConfigs[1].relabel_configs[2].target_label | string | `"hostname"` |  |
| promtail.extraVolumeMounts[0].mountPath | string | `"/var/log/journal"` |  |
| promtail.extraVolumeMounts[0].name | string | `"journal"` |  |
| promtail.extraVolumeMounts[0].readOnly | bool | `true` |  |
| promtail.extraVolumes[0].hostPath.path | string | `"/var/log/journal"` |  |
| promtail.extraVolumes[0].name | string | `"journal"` |  |
| promtail.syslogService.enabled | bool | `true` |  |
| promtail.syslogService.port | int | `1514` |  |
| promtail.syslogService.type | string | `"LoadBalancer"` |  |
| tempo.enabled | bool | `true` |  |
| tempo.serviceMonitor.enabled | bool | `true` |  |
| tempo.tempo.ingester.max_block_bytes | int | `1000000` |  |
| tempo.tempo.ingester.max_block_duration | string | `"5m"` |  |
| tempo.tempo.ingester.trace_idle_period | string | `"10s"` |  |
| tempo.tempo.metricsGenerator.enabled | bool | `true` |  |
| tempo.tempo.metricsGenerator.remoteWriteUrl | string | `"http://prometheus-operated.prometheus-operator:9090/api/v1/write"` |  |
| tempo.tempo.querier.max_concurrent_queries | int | `100` |  |
| tempo.tempo.querier.search.prefer_self | int | `50` |  |
| tempo.tempo.query_frontend.max_outstanding_per_tenant | int | `2000` |  |
| tempo.tempo.query_frontend.search.concurrent_jobs | int | `2000` |  |
| tempo.tempo.query_frontend.search.target_bytes_per_job | int | `400000000` |  |
| tempo.tempo.searchEnabled | bool | `true` |  |
| tempo.tempo.storage.trace.backend | string | `"local"` |  |
| tempo.tempo.storage.trace.block.bloom_filter_false_positive | float | `0.05` |  |
| tempo.tempo.storage.trace.block.encoding | string | `"zstd"` |  |
| tempo.tempo.storage.trace.block.index_downsample_bytes | int | `1000` |  |
| tempo.tempo.storage.trace.block.version | string | `"vParquet"` |  |
| tempo.tempo.storage.trace.local.path | string | `"/var/tempo/traces"` |  |
| tempo.tempo.storage.trace.pool.max_workers | int | `100` |  |
| tempo.tempo.storage.trace.pool.queue_depth | int | `10000` |  |
| tempo.tempo.storage.trace.wal.path | string | `"/var/tempo/wal"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
