# BENCHMARK

### Tool

- [Percona-Lab/sysbench-tpcc](https://github.com/Percona-Lab/sysbench-tpcc)

#### 1. PostgreSQL: prepare data and tables

```shell
./tpcc.lua \
  --pgsql-host=localhost \
  --pgsql-port=5432 \
  --pgsql-user=postgres \
  --pgsql-password=shortlink \
  --pgsql-db=shortlink \
  --threads=2 \
  --tables=10 \
  --scale=100 \
  --trx_level=RC \
  --db-ps-mode=auto \
  --db-driver=pgsql \
  prepare
```

#### 2. PostgreSQL: Run benchmark

```shell
./tpcc.lua \
  --pgsql-host=localhost \
  --pgsql-port=5432 \
  --pgsql-user=postgres \
  --pgsql-password=shortlink \
  --pgsql-db=shortlink \
  --threads=2 \
  --tables=10 \
  --scale=100 \
  --trx_level=RC \
  --db-ps-mode=auto \
  --db-driver=pgsql \
  --time=3000 \
  --report-interval=1 \
  run
```

#### 3. PostgreSQL: Cleanup

```shell
./tpcc.lua \
  --pgsql-host=localhost \
  --pgsql-port=5432 \
  --pgsql-user=postgres \
  --pgsql-password=shortlink \
  --pgsql-db=shortlink \
  --threads=2 \
  --tables=10 \
  --scale=100 \
  --trx_level=RC \
  --db-driver=pgsql \
  cleanup
```
