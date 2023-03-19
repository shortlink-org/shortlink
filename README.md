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

<div style="align-items: center; display: flex;">
  <a href="https://www.producthunt.com/posts/shortlink-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-shortlink&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=374140&theme=light" alt="ShortLink - Get&#0032;ready&#0032;to&#0032;share&#0032;your&#0032;links&#0032;with&#0032;ease&#0033; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>
  <img height="100px" src="https://slsa.dev/images/SLSA-Badge-full-level1.svg" alt="SLSA">
</div>

</div>

<hr />

### High Level Architecture 🚀

![shortlink-architecture](./docs/shortlink-architecture.png)
_Please [star ⭐](https://github.com/shortlink-org/shortlink/stargazers) the repo if you want us to continue developing and improving ShortLink! 😀_


### Architecture decision records (ADR)

An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.

+ [Decisions](./docs/ADR/decisions)
+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)

### Architecture

You can find the architecture documentation [here](./docs/ADR/decisions/0011-application-architecture-documentation.md).

### Services

<details><summary>DETAILS</summary>
<p>

| Service          | Description                                                         | Language/Framework | Docs                                                  | Status                                                                                                                                                              |
|------------------|---------------------------------------------------------------------|--------------------|-------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| api              | Internal GateWay                                                    | Go                 | [docs](./internal/services/api/README.md)             | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-api&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-api)               |
| billing          | Billing service                                                     | Go                 | [docs](./internal/services/billing/README.md)         | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-billing&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-billing)       |
| bot              | Telegram bot                                                        | JAVA               | [docs](./internal/services/bot/README.md)             | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-bot&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-bot)               |                                                                    
| chat             | Chat service                                                        | Elixir (Phoenix)   | [docs](./internal/services/chat/README.md)            | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-chat&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-chat)             |                                                                   
| chrome-extension | Chrome extension                                                    | JavaScript         | [docs](internal/extension/chrome-extension/README.md) |                                                                                                                                                                     |                                                                         
| shortdb          | Custom database                                                     | Go                 | [docs](./pkg/shortdb/README.md)                       | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortldb&revision=true)](https://shortlink.best/argo/cd/applications/shortldb)                         |                                                                          
| shortdb-operator | Kubernetes Operator for [shortdb](./pkg/shortdb/README.md) database | Go                 | [docs](./pkg/shortdb-operator/README.md)              | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortldb-operator&revision=true)](https://shortlink.best/argo/cd/applications/shortldb-operator)       |                                                                 
| csi              | CSI example                                                         | Go                 | [docs](./internal/services/csi/README.md)             | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-csi&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-csi)               |                                                                     
| link             | Link service                                                        | Go                 | [docs](./internal/services/api/README.md)             | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-link&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-link)             |                                                                    
| logger           | Logger service                                                      | Go                 | [docs](./internal/services/logger/README.md)          | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-logger&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-logger)         |                                                                  
| merch            | Merch store                                                         | Coming soon        | [docs](./internal/services/merch/README.md)           | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-merch&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-merch)           |                                                                   
| metadata         | Parser site by API                                                  | Go                 | [docs](./internal/services/metadata/README.md)        | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-metadata&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-metadata)     |                                                                
| newsletter       | Newsletter service                                                  | Rust               | [docs](./internal/services/newsletter/README.md)      | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-newsletter&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-newsletter) |                                                              
| notify           | Send notify to smtp, slack, telegram                                | Go                 | [docs](./internal/services/notify/README.md)          | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-notify&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-notify)         |                                                                  
| proxy            | Proxy service for redirect to original URL                          | TypeScript         | [docs](./internal/services/proxy/README.md)           | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-proxy&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-proxy)           |                                                                   
| referral         | Referral program                                                    | Python             | [docs](./internal/services/referral/README.md)        | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-referral&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-referral)     |                                                                
| search           | Search service                                                      | Coming soon        | [docs](./internal/services/search/README.md)          | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-search&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-search)         |                                                                  
| shortctl         | Shortlink CLI                                                       | Go                 | [docs](./internal/services/cli/README.md)             |                                                                                                                                                                     |                                                                   
| stats            | Stats server                                                        | CPP                | [docs](./internal/services/stats/README.md)           | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-stats&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-stats)           |                                                                   
| support          | Support service                                                     | PHP                | [docs](./internal/services/support/README.md)         | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-support&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-support)       |                                                                 
| wallet           | Wallet service                                                      | Go (Solidity)      | [docs](./internal/services/wallet/README.md)          |                                                                                                                                                                     |
| ws               | Websocket service                                                   | Go                 | [docs](./internal/services/ws/README.md)              | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-ws&revision=true)                                                                            |

</p>
</details>

### Third-party Service

<details><summary>DETAILS</summary>
<p>

| Service       | Description                                                             | Language/Framework | Docs                                                | Status                                                                                                                                                |
|---------------|-------------------------------------------------------------------------|--------------------|-----------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| ory/kratos    | User management service                                                 | Go                 | [docs](https://www.ory.sh/kratos/docs/)             | [![App Status](https://shortlink.best/argo/cd/api/badge?name=auth&revision=true)](https://shortlink.best/argo/cd/applications/auth)                   |          
| ory/hydra     | OAuth 2.0 Provider                                                      | Go                 | [docs](https://www.ory.sh/keto/docs/)               | [![App Status](https://shortlink.best/argo/cd/api/badge?name=auth&revision=true)](https://shortlink.best/argo/cd/applications/auth)                   |          
| backstage     | Backstage is an open platform for building developer portals.           | TypeScript         | [docs](https://backstage.io/docs/)                  | [![App Status](https://shortlink.best/argo/cd/api/badge?name=backstage&revision=true)](https://shortlink.best/argo/cd/applications/backstage)         |    
| grafana       | Grafana is the open source analytics & monitoring solution for          | More               | [docs](https://grafana.com/docs/)                   | [![App Status](https://shortlink.best/argo/cd/api/badge?name=grafana&revision=true)](https://shortlink.best/argo/cd/applications/grafana)             |       
| cert-manager  | Automatically provision and manage TLS certificates in Kubernetes       | Go                 | [docs](https://cert-manager.io/docs/)               | [![App Status](https://shortlink.best/argo/cd/api/badge?name=cert-manager&revision=true)](https://shortlink.best/argo/cd/applications/cert-manager)   |  
| istio         | Istio is an open platform to connect, manage, and secure microservices. | Go                 | [docs](https://istio.io/latest/docs/)               | [![App Status](https://shortlink.best/argo/cd/api/badge?name=istio&revision=true)](https://shortlink.best/argo/cd/applications/istio)                 |         
| nginx-ingress | Ingress controller for Kubernetes using NGINX                           | Go                 | [docs](https://kubernetes.github.io/ingress-nginx/) | [![App Status](https://shortlink.best/argo/cd/api/badge?name=nginx-ingress&revision=true)](https://shortlink.best/argo/cd/applications/nginx-ingress) | 
| kafka         | Kafka is used as a message broker                                       | Java               | [docs](https://kafka.apache.org/)                   | [![App Status](https://shortlink.best/argo/cd/api/badge?name=kafka&revision=true)](https://shortlink.best/argo/cd/applications/kafka)                 |         
| keycloak      | Keycloak is an open source identity and access management solution      | Java               | [docs](https://www.keycloak.org/documentation.html) | [![App Status](https://shortlink.best/argo/cd/api/badge?name=keycloak&revision=true)](https://shortlink.best/argo/cd/applications/keycloak)           | 

</p>
</details>

### UI

<details><summary>DETAILS</summary>
<p>

| Service   | Description                          | Language/Framework        | Docs                           | Status                                                                                                                                                        |
|-----------|--------------------------------------|---------------------------|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| landing   | Welcome page                         | JS/NextJS                 | [docs](./ui/landing/README.md) | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-landing&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-landing) |
| next      | UI service                           | JS/NextJS                 | [docs](./ui/next/README.md)    | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-next&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-next)       |
| ui-kit    | UI kit for ShortLink                 | JS/React                  | [docs](./ui/ui-kit/README.md)  | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-ui-kit&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-ui-kit)   |

</p>
</details>

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

<details><summary>DETAILS</summary>
<p>

###### For run
```
make minikube-up
make helm-shortlink-up
```

###### For down
```
make minikube-down
```

</p>
</details>

##### Skaffold [(link)](https://skaffold.dev/)

<details><summary>DETAILS</summary>
<p>

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

</p>
</details>

### Mobile

+ `Hello World` on flutter ;-)

### License

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=large)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_large)

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/shortlink-org/shortlink&style=flat
