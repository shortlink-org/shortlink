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

#### Example work with repl

```
> create table users (id integer, name string, active bool);
> insert into users ('id', 'name', 'active') VALUES ('1', 'Ivan', 'false');
> select id, name, active from users;
```

#### Reference

- Parser
  - [Let's build a SQL parser in Go!](https://marianogappa.github.io/software/2019/06/05/lets-build-a-sql-parser-in-go/)
  - [Simple SQL parser meant for querying CSV files on golang](https://github.com/marianogappa/sqlparser) 
- Database
  - [Let's Build a Simple Database](https://cstack.github.io/db_tutorial/)
