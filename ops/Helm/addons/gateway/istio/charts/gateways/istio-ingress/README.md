# istio-ingress

![Version: 1.11.1](https://img.shields.io/badge/Version-1.11.1-informational?style=flat-square)

Helm chart for deploying Istio gateways

## Source Code

* <http://github.com/istio/istio>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| gateways.istio-ingressgateway.additionalContainers | list | `[]` |  |
| gateways.istio-ingressgateway.autoscaleEnabled | bool | `true` |  |
| gateways.istio-ingressgateway.autoscaleMax | int | `5` |  |
| gateways.istio-ingressgateway.autoscaleMin | int | `1` |  |
| gateways.istio-ingressgateway.configVolumes | list | `[]` |  |
| gateways.istio-ingressgateway.cpu.targetAverageUtilization | int | `80` |  |
| gateways.istio-ingressgateway.customService | bool | `false` |  |
| gateways.istio-ingressgateway.env.ISTIO_META_ROUTER_MODE | string | `"standard"` |  |
| gateways.istio-ingressgateway.externalTrafficPolicy | string | `""` |  |
| gateways.istio-ingressgateway.ingressPorts | list | `[]` |  |
| gateways.istio-ingressgateway.injectionTemplate | string | `""` |  |
| gateways.istio-ingressgateway.labels.app | string | `"istio-ingressgateway"` |  |
| gateways.istio-ingressgateway.labels.istio | string | `"ingressgateway"` |  |
| gateways.istio-ingressgateway.loadBalancerIP | string | `""` |  |
| gateways.istio-ingressgateway.loadBalancerSourceRanges | list | `[]` |  |
| gateways.istio-ingressgateway.name | string | `"istio-ingressgateway"` |  |
| gateways.istio-ingressgateway.nodeSelector | object | `{}` |  |
| gateways.istio-ingressgateway.podAnnotations | object | `{}` |  |
| gateways.istio-ingressgateway.podAntiAffinityLabelSelector | list | `[]` |  |
| gateways.istio-ingressgateway.podAntiAffinityTermLabelSelector | list | `[]` |  |
| gateways.istio-ingressgateway.ports[0].name | string | `"status-port"` |  |
| gateways.istio-ingressgateway.ports[0].port | int | `15021` |  |
| gateways.istio-ingressgateway.ports[0].protocol | string | `"TCP"` |  |
| gateways.istio-ingressgateway.ports[0].targetPort | int | `15021` |  |
| gateways.istio-ingressgateway.ports[1].name | string | `"http2"` |  |
| gateways.istio-ingressgateway.ports[1].port | int | `80` |  |
| gateways.istio-ingressgateway.ports[1].protocol | string | `"TCP"` |  |
| gateways.istio-ingressgateway.ports[1].targetPort | int | `8080` |  |
| gateways.istio-ingressgateway.ports[2].name | string | `"https"` |  |
| gateways.istio-ingressgateway.ports[2].port | int | `443` |  |
| gateways.istio-ingressgateway.ports[2].protocol | string | `"TCP"` |  |
| gateways.istio-ingressgateway.ports[2].targetPort | int | `8443` |  |
| gateways.istio-ingressgateway.ports[3].name | string | `"grpc"` |  |
| gateways.istio-ingressgateway.ports[3].port | int | `50051` |  |
| gateways.istio-ingressgateway.ports[3].protocol | string | `"TCP"` |  |
| gateways.istio-ingressgateway.ports[3].targetPort | int | `50051` |  |
| gateways.istio-ingressgateway.resources.limits.cpu | string | `"2000m"` |  |
| gateways.istio-ingressgateway.resources.limits.memory | string | `"1024Mi"` |  |
| gateways.istio-ingressgateway.resources.requests.cpu | string | `"50m"` |  |
| gateways.istio-ingressgateway.resources.requests.memory | string | `"64Mi"` |  |
| gateways.istio-ingressgateway.rollingMaxSurge | string | `"100%"` |  |
| gateways.istio-ingressgateway.rollingMaxUnavailable | string | `"25%"` |  |
| gateways.istio-ingressgateway.runAsRoot | bool | `false` |  |
| gateways.istio-ingressgateway.secretVolumes[0].mountPath | string | `"/etc/istio/ingressgateway-certs"` |  |
| gateways.istio-ingressgateway.secretVolumes[0].name | string | `"ingressgateway-certs"` |  |
| gateways.istio-ingressgateway.secretVolumes[0].secretName | string | `"istio-ingressgateway-certs"` |  |
| gateways.istio-ingressgateway.secretVolumes[1].mountPath | string | `"/etc/istio/ingressgateway-ca-certs"` |  |
| gateways.istio-ingressgateway.secretVolumes[1].name | string | `"ingressgateway-ca-certs"` |  |
| gateways.istio-ingressgateway.secretVolumes[1].secretName | string | `"istio-ingressgateway-ca-certs"` |  |
| gateways.istio-ingressgateway.serviceAnnotations | object | `{}` |  |
| gateways.istio-ingressgateway.tolerations | list | `[]` |  |
| gateways.istio-ingressgateway.type | string | `"ClusterIP"` |  |
| gateways.istio-ingressgateway.zvpn.enabled | bool | `false` |  |
| gateways.istio-ingressgateway.zvpn.suffix | string | `"global"` |  |
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
| global.multiCluster.globalDomainSuffix | string | `"global"` |  |
| global.multiCluster.includeEnvoyFilter | bool | `true` |  |
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
| meshConfig.defaultConfig.sampling | float | `100` |  |
| meshConfig.defaultConfig.tracing.zipkin.address | string | `"zipkin.jaeger-operator:9411"` |  |
| meshConfig.enablePrometheusMerge | bool | `true` |  |
| meshConfig.trustDomain | string | `"cluster.local"` |  |
| ownerName | string | `""` |  |
| revision | string | `""` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
