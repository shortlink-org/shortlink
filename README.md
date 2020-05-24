# shortlink

[![go-doc](https://godoc.org/github.com/batazor/shortlink?status.svg)](https://godoc.org/github.com/batazor/shortlink)
[![codecov](https://codecov.io/gh/batazor/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)
[![Releases](https://img.shields.io/github/release-pre/batazor/shortlink.svg)](https://github.com/batazor/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/batazor/shortlink.svg)](https://github.com/batazor/shortlink/blob/master/LICENSE)
![GitHub last commit](https://img.shields.io/github/last-commit/batazor/shortlink)
![GitHub contributors](https://img.shields.io/github/contributors/batazor/shortlink)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)
[![Mergify Status][mergify-status]][mergify]

Shortlink service

### High Level Architecture

![shortlink-arhitecture](./docs/shortlink-arhitecture.png)

##### Requirements

- docker
- docker-compose
- protoc 3.7.1+

##### Services

| Service     | Description                       | Language/Framework |
|-------------|-----------------------------------|--------------------|
| shortlink   | Shortlink service                 | Go                 |
| logger      | Logger service                    | Go                 |

### Run

```
make run
```

##### DNS/HTTP

+ `ui-nuxt.local`
+ `shortlink.local`

Add `127.0.0.1 *.local` to your `/etc/hosts`

### HTTP API

![arhitecture.json](./docs/arhitecture.png)

+ Import [Postman link](./docs/shortlink.postman_collection.json) for
  test HTTP API
+ Swagger [docs](https://shortlink1.gitlab.io/shortlink)

###### Support HTTP REST API:

- HTTP (chi)
- gRPC-gateway
- go-kit
- GraphQL
- [CloudEvents](https://cloudevents.io/)

### MQ

+ [Kafka](https://kafka.apache.org/)

### Store provider

+ RAM
+ Redis
+ MongoDB
+ MySQL
+ Postgres
+ DGraph
+ LevelDB
+ Badger
+ SQLite
+ Scylla
+ Сassandra (via: Scylla driver)

### Cloud-Native

+ Prometheus
  + prometheus-operator
+ HealthCheck
+ Support K8S
  + Helm Chart
    + [pingcap/chaos-meshh](https://github.com/pingcap/chaos-mesh)
  + Minikube
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

| Name                | Default                                                     | Description                                                                                    |
|:--------------------|:------------------------------------------------------------|:-----------------------------------------------------------------------------------------------|
| STORE_TYPE          | ram                                                         | Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra |
| STORE_MONGODB_URI   | mongodb://localhost:27017                                   | MongoDB URI                                                                                    |
| STORE_MYSQL_URI     | shortlink:shortlink@localhost:3306/shortlink?parseTime=true | MySQL URI                                                                                      |
| STORE_BADGER_PATH   | /tmp/links.badger                                           | Badger path to file                                                                            |
| STORE_DGRAPH_URI    | localhost:9080                                              | DGRAPH link                                                                                    |
| STORE_LEVELDB_PATH  | /tmp/links.db                                               | LevelDB path to file                                                                           |
| STORE_POSTGRES_URI  | postgres://shortlink:shortlink@localhost:5432/shortlink     | Postgres URI                                                                                   |
| STORE_REDIS_URI     | localhost:6379                                              | Redis URI                                                                                      |
| STORE_SQLITE_PATH   | /tmp/links.sqlite                                           | SQLite URI                                                                                     |
| STORE_CASSANDRA_URI | localhost:9042                                              | Cassandra URI                                                                                  |
| STORE_SCYLLA_URI    | localhost:9042                                              | Scylla URI                                                                                     |
| LOG_LEVEL           | 3                                                           | Log level. Select 0-4 (Fatal->Debug)                                                           |
| LOG_TIME_FORMAT     | 2006-01-02T15:04:05.999999999Z07:00                         | Log time format (golang time format)                                                           |
| TRACER_SERVICE_NAME | ShortLink                                                   | Service Name                                                                                   |
| TRACER_URI          | localhost:6831                                              | Tracing addr:host                                                                              |
| API_TYPE            | http-chi                                                    | Select: http-chi, gRPC-web, graphql, cloudevents, go-kit                                       |
| API_PORT            | 7070                                                        | API port                                                                                       |
| API_TIMEOUT         | 60                                                          | Request Timeout                                                                                |
| MQ_ENABLED          | false                                                       | Enabled MQ-service                                                                             |
| MQ_TYPE             | kafka                                                       | Select: kafka, nats                                                                            |
| MQ_KAFKA_URI        | localhost:9092                                              | Kafka URI                                                                                      |
| SENTRY_DSN          | ___DSN___                                                   | key for sentry                                                                                 |


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

+ `ui-nuxt.shortlink.minikube`
+ `api.shortlink.minikube`
+ `grafana.minikube`
+ `jaeger.minikube`
+ `prometheus.minikube`

</p>
</details>

### GITLAB CI

![](./docs/gitlab-pipeline.png)

##### GitLab Variable

- SURGE_LOGIN
- SURGE_TOKEN
- DANGER_GITLAB_API_TOKEN - `API TOKEN` for danger

##### Support environment

- Minikube
- [Yandex Cloud](https://cloud.yandex.ru/)
- [AWS EKS](https://aws.amazon.com/eks/)

## -~- THE END -~-

[mergify]: https://mergify.io
[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/batazor/shortlink&style=flat
