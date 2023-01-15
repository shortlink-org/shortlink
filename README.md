<div align="center">

# shortlink 

Shortlink is an open-source project that provides a pretty user interface and respects GDPR.   
We use edge technologies and have many years of experience.  

We're constantly researching the best solutions on the market so that we can benefit  
our community and solve a problem for millions of people.

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/shortlink)](https://artifacthub.io/packages/search?repo=shortlink)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/shortlink-org/shortlink)](https://pkg.go.dev/mod/github.com/shortlink-org/shortlink)
[![codecov](https://codecov.io/gh/shortlink-org/shortlink/branch/main/graph/badge.svg?token=Wxz5bI4QzF)](https://codecov.io/gh/shortlink-org/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/shortlink-org/shortlink)](https://goreportcard.com/report/github.com/shortlink-org/shortlink)
[![Releases](https://img.shields.io/github/release-pre/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/blob/main/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)
[![StackShare](http://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/shortlink-org/shortlink)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=shield)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_shield)

<hr />

<a href="https://www.producthunt.com/posts/shortlink-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-shortlink&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=374140&theme=light" alt="ShortLink - Get&#0032;ready&#0032;to&#0032;share&#0032;your&#0032;links&#0032;with&#0032;ease&#0033; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

</div>

<hr />

### High Level Architecture 🚀

![shortlink-architecture](./docs/shortlink-architecture.png)

### Architecture decision records (ADR)

An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.

+ [Decisions](./docs/ADR/decisions)
+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)

### Architecture

You can find the architecture documentation [here](./docs/ADR/decisions/0011-application-architecture-documentation.md).

### Services

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

| Service     | Description                                                     | Language/Framework | Docs                                    |
|-------------|-----------------------------------------------------------------|--------------------|-----------------------------------------|
| ory/kratos  | User management service                                         | Go                 | [docs](https://www.ory.sh/kratos/docs/) |
| ory/hydra   | OAuth 2.0 Provider                                              | Go                 | [docs](https://www.ory.sh/keto/docs/)   |
| backstage   | Backstage is an open platform for building developer portals.   | TypeScript         | [docs](https://backstage.io/docs/)      |

### UI

| Service   | Description                          | Language/Framework        | Docs                                           |
|-----------|--------------------------------------|---------------------------|------------------------------------------------|
| landing   | Welcome page                         | JS/NextJS                 | [docs](./ui/landing/README.md)                 |
| next      | UI service                           | JS/NextJS                 | [docs](./ui/next/README.md)                    |
| ui-kit    | UI kit for ShortLink                 | JS/React                  | [docs](./ui/ui-kit/README.md)                  |

### Docs

- [Ops](./ops/README.md)

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

### Observability

<details><summary>DETAILS</summary>
<p>

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

</p>
</details>
  
### Mobile

+ `Hello World` on flutter ;-)

### License

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=large)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_large)

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/shortlink-org/shortlink&style=flat
