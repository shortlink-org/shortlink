Feature: SQL Parser
  We should be able to parse the SQL expression

  Scenario Outline: We got an result on an SQL query
    Given <SQL> expression
    When we get the query
    Then the response <result>

    Examples:
      | SQL                                                                   | result                                                  |
      | ""                                                                    | query type cannot be empty                              |
      | SELECT                                                                | table name cannot be empty                              |
      | SELECT FROM 'a'                                                       | at SELECT: expected field to SELECT                     |
      | SELECT b, FROM 'a'                                                    | at SELECT: expected field to SELECT                     |
      | SELECT a FROM 'b'                                                     | nil                                                     |
      | select a fRoM 'c'                                                     | nil                                                     |
      | SELECT a, c, d FROM 'b10'                                             | nil                                                     |
      | SELECT a as z, b as y, c FROM 'b0'                                    | nil                                                     |
      | SELECT a, c, d FROM 'b1' WHERE                                        | at WHERE: empty WHERE clause                            |
      | SELECT a, c, d FROM 'b2' WHERE a                                      | at WHERE: condition without operator                    |
      | SELECT a, c, d FROM 'b3' WHERE a = ''                                 | nil                                                     |
      | SELECT a, c, d FROM 'b4' WHERE a < '1'                                | nil                                                     |
      | SELECT a, c, d FROM 'b5' WHERE a <= '1'                               | nil                                                     |
      | SELECT a, c, d FROM 'b6' WHERE a > '1'                                | nil                                                     |
      | SELECT a, c, d FROM 'b7' WHERE a >= '1'                               | nil                                                     |
      | SELECT a, c, d FROM 'b8' WHERE a != '1'                               | nil                                                     |
      | SELECT a, c, d FROM 'b9' WHERE a != b                                 | nil                                                     |
      | SELECT * FROM 'b'                                                     | nil                                                     |
      | SELECT a, * FROM 'b'                                                  | nil                                                     |
      | SELECT a, c, d FROM 'b' WHERE a != '1' AND b = '2'                    | nil                                                     |
      | UPDATE                                                                | table name cannot be empty                              |
      | UPDATE 'a1'                                                           | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a2' SET                                                       | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | UPDATE 'a3' SET b WHERE                                               | at UPDATE: expected '='                                 |
      | UPDATE 'a4' SET b = WHERE                                             | at UPDATE: expected quoted value                        |
      | UPDATE 'a5' SET b = 'hello' WHERE                                     | at WHERE: empty WHERE clause                            |
      | UPDATE 'a6' SET b = 'hello' WHERE a                                   | at WHERE: condition without operator                    |
      | UPDATE 'a7' SET b = 'hello' WHERE a = '1'                             | nil                                                     |
      | UPDATE 'a8' SET b = 'hello\\'world' WHERE a = '1'                     | nil                                                     |
      | UPDATE 'a9' SET b = 'hello', c = 'bye' WHERE a = '1'                  | nil                                                     |
      | UPDATE 'a0' SET b = 'hello', c = 'bye' WHERE a = '1' AND b = '789'    | nil                                                     |
      | DELETE FROM                                                           | table name cannot be empty                              |
      | DELETE FROM 'a1'                                                      | at WHERE: WHERE clause is mandatory for UPDATE & DELETE |
      | DELETE FROM 'a2' WHERE                                                | at WHERE: empty WHERE clause                            |
      | DELETE FROM 'a3' WHERE b                                              | at WHERE: condition without operator                    |
      | DELETE FROM 'a4' WHERE b = '1'                                        | nil                                                     |
      | INSERT INTO                                                           | table name cannot be empty                              |
      | INSERT INTO 'a1'                                                      | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a2' (                                                    | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a3' (b                                                   | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a4' (b)                                                  | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a5' (b) VALUES                                           | at INSERT INTO: need at least one row to insert         |
      | INSERT INTO 'a6' (b) VALUES (                                         | at INSERT INTO: value count doesn't match field count   |
      | INSERT INTO 'a7' (b) VALUES ('1')                                     | nil                                                     |
      | INSERT INTO 'a8' (*) VALUES ('1')                                     | at INSERT INTO: expected at least one field to insert   |
      | INSERT INTO 'a9' (b,c,d) VALUES ('1','2','3')                         | nil                                                     |
      | INSERT INTO 'a0' (b,c,d) VALUES ('1','2' ,'3'),('4','5','6')          | nil                                                     |
