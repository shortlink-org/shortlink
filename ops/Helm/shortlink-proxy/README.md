# shortlink-proxy

![Version: 0.1.3](https://img.shields.io/badge/Version-0.1.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink service for get proxy by URL

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
| deploy.env.API_LINK_SERVICE | string | `"http://api-shortlink-api.shortlink:7070"` |  |
| deploy.env.GRPC_CLIENT_HOST | string | `"istio-ingressgateway.istio-system"` |  |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.MQ_RABBIT_URI | string | `"amqp://admin:admin@rabbitmq.rabbitmq:5672"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
| deploy.env.STORE_POSTGRES_URI | string | `"postgres://postgres:shortlink@postgres-postgresql.postgres:5432/shortlink?sslmode=disable"` |  |
| deploy.env.TRACER_URI | string | `"jaeger-agent.jaeger-operator:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` |  |
| deploy.image.repository | string | `"batazor/shortlink-proxy"` |  |
| deploy.image.tag | string | `"latest"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe.failureThreshold | int | `1` |  |
| deploy.livenessProbe.httpGet.path | string | `"/ready"` |  |
| deploy.livenessProbe.httpGet.port | int | `3030` |  |
| deploy.livenessProbe.initialDelaySeconds | int | `5` |  |
| deploy.livenessProbe.periodSeconds | int | `5` |  |
| deploy.livenessProbe.successThreshold | int | `1` |  |
| deploy.nodeSelector | list | `[]` |  |
| deploy.podSecurityContext.fsGroup | int | `1000` |  |
| deploy.readinessProbe.failureThreshold | int | `30` |  |
| deploy.readinessProbe.httpGet.path | string | `"/ready"` |  |
| deploy.readinessProbe.httpGet.port | int | `3030` |  |
| deploy.readinessProbe.initialDelaySeconds | int | `5` |  |
| deploy.readinessProbe.periodSeconds | int | `5` |  |
| deploy.readinessProbe.successThreshold | int | `1` |  |
| deploy.replicaCount | int | `1` |  |
| deploy.resources.limits.cpu | string | `"100m"` |  |
| deploy.resources.limits.memory | string | `"1024Mi"` |  |
| deploy.resources.requests.cpu | string | `"10m"` |  |
| deploy.resources.requests.memory | string | `"64Mi"` |  |
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
| enabled | bool | `true` |  |
| fullnameOverride | string | `""` |  |
| host | string | `"shortlink.ddns.net"` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$2"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| ingress.enabled | bool | `false` |  |
| ingress.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| ingress.type | string | `"nginx"` |  |
| nameOverride | string | `""` |  |
| secret.enabled | bool | `false` |  |
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nYour CA...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nYour cert...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nYour key...\n-----END EC PRIVATE KEY-----\n"` |  |
| service.port | int | `3030` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"shortlink"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)
