# Link services

Service for work with link-domain (CRUD)

### Store provider

![URI FORMAT](./docs/URI_FORMAT.png)

> support - enabled batch mode; filter, etc...  
> scale - scalability/single mode

| Name                            | Support   | Scale    |
|---------------------------------|-----------|----------|
| RAM                             | ✅         | ❌       |
| MongoDB                         | ✅         | ✅       |
| Postgres                        | ✅         | ✅       |
| Redis                           | ❌         | ✅       |
| MySQL                           | ❌         | ✅       |
| LevelDB                         | ❌         | ❌       |
| Badger                          | ❌         | ❌       |
| SQLite                          | ❌         | ❌       |
| Scylla                          | ❌         | ✅       |
| Сassandra (via: Scylla driver)  | ❌         | ✅       |
| RethinkDB                       | ❌         | ✅       |
| DGraph                          | ❌         | ✅       |

### Example request

##### Add new link
```
grpcurl -cacert ./ops/cert/intermediate_ca.pem -d '{"url": "http://google.com"}' localhost:50052 link_rpc.Link/Add
```
