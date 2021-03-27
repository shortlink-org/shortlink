# shortlink-metadata

![Version: 0.5.7](https://img.shields.io/badge/Version-0.5.7-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink service for get metadata by URL

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
| deploy.env.SERVICE_NAME | string | `"metadata"` |  |
| deploy.env.TRACER_URI | string | `"jaeger-agent.jaeger-operator:6831"` |  |
| deploy.image.pullPolicy | string | `"IfNotPresent"` |  |
| deploy.image.repository | string | `"batazor/shortlink-metadata"` |  |
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
| secret.grpcIntermediateCA | string | `"-----BEGIN CERTIFICATE-----\nMIICljCCAjugAwIBAgIUZhuY8pa+aFn96PpXHKoFxgW9WsQwCgYIKoZIzj0EAwIw\ngYwxCzAJBgNVBAYTAlJVMQ8wDQYDVQQIEwZNb3Njb3cxDzANBgNVBAcTBk1vc2Nv\ndzESMBAGA1UEChMJU2hvcnRsaW5rMSswKQYDVQQLEyJFeGFtcGxlIFJvb3QgQ2Vy\ndGlmaWNhdGUgQXV0aG9yaXR5MRowGAYDVQQDExFTaG9ydGxpbmsgUm9vdCBDQTAe\nFw0yMDExMjExODAyMDBaFw0yMTExMjExODAyMDBaMH8xCzAJBgNVBAYTAlJVMQ8w\nDQYDVQQIEwZNb3Njb3cxDzANBgNVBAcTBk1vc2NvdzESMBAGA1UEChMJU2hvcnRs\naW5rMQ8wDQYDVQQLEwZNb3Njb3cxKTAnBgNVBAMTIEN1c3RvbSBTaG9ydGxpbmsg\nSW50ZXJtZWRpYXRlIENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsMwqTmk0\nbvNChfE79Ljr/mnkw90XVe4J45GgYYZZ83eUqetg/dnT+0h/Mdw1uEABYtbmRG4Q\nyGdNIcSCsS8tf6OBhjCBgzAOBgNVHQ8BAf8EBAMCAaYwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCMBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFBiZ\nymfiD4U/jz6qSNvU26XMCC9oMB8GA1UdIwQYMBaAFA93UcPMOw3jdtWLxuCopshq\nK9FrMAoGCCqGSM49BAMCA0kAMEYCIQDGHwhl3IrIgD75cvqBqvitltzEDqBlnGMi\nM3FEoCXGhwIhAIFPuVTuk16zNNJZNlY+027k0pg0SOfNcw0qcNyFtOvC\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerCert | string | `"-----BEGIN CERTIFICATE-----\nMIICkTCCAjigAwIBAgIUdo/zgCCySxFfOxrYpympLLN0mvcwCgYIKoZIzj0EAwIw\nfzELMAkGA1UEBhMCUlUxDzANBgNVBAgTBk1vc2NvdzEPMA0GA1UEBxMGTW9zY293\nMRIwEAYDVQQKEwlTaG9ydGxpbmsxDzANBgNVBAsTBk1vc2NvdzEpMCcGA1UEAxMg\nQ3VzdG9tIFNob3J0bGluayBJbnRlcm1lZGlhdGUgQ0EwHhcNMjAxMTIxMTgwMjAw\nWhcNMjExMTIxMTgwMjAwWjBoMQswCQYDVQQGEwJSVTEPMA0GA1UECBMGTW9zY293\nMQ8wDQYDVQQHEwZNb3Njb3cxEjAQBgNVBAoTCVNob3J0bGluazEPMA0GA1UECxMG\nTW9zY293MRIwEAYDVQQDEwlzaG9ydGxpbmswWTATBgcqhkjOPQIBBggqhkjOPQMB\nBwNCAARXdzYwc4cLaba2/9zxd0aT0wGSOy40s47jT7fkGwCuOvNB7Yl80ed/jEht\n+BJJgT87MOVOHLBXT9SEa2O/8Iw6o4GoMIGlMA4GA1UdDwEB/wQEAwIFoDATBgNV\nHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQvXJcdbHow\nDJoiXyvryuTo1NFAtjAfBgNVHSMEGDAWgBQYmcpn4g+FP48+qkjb1NulzAgvaDAw\nBgNVHREEKTAngglsb2NhbGhvc3SCBWxvY2FsggcqLmxvY2FshwR/AAABhwQAAAAA\nMAoGCCqGSM49BAMCA0cAMEQCIDTXPLlM1YpK5Iwe80imyysmJAkbA+jKSWW0CBvd\nrUQDAiB71ah7iKjM1P9lOzWfD3nm6DYZSdWLmUCXnrjBt6xYEQ==\n-----END CERTIFICATE-----\n"` |  |
| secret.grpcServerKey | string | `"-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIAm9fkeHAhonIbVt3LQHgibo7x5+5RkMEW6a1qm00KUkoAoGCCqGSM49\nAwEHoUQDQgAEV3c2MHOHC2m2tv/c8XdGk9MBkjsuNLOO40+35BsArjrzQe2JfNHn\nf4xIbfgSSYE/OzDlThywV0/UhGtjv/CMOg==\n-----END EC PRIVATE KEY-----\n"` |  |
| service.port | int | `7070` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `"shortlink"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
