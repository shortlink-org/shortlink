# shortlink

Shortlink service

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

+ [Postman link](./docs/shortlink.postman_collection.json)

- GET /:URL #=> 200 or 404
- POST / {"url":"example.com"} #=> 201
- DELETE / {"url":"example.com"} #=> 200
- GET /s/:URL #=> 301

### Store provider

+ RAM
+ Redis
+ MongoDB
+ Postgres
+ DGraph
