## Database `ShortDB`

This database with support SQL language.

#### Docs

- [Architectura](https://miro.com/app/board/uXjVOaNZlsE=/?invite_link_id=250172061542)
- [Roadmaps](./ROADMAP.md)
- [CLI](./docs/dbctl.md)

#### Feature

- REPL interface
  - save history (by default last 100 commands)
  - navigation by history commands
- SQL-parser
  - SELECT, INSERT INTO, CREATE TABLE, etc...
- Engine:
  - file
- API
  - gRPC
  - [custom protocol](./protocol/README.md)

#### Example work with REPL

```
> create table users (id integer, name string, active bool);
> insert into users ('id', 'name', 'active') VALUES ('1', 'Ivan', 'false');
> select id, name, active from users;
```

[![asciicast](https://asciinema.org/a/ElqLr756zSjpwFCuAQgSbXxBB.svg)](https://asciinema.org/a/ElqLr756zSjpwFCuAQgSbXxBB)

#### Docker build

```bash
$> docker buildx build --platform=linux/amd64,linux/arm64 --load -t shortdb -f ops/dockerfile/shortdb.Dockerfile .
```

#### Reference

- Parser
  - [Let's build a SQL parser in Go!](https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/)
  - [Simple SQL parser meant for querying CSV files on golang](https://github.com/marianogappa/sqlparser) 
  - [LL(1) parser](https://en.wikipedia.org/wiki/LL_parser)
- Database
  - [Let's Build a Simple Database](https://cstack.github.io/db_tutorial/)

### Benchmark Engine

> cpu: Intel(R) Core(TM) i3-7300 CPU @ 4.00GHz

| Name                                        | Count |   ns/op |
|:--------------------------------------------|------:|--------:|
| **CREATE_TABLE**                            |       |         |
| CREATE_TABLE-4                              |  8199 |  143758 |
| **INSERT INTO**                             |       |         |
| INSERT_INTO_USERS-4                         | 14222 |   83524 |
| **SELECT USERS**                            |       |         |
| SELECT_USERS-4                              | 13066 |   91571 |
| SELECT_USERS_WITH_WHERE_id=99_AND_LIMIT_2-4 |  6096 |  210703 |
| SELECT_USERS_FULL_SCAN-4                    |   202 | 5813380 |

### Benchmark Parser

> cpu: Intel(R) Core(TM) i3-7300 CPU @ 4.00GHz

| Name             | Count |  ns/op |
|:-----------------|------:|-------:|
| **CREATE_TABLE** |       |        |
| CREATE_TABLE-4   | 10000 | 105080 |
| **SELECT**       |       |        |
| SELECT-4         | 11912 |  90804 |
| **INSERT INTO**  |       |        |
| INSERT_INTO-4    | 16987 |  70963 |
