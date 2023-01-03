# 4. Use UUID as Primary Keys

Date: 2022-12-10

## Status

Accepted

## Context

We need some identifier for entities that will be unique throughout the service.
A more detailed description of the problem can be found in the issue [#762](https://github.com/shortlink-org/shortlink/issues/762)

##### Docs

- [UUID or GUID as Primary Keys? Be Careful!](https://tomharrisonjr.com/uuid-or-guid-as-primary-keys-be-careful-7b2aa3dcb439)
- [Primary Keys: IDs versus GUIDs](https://blog.codinghorror.com/primary-keys-ids-versus-guids/)
- [UUID vs. Sequences](https://blog.josephscott.org/2005/07/22/uuid-vs-sequences/)

- **MySQL**
    - [MySQL 8.0.13: Default Value as uuid](https://stackoverflow.com/questions/60462208/mysql-8-0-13-default-value-as-uuid-not-working)
- **PostgreSQL**
    - [Default value for UUID column in Postgres](https://dba.stackexchange.com/questions/122623)

## Decision

We will use UUID as primary keys

## Consequences

- We will have a unique identifier for each entity
