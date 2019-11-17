# shortlink

[![go-doc](https://godoc.org/github.com/batazor/shortlink?status.svg)](https://godoc.org/github.com/batazor/shortlink)
[![codecov](https://codecov.io/gh/batazor/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/batazor/shortlink)
[![Build Status](https://travis-ci.org/batazor/shortlink.svg?branch=master)](https://travis-ci.org/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)
[![Releases](https://img.shields.io/github/release-pre/batazor/shortlink.svg)](https://github.com/batazor/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/batazor/shortlink.svg)](https://github.com/batazor/shortlink/blob/master/LICENSE)
![GitHub last commit](https://img.shields.io/github/last-commit/batazor/shortlink)
![GitHub contributors](https://img.shields.io/github/contributors/batazor/shortlink)

Shortlink service

### High Level Architecture

![shortlink-arhitecture](./docs/shortlink-arhitecture.png)

##### Install dependencies

```
go get -u moul.io/protoc-gen-gotemplate
go get -u github.com/jteeuwen/go-bindata

make
```

### Run

```
docker-compose \
    -f docker-compose.yaml \
    -f ops/docker-compose/database/[DATABASE_NAME] \
    -f ops/docker-compose/gateway/[GATEWAY_NAME] \
    up -d 
```

##### As example

```
docker-compose \
    -f docker-compose.yaml \
    -f ops/docker-compose/database/postgres.yaml \
    -f ops/docker-compose/gateway/traefik.yaml \
    -f ops/docker-compose/tooling/opentracing.yaml \
    up -d 
```

### HTTP API

+ Import [Postman link](./docs/shortlink.postman_collection.json) for test HTTP API

###### Support HTTP REST API:

- HTTP (chi)
- gRPC-gateway
- GraphQL

### Store provider

+ RAM
+ Redis
+ MongoDB
+ Postgres
+ DGraph
+ LevelDB
+ Badger
+ SQLite

### Cloud-Native

+ Prometheus
+ HealthCheck
+ Support K8S (Helm Chart)
