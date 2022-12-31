<div align="center">

# shortlink

Shortlink is an open-source project that provides a pretty user interface and respects GDPR.   
We use edge technologies and have many years of experience.  

We're constantly researching the best solutions on the market so that we can benefit  
our community and solve a problem for millions of people.

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/batazor/shortlink)](https://pkg.go.dev/mod/github.com/batazor/shortlink)
[![codecov](https://codecov.io/gh/batazor/shortlink/branch/main/graph/badge.svg?token=Wxz5bI4QzF)](https://codecov.io/gh/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)
[![Releases](https://img.shields.io/github/release-pre/batazor/shortlink.svg)](https://github.com/batazor/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/batazor/shortlink.svg)](https://github.com/batazor/shortlink/blob/main/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)

</div>

<hr />

### High Level Architecture 🚀

![shortlink-architecture](./docs/shortlink-architecture.png)

### Architecture (miro.com)

- [Low-level schema](https://miro.com/app/board/o9J_laImQpo=/)
- [Auth](https://miro.com/app/board/o9J_lA5Wmhg=/)
- [Event Sourcing](https://miro.com/app/board/o9J_l-6o1U0=/)
- [C4](./docs/c4)

### Architecture decision records (ADR)

An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.

+ [Decisions](./docs/ADR/decisions)
+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)

### Techradar

[shortlink-techradar](https://radar.thoughtworks.com/?sheetId=https://raw.githubusercontent.com/batazor/shortlink/main/docs/thoughtworks.radar.csv)

##### Services

| Service           | Description                                                           | Language/Framework | Docs                                                  |
|-------------------|-----------------------------------------------------------------------|--------------------|-------------------------------------------------------|
| api               | Internal GateWay                                                      | Go                 | [docs](./internal/services/api/README.md)             |
| billing           | Billing service                                                       | Go                 | [docs](./internal/services/billing/README.md)         |
| bot               | Telegram bot                                                          | JAVA               | [docs](./internal/services/bot/README.md)             |
| chat              | Chat service                                                          | Elixir (Phoenix)   | [docs](./internal/services/chat/README.md)            |
| chrome-extension  | Chrome extension                                                      | JavaScript         | [docs](internal/extension/chrome-extension/README.md) |
| shortdb           | Custom database                                                       | Go                 | [docs](./pkg/shortdb/README.md)                       |
| shortdb-operator  | Kubernetes Operator for [shortdb]((./pkg/shortdb/README.md)) database | Go                 | [docs](./pkg/shortdb-operator/README.md)              |
| csi               | CSI example                                                           | Go                 | [docs](./internal/services/csi/README.md)             |
| link              | Link service                                                          | Go                 | [docs](./internal/services/api/README.md)             |
| logger            | Logger service                                                        | Go                 | [docs](./internal/services/logger/README.md)          |
| merch             | Merch store                                                           | Coming soon        | [docs](./internal/services/merch/README.md)           |
| metadata          | Parser site by API                                                    | Go                 | [docs](./internal/services/metadata/README.md)        |
| newsletter        | Newsletter service                                                    | Rust               | [docs](./internal/services/newsletter/README.md)      |
| notify            | Send notify to smtp, slack, telegram                                  | Go                 | [docs](./internal/services/notify/README.md)          |
| proxy             | Proxy service for redirect to original URL                            | TypeScript         | [docs](./internal/services/proxy/README.md)           |
| referral          | Referral program                                                      | Python             | [docs](./internal/services/referral/README.md)        |
| search            | Search service                                                        | Coming soon        | [docs](./internal/services/search/README.md)          |
| shortctl          | Shortlink CLI                                                         | Go                 | [docs](./internal/services/cli/README.md)             |
| stats             | Stats server                                                          | CPP                | [docs](./internal/services/stats/README.md)           |
| support           | Support service                                                       | PHP                | [docs](./internal/services/support/README.md)         |
| wallet            | Wallet service                                                        | Go (Solidity)      | [docs](./internal/services/wallet/README.md)          |
| ws                | Webscoket service                                                     | Go                 | [docs](./internal/services/ws/README.md)              |

### Third-party Service

| Service                 | Description                                                           | Language/Framework        | Docs                                             |
|-------------------------|-----------------------------------------------------------------------|---------------------------|--------------------------------------------------|
| ory/kratos              | User management service                                               | Go                        | [docs](https://www.ory.sh/kratos/docs/)          |
| ory/hydra               | OAuth 2.0 Provider                                                    | Go                        | [docs](https://www.ory.sh/keto/docs/)            |

### Run

<details><summary>DETAILS</summary>
<p>

##### Require

###### Install GIT sub-repository

```
git submodule update --init --recursive
```

##### docker compose

###### For run
```
make run
```

###### For down
```
make down
```


##### Kubernetes (1.21+)

###### For run
```
make minikube-up
make helm-shortlink-up
```

###### For down
```
make minikube-down
```

##### Skaffold [(link)](https://skaffold.dev/)

###### For run
```
make skaffold-init
make skaffold-up
```

###### For down
```
make skaffold-down
```

###### Debug mode
```
make skaffold-debug
```

</p>
</details>

### Configuration

<details><summary>DETAILS</summary>
<p>

##### [12 factors: ENV](https://12factor.net/config)

[View ENV Variables](./docs/env.md)

</p>
</details>

### OpenTracing

<details><summary>DETAILS</summary>
<p>

![http_add_link](./docs/opentracing_add_link.png)

</p>
</details>

### UI

| Service   | Description                          | Language/Framework        | Docs                                           |
|-----------|--------------------------------------|---------------------------|------------------------------------------------|
| landing   | Welcome page                         | JS/NextJS                 | [docs](./ui/landing/README.md)                 |
| next      | UI service                           | JS/NextJS                 | [docs](./ui/next/README.md)                    |
| ui-kit    | UI kit for ShortLink                 | JS/React                  | [docs](./ui/ui-kit/README.md)                  |

### MQ

| [Kafka](https://kafka.apache.org/) | [RabbitMQ](https://www.rabbitmq.com/) | [NATS](https://nats.io/) |
|------------------------------------|---------------------------------------|--------------------------|

### Cloud-Native stack

+ Development
  + [Skaffold](https://skaffold.dev/)
  + Telepresence
+ Security
  + SOPS
+ HealthCheck
+ Support K8S
  + Helm Chart
    + [pingcap/chaos-meshh](https://github.com/pingcap/chaos-mesh)
  + Minikube
  + Backup/Restore [(Velero)](https://velero.io/)
  + Custom CSI driver (fork [csi-driver-host-pat](https://github.com/kubernetes-csi/csi-driver-host-path))
+ MetalLB
+ [kyverno](https://kyverno.io/) - Kubernetes Native Policy Management
+ Storage
    + [rook-ceph](https://rook.io/)
        + ceph cluster (3 node)
        + grafana dashboard
        + prometheus metrics
+ Ingress (Gateway)
    + Istio
        + [kiali](https://kiali.io/) - The Console for Istio Service Mesh
    + Nginx
    + Traefik

### Observability

+ **[Prometheus](https://prometheus.io/)** - Monitoring system
    + prometheus-operator
        + notify: slack, email, telegram

---

+ **Grafana stack (LGTM)**
    * [Grafana](https://github.com/grafana/grafana), the open and composable observability and data visualization
      platform.
    * [Loki](https://github.com/grafana/loki), like Prometheus, but for logs.
        + [docs](docs/tutorial/logger.md)

    + [Tempo](https://grafana.com/docs/tempo/latest/), a high volume, high throughput distributed tracing system.

    * [Grafana](https://github.com/grafana/grafana), the open and composable observability and data visualization
      platform.

    + [OnCall](https://grafana.com/oss/oncall/) - On-call scheduling
    + [Phlare](https://grafana.com/oss/phlare/) - Profiling and flame graphs

### Mobile

+ `Hello World` on flutter ;-)

### CI/CD

| [GitLab CI](./ops/gitlab/README.md) | [GitHub CI](./.github/DOCS.md) | [ArgoCD](./ops/argocd/README.md) |
|-------------------------------------|--------------------------------|----------------------------------|

## -~- THE END -~-

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/batazor/shortlink&style=flat
