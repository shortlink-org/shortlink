# shortlink-auth

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink auth service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | batazor111@gmail.com | batazor.ru |

## Source Code

* <https://github.com/batazor/shortlink>

## Requirements

Kubernetes: `>= 1.19.0 || >= v1.19.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://k8s.ory.sh/helm/charts | kratos | 0.21.8 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| kratos.enabled | bool | `true` |  |
| kratos.image.tag | string | `"v0.6.3-sqlite"` |  |
| kratos.kratos.autoMigrate | bool | `false` |  |
| kratos.kratos.config.courier.smtp.connection_uri | string | `"smtps://test:test@mailslurper:1025/?skip_ssl_verify=true"` |  |
| kratos.kratos.config.courier.smtp.from_address | string | `"no-reply@shortlink.com"` |  |
| kratos.kratos.config.dsn | string | `"memory"` |  |
| kratos.kratos.config.hashers.argon2.iterations | int | `2` |  |
| kratos.kratos.config.hashers.argon2.key_length | int | `16` |  |
| kratos.kratos.config.hashers.argon2.memory | string | `"128MB"` |  |
| kratos.kratos.config.hashers.argon2.parallelism | int | `1` |  |
| kratos.kratos.config.hashers.argon2.salt_length | int | `16` |  |
| kratos.kratos.config.identity.default_schema_url | string | `"file://identity.schema.json"` |  |
| kratos.kratos.config.log.format | string | `"json"` |  |
| kratos.kratos.config.log.leak_sensitive_values | bool | `true` |  |
| kratos.kratos.config.log.level | string | `"debug"` |  |
| kratos.kratos.config.secrets.cookie[0] | string | `"PLEASE-CHANGE-ME-I-AM-VERY-INSECURE"` |  |
| kratos.kratos.config.selfservice.default_browser_return_url | string | `"http://127.0.0.1:3000/next/"` |  |
| kratos.kratos.config.selfservice.flows.error.ui_url | string | `"http://127.0.0.1:3000/next/error"` |  |
| kratos.kratos.config.selfservice.flows.login.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.login.ui_url | string | `"http://127.0.0.1:3000/next/auth/login"` |  |
| kratos.kratos.config.selfservice.flows.logout.after.default_browser_return_url | string | `"http://127.0.0.1:3000/next/auth/login"` |  |
| kratos.kratos.config.selfservice.flows.recovery.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.recovery.ui_url | string | `"http://127.0.0.1:3000/next/auth/recovery"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.oidc.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.after.password.hooks[0].hook | string | `"session"` |  |
| kratos.kratos.config.selfservice.flows.registration.lifespan | string | `"10m"` |  |
| kratos.kratos.config.selfservice.flows.registration.ui_url | string | `"http://127.0.0.1:3000/next/auth/registration"` |  |
| kratos.kratos.config.selfservice.flows.settings.privileged_session_max_age | string | `"15m"` |  |
| kratos.kratos.config.selfservice.flows.settings.ui_url | string | `"http://127.0.0.1:3000/next/auth/profile"` |  |
| kratos.kratos.config.selfservice.flows.verification.after.default_browser_return_url | string | `"http://127.0.0.1:3000/next/"` |  |
| kratos.kratos.config.selfservice.flows.verification.enabled | bool | `true` |  |
| kratos.kratos.config.selfservice.flows.verification.ui_url | string | `"http://127.0.0.1:3000/next/auth/verify"` |  |
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
| kratos.kratos.config.selfservice.whitelisted_return_urls[0] | string | `"http://127.0.0.1:3000"` |  |
| kratos.kratos.config.serve.admin.base_url | string | `"http://127.0.0.1:4434/"` |  |
| kratos.kratos.config.serve.public.base_url | string | `"http://127.0.0.1:4433/"` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[0] | string | `"Authorization"` |  |
| kratos.kratos.config.serve.public.cors.allowed_headers[1] | string | `"Cookie"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[0] | string | `"POST"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[1] | string | `"GET"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[2] | string | `"PUT"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[3] | string | `"PATCH"` |  |
| kratos.kratos.config.serve.public.cors.allowed_methods[4] | string | `"DELETE"` |  |
| kratos.kratos.config.serve.public.cors.allowed_origins[0] | string | `"http://127.0.0.1:3000"` |  |
| kratos.kratos.config.serve.public.cors.debug | bool | `true` |  |
| kratos.kratos.config.serve.public.cors.enabled | bool | `true` |  |
| kratos.kratos.config.serve.public.cors.exposed_headers[0] | string | `"Content-Type"` |  |
| kratos.kratos.config.serve.public.cors.exposed_headers[1] | string | `"Set-Cookie"` |  |
| kratos.kratos.config.session.cookie.domain | string | `"http://127.0.0.1:3000"` |  |
| kratos.kratos.config.session.cookie.same_site | string | `"Lax"` |  |
| kratos.kratos.config.session.lifespan | string | `"720h"` |  |
| kratos.kratos.config.version | string | `"v0.6.3-alpha.1"` |  |
| kratos.kratos.development | bool | `true` |  |
| kratos.kratos.identitySchemas."identity.schema.json" | string | `"{\n  \"$id\": \"https://schemas.ory.sh/presets/kratos/quickstart/email-password/identity.schema.json\",\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"title\": \"Person\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"traits\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"email\",\n          \"title\": \"E-Mail\",\n          \"minLength\": 3,\n          \"ory.sh/kratos\": {\n            \"credentials\": {\n              \"password\": {\n                \"identifier\": true\n              }\n            },\n            \"verification\": {\n              \"via\": \"email\"\n            },\n            \"recovery\": {\n              \"via\": \"email\"\n            }\n          }\n        },\n        \"name\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"first\": {\n              \"title\": \"First Name\",\n              \"type\": \"string\"\n            },\n            \"last\": {\n              \"title\": \"Last Name\",\n              \"type\": \"string\"\n            }\n          }\n        }\n      },\n      \"required\": [\n        \"email\"\n      ],\n      \"additionalProperties\": false\n    }\n  }\n}\n"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)
