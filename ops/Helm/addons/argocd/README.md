# argo

![Version: 0.3.14](https://img.shields.io/badge/Version-0.3.14-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 5.19.4](https://img.shields.io/badge/AppVersion-5.19.4-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.24.0 || >= v1.24.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-cd | 5.34.1 |
| https://argoproj.github.io/argo-helm | argo-events | 2.3.3 |
| https://argoproj.github.io/argo-helm | argo-rollouts | 2.28.0 |
| https://argoproj.github.io/argo-helm | argo-workflows | 0.27.0 |
| https://argoproj.github.io/argo-helm | argocd-apps | 1.0.1 |
| https://argoproj.github.io/argo-helm | argocd-image-updater | 0.9.1 |
| oci://registry-1.docker.io/bitnamicharts | redis | 17.11.2 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| argo-cd.applicationSet.metrics.enabled | bool | `true` |  |
| argo-cd.applicationSet.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.applicationSet.resources.limits.cpu | string | `"150m"` |  |
| argo-cd.applicationSet.resources.limits.memory | string | `"256Mi"` |  |
| argo-cd.applicationSet.resources.requests.cpu | string | `"50m"` |  |
| argo-cd.applicationSet.resources.requests.memory | string | `"64Mi"` |  |
| argo-cd.configs.repositories.shortlink.name | string | `"shortlink"` |  |
| argo-cd.configs.repositories.shortlink.type | string | `"git"` |  |
| argo-cd.configs.repositories.shortlink.url | string | `"https://github.com/shortlink-org/shortlink"` |  |
| argo-cd.controller.metrics.applicationLabels.enabled | bool | `true` |  |
| argo-cd.controller.metrics.enabled | bool | `true` |  |
| argo-cd.controller.metrics.serviceMonitor.enabled | bool | `true` |  |
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
| argo-cd.dex.resources.limits.cpu | string | `"50m"` |  |
| argo-cd.dex.resources.limits.memory | string | `"64Mi"` |  |
| argo-cd.dex.resources.requests.cpu | string | `"10m"` |  |
| argo-cd.dex.resources.requests.memory | string | `"32Mi"` |  |
| argo-cd.enabled | bool | `true` |  |
| argo-cd.externalRedis.host | string | `"redis-master.argocd"` |  |
| argo-cd.fullnameOverride | string | `"argocd"` |  |
| argo-cd.global.logging.format | string | `"json"` |  |
| argo-cd.global.logging.level | string | `"warn"` |  |
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
| argo-cd.repoServer.initContainers[0].env[0].value | string | `"4.4.2"` |  |
| argo-cd.repoServer.initContainers[0].env[1].name | string | `"KUBECTL_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[1].value | string | `"1.27.0"` |  |
| argo-cd.repoServer.initContainers[0].env[2].name | string | `"VALS_VERSION"` |  |
| argo-cd.repoServer.initContainers[0].env[2].value | string | `"0.24.0"` |  |
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
| argo-cd.server.config.url | string | `"https://shortlink.best/argo/cd"` |  |
| argo-cd.server.configAnnotations | object | `{}` |  |
| argo-cd.server.extensions.enabled | bool | `true` |  |
| argo-cd.server.extraArgs[0] | string | `"--rootpath"` |  |
| argo-cd.server.extraArgs[1] | string | `"/argo/cd"` |  |
| argo-cd.server.extraArgs[2] | string | `"--basehref"` |  |
| argo-cd.server.extraArgs[3] | string | `"/argo/cd"` |  |
| argo-cd.server.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTPS"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet" | string | `"proxy_ssl_server_name on;\nproxy_ssl_name $host;"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/secure-backends" | string | `"true"` |  |
| argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-redirect" | string | `"true"` |  |
| argo-cd.server.ingress.enabled | bool | `true` |  |
| argo-cd.server.ingress.hosts[0] | string | `"shortlink.best"` |  |
| argo-cd.server.ingress.https | bool | `true` |  |
| argo-cd.server.ingress.ingressClassName | string | `"nginx"` |  |
| argo-cd.server.ingress.paths[0] | string | `"/argo/cd(/|$)(.*)"` |  |
| argo-cd.server.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| argo-cd.server.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| argo-cd.server.metrics.enabled | bool | `true` |  |
| argo-cd.server.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-cd.server.rbacConfig."policy.csv" | string | `"p, role:org-admin, applications, *, */*, allow\np, role:org-admin, clusters, get, *, allow\np, role:org-admin, repositories, get, *, allow\np, role:org-admin, repositories, create, *, allow\np, role:org-admin, repositories, update, *, allow\np, role:org-admin, repositories, delete, *, allow\ng, shortlink-org:devops, role:org-admin\n"` |  |
| argo-cd.server.rbacConfig."policy.default" | string | `"role:readonly"` |  |
| argo-events.enabled | bool | `false` |  |
| argo-events.fullnameOverride | string | `"argo-events"` |  |
| argo-rollouts.controller.resources.limits.cpu | string | `"100m"` |  |
| argo-rollouts.controller.resources.limits.memory | string | `"128Mi"` |  |
| argo-rollouts.controller.resources.requests.cpu | string | `"20m"` |  |
| argo-rollouts.controller.resources.requests.memory | string | `"16Mi"` |  |
| argo-rollouts.dashboard.enabled | bool | `true` |  |
| argo-rollouts.dashboard.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTP"` |  |
| argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-rollouts.dashboard.ingress.enabled | bool | `true` |  |
| argo-rollouts.dashboard.ingress.hosts[0] | string | `"shortlink.best"` |  |
| argo-rollouts.dashboard.ingress.ingressClassName | string | `"nginx"` |  |
| argo-rollouts.dashboard.ingress.paths[0] | string | `"/argo/rollouts"` |  |
| argo-rollouts.dashboard.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| argo-rollouts.dashboard.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| argo-rollouts.dashboard.readonly | bool | `true` |  |
| argo-rollouts.dashboard.resources.limits.cpu | string | `"100m"` |  |
| argo-rollouts.dashboard.resources.limits.memory | string | `"256Mi"` |  |
| argo-rollouts.dashboard.resources.requests.cpu | string | `"20m"` |  |
| argo-rollouts.dashboard.resources.requests.memory | string | `"16Mi"` |  |
| argo-rollouts.enabled | bool | `true` |  |
| argo-rollouts.fullnameOverride | string | `"argo-rollouts"` |  |
| argo-rollouts.metrics.enabled | bool | `true` |  |
| argo-rollouts.metrics.serviceMonitor.enabled | bool | `true` |  |
| argo-workflows.controller.metricsConfig.enabled | bool | `true` |  |
| argo-workflows.controller.serviceMonitor.enabled | bool | `true` |  |
| argo-workflows.controller.telemetryConfig.enabled | bool | `true` |  |
| argo-workflows.controller.workflowNamespaces | list | `[]` |  |
| argo-workflows.enabled | bool | `false` |  |
| argo-workflows.fullnameOverride | string | `"argo-workflows"` |  |
| argo-workflows.server.baseHref | string | `"/argo/workflows/"` |  |
| argo-workflows.server.extraArgs[0] | string | `"--basehref"` |  |
| argo-workflows.server.extraArgs[1] | string | `"/argo/workflows/"` |  |
| argo-workflows.server.extraArgs[2] | string | `"--auth-mode=server"` |  |
| argo-workflows.server.extraEnv[0].name | string | `"ARGO_BASE_HREF"` |  |
| argo-workflows.server.extraEnv[0].value | string | `"/argo/workflows"` |  |
| argo-workflows.server.ingress.annotations."cert-manager.io/cluster-issuer" | string | `"cert-manager-production"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol" | string | `"HTTP"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules" | string | `"true"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target" | string | `"/$1"` |  |
| argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/use-regex" | string | `"true"` |  |
| argo-workflows.server.ingress.enabled | bool | `true` |  |
| argo-workflows.server.ingress.hosts[0] | string | `"shortlink.best"` |  |
| argo-workflows.server.ingress.ingressClassName | string | `"nginx"` |  |
| argo-workflows.server.ingress.paths[0] | string | `"/argo/workflows/?(.*)"` |  |
| argo-workflows.server.ingress.tls[0].hosts[0] | string | `"shortlink.best"` |  |
| argo-workflows.server.ingress.tls[0].secretName | string | `"shortlink-ingress-tls"` |  |
| argocd-apps.applications | list | `[]` (See [values.yaml]) | Deploy Argo CD Applications within this helm release # Ref: https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/ |
| argocd-apps.enabled | bool | `true` |  |
| argocd-apps.projects | list | `[]` (See [values.yaml]) | Deploy Argo CD Projects within this helm release # Ref: https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/ |
| argocd-image-updater.config.applicationsAPIKind | string | `"kubernetes"` |  |
| argocd-image-updater.config.gitCommitMail | string | `"argocd@shortlink.best"` |  |
| argocd-image-updater.config.gitCommitTemplate | string | `"build: automatic update of {{ .AppName }}\n\n{{ range .AppChanges -}}\nupdates image {{ .Image }} tag '{{ .OldTag }}' to '{{ .NewTag }}'\n{{ end -}}\n\nSigned-off-by: argocd <argocd@shortlink.best>\n"` |  |
| argocd-image-updater.config.gitCommitUser | string | `"argocd-image-updater"` |  |
| argocd-image-updater.enabled | bool | `false` |  |
| argocd-image-updater.fullnameOverride | string | `"argocd-image-updater"` |  |
| argocd-image-updater.metrics.enabled | bool | `true` |  |
| argocd-image-updater.metrics.serviceMonitor.enabled | bool | `true` |  |
| argocd-image-updater.registries[0].api_url | string | `"https://registry.gitlab.com"` |  |
| argocd-image-updater.registries[0].default | bool | `true` |  |
| argocd-image-updater.registries[0].name | string | `"GitLab"` |  |
| argocd-image-updater.registries[0].ping | string | `"yes"` |  |
| argocd-image-updater.resources.limits.cpu | string | `"100m"` |  |
| argocd-image-updater.resources.limits.memory | string | `"128Mi"` |  |
| argocd-image-updater.resources.requests.cpu | string | `"10m"` |  |
| argocd-image-updater.resources.requests.memory | string | `"16Mi"` |  |
| argocd-image-updater.updateStrategy.type | string | `"Recreate"` |  |
| redis.auth.enabled | bool | `false` |  |
| redis.auth.password | string | `""` |  |
| redis.fullnameOverride | string | `"redis"` |  |
| redis.master.kind | string | `"Deployment"` |  |
| redis.master.persistence.enabled | bool | `false` |  |
| redis.master.persistence.storageClass | string | `"local-path"` |  |
| redis.master.resources.limits.cpu | string | `"500m"` |  |
| redis.master.resources.limits.memory | string | `"350Mi"` |  |
| redis.master.resources.requests.cpu | string | `"10m"` |  |
| redis.master.resources.requests.memory | string | `"64Mi"` |  |
| redis.metrics.enabled | bool | `true` |  |
| redis.metrics.prometheusRule.enabled | bool | `true` |  |
| redis.metrics.serviceMonitor.enabled | bool | `true` |  |
| redis.replica.persistence.enabled | bool | `false` |  |
| redis.replica.replicaCount | int | `0` |  |
| redis.replica.resources.limits.cpu | string | `"500m"` |  |
| redis.replica.resources.limits.memory | string | `"256Mi"` |  |
| redis.replica.resources.requests.cpu | string | `"10m"` |  |
| redis.replica.resources.requests.memory | string | `"64Mi"` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
