## X/Open XA

### Advantages and Disadvantages

Compared to other patterns like SAGA and TCC, the advantages of XA global transactions are:

+ Simple and easy to understand
+ Automatic rollback of business, no need to write compensation manually

The disadvantages of XA are:

+ Need the XA transaction from an underlying database
+ Data is locked from data modification until the commitment, much longer than other patterns. It is not suitable for highly concurrent business.

### Docs

- [DTM](https://en.dtm.pub/)
- [Understanding XA Transactions With Practical Examples in Go](https://betterprogramming.pub/understanding-xa-transactions-with-practical-examples-in-go-67e99fd333db)
- [Postgres: PREPARE TRANSACTION](https://postgrespro.ru/docs/postgresql/16/sql-prepare-transaction?lang=ru-en)
