# shortlink-api

![Version: 0.10.0](https://img.shields.io/badge/Version-0.10.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink API service

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
| file://../shortlink-common | shortlink-common | 0.2.4 |
| https://k8s.ory.sh/helm/charts | kratos | 0.26.5 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| NetworkPolicy.enabled | bool | `false` |  |
| commonAnnotations | object | `{}` | Add annotations to all the deployed resources |
| commonLabels | object | `{}` | Add labels to all the deployed resources |
| deploy.env.GRPC_CLIENT_HOST | string | `"istio-ingress.istio-ingress.svc.cluster.local"` |  |
| deploy.env.MQ_ENABLED | string | `"false"` |  |
| deploy.env.MQ_RABBIT_URI | string | `"amqp://admin:admin@shortlink.rabbitmq:5672"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
| deploy.env.STORE_POSTGRES_URI | string | `"postgres://postgres:shortlink@postgresql.postgresql:5432/shortlink?sslmode=disable"` | Default store config |
| deploy.env.TRACER_URI | string | `"grafana-tempo.grafana:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` | Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images |
| deploy.image.repository | string | `"registry.gitlab.com/shortlink-org/shortlink/api"` |  |
| deploy.image.tag | string | `"0.13.5"` |  |
| deploy.ports[0].containerPort | int | `7070` |  |
| deploy.ports[0].name | string | `"http"` |  |
| deploy.ports[0].protocol | string | `"TCP"` |  |
| deploy.ports[1].containerPort | int | `9090` |  |
| deploy.ports[1].name | string | `"metrics"` |  |
| deploy.ports[1].protocol | string | `"TCP"` |  |
| deploy.replicaCount | int | `1` |  |
| external_database | object | `{"enable":false,"ip":"192.168.0.101","port":6379}` | If you want to use an external database |
| fullnameOverride | string | `""` |  |
| host | string | `"shortlink.best"` |  |
| ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| ingress.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| ingress.enabled | bool | `true` |  |
| ingress.hostname | string | `"shortlink.best"` |  |
| ingress.path | string | `"/api/?(.*)"` |  |
| ingress.service.name | string | `"shortlink-api"` |  |
| ingress.service.port | int | `7070` |  |
| ingress.type | string | `"nginx"` |  |
| kratos.enabled | bool | `true` |  |
| kratos.fullnameOverride | string | `"kratos"` |  |
| kratos.ingress.admin.className | string | `"nginx"` |  |
| kratos.ingress.admin.enabled | bool | `false` |  |
| kratos.ingress.public.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| kratos.ingress.public.annotations."kubernetes.io/tls-acme" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| kratos.ingress.public.className | string | `"nginx"` |  |
| kratos.ingress.public.enabled | bool | `true` |  |
| kratos.ingress.public.hosts[0].host | string | `"shortlink.best"` |  |
| kratos.ingress.public.hosts[0].paths[0].path | string | `"/api/auth\\/(.*)"` |  |
| kratos.ingress.public.hosts[0].paths[0].pathType | string | `"Prefix"` |  |
| kratos.ingress.public.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| kratos.ingress.public.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| kratos.kratos.automigration | object | `{"enabled":true,"type":"job"}` | Enables database migration |
| kratos.kratos.automigration.type | string | `"job"` | Configure the way to execute database migration. Possible values: job, initContainer When set to job, the migration will be executed as a job on release or upgrade. When set to initContainer, the migration will be executed when kratos pod is created Defaults to job |
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
| kratos.kratos.config.log.level | string | `"info"` |  |
| kratos.kratos.config.secrets.cookie[0] | string | `"PLEASE-CHANGE-ME-I-AM-VERY-INSECURE"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[0] | string | `"*"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[1] | string | `"http://*"` |  |
| kratos.kratos.config.selfservice.allowed_return_urls[2] | string | `"https://*"` |  |
| kratos.kratos.config.selfservice.default_browser_return_url | string | `"https://shortlink.best"` |  |
| kratos.kratos.config.selfservice.flows.error.ui_url | string | `"https://shortlink.best/next/error"` |  |
| kratos.kratos.config.selfservice.flows.login.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.login.ui_url | string | `"https://shortlink.best/next"` |  |
| kratos.kratos.config.selfservice.flows.logout.after.default_browser_return_url | string | `"https://shortlink.best/next/auth/login"` |  |
| kratos.kratos.config.selfservice.flows.recovery.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.recovery.ui_url | string | `"https://shortlink.best/next/auth/recovery"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.oidc.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.password.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.registration.ui_url | string | `"https://shortlink.best/next/auth/registration"` |  |
| kratos.kratos.config.selfservice.flows.settings.privileged_session_max_age | string | `"15m"` |  |
| kratos.kratos.config.selfservice.flows.settings.ui_url | string | `"https://shortlink.best/next/user/profile"` |  |
| kratos.kratos.config.selfservice.flows.verification.after.default_browser_return_url | string | `"https://shortlink.best/next"` |  |
| kratos.kratos.config.selfservice.flows.verification.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.verification.ui_url | string | `"https://shortlink.best/next/auth/verification"` |  |
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
| kratos.kratos.config.serve.public.base_url | string | `"https://shortlink.best/api/auth"` |  |
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
| kratos.kratos.config.serve.public.cors.allowed_origins[1] | string | `"https://shortlink.best"` |  |
| kratos.kratos.config.serve.public.cors.debug | bool | `true` |  |
| kratos.kratos.config.serve.public.cors.enabled | bool | `true` |  |
| kratos.kratos.config.session.cookie.domain | string | `"https://shortlink.best"` |  |
| kratos.kratos.config.session.cookie.same_site | string | `"Lax"` |  |
| kratos.kratos.config.session.lifespan | string | `"720h"` |  |
| kratos.kratos.development | bool | `true` |  |
| kratos.kratos.identitySchemas."identity.default.schema.json" | string | `"{\n  \"$id\": \"https://schemas.ory.sh/presets/kratos/quickstart/email-password/identity.schema.json\",\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"title\": \"Person\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"traits\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"email\",\n          \"title\": \"E-Mail\",\n          \"minLength\": 3,\n          \"ory.sh/kratos\": {\n            \"credentials\": {\n              \"password\": {\n                \"identifier\": true\n              }\n            },\n            \"verification\": {\n              \"via\": \"email\"\n            },\n            \"recovery\": {\n              \"via\": \"email\"\n            }\n          }\n        },\n        \"name\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"first\": {\n              \"title\": \"First Name\",\n              \"type\": \"string\"\n            },\n            \"last\": {\n              \"title\": \"Last Name\",\n              \"type\": \"string\"\n            }\n          }\n        }\n      },\n      \"required\": [\n        \"email\"\n      ],\n      \"additionalProperties\": false\n    }\n  }\n}\n"` |  |
| kratos.kratos.serviceMonitor | object | `{"enabled":true}` | Parameters for the Prometheus ServiceMonitor objects. Reference: https://docs.openshift.com/container-platform/4.6/rest_api/monitoring_apis/servicemonitor-monitoring-coreos-com-v1.html |
| kratos.kratos.serviceMonitor.enabled | bool | `true` | switch to false to prevent creating the ServiceMonitor |
| monitoring.enabled | bool | `true` | Creates a Prometheus Operator ServiceMonitor |
| monitoring.jobLabel | string | `""` | The label to use to retrieve the job name from. |
| monitoring.labels | object | `{"release":"prometheus-operator"}` | Additional labels that can be used so PodMonitor will be discovered by Prometheus |
| nameOverride | string | `""` |  |
| secret.enabled | bool | `false` |  |
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nYour CA...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nYour cert...\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nYour key...\n-----END EC PRIVATE KEY-----\n"` |  |
| service.port | int | `7070` |  |
| service.type | string | `"ClusterIP"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
