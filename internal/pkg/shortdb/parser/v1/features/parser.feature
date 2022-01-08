Feature: SQL Parser
  We should be able to parse the SQL expression

  Scenario Outline: We got an result on an SQL query
    Given "<SQL>" expression
    Then the response "<result>"

    Examples:
      | SQL                                                                   | result                                                  |
      |                                                                       | query type cannot be empty                              |
      | SELECT                                                                | table name cannot be empty                              |
      | SELECT FROM 'a'                                                       | at SELECT: expected field to SELECT                     |
      | SELECT b, FROM 'a'                                                    | at SELECT: expected field to SELECT                     |
      | SELECT a FROM 'b'                                                     |                                                         |
      | select a fRoM 'c'                                                     |                                                         |
      | SELECT a, c, d FROM 'b10'                                             |                                                         |
      | SELECT a as z, b as y, c FROM 'b'                                     |                                                         |
      | SELECT a, c, d FROM 'b' WHERE                                         | at WHERE: empty WHERE clause                            |
      | SELECT a, c, d FROM 'b' WHERE a                                       | at WHERE: condition without operator                    |
      | SELECT a, c, d FROM 'b' WHERE a = ''                                  |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a < '1'                                 |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a <= '1'                                | at WHERE: expected quoted value                         |
      | SELECT a, c, d FROM 'b' WHERE a > '1'                                 |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a >= '1'                                |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a != '1'                                |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a != b                                  |                                                         |
      | SELECT * FROM 'b'                                                     | at SELECT: expected field to SELECT                     |
      | SELECT a, * FROM 'b'                                                  | at SELECT: expected field to SELECT                     |
      | SELECT a, c, d FROM 'b' WHERE a != '1' AND b = '2'                    |                                                         |
      | SELECT start as s, middle as m, end as e FROM there join what on there.it != what.what and there.who = what.shit left join whoot on whoot.tweet <= what.what where this = that order by start desc, end, middle asc | at ON: expected <tablename>.<fieldname> |
      | UPDATE                                                                | table name cannot be empty                              |
      | UPDATE 'a'                                                            | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a' SET                                                        | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a' SET b WHERE                                                | at UPDATE: expected '='                                 |
      | UPDATE 'a' SET b = WHERE                                              | at UPDATE: expected quoted value                        |
      | UPDATE 'a' SET b = 'hello' WHERE                                      | at WHERE: empty WHERE clause                            |
      | UPDATE 'a' SET b = 'hello' WHERE a                                    | at WHERE: condition without operator                    |
      | UPDATE 'a' SET b = 'hello' WHERE a = '1'                              |                                                         |
      | UPDATE 'a' SET b = 'hello\\'world' WHERE a = '1'                      |                                                         |
      | UPDATE 'a' SET b = 'hello', c = 'bye' WHERE a = '1'                   |                                                         |
      | UPDATE 'a' SET b = 'hello', c = 'bye' WHERE a = '1' AND b = '789'     |                                                         |
      | DELETE FROM                                                           | table name cannot be empty                              |
      | DELETE FROM 'a'                                                       | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | DELETE FROM 'a' WHERE                                                 | at UPDATE: expected 'SET'                               |
      | DELETE FROM 'a' WHERE b                                               | at UPDATE: expected 'SET'                               |
      | DELETE FROM 'a' WHERE b = '1'                                         | at UPDATE: expected 'SET'                               |
      | INSERT INTO                                                           | table name cannot be empty                              |
      | INSERT INTO 'a'                                                       | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (                                                     | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b                                                    | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b)                                                   | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b) VALUES                                            | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b) VALUES (                                          | at INSERT INTO: value count doesn't match field count   |
      | INSERT INTO 'a' (b) VALUES ('1')                                      |                                                         |
      | INSERT INTO 'a' (*) VALUES ('1')                                      | at INSERT INTO: expected at least one field to insert   |
      | INSERT INTO 'a' (b,c,d) VALUES ('1','2','3')                          |                                                         |
      | INSERT INTO 'a' (b,c,d) VALUES ('1','2','3'),('4','5','6')            |                                                         |
      | 123                                                                   | incorrect sql-expression                                |
      | CREATE TABLE;                                                         | at CREATE TABLE: table name cannot be empty             |
      | CREATE TABLE users;                                                   | at CREATE TABLE: expected opening parens                |
      | CREATE TABLE users (;                                                 | at CREATE TABLE: expected at least one field to create table |
      | CREATE TABLE users ( id nontype );                                    | at CREATE TABLE: unsupported type of field              |
      | CREATE TABLE users ( id integer );                                    |                                                         |
      | CREATE TABLE users ( id integer, );                                   | at CREATE TABLE: expected at least one field to create table |
      | CREATE TABLE users ( id integer,;                                     | at CREATE TABLE: expected at least one field to create table |
      | CREATE TABLE users ( id integer, name text );                         |                                                         |
      | DROP TABLE                                                            | table name cannot be empty                              |
      | DROP TABLE users                                                      |                                                         |
