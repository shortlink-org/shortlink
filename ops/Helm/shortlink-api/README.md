# shortlink-api

![Version: 0.8.3](https://img.shields.io/badge/Version-0.8.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink API service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Source Code

* <https://github.com/batazor/shortlink>

## Requirements

Kubernetes: `>= 1.21.0 || >= v1.21.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://k8s.ory.sh/helm/charts | kratos | 0.23.1 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| NetworkPolicy.enabled | bool | `false` |  |
| database.postgres.enable | bool | `true` |  |
| deploy.affinity | list | `[]` |  |
| deploy.annotations | object | `{}` |  |
| deploy.env.GRPC_CLIENT_HOST | string | `"istio-ingressgateway.istio-system.svc.cluster.local"` |  |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.MQ_RABBIT_URI | string | `"amqp://admin:admin@rabbitmq.rabbitmq:5672"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
| deploy.env.STORE_POSTGRES_URI | string | `"postgres://postgres:shortlink@postgresql.postgresql:5432/shortlink?sslmode=disable"` |  |
| deploy.env.TRACER_URI | string | `"grafana-tempo.grafana:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` |  |
| deploy.image.repository | string | `"batazor/shortlink-api"` |  |
| deploy.image.tag | string | `"latest"` |  |
| deploy.imagePullSecrets | list | `[]` |  |
| deploy.livenessProbe.failureThreshold | int | `1` |  |
| deploy.livenessProbe.httpGet.path | string | `"/live"` |  |
| deploy.livenessProbe.httpGet.port | int | `9090` |  |
| deploy.livenessProbe.initialDelaySeconds | int | `5` |  |
| deploy.livenessProbe.periodSeconds | int | `5` |  |
| deploy.livenessProbe.successThreshold | int | `1` |  |
| deploy.nodeSelector | list | `[]` |  |
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
| deploy.resources.requests.cpu | string | `"10m"` |  |
| deploy.resources.requests.memory | string | `"32Mi"` |  |
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
| external_database.enable | bool | `false` |  |
| external_database.ip | string | `"192.168.0.101"` |  |
| external_database.port | int | `6379` |  |
| fullnameOverride | string | `""` |  |
| host | string | `"shortlink.ddns.net"` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| ingress.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.enabled | bool | `false` |  |
| ingress.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| ingress.type | string | `"nginx"` |  |
| kratos.enabled | bool | `true` |  |
| kratos.fullnameOverride | string | `"kratos"` |  |
| kratos.image.tag | string | `"v0.9.0-alpha.3"` |  |
| kratos.ingress.admin.enabled | bool | `false` |  |
| kratos.ingress.public.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| kratos.ingress.public.annotations."kubernetes.io/ingress.class" | string | `"nginx"` |  |
| kratos.ingress.public.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| kratos.ingress.public.enabled | bool | `true` |  |
| kratos.ingress.public.hosts[0].host | string | `"shortlink.ddns.net"` |  |
| kratos.ingress.public.hosts[0].paths[0].path | string | `"/api/auth\\/(.*)"` |  |
| kratos.ingress.public.hosts[0].paths[0].pathType | string | `"Prefix"` |  |
| kratos.ingress.public.tls[0].hosts[0] | string | `"shortlink.ddns.net"` |  |
| kratos.ingress.public.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| kratos.kratos.autoMigrate | bool | `false` |  |
| kratos.kratos.config.courier.smtp.connection_uri | string | `"smtps://test:test@mailslurper:1025/?skip_ssl_verify=true"` |  |
| kratos.kratos.config.courier.smtp.from_address | string | `"no-reply@shortlink.com"` |  |
| kratos.kratos.config.dsn | string | `"memory"` |  |
| kratos.kratos.config.hashers.argon2.iterations | int | `2` |  |
| kratos.kratos.config.hashers.argon2.key_length | int | `16` |  |
| kratos.kratos.config.hashers.argon2.memory | string | `"128MB"` |  |
| kratos.kratos.config.hashers.argon2.parallelism | int | `1` |  |
| kratos.kratos.config.hashers.argon2.salt_length | int | `16` |  |
| kratos.kratos.config.identity.default_schema_id | string | `"default"` |  |
| kratos.kratos.config.identity.schemas[0].id | string | `"default"` |  |
| kratos.kratos.config.identity.schemas[0].url | string | `"file:///etc/config/identity.default.schema.json"` |  |
| kratos.kratos.config.log.format | string | `"json"` |  |
| kratos.kratos.config.log.leak_sensitive_values | bool | `true` |  |
| kratos.kratos.config.log.level | string | `"debug"` |  |
| kratos.kratos.config.secrets.cookie[0] | string | `"PLEASE-CHANGE-ME-I-AM-VERY-INSECURE"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[0] | string | `"*"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[1] | string | `"http://*"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[2] | string | `"https://*"` |  |
| kratos.kratos.config.selfservice.default_browser_return_url | string | `"https://shortlink.ddns.net"` |  |
| kratos.kratos.config.selfservice.flows.error.ui_url | string | `"https://shortlink.ddns.net/next/error"` |  |
| kratos.kratos.config.selfservice.flows.login.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.login.ui_url | string | `"https://shortlink.ddns.net/next"` |  |
| kratos.kratos.config.selfservice.flows.logout.after.default_browser_return_url | string | `"https://shortlink.ddns.net/next/auth/login"` |  |
| kratos.kratos.config.selfservice.flows.recovery.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.recovery.ui_url | string | `"https://shortlink.ddns.net/next/auth/recovery"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.oidc.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.password.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.registration.ui_url | string | `"https://shortlink.ddns.net/next/auth/registration"` |  |
| kratos.kratos.config.selfservice.flows.settings.privileged_session_max_age | string | `"15m"` |  |
| kratos.kratos.config.selfservice.flows.settings.ui_url | string | `"https://shortlink.ddns.net/next/user/profile"` |  |
| kratos.kratos.config.selfservice.flows.verification.after.default_browser_return_url | string | `"https://shortlink.ddns.net/next"` |  |
| kratos.kratos.config.selfservice.flows.verification.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.verification.ui_url | string | `"https://shortlink.ddns.net/next/auth/verify"` |  |
| kratos.kratos.config.selfservice.methods.link.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].client_id | string | `"...."` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].client_secret | string | `"...."` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].id | string | `"github"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].mapper_url | string | `"file:///etc/config/kratos/oidc.github.jsonnet"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].provider | string | `"github"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[0].scope[0] | string | `"user:email"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].client_id | string | `"...."` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].client_secret | string | `"...."` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].id | string | `"gitlab"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].mapper_url | string | `"file:///etc/config/kratos/oidc.gitlab.jsonnet"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].provider | string | `"gitlab"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].scope[0] | string | `"read_user"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].scope[1] | string | `"openid"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].scope[2] | string | `"profile"` |  |
| kratos.kratos.config.selfservice.methods.oidc.config.providers[1].scope[3] | string | `"email"` |  |
| kratos.kratos.config.selfservice.methods.oidc.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.password.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.profile.enabled | bool | `true` |  |
| kratos.kratos.config.serve.admin.base_url | string | `"http://127.0.0.1:4434/"` |  |
| kratos.kratos.config.serve.public.base_url | string | `"https://shortlink.ddns.net/api/auth"` |  |
| kratos.kratos.config.serve.public.cors.allow_credentials | bool | `true` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[0] | string | `"Authorization"` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[1] | string | `"Cookie"` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[2] | string | `"Content-Type"` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[3] | string | `"Set-Cookie"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[0] | string | `"POST"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[1] | string | `"GET"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[2] | string | `"PUT"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[3] | string | `"PATCH"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[4] | string | `"DELETE"` |  |
| kratos.kratos.config.serve.public.cors.allowed_origins[0] | string | `"http://127.0.0.1:3000"` |  |
| kratos.kratos.config.serve.public.cors.allowed_origins[1] | string | `"https://shortlink.ddns.net"` |  |
| kratos.kratos.config.serve.public.cors.debug | bool | `true` |  |
| kratos.kratos.config.serve.public.cors.enabled | bool | `true` |  |
| kratos.kratos.config.session.cookie.domain | string | `"https://shortlink.ddns.net"` |  |
| kratos.kratos.config.session.cookie.same_site | string | `"Lax"` |  |
| kratos.kratos.config.session.lifespan | string | `"720h"` |  |
| kratos.kratos.development | bool | `true` |  |
| kratos.kratos.identitySchemas."identity.default.schema.json" | string | `"{\n  \"$id\": \"https://schemas.ory.sh/presets/kratos/quickstart/email-password/identity.schema.json\",\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"title\": \"Person\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"traits\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"email\",\n          \"title\": \"E-Mail\",\n          \"minLength\": 3,\n          \"ory.sh/kratos\": {\n            \"credentials\": {\n              \"password\": {\n                \"identifier\": true\n              }\n            },\n            \"verification\": {\n              \"via\": \"email\"\n            },\n            \"recovery\": {\n              \"via\": \"email\"\n            }\n          }\n        },\n        \"name\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"first\": {\n              \"title\": \"First Name\",\n              \"type\": \"string\"\n            },\n            \"last\": {\n              \"title\": \"Last Name\",\n              \"type\": \"string\"\n            }\n          }\n        }\n      },\n      \"required\": [\n        \"email\"\n      ],\n      \"additionalProperties\": false\n    }\n  }\n}\n"` |  |
| monitoring.enabled | bool | `true` |  |
| nameOverride | string | `""` |  |
| secret.enabled | bool | `false` |  |
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nYour CA...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nYour cert...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nYour key...\n-----END EC PRIVATE KEY-----\n"` |  |
| service.port | int | `7070` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `"shortlink"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.8.1](https://github.com/norwoodj/helm-docs/releases/v1.8.1)
