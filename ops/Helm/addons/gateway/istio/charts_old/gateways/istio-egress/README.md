# istio-egress

![Version: 1.11.1](https://img.shields.io/badge/Version-1.11.1-informational?style=flat-square)

Helm chart for deploying Istio gateways

## Source Code

* <http://github.com/istio/istio>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| gateways.istio-egressgateway.additionalContainers | list | `[]` |  |
| gateways.istio-egressgateway.autoscaleEnabled | bool | `true` |  |
| gateways.istio-egressgateway.autoscaleMax | int | `5` |  |
| gateways.istio-egressgateway.autoscaleMin | int | `1` |  |
| gateways.istio-egressgateway.configVolumes | list | `[]` |  |
| gateways.istio-egressgateway.cpu.targetAverageUtilization | int | `80` |  |
| gateways.istio-egressgateway.env.ISTIO_META_ROUTER_MODE | string | `"standard"` |  |
| gateways.istio-egressgateway.injectionTemplate | string | `""` |  |
| gateways.istio-egressgateway.labels.app | string | `"istio-egressgateway"` |  |
| gateways.istio-egressgateway.labels.istio | string | `"egressgateway"` |  |
| gateways.istio-egressgateway.name | string | `"istio-egressgateway"` |  |
| gateways.istio-egressgateway.nodeSelector | object | `{}` |  |
| gateways.istio-egressgateway.podAnnotations | object | `{}` |  |
| gateways.istio-egressgateway.podAntiAffinityLabelSelector | list | `[]` |  |
| gateways.istio-egressgateway.podAntiAffinityTermLabelSelector | list | `[]` |  |
| gateways.istio-egressgateway.ports[0].name | string | `"http2"` |  |
| gateways.istio-egressgateway.ports[0].port | int | `80` |  |
| gateways.istio-egressgateway.ports[0].protocol | string | `"TCP"` |  |
| gateways.istio-egressgateway.ports[0].targetPort | int | `8080` |  |
| gateways.istio-egressgateway.ports[1].name | string | `"https"` |  |
| gateways.istio-egressgateway.ports[1].port | int | `443` |  |
| gateways.istio-egressgateway.ports[1].protocol | string | `"TCP"` |  |
| gateways.istio-egressgateway.ports[1].targetPort | int | `8443` |  |
| gateways.istio-egressgateway.resources.limits.cpu | string | `"2000m"` |  |
| gateways.istio-egressgateway.resources.limits.memory | string | `"1024Mi"` |  |
| gateways.istio-egressgateway.resources.requests.cpu | string | `"100m"` |  |
| gateways.istio-egressgateway.resources.requests.memory | string | `"128Mi"` |  |
| gateways.istio-egressgateway.rollingMaxSurge | string | `"100%"` |  |
| gateways.istio-egressgateway.rollingMaxUnavailable | string | `"25%"` |  |
| gateways.istio-egressgateway.runAsRoot | bool | `false` |  |
| gateways.istio-egressgateway.secretVolumes[0].mountPath | string | `"/etc/istio/egressgateway-certs"` |  |
| gateways.istio-egressgateway.secretVolumes[0].name | string | `"egressgateway-certs"` |  |
| gateways.istio-egressgateway.secretVolumes[0].secretName | string | `"istio-egressgateway-certs"` |  |
| gateways.istio-egressgateway.secretVolumes[1].mountPath | string | `"/etc/istio/egressgateway-ca-certs"` |  |
| gateways.istio-egressgateway.secretVolumes[1].name | string | `"egressgateway-ca-certs"` |  |
| gateways.istio-egressgateway.secretVolumes[1].secretName | string | `"istio-egressgateway-ca-certs"` |  |
| gateways.istio-egressgateway.serviceAnnotations | object | `{}` |  |
| gateways.istio-egressgateway.tolerations | list | `[]` |  |
| gateways.istio-egressgateway.type | string | `"ClusterIP"` |  |
| gateways.istio-egressgateway.zvpn.enabled | bool | `false` |  |
| gateways.istio-egressgateway.zvpn.suffix | string | `"global"` |  |
| global.arch.amd64 | int | `2` |  |
| global.arch.ppc64le | int | `2` |  |
| global.arch.s390x | int | `2` |  |
| global.caAddress | string | `""` |  |
| global.defaultConfigVisibilitySettings | list | `[]` |  |
| global.defaultNodeSelector | object | `{}` |  |
| global.defaultPodDisruptionBudget.enabled | bool | `true` |  |
| global.defaultResources.requests.cpu | string | `"10m"` |  |
| global.defaultTolerations | list | `[]` |  |
| global.hub | string | `"docker.io/istio"` |  |
| global.imagePullPolicy | string | `""` |  |
| global.imagePullSecrets | list | `[]` |  |
| global.istioNamespace | string | `"istio-system"` |  |
| global.jwtPolicy | string | `"third-party-jwt"` |  |
| global.logAsJson | bool | `false` |  |
| global.logging.level | string | `"default:info"` |  |
| global.meshID | string | `""` |  |
| global.mountMtlsCerts | bool | `false` |  |
| global.multiCluster.clusterName | string | `""` |  |
| global.multiCluster.enabled | bool | `false` |  |
| global.network | string | `""` |  |
| global.pilotCertProvider | string | `"istiod"` |  |
| global.priorityClassName | string | `""` |  |
| global.proxy.clusterDomain | string | `"cluster.local"` |  |
| global.proxy.componentLogLevel | string | `"misc:error"` |  |
| global.proxy.enableCoreDump | bool | `false` |  |
| global.proxy.image | string | `"proxyv2"` |  |
| global.proxy.logLevel | string | `"warning"` |  |
| global.sds.token.aud | string | `"istio-ca"` |  |
| global.sts.servicePort | int | `0` |  |
| global.tag | string | `"1.11.1"` |  |
| meshConfig.defaultConfig.proxyMetadata | object | `{}` |  |
| meshConfig.defaultConfig.tracing | string | `nil` |  |
| meshConfig.enablePrometheusMerge | bool | `true` |  |
| meshConfig.trustDomain | string | `"cluster.local"` |  |
| ownerName | string | `""` |  |
| revision | string | `""` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.7.0](https://github.com/norwoodj/helm-docs/releases/v1.7.0)
