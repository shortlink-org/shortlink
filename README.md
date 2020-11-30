# shortlink

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/batazor/shortlink)](https://pkg.go.dev/mod/github.com/batazor/shortlink)
[![codecov](https://codecov.io/gh/batazor/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)
[![Releases](https://img.shields.io/github/release-pre/batazor/shortlink.svg)](https://github.com/batazor/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/batazor/shortlink.svg)](https://github.com/batazor/shortlink/blob/master/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)

Shortlink service

### High Level Architecture

![shortlink-arhitecture](./docs/shortlink-arhitecture.png)

##### Requirements

- docker
- docker-compose
- protoc 3.7.1+
- ansible
- ansible-galaxy

##### Services

| Service     | Description                          | Language/Framework | Docs                       |
|-------------|--------------------------------------|--------------------|----------------------------|
| shortlink   | Shortlink service                    | Go                 |                            |
| logger      | Logger service                       | Go                 |                            |
| bot         | Send notify to smtp, slack, telegram | Go                 |                            |
| shortctl    | Shortlink CLI                        | Go                 | [docs](./docs/shortctl.md) |

### Run

```
make run
```

##### Kuberentes

```
make minikube-up
make helm-shortlink-up
```

##### DNS/HTTP

services:
  + `ui-nuxt.local`
  + `shortlink.local`

Add `127.0.0.1 [name services].local` to your `/etc/hosts`

### HTTP API

![arhitecture.json](./docs/arhitecture.png)

+ Import [Postman link](./docs/shortlink.postman_collection.json) for
  test HTTP API
+ Swagger [docs](https://shortlink-org.gitlab.io/shortlink)

###### Support HTTP REST API:

- HTTP (chi)
- gRPC-gateway
- GraphQL
- ***Optional***
    - go-kit
    - [CloudEvents](https://cloudevents.io/)

### MQ

+ [Kafka](https://kafka.apache.org/)
+ NATS
+ RabbitMQ

### Store provider

+ RAM
+ MongoDB
+ Postgres
+ ***Optional***
    + Redis
    + MySQL
    + LevelDB
    + Badger
    + SQLite
    + Scylla
    + Сassandra (via: Scylla driver)
    + RethinkDB
    + DGraph

### Cloud-Native

+ Logger
  + [Loki](./docs/logger.md)
+ Prometheus
  + prometheus-operator
    + notify: slack, email, telegram
+ HealthCheck
+ Support K8S
  + Helm Chart
    + [pingcap/chaos-meshh](https://github.com/pingcap/chaos-mesh)
  + Minikube
  + Backup/Restore ([Velero](https://velero.io/)
  + Custom CSI driver (fork [csi-driver-host-pat](https://github.com/kubernetes-csi/csi-driver-host-path))
+ Istio
+ MetalLB

### Gateway

+ Traefik
+ Nginx

### UI service

+ Nuxt: [demo UI](http://shortlink.surge.sh/)
+ Next

| Service     | Description                       | Language/Framework |
|-------------|-----------------------------------|--------------------|
| next        | UI service                        | JS/ReactJS         |
| nuxt        | UI service                        | JS/VueJS           |

##### ENV for UI

Use `.env` file in `pkg/ui/[nuxt/next/etc]` directories for setting your UI


| Name                | Default                                                     | Description                                                                                    |
|:--------------------|:------------------------------------------------------------|:-----------------------------------------------------------------------------------------------|
| NODE_ENV            | -                                                           | Select: production, development, etc...                                                        |
| SENTRY_DSN          | -                                                           | Your sentry DSN                                                                                |
| API_URL_HTTP        | http://localhost:7070                                       | HTTP API Endpoint                                                                              |

#### UI Screenshot

| Describe                | Screenshot                           |
|-------------------------|--------------------------------------|
| Link Table              | ![link table](./docs/next-js-ui.png) |

### Configuration

<details><summary>DETAILS</summary>
<p>

##### [12 factors: ENV](https://12factor.net/config)

[View ENV Variables](./docs/env.md)

</p>
</details>

### CoreDNS IP table

| Service | Ip address | Description                                    |
|:--------|:-----------|:-----------------------------------------------|
| store   | 10.5.0.100 | Main database (postgres/mongo/cassandra/redis) |

##### troubleshooting

Sometimes a container without a specified ip may occupy a binding
address of another service, which will result in `Address already in
use`.

### Ansible

<details><summary>DETAILS</summary>
<p>

##### Vagrant

```
cd ops/vagrant
vagrant up

cd ops/ansible
ansible-playbook playbooks/playbook.yml
```

##### DNS/HTTP

+ `ui-nuxt.shortlink.vagrant:8081`

</p>
</details>

### Kubernetes

<details><summary>DETAILS</summary>
<p>

##### HELM

+ **common** - run common tools (ingress)
+ **shortlink-\*** - run shortlink applications (api, logger, ui)
+ **chaos** - run chaos daemon
+ **ingress** - run ingress ;-)

##### DNS

+ `ui-nuxt.local`
+ `grafana.local`
+ `jaeger.local`
+ `prometheus.local`

</p>
</details>

### GITLAB CI

![](./docs/gitlab-pipeline.png)

![](./docs/gitlab-ci.png)

##### GitLab Variable

- SURGE_LOGIN
- SURGE_TOKEN
- DANGER_GITLAB_API_TOKEN - `API TOKEN` for danger

##### Support environment

- Minikube
- [Yandex Cloud](https://cloud.yandex.ru/)
- [AWS EKS](https://aws.amazon.com/eks/)

### GitHub CI

- Travis action (backend/frontend)
- Docker action
- Helm chart

### [Error code](./internal/error/status/exit.go)

| CODE | Describe     |
|------|--------------|
| 0    | SUCCESS      |
| 1    | ERROR_CONFIG |

## -~- THE END -~-

[mergify]: https://mergify.io
[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/batazor/shortlink&style=flat
