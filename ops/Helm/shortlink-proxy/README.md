# shortlink-proxy

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink service for get proxy by URL

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
| deploy.env.API_LINK_SERVICE | string | `"http://shortlink-api.shortlink:7070"` |  |
| deploy.env.GRPC_CLIENT_HOST | string | `"istio-ingress.istio-ingress"` |  |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
| deploy.env.STORE_POSTGRES_URI | string | `"postgres://postgres:shortlink@postgresql.postgresql:5432/shortlink?sslmode=disable"` | Default store config |
| deploy.env.TRACER_URI | string | `"grafana-tempo.grafana:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` | Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images |
| deploy.image.repository | string | `"registry.gitlab.com/shortlink-org/shortlink/proxy"` |  |
| deploy.image.tag | string | `"0.13.88"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe | object | `{"httpGet":{"path":"/ready","port":3020}}` | define a liveness probe that checks every 5 seconds, starting after 5 seconds |
| deploy.nodeSelector | list | `[]` | Node labels and tolerations for pod assignment ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#taints-and-tolerations-beta-feature |
| deploy.podSecurityContext.fsGroup | int | `1000` | fsGroup is the group ID associated with the container |
| deploy.readinessProbe | object | `{"httpGet":{"path":"/ready","port":3020}}` | define a readiness probe that checks every 5 seconds, starting after 5 seconds |
| deploy.replicaCount | int | `1` |  |
| deploy.resources.limits | object | `{"cpu":"100m","memory":"1024Mi"}` | We usually recommend not to specify default resources and to leave this as a conscious choice for the user. This also increases chances charts run on environments with little resources, such as Minikube. If you do want to specify resources, uncomment the following lines, adjust them as necessary, and remove the curly braces after 'resources:'. |
| deploy.resources.requests.cpu | string | `"10m"` |  |
| deploy.resources.requests.memory | string | `"64Mi"` |  |
| deploy.securityContext | object | `{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":"true","runAsGroup":1000,"runAsNonRoot":true,"runAsUser":1000}` | Security Context policies for controller pods See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls |
| deploy.strategy.rollingUpdate.maxSurge | int | `1` |  |
| deploy.strategy.rollingUpdate.maxUnavailable | int | `0` |  |
| deploy.strategy.type | string | `"RollingUpdate"` |  |
| deploy.terminationGracePeriodSeconds | int | `90` |  |
| deploy.tolerations | list | `[]` |  |
| enabled | bool | `true` |  |
| fullnameOverride | string | `""` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"false"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"false"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"false"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/s/$2"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hostname | string | `"shortlink.best"` |  |
| ingress.path | string | `"/s(/|$)(.*)"` |  |
| ingress.service.name | string | `"shortlink-proxy"` |  |
| ingress.service.port | int | `3020` |  |
| ingress.type | string | `"nginx"` |  |
| monitoring.enabled | bool | `true` | Creates a Prometheus Operator ServiceMonitor |
| monitoring.jobLabel | string | `""` | The label to use to retrieve the job name from. |
| monitoring.labels | object | `{"release":"prometheus-operator"}` | Additional labels that can be used so PodMonitor will be discovered by Prometheus |
| nameOverride | string | `""` |  |
| secret.enabled | bool | `false` |  |
| service.ports[0].name | string | `"http"` |  |
| service.ports[0].port | int | `3020` |  |
| service.ports[0].protocol | string | `"TCP"` |  |
| service.ports[0].public | bool | `true` |  |
| service.ports[0].targetPort | int | `3020` |  |
| service.type | string | `"ClusterIP"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
