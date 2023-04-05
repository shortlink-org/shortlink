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
| LevelDB                         | ❌         | ❌       |
| Badger                          | ❌         | ❌       |
| SQLite                          | ❌         | ❌       |
| DGraph                          | ❌         | ✅       |

### Example request

We support reflection for request. You can use [Postman](https://www.postman.com/) or [grpcurl](https://github.com/fullstorydev/grpcurl) for test.

![postman](https://blog.postman.com/wp-content/uploads/2022/01/grpc-author-msg.gif)

### Changelog

- [19.09.2022] Drop support database: MySQL
- [04.08.2021] Drop support database: scylla, cassandra
