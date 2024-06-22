## X/Open XA

### Advantages and Disadvantages

Compared to other patterns like SAGA and TCC, the advantages of XA global transactions are:

+ Simple and easy to understand
+ Automatic rollback of business, no need to write compensation manually

The disadvantages of XA are:

+ Need the XA transaction from an underlying database
+ Data is locked from data modification until the commitment, much longer than other patterns. It is not suitable for highly concurrent business.

### Docs

- [GitHub: X/Open XA](https://github.com/topics/xa)
  - [Apache Seata(incubating)](https://seata.apache.org/)
  - [DTM](https://en.dtm.pub/)
- DataBase
  - [Postgres: PREPARE TRANSACTION](https://postgrespro.ru/docs/postgresql/16/sql-prepare-transaction?lang=ru-en)
- Blogs
  - [Understanding XA Transactions With Practical Examples in Go](https://betterprogramming.pub/understanding-xa-transactions-with-practical-examples-in-go-67e99fd333db)
  - [dtm blog](https://medium.com/@dongfuye)
  - [Распределенные транзакции (XA) с помощью JTA в JavaSE (на примере Spring + Atomikos)](https://samolisov.blogspot.com/2011/02/xa-jta-javase-spring-atomikos-2.html)

### Examples

- [incubator-seata-go-samples](https://github.com/apache/incubator-seata-go-samples)
- [dtm-example](./dtm-example/README.md)
