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
      | SELECT a FROM 'b'                                                     | nil                                                     |
      | select a fRoM 'c'                                                     | nil                                                     |
      | SELECT a, c, d FROM 'b10'                                             | nil                                                     |
      | SELECT a as z, b as y, c FROM 'b'                                     | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE                                         | at WHERE: empty WHERE clause                            |
      | SELECT a, c, d FROM 'b' WHERE a                                       | at WHERE: condition without operator                    |
      | SELECT a, c, d FROM 'b' WHERE a = ''                                  | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a < '1'                                 | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a <= '1'                                | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a > '1'                                 | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a >= '1'                                | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a != '1'                                | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a != b                                  | nil                                                     |
      | SELECT * FROM 'b'                                                     | nil                                                     |
      | SELECT a, * FROM 'b'                                                  | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a != '1' AND b = '2'                    | nil                                                     |
      | UPDATE                                                                | table name cannot be empty                              |
      | UPDATE 'a'                                                            | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a' SET                                                        | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a' SET b WHERE                                                | at UPDATE: expected '='                                 |
      | UPDATE 'a' SET b = WHERE                                              | at UPDATE: expected quoted value                        |
      | UPDATE 'a' SET b = 'hello' WHERE                                      | at WHERE: empty WHERE clause                            |
      | UPDATE 'a' SET b = 'hello' WHERE a                                    | at WHERE: condition without operator                    |
      | UPDATE 'a' SET b = 'hello' WHERE a = '1'                              | nil                                                     |
      | UPDATE 'a' SET b = 'hello\\'world' WHERE a = '1'                      | nil                                                     |
      | UPDATE 'a' SET b = 'hello', c = 'bye' WHERE a = '1'                   | nil                                                     |
      | UPDATE 'a' SET b = 'hello', c = 'bye' WHERE a = '1' AND b = '789'     | nil                                                     |
      | DELETE FROM                                                           | table name cannot be empty                              |
      | DELETE FROM 'a'                                                       | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | DELETE FROM 'a' WHERE                                                 | at WHERE: empty WHERE clause                            |
      | DELETE FROM 'a' WHERE b                                               | at WHERE: condition without operator                    |
      | DELETE FROM 'a' WHERE b = '1'                                         | nil                                                     |
      | INSERT INTO                                                           | table name cannot be empty                              |
      | INSERT INTO 'a'                                                       | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (                                                     | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b                                                    | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b)                                                   | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b) VALUES                                            | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a' (b) VALUES (                                          | at INSERT INTO: value count doesn't match field count   |
      | INSERT INTO 'a' (b) VALUES ('1')                                      | nil                                                     |
      | INSERT INTO 'a' (*) VALUES ('1')                                      | at INSERT INTO: expected at least one field to insert   |
      | INSERT INTO 'a' (b,c,d) VALUES ('1','2','3')                          | nil                                                     |
      | INSERT INTO 'a' (b,c,d) VALUES ('1','2','3'),('4','5','6')            | nil                                                     |
