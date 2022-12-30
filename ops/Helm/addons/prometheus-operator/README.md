# prometheus-operator

![Version: 0.5.0](https://img.shields.io/badge/Version-0.5.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.22.0 || >= v1.22.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://prometheus-community.github.io/helm-charts | kube-prometheus-stack | 43.2.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| kube-prometheus-stack.alertmanager.alertmanagerSpec.alertmanagerConfigSelector.matchLabels.alertmanagerConfig | string | `"shortlink"` |  |
| kube-prometheus-stack.alertmanager.alertmanagerSpec.externalUrl | string | `"https://shortlink.best/alertmanager"` |  |
| kube-prometheus-stack.alertmanager.alertmanagerSpec.logFormat | string | `"json"` |  |
| kube-prometheus-stack.alertmanager.alertmanagerSpec.routePrefix | string | `"/alertmanager/"` |  |
| kube-prometheus-stack.alertmanager.config.global.resolve_timeout | string | `"5m"` |  |
| kube-prometheus-stack.alertmanager.config.receivers[0].name | string | `"null"` |  |
| kube-prometheus-stack.alertmanager.config.route.group_by[0] | string | `"job"` |  |
| kube-prometheus-stack.alertmanager.config.route.group_interval | string | `"5m"` |  |
| kube-prometheus-stack.alertmanager.config.route.group_wait | string | `"30s"` |  |
| kube-prometheus-stack.alertmanager.config.route.receiver | string | `"null"` |  |
| kube-prometheus-stack.alertmanager.config.route.repeat_interval | string | `"12h"` |  |
| kube-prometheus-stack.alertmanager.config.route.routes[0].match.alertname | string | `"Watchdog"` |  |
| kube-prometheus-stack.alertmanager.config.route.routes[0].receiver | string | `"null"` |  |
| kube-prometheus-stack.alertmanager.config.templates[0] | string | `"/etc/alertmanager/config/*.tmpl"` |  |
| kube-prometheus-stack.alertmanager.enabled | bool | `true` |  |
| kube-prometheus-stack.alertmanager.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| kube-prometheus-stack.alertmanager.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| kube-prometheus-stack.alertmanager.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"false"` |  |
| kube-prometheus-stack.alertmanager.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| kube-prometheus-stack.alertmanager.ingress.enabled | bool | `true` |  |
| kube-prometheus-stack.alertmanager.ingress.hosts[0] | string | `"shortlink.best"` |  |
| kube-prometheus-stack.alertmanager.ingress.ingressClassName | string | `"nginx"` |  |
| kube-prometheus-stack.alertmanager.ingress.pathType | string | `"Prefix"` |  |
| kube-prometheus-stack.alertmanager.ingress.paths[0] | string | `"/alertmanager"` |  |
| kube-prometheus-stack.alertmanager.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| kube-prometheus-stack.alertmanager.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| kube-prometheus-stack.coreDns.enabled | bool | `true` |  |
| kube-prometheus-stack.enabled | bool | `true` |  |
| kube-prometheus-stack.fullnameOverride | string | `"prometheus"` |  |
| kube-prometheus-stack.grafana.enabled | bool | `false` |  |
| kube-prometheus-stack.kubeDns.enabled | bool | `false` |  |
| kube-prometheus-stack.prometheus.enabled | bool | `true` |  |
| kube-prometheus-stack.prometheus.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| kube-prometheus-stack.prometheus.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| kube-prometheus-stack.prometheus.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"false"` |  |
| kube-prometheus-stack.prometheus.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| kube-prometheus-stack.prometheus.ingress.enabled | bool | `true` |  |
| kube-prometheus-stack.prometheus.ingress.hosts[0] | string | `"shortlink.best"` |  |
| kube-prometheus-stack.prometheus.ingress.ingressClassName | string | `"nginx"` |  |
| kube-prometheus-stack.prometheus.ingress.pathType | string | `"Prefix"` |  |
| kube-prometheus-stack.prometheus.ingress.paths[0] | string | `"/prometheus"` |  |
| kube-prometheus-stack.prometheus.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| kube-prometheus-stack.prometheus.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues | bool | `false` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.retention | string | `"3d"` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.routePrefix | string | `"/prometheus/"` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.serviceMonitorNamespaceSelector | object | `{}` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.serviceMonitorSelector | object | `{}` |  |
| kube-prometheus-stack.prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues | bool | `false` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
