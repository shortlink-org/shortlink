# kratos

![Version: 0.2.2](https://img.shields.io/badge/Version-0.2.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.24.0 || >= v1.24.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://k8s.ory.sh/helm/charts | kratos | 0.33.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| kratos.enabled | bool | `true` |  |
| kratos.fullnameOverride | string | `"kratos"` |  |
| kratos.ingress.admin.className | string | `"nginx"` |  |
| kratos.ingress.admin.enabled | bool | `false` |  |
| kratos.ingress.public.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"false"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| kratos.ingress.public.className | string | `"nginx"` |  |
| kratos.ingress.public.enabled | bool | `true` |  |
| kratos.ingress.public.hosts[0].host | string | `"shortlink.best"` |  |
| kratos.ingress.public.hosts[0].paths[0].path | string | `"/api/auth/?(.*)"` |  |
| kratos.ingress.public.hosts[0].paths[0].pathType | string | `"Prefix"` |  |
| kratos.ingress.public.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| kratos.ingress.public.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| kratos.kratos.automigration | object | `{"enabled":false,"type":"job"}` | Enables database migration |
| kratos.kratos.automigration.type | string | `"job"` | Configure the way to execute database migration. Possible values: job, initContainer When set to job, the migration will be executed as a job on release or upgrade. When set to initContainer, the migration will be executed when kratos pod is created Defaults to job |
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
| kratos.kratos.config.selfservice.methods.lookup_secret.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.oidc.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.password.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.profile.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.methods.totp.config.issuer | string | `"shortlink.best"` |  |
| kratos.kratos.config.selfservice.methods.totp.enabled | bool | `true` |  |
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
| kratos.kratos.identitySchemas."identity.default.schema.json" | string | `"{\n  \"$id\": \"https://schemas.ory.sh/presets/kratos/quickstart/email-password/identity.schema.json\",\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"title\": \"Person\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"traits\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"email\",\n          \"title\": \"E-Mail\",\n          \"minLength\": 3,\n          \"ory.sh/kratos\": {\n            \"credentials\": {\n              \"password\": {\n                \"identifier\": true\n              },\n              \"totp\": {\n                \"account_name\": true\n              }\n            },\n            \"verification\": {\n              \"via\": \"email\"\n            },\n            \"recovery\": {\n              \"via\": \"email\"\n            }\n          }\n        },\n        \"name\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"first\": {\n              \"title\": \"First Name\",\n              \"type\": \"string\"\n            },\n            \"last\": {\n              \"title\": \"Last Name\",\n              \"type\": \"string\"\n            }\n          }\n        }\n      },\n      \"required\": [\n        \"email\"\n      ],\n      \"additionalProperties\": false\n    }\n  }\n}\n"` |  |
| kratos.kratos.identitySchemas."oidc.github.jsonnet" | string | `"local claims = {\n  email_verified: false,\n} + std.extVar('claims');\n\n{\n  identity: {\n    traits: {\n      // Allowing unverified email addresses enables account\n      // enumeration attacks, especially if the value is used for\n      // e.g. verification or as a password login identifier.\n      //\n      // Therefore we only return the email if it (a) exists and (b) is marked verified\n      // by GitHub.\n      [if 'email' in claims && claims.email_verified then 'email' else null]: claims.email,\n    },\n    metadata_public: {\n      github_username: claims.username,\n    }\n  },\n}\n"` |  |
| kratos.kratos.identitySchemas."oidc.gitlab.jsonnet" | string | `"local claims = {\n  email_verified: false,\n} + std.extVar('claims');\n{\n  identity: {\n    traits: {\n      // Allowing unverified email addresses enables account\n      // enumeration attacks,  if the value is used for\n      // verification or as a password login identifier.\n      //\n      // Therefore we only return the email if it (a) exists and (b) is marked verified\n      // by GitLab.\n      [if 'email' in claims && claims.email_verified then 'email' else null]: claims.email,\n    },\n  },\n}\n"` |  |
| kratos.secret.hashSumEnabled | bool | `false` |  |
| kratos.serviceMonitor.enabled | bool | `true` |  |
| kratos.serviceMonitor.labels.release | string | `"prometheus-operator"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
