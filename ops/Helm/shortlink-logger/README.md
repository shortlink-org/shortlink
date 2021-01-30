# shortlink-logger

![Version: 0.5.5](https://img.shields.io/badge/Version-0.5.5-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink logger service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Source Code

* <https://github.com/batazor/shortlink>

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| deploy.affinity | list | `[]` |  |
| deploy.annotations | object | `{}` |  |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.SERVICE_NAME | string | `"Logger"` |  |
| deploy.env.TRACER_URI | string | `"jaeger-agent.jaeger-operator:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` |  |
| deploy.image.repository | string | `"batazor/shortlink-logger"` |  |
| deploy.image.tag | string | `"latest"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe.failureThreshold | int | `1` |  |
| deploy.livenessProbe.httpGet.path | string | `"/live"` |  |
| deploy.livenessProbe.httpGet.port | int | `9090` |  |
| deploy.livenessProbe.initialDelaySeconds | int | `5` |  |
| deploy.livenessProbe.periodSeconds | int | `5` |  |
| deploy.livenessProbe.successThreshold | int | `1` |  |
| deploy.nodeSelector | object | `{}` |  |
| deploy.podSecurityContext.fsGroup | int | `1000` |  |
| deploy.readinessProbe.failureThreshold | int | `30` |  |
| deploy.readinessProbe.httpGet.path | string | `"/ready"` |  |
| deploy.readinessProbe.httpGet.port | int | `9090` |  |
| deploy.readinessProbe.initialDelaySeconds | int | `5` |  |
| deploy.readinessProbe.periodSeconds | int | `5` |  |
| deploy.readinessProbe.successThreshold | int | `1` |  |
| deploy.replicaCount | int | `1` |  |
| deploy.resources.limits.cpu | string | `"100m"` |  |
| deploy.resources.limits.memory | string | `"128Mi"` |  |
| deploy.resources.requests.cpu | string | `"100m"` |  |
| deploy.resources.requests.memory | string | `"128Mi"` |  |
| deploy.securityContext.allowPrivilegeEscalation | bool | `false` |  |
| deploy.securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| deploy.securityContext.readOnlyRootFilesystem | bool | `true` |  |
| deploy.securityContext.runAsGroup | int | `1000` |  |
| deploy.securityContext.runAsNonRoot | bool | `true` |  |
| deploy.securityContext.runAsUser | int | `1000` |  |
| deploy.strategy.rollingUpdate.maxSurge | int | `1` |  |
| deploy.strategy.rollingUpdate.maxUnavailable | int | `0` |  |
| deploy.strategy.type | string | `"RollingUpdate"` |  |
| deploy.terminationGracePeriodSeconds | int | `90` |  |
| deploy.tolerations | list | `[]` |  |
| fullnameOverride | string | `""` |  |
| nameOverride | string | `""` |  |
| service.port | int | `7070` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"shortlink"` |  |

