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
| deploy.env.MQ_RABBIT_URI | string | `"amqp://admin:admin@rabbitmq.rabbitmq:5672"` |  |
| deploy.env.MQ_TYPE | string | `"rabbitmq"` |  |
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
| fullnameOverride | string | `""` |  |
| nameOverride | string | `""` |  |
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nMIIClTCCAjugAwIBAgIUOCl/NAZsR4Qy+78LBDy9l/1MN+IwCgYIKoZIzj0EAwIw\ngYwxCzAJBgNVBAYTAlJVMQ8wDQYDVQQIEwZNb3Njb3cxDzANBgNVBAcTBk1vc2Nv\ndzESMBAGA1UEChMJU2hvcnRsaW5rMSswKQYDVQQLEyJFeGFtcGxlIFJvb3QgQ2Vy\ndGlmaWNhdGUgQXV0aG9yaXR5MRowGAYDVQQDExFTaG9ydGxpbmsgUm9vdCBDQTAe\nFw0yMTA0MTgxMzUwMDBaFw0yMjA0MTgxMzUwMDBaMH8xCzAJBgNVBAYTAlJVMQ8w\nDQYDVQQIEwZNb3Njb3cxDzANBgNVBAcTBk1vc2NvdzESMBAGA1UEChMJU2hvcnRs\naW5rMQ8wDQYDVQQLEwZNb3Njb3cxKTAnBgNVBAMTIEN1c3RvbSBTaG9ydGxpbmsg\nSW50ZXJtZWRpYXRlIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErGrPBmwM\n+FlQhRgAB2a+76LCLR9KF772wff9MxgXiSaD1cQZB8iiIPQzWObkKb0nUxBBS6u2\nulXdhj0Po32A46OBhjCBgzAOBgNVHQ8BAf8EBAMCAaYwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCMBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFHqZ\niQcOHCuTdL6qkOTmemcjUeG5MB8GA1UdIwQYMBaAFKYsspmh6iQcM0XotbiZoZEx\n3jyeMAoGCCqGSM49BAMCA0gAMEUCIQC536u40hdoh4EU2Wcfp0Wz/MDQI01Ol16I\nFTfi/rQgPgIgS1XwXKSpJjE6qStUO308w9BVEV/CT/KJ9kpqTVlhm4M=\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nMIICnzCCAkWgAwIBAgIUEuCWhGqFAMdEFamP0LohiAxdEQYwCgYIKoZIzj0EAwIw\nfzELMAkGA1UEBhMCUlUxDzANBgNVBAgTBk1vc2NvdzEPMA0GA1UEBxMGTW9zY293\nMRIwEAYDVQQKEwlTaG9ydGxpbmsxDzANBgNVBAsTBk1vc2NvdzEpMCcGA1UEAxMg\nQ3VzdG9tIFNob3J0bGluayBJbnRlcm1lZGlhdGUgQ0EwHhcNMjEwNDE4MTM1MDAw\nWhcNMjIwNDE4MTM1MDAwWjBoMQswCQYDVQQGEwJSVTEPMA0GA1UECBMGTW9zY293\nMQ8wDQYDVQQHEwZNb3Njb3cxEjAQBgNVBAoTCVNob3J0bGluazEPMA0GA1UECxMG\nTW9zY293MRIwEAYDVQQDEwlzaG9ydGxpbmswWTATBgcqhkjOPQIBBggqhkjOPQMB\nBwNCAARjxcJl5q0fVcUU51TeEBKDNfyn4I59JpbUtfyQShpVJxZQ0AP8XYEx9l40\nj9QWkB/QP7+4wj+O+so0TU8W1V4ho4G1MIGyMA4GA1UdDwEB/wQEAwIFoDATBgNV\nHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBR5Bu+KogPg\n5cjtR+ZQWYT3tdqRNzAfBgNVHSMEGDAWgBR6mYkHDhwrk3S+qpDk5npnI1HhuTA9\nBgNVHREENjA0gglsb2NhbGhvc3SCBWxvY2FsggcqLmxvY2FsggsqLnNob3J0bGlu\na4cEfwAAAYcEAAAAADAKBggqhkjOPQQDAgNIADBFAiEA+l1fywhag1A2ozb4xmV3\nvUwKpDrtJKL5hfat4XbVgHsCIB6jD+t+qOi3GjbB04kxWckLLMdMuQGUhA26/MW/\nhIkH\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIEOFKH5IWvobNXqUU5LF64NNh3o01fa1eSyNrN+8LwKjoAoGCCqGSM49\nAwEHoUQDQgAEY8XCZeatH1XFFOdU3hASgzX8p+COfSaW1LX8kEoaVScWUNAD/F2B\nMfZeNI/UFpAf0D+/uMI/jvrKNE1PFtVeIQ==\n-----END EC PRIVATE KEY-----\n"` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"shortlink"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
