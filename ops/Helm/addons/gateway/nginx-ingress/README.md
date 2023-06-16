# nginx-ingress

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.24.0 || >= v1.24.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kubernetes.github.io/ingress-nginx | ingress-nginx | 4.6.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| ingress-nginx.controller.admissionWebhooks.enabled | bool | `false` |  |
| ingress-nginx.controller.config.enable-opentracing | string | `"true"` |  |
| ingress-nginx.controller.config.jaeger-collector-host | string | `"grafana-tempo.grafana"` |  |
| ingress-nginx.controller.config.jaeger-service-name | string | `"nginx-ingress"` |  |
| ingress-nginx.controller.hostNetwork | bool | `false` |  |
| ingress-nginx.controller.ingressClassResource.default | bool | `true` |  |
| ingress-nginx.controller.kind | string | `"DaemonSet"` |  |
| ingress-nginx.controller.metrics.enabled | bool | `true` |  |
| ingress-nginx.controller.metrics.prometheusRule.additionalLabels.app | string | `"kube-prometheus-stack"` |  |
| ingress-nginx.controller.metrics.prometheusRule.additionalLabels.release | string | `"prometheus-operator"` |  |
| ingress-nginx.controller.metrics.prometheusRule.enabled | bool | `true` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].alert | string | `"NGINXConfigFailed"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].annotations.description | string | `"bad ingress config - nginx config test failed"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].annotations.summary | string | `"uninstall the latest ingress changes to allow config reloads to resume"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].expr | string | `"count(nginx_ingress_controller_config_last_reload_successful == 0) > 0"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].for | string | `"1s"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[0].labels.severity | string | `"critical"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].alert | string | `"NGINXCertificateExpiry"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].annotations.description | string | `"ssl certificate(s) will expire in less then a week"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].annotations.summary | string | `"renew expiring certificates to avoid downtime"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].expr | string | `"(avg(nginx_ingress_controller_ssl_expire_time_seconds) by (host) - time()) < 604800"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].for | string | `"1s"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[1].labels.severity | string | `"critical"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].alert | string | `"NGINXTooMany500s"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].annotations.description | string | `"Too many 5XXs"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].annotations.summary | string | `"More than 5% of all requests returned 5XX, this requires your attention"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].expr | string | `"100 * ( sum( nginx_ingress_controller_requests{status=~\"5.+\"} ) / sum(nginx_ingress_controller_requests) ) > 5"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].for | string | `"1m"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[2].labels.severity | string | `"warning"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].alert | string | `"NGINXTooMany400s"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].annotations.description | string | `"Too many 4XXs"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].annotations.summary | string | `"More than 5% of all requests returned 4XX, this requires your attention"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].expr | string | `"100 * ( sum( nginx_ingress_controller_requests{status=~\"4.+\"} ) / sum(nginx_ingress_controller_requests) ) > 5"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].for | string | `"1m"` |  |
| ingress-nginx.controller.metrics.prometheusRule.rules[3].labels.severity | string | `"warning"` |  |
| ingress-nginx.controller.metrics.serviceMonitor.additionalLabels.release | string | `"prometheus-operator"` |  |
| ingress-nginx.controller.metrics.serviceMonitor.enabled | bool | `true` |  |
| ingress-nginx.controller.metrics.serviceMonitor.namespaceSelector.matchNames[0] | string | `"nginx-ingress"` |  |
| ingress-nginx.controller.podSecurityContext.fsGroup | int | `1001` |  |
| ingress-nginx.controller.service.nodePorts.http | int | `80` |  |
| ingress-nginx.controller.service.nodePorts.https | int | `443` |  |
| ingress-nginx.controller.service.type | string | `"NodePort"` |  |
| ingress-nginx.defaultBackend.enabled | bool | `true` |  |
| ingress-nginx.enabled | bool | `true` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
