# shortlink-bot

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink service for get bot by URL

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Source Code

* <https://github.com/shortlink-org/shortlink>

## Requirements

Kubernetes: `>= 1.22.0 || >= v1.22.0-0`

| Repository | Name | Version |
|------------|------|---------|
| file://../shortlink-common | shortlink-common | 0.2.25 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| commonAnnotations | object | `{}` | Add annotations to all the deployed resources |
| commonLabels | object | `{}` | Add labels to all the deployed resources |
| deploy.affinity | list | `[]` |  |
| deploy.annotations | object | `{}` | Annotations to be added to controller pods |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
| deploy.env.STORE_POSTGRES_URI | string | `"postgres://postgres:shortlink@postgresql.postgresql:5432/shortlink?sslmode=disable"` | Default store config |
| deploy.env.TELEGRAM_BOT_TOKEN | string | `"YOUR_TELEGRAM_TOKEN"` |  |
| deploy.env.TELEGRAM_BOT_USERNAME | string | `"shortlink_my_bot"` |  |
| deploy.env.TRACER_URI | string | `"grafana-tempo.grafana:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` | Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images |
| deploy.image.repository | string | `"registry.gitlab.com/shortlink-org/shortlink/bot"` |  |
| deploy.image.tag | string | `"0.13.88"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe | object | `{"failureThreshold":1,"httpGet":{"path":"/live","port":9090},"initialDelaySeconds":5,"periodSeconds":5,"successThreshold":1}` | define a liveness probe that checks every 5 seconds, starting after 5 seconds |
| deploy.nodeSelector | list | `[]` | Node labels and tolerations for pod assignment ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#taints-and-tolerations-beta-feature |
| deploy.podSecurityContext.fsGroup | int | `1000` | fsGroup is the group ID associated with the container |
| deploy.readinessProbe | object | `{"failureThreshold":30,"httpGet":{"path":"/ready","port":9090},"initialDelaySeconds":5,"periodSeconds":5,"successThreshold":1}` | define a readiness probe that checks every 5 seconds, starting after 5 seconds |
| deploy.replicaCount | int | `1` |  |
| deploy.resources.limits | object | `{"cpu":"100m","memory":"128Mi"}` | We usually recommend not to specify default resources and to leave this as a conscious choice for the user. This also increases chances charts run on environments with little resources, such as Minikube. If you do want to specify resources, uncomment the following lines, adjust them as necessary, and remove the curly braces after 'resources:'. |
| deploy.resources.requests.cpu | string | `"10m"` |  |
| deploy.resources.requests.memory | string | `"32Mi"` |  |
| deploy.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":"true","runAsGroup":1000,"runAsNonRoot":true,"runAsUser":1000}` | Security Context policies for controller pods See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls |
| deploy.strategy.rollingUpdate.maxSurge | int | `1` |  |
| deploy.strategy.rollingUpdate.maxUnavailable | int | `0` |  |
| deploy.strategy.type | string | `"RollingUpdate"` |  |
| deploy.terminationGracePeriodSeconds | int | `90` |  |
| deploy.tolerations | list | `[]` |  |
| fullnameOverride | string | `""` |  |
| monitoring.enabled | bool | `true` | Creates a Prometheus Operator ServiceMonitor |
| monitoring.jobLabel | string | `""` | The label to use to retrieve the job name from. |
| monitoring.labels | object | `{"release":"prometheus-operator"}` | Additional labels that can be used so PodMonitor will be discovered by Prometheus |
| nameOverride | string | `""` |  |
| secret.enabled | bool | `false` |  |
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nYour CA...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nYour cert...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nYour key...\n-----END EC PRIVATE KEY-----\n"` |  |
| service.ports | list | `[]` |  |
| service.type | string | `"ClusterIP"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
