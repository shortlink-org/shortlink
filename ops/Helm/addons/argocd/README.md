# argo

![Version: 0.3.0](https://img.shields.io/badge/Version-0.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.22.0 || >= v1.22.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-cd | 5.3.6 |
| https://argoproj.github.io/argo-helm | argo-events | 2.0.4 |
| https://argoproj.github.io/argo-helm | argo-rollouts | 2.20.0 |
| https://argoproj.github.io/argo-helm | argo-workflows | 0.17.1 |
| https://argoproj.github.io/argo-helm | argocd-image-updater | 0.8.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| argo-cd.applicationSet.metrics.enabled | bool | `true` |  |
| argo-cd.applicationSet.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.configs.repositories.shortlink.name | string | `"shortlink"` |  |
| argo-cd.configs.repositories.shortlink.type | string | `"git"` |  |
| argo-cd.configs.repositories.shortlink.url | string | `"https://github.com/batazor/shortlink"` |  |
| argo-cd.controller.metrics.applicationLabels.enabled | bool | `true` |  |
| argo-cd.controller.metrics.enabled | bool | `true` |  |
| argo-cd.controller.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.controller.rules.enabled | bool | `true` |  |
| argo-cd.dex.enabled | bool | `true` |  |
| argo-cd.dex.metrics.enabled | bool | `true` |  |
| argo-cd.dex.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.externalRedis.host | string | `"shortlink-redis-master.redis"` |  |
| argo-cd.fullnameOverride | string | `"argocd"` |  |
| argo-cd.global.logging.format | string | `"json"` |  |
| argo-cd.notifications.argocdUrl | string | `"https://architecture.ddns.net/argocd"` |  |
| argo-cd.notifications.metrics.enabled | bool | `true` |  |
| argo-cd.notifications.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.notifications.notifiers."service.slack" | string | `"token: $slack-token\nusername: argocd # optional username\nicon: :dart: # optional icon for the message (supports both emoij and url notation)\n"` |  |
| argo-cd.notifications.secret.items.slack-token | string | `"<SECRET>"` |  |
| argo-cd.redis.enabled | bool | `false` |  |
| argo-cd.repoServer.metrics.enabled | bool | `true` |  |
| argo-cd.repoServer.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.repoServer.volumes[0].emptyDir | object | `{}` |  |
| argo-cd.repoServer.volumes[0].name | string | `"custom-tools"` |  |
| argo-cd.server.config.url | string | `"https://architecture.ddns.net/argocd"` |  |
| argo-cd.server.configAnnotations | object | `{}` |  |
| argo-cd.server.extensions.enabled | bool | `true` |  |
| argo-cd.server.extraArgs[0] | string | `"--rootpath"` |  |
| argo-cd.server.extraArgs[1] | string | `"/argocd"` |  |
| argo-cd.server.extraArgs[2] | string | `"--basehref"` |  |
| argo-cd.server.extraArgs[3] | string | `"/argocd"` |  |
| argo-cd.server.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTPS"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"proxy_ssl_server_name on;\nproxy_ssl_name $host;"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/secure-backends" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-redirect" | string | `"true"` |  |
| argo-cd.server.ingress.enabled | bool | `true` |  |
| argo-cd.server.ingress.hosts[0] | string | `"architecture.ddns.net"` |  |
| argo-cd.server.ingress.https | bool | `true` |  |
| argo-cd.server.ingress.ingressClassName | string | `"nginx"` |  |
| argo-cd.server.ingress.paths[0] | string | `"/argocd(/|$)(.*)"` |  |
| argo-cd.server.ingress.tls[0].hosts[0] | string | `"architecture.ddns.net"` |  |
| argo-cd.server.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| argo-cd.server.metrics.enabled | bool | `true` |  |
| argo-cd.server.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.server.rbacConfig."policy.csv" | string | `"p, role:org-admin, applications, *, */*, allow\np, role:org-admin, clusters, get, *, allow\np, role:org-admin, repositories, get, *, allow\np, role:org-admin, repositories, create, *, allow\np, role:org-admin, repositories, update, *, allow\np, role:org-admin, repositories, delete, *, allow\ng, devops, role:admin\ng, gitlab, role:org-admin\n"` |  |
| argo-cd.server.rbacConfig."policy.default" | string | `"role:readonly"` |  |
| argo-events.fullnameOverride | string | `"argo-events"` |  |
| argo-workflows.controller.metricsConfig.enabled | bool | `true` |  |
| argo-workflows.controller.serviceMonitor.enabled | bool | `true` |  |
| argo-workflows.controller.telemetryConfig.enabled | bool | `true` |  |
| argo-workflows.fullnameOverride | string | `"argo-workflows"` |  |
| argo-workflows.server.extraArgs[0] | string | `"--basehref"` |  |
| argo-workflows.server.extraArgs[1] | string | `"/argoworkflows/"` |  |
| argo-workflows.server.extraArgs[2] | string | `"--auth-mode=server"` |  |
| argo-workflows.server.extraEnv[0].name | string | `"BASE_HREF"` |  |
| argo-workflows.server.extraEnv[0].value | string | `"/argoworkflows"` |  |
| argo-workflows.server.extraEnv[1].name | string | `"ARGO_BASE_HREF"` |  |
| argo-workflows.server.extraEnv[1].value | string | `"/argoworkflows"` |  |
| argo-workflows.server.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTP"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentracing" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| argo-workflows.server.ingress.enabled | bool | `true` |  |
| argo-workflows.server.ingress.hosts[0] | string | `"architecture.ddns.net"` |  |
| argo-workflows.server.ingress.ingressClassName | string | `"nginx"` |  |
| argo-workflows.server.ingress.paths[0] | string | `"/argoworkflows/?(.*)"` |  |
| argo-workflows.server.ingress.tls[0].hosts[0] | string | `"architecture.ddns.net"` |  |
| argo-workflows.server.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| argo.enabled | bool | `true` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
