# argo

![Version: 0.3.14](https://img.shields.io/badge/Version-0.3.14-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.19.4](https://img.shields.io/badge/AppVersion-5.19.4-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-cd | 5.46.2 |
| https://argoproj.github.io/argo-helm | argo-events | 2.4.1 |
| https://argoproj.github.io/argo-helm | argo-rollouts | 2.32.0 |
| https://argoproj.github.io/argo-helm | argo-workflows | 0.33.3 |
| https://argoproj.github.io/argo-helm | argocd-apps | 1.4.1 |
| https://argoproj.github.io/argo-helm | argocd-image-updater | 0.9.1 |
| oci://registry-1.docker.io/bitnamicharts | redis | 18.0.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| argo-cd.applicationSet.metrics.enabled | bool | `true` |  |
| argo-cd.applicationSet.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.configs.repositories.shortlink.name | string | `"shortlink"` |  |
| argo-cd.configs.repositories.shortlink.type | string | `"git"` |  |
| argo-cd.configs.repositories.shortlink.url | string | `"https://github.com/shortlink-org/shortlink"` |  |
| argo-cd.controller.metrics.applicationLabels.enabled | bool | `true` |  |
| argo-cd.controller.metrics.enabled | bool | `true` |  |
| argo-cd.controller.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.controller.replicas | int | `2` |  |
| argo-cd.controller.rules.enabled | bool | `true` |  |
| argo-cd.controller.rules.spec[0].alert | string | `"ArgoAppMissing"` |  |
| argo-cd.controller.rules.spec[0].annotations.description | string | `"Argo CD has not reported any applications data for the past 15 minutes which means that it must be down or not functioning properly.  This needs to be resolved for this cloud to continue to maintain state.\n"` |  |
| argo-cd.controller.rules.spec[0].annotations.summary | string | `"[Argo CD] No reported applications"` |  |
| argo-cd.controller.rules.spec[0].expr | string | `"absent(argocd_app_info)\n"` |  |
| argo-cd.controller.rules.spec[0].for | string | `"15m"` |  |
| argo-cd.controller.rules.spec[0].labels.severity | string | `"critical"` |  |
| argo-cd.controller.rules.spec[1].alert | string | `"ArgoAppNotSynced"` |  |
| argo-cd.controller.rules.spec[1].annotations.description | string | `"The application [{{`{{$labels.name}}`}} has not been synchronized for over\n 12 hours which means that the state of this cloud has drifted away from the\n state inside Git.\n"` |  |
| argo-cd.controller.rules.spec[1].annotations.summary | string | `"[{{`{{$labels.name}}`}}] Application not synchronized"` |  |
| argo-cd.controller.rules.spec[1].expr | string | `"argocd_app_info{sync_status!=\"Synced\"} == 1\n"` |  |
| argo-cd.controller.rules.spec[1].for | string | `"12h"` |  |
| argo-cd.controller.rules.spec[1].labels.severity | string | `"warning"` |  |
| argo-cd.dex.enabled | bool | `true` |  |
| argo-cd.dex.metrics.enabled | bool | `true` |  |
| argo-cd.dex.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.dex.resources.limits.cpu | string | `"300m"` |  |
| argo-cd.dex.resources.limits.memory | string | `"2Gi"` |  |
| argo-cd.dex.resources.requests.cpu | string | `"10m"` |  |
| argo-cd.dex.resources.requests.memory | string | `"32Mi"` |  |
| argo-cd.enabled | bool | `true` |  |
| argo-cd.externalRedis.host | string | `"redis-master.argocd"` |  |
| argo-cd.fullnameOverride | string | `"argocd"` |  |
| argo-cd.global.image | string | `nil` |  |
| argo-cd.global.logging.format | string | `"json"` |  |
| argo-cd.global.logging.level | string | `"warn"` |  |
| argo-cd.notifications.metrics.enabled | bool | `true` |  |
| argo-cd.notifications.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.notifications.resources.limits.cpu | string | `"300m"` |  |
| argo-cd.notifications.resources.limits.memory | string | `"2Gi"` |  |
| argo-cd.notifications.resources.requests.cpu | string | `"5m"` |  |
| argo-cd.notifications.resources.requests.memory | string | `"64Mi"` |  |
| argo-cd.redis.enabled | bool | `false` |  |
| argo-cd.repoServer.env[0].name | string | `"HELM_PLUGINS"` |  |
| argo-cd.repoServer.env[0].value | string | `"/custom-tools/helm-plugins/"` |  |
| argo-cd.repoServer.env[1].name | string | `"HELM_SECRETS_SOPS_PATH"` |  |
| argo-cd.repoServer.env[1].value | string | `"/custom-tools/sops"` |  |
| argo-cd.repoServer.env[2].name | string | `"HELM_SECRETS_VALS_PATH"` |  |
| argo-cd.repoServer.env[2].value | string | `"/custom-tools/vals"` |  |
| argo-cd.repoServer.env[3].name | string | `"HELM_SECRETS_KUBECTL_PATH"` |  |
| argo-cd.repoServer.env[3].value | string | `"/custom-tools/kubectl"` |  |
| argo-cd.repoServer.env[4].name | string | `"HELM_SECRETS_CURL_PATH"` |  |
| argo-cd.repoServer.env[4].value | string | `"/custom-tools/curl"` |  |
| argo-cd.repoServer.env[5].name | string | `"HELM_SECRETS_VALUES_ALLOW_SYMLINKS"` |  |
| argo-cd.repoServer.env[5].value | string | `"false"` |  |
| argo-cd.repoServer.env[6].name | string | `"HELM_SECRETS_VALUES_ALLOW_ABSOLUTE_PATH"` |  |
| argo-cd.repoServer.env[6].value | string | `"false"` |  |
| argo-cd.repoServer.env[7].name | string | `"HELM_SECRETS_VALUES_ALLOW_PATH_TRAVERSAL"` |  |
| argo-cd.repoServer.env[7].value | string | `"false"` |  |
| argo-cd.repoServer.initContainers[0].args[0] | string | `"mkdir -p /custom-tools/helm-plugins\nwget -qO- https://github.com/jkroepke/helm-secrets/releases/download/v${HELM_SECRETS_VERSION}/helm-secrets.tar.gz | tar -C /custom-tools/helm-plugins -xzf-;\n\nwget -qO /custom-tools/sops https://github.com/mozilla/sops/releases/download/v${SOPS_VERSION}/sops-v${SOPS_VERSION}.linux\nwget -qO /custom-tools/kubectl https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl\n\nwget -qO- https://github.com/variantdev/vals/releases/download/v${VALS_VERSION}/vals_${VALS_VERSION}_linux_amd64.tar.gz | tar -xzf- -C /custom-tools/ vals;\n\n# helm secrets wrapper mode installation (optional)\n# RUN printf '#!/usr/bin/env sh\\nexec %s secrets \"$@\"' \"${HELM_SECRETS_HELM_PATH}\" >\"/usr/local/sbin/helm\" && chmod +x \"/custom-tools/helm\"\n\nchmod +x /custom-tools/*\n"` |  |
| argo-cd.repoServer.initContainers[0].command[0] | string | `"sh"` |  |
| argo-cd.repoServer.initContainers[0].command[1] | string | `"-ec"` |  |
| argo-cd.repoServer.initContainers[0].env[0].name | string | `"HELM_SECRETS_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[0].value | string | `"4.5.0"` |  |
| argo-cd.repoServer.initContainers[0].env[1].name | string | `"KUBECTL_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[1].value | string | `"1.28.0"` |  |
| argo-cd.repoServer.initContainers[0].env[2].name | string | `"VALS_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[2].value | string | `"0.27.1"` |  |
| argo-cd.repoServer.initContainers[0].env[3].name | string | `"SOPS_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[3].value | string | `"3.7.3"` |  |
| argo-cd.repoServer.initContainers[0].image | string | `"alpine:latest"` |  |
| argo-cd.repoServer.initContainers[0].name | string | `"download-tools"` |  |
| argo-cd.repoServer.initContainers[0].volumeMounts[0].mountPath | string | `"/custom-tools"` |  |
| argo-cd.repoServer.initContainers[0].volumeMounts[0].name | string | `"custom-tools"` |  |
| argo-cd.repoServer.metrics.enabled | bool | `true` |  |
| argo-cd.repoServer.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.repoServer.rbac[0].apiGroups[0] | string | `""` |  |
| argo-cd.repoServer.rbac[0].resources[0] | string | `"secrets"` |  |
| argo-cd.repoServer.rbac[0].verbs[0] | string | `"get"` |  |
| argo-cd.repoServer.serviceAccount.create | bool | `true` |  |
| argo-cd.repoServer.serviceAccount.name | string | `"argocd-repo-server"` |  |
| argo-cd.repoServer.volumeMounts[0].mountPath | string | `"/custom-tools"` |  |
| argo-cd.repoServer.volumeMounts[0].name | string | `"custom-tools"` |  |
| argo-cd.repoServer.volumeMounts[1].mountPath | string | `"/sops-gpg/"` |  |
| argo-cd.repoServer.volumeMounts[1].name | string | `"sops-gpg"` |  |
| argo-cd.repoServer.volumes[0].emptyDir | object | `{}` |  |
| argo-cd.repoServer.volumes[0].name | string | `"custom-tools"` |  |
| argo-cd.repoServer.volumes[1].name | string | `"sops-gpg"` |  |
| argo-cd.repoServer.volumes[1].secret.secretName | string | `"sops-gpg"` |  |
| argo-cd.server.config."helm.valuesFileSchemes" | string | `"secrets+gpg-import, secrets+gpg-import-kubernetes, secrets+age-import, secrets+age-import-kubernetes, secrets,secrets+literal, https"` |  |
| argo-cd.server.config."statusbadge.enabled" | string | `"true"` |  |
| argo-cd.server.config.url | string | `"https://argo.shortlink.best"` |  |
| argo-cd.server.configAnnotations | object | `{}` |  |
| argo-cd.server.extensions.enabled | bool | `true` |  |
| argo-cd.server.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTPS"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"proxy_ssl_server_name on;\nproxy_ssl_name $host;"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"false"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/secure-backends" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-redirect" | string | `"true"` |  |
| argo-cd.server.ingress.enabled | bool | `true` |  |
| argo-cd.server.ingress.hosts[0] | string | `"argo.shortlink.best"` |  |
| argo-cd.server.ingress.https | bool | `true` |  |
| argo-cd.server.ingress.ingressClassName | string | `"nginx"` |  |
| argo-cd.server.ingress.tls[0].hosts[0] | string | `"argo.shortlink.best"` |  |
| argo-cd.server.ingress.tls[0].secretName | string | `"argo-ingress-tls"` |  |
| argo-cd.server.metrics.enabled | bool | `true` |  |
| argo-cd.server.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.server.rbacConfig."policy.csv" | string | `"p, role:org-admin, applications, *, */*, allow\np, role:org-admin, clusters, get, *, allow\np, role:org-admin, repositories, get, *, allow\np, role:org-admin, repositories, create, *, allow\np, role:org-admin, repositories, update, *, allow\np, role:org-admin, repositories, delete, *, allow\ng, shortlink-org:devops, role:org-admin\n"` |  |
| argo-cd.server.rbacConfig."policy.default" | string | `"role:readonly"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
