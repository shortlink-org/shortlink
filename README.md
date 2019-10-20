# shortlink

[![go-doc](https://godoc.org/github.com/batazor/shortlink?status.svg)](https://godoc.org/github.com/batazor/shortlink)
[![codecov](https://codecov.io/gh/batazor/shortlink/branch/master/graph/badge.svg)](https://codecov.io/gh/batazor/shortlink)
[![Build Status](https://travis-ci.org/batazor/shortlink.svg?branch=master)](https://travis-ci.org/batazor/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/batazor/shortlink)](https://goreportcard.com/report/github.com/batazor/shortlink)
[![Releases](https://img.shields.io/github/release-pre/batazor/shortlink.svg)](https://github.com/batazor/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/batazor/shortlink.svg)](https://github.com/batazor/shortlink/blob/master/LICENSE)

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
    up -d 
```

##### As example

```
docker-compose \
    -f docker-compose.yaml \
    -f ops/docker-compose/database/dgraph.yaml \
    up -d 
```

### HTTP API

+ Import [Postman link](./docs/shortlink.postman_collection.json) for test HTTP API

- GET /:URL #=> 200 or 404
- POST / {"url":"example.com"} #=> 201
- DELETE / {"url":"example.com"} #=> 200
- GET /s/:URL #=> 301

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
