# kyverno

![Version: 0.2.1](https://img.shields.io/badge/Version-0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kyverno.github.io/kyverno | kyverno | 3.0.5 |
| https://kyverno.github.io/kyverno | kyverno-policies | 3.0.4 |
| https://kyverno.github.io/policy-reporter | policy-reporter | 2.20.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| kyverno-policies.background | bool | `false` |  |
| kyverno-policies.enabled | bool | `true` |  |
| kyverno-policies.failurePolicy | string | `"Ignore"` |  |
| kyverno-policies.podSecuritySeverity | string | `"low"` |  |
| kyverno-policies.validationFailureActionByPolicy.disallow-capabilities-strict | string | `"audit"` |  |
| kyverno-policies.validationFailureActionByPolicy.disallow-host-path | string | `"audit"` |  |
| kyverno-policies.validationFailureActionByPolicy.disallow-host-ports | string | `"audit"` |  |
| kyverno.admissionController.hostNetwork | bool | `false` |  |
| kyverno.admissionController.serviceMonitor.additionalLabels.release | string | `"prometheus-operator"` |  |
| kyverno.admissionController.serviceMonitor.enabled | bool | `true` |  |
| kyverno.admissionController.tracing.address | string | `"grafana-tempo.grafana"` |  |
| kyverno.admissionController.tracing.enabled | bool | `true` |  |
| kyverno.admissionController.tracing.port | int | `4317` |  |
| kyverno.backgroundController.enabled | bool | `true` |  |
| kyverno.backgroundController.serviceMonitor.additionalLabels.release | string | `"prometheus-operator"` |  |
| kyverno.backgroundController.serviceMonitor.enabled | bool | `true` |  |
| kyverno.backgroundController.tracing.address | string | `"grafana-tempo.grafana"` |  |
| kyverno.backgroundController.tracing.enabled | bool | `true` |  |
| kyverno.backgroundController.tracing.port | int | `4317` |  |
| kyverno.cleanupController.enabled | bool | `true` |  |
| kyverno.cleanupController.logging.format | string | `"json"` |  |
| kyverno.cleanupController.networkPolicy.enabled | bool | `true` |  |
| kyverno.cleanupController.serviceMonitor.additionalLabels.release | string | `"prometheus-operator"` |  |
| kyverno.cleanupController.serviceMonitor.enabled | bool | `true` |  |
| kyverno.cleanupController.tracing.address | string | `"grafana-tempo.grafana"` |  |
| kyverno.cleanupController.tracing.enabled | bool | `true` |  |
| kyverno.cleanupController.tracing.port | int | `4317` |  |
| kyverno.enabled | bool | `true` |  |
| kyverno.reportsController.enabled | bool | `true` |  |
| kyverno.reportsController.serviceMonitor.additionalLabels.release | string | `"prometheus-operator"` |  |
| kyverno.reportsController.serviceMonitor.enabled | bool | `true` |  |
| kyverno.reportsController.tracing.address | string | `"grafana-tempo.grafana"` |  |
| kyverno.reportsController.tracing.enabled | bool | `true` |  |
| kyverno.reportsController.tracing.port | int | `4317` |  |
| kyverno.webhooksCleanup.enabled | bool | `false` |  |
| policy-reporter.enabled | bool | `true` |  |
| policy-reporter.global.plugins.kyverno | bool | `true` |  |
| policy-reporter.grafana.folder.annotation | string | `"grafana_dashboard_folder"` |  |
| policy-reporter.grafana.folder.name | string | `"Security"` |  |
| policy-reporter.kyvernoPlugin.enabled | bool | `true` |  |
| policy-reporter.metrics.enabled | bool | `true` |  |
| policy-reporter.monitoring.enabled | bool | `true` |  |
| policy-reporter.networkPolicy.enabled | bool | `false` |  |
| policy-reporter.resources.limits.cpu | string | `"100m"` |  |
| policy-reporter.resources.limits.memory | string | `"128Mi"` |  |
| policy-reporter.resources.requests.cpu | string | `"5m"` |  |
| policy-reporter.resources.requests.memory | string | `"75Mi"` |  |
| policy-reporter.rest.enabled | bool | `true` |  |
| policy-reporter.target.loki.host | string | `"http://grafana-loki.grafana:3100"` |  |
| policy-reporter.target.loki.minimumPriority | string | `"warning"` |  |
| policy-reporter.target.loki.skipExistingOnStartup | bool | `true` |  |
| policy-reporter.target.loki.sources[0] | string | `"kyverno"` |  |
| policy-reporter.ui.enabled | bool | `true` |  |
| policy-reporter.ui.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"false"` |  |
| policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| policy-reporter.ui.ingress.className | string | `"nginx"` |  |
| policy-reporter.ui.ingress.enabled | bool | `true` |  |
| policy-reporter.ui.ingress.hosts[0].host | string | `"shortlink.best"` |  |
| policy-reporter.ui.ingress.hosts[0].paths[0].path | string | `"/kyverno/?(.*)"` |  |
| policy-reporter.ui.ingress.hosts[0].paths[0].pathType | string | `"Prefix"` |  |
| policy-reporter.ui.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| policy-reporter.ui.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| policy-reporter.ui.plugins.kyverno | bool | `true` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
