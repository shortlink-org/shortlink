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
#      | SELECT a, c, d FROM 'b' WHERE a <= '1'                                |  7                                                       |
      | SELECT a, c, d FROM 'b' WHERE a > '1'                                 |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a >= '1'                                |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a != '1'                                |                                                         |
      | SELECT a, c, d FROM 'b' WHERE a != b                                  |                                                         |
#      | SELECT * FROM 'b'                                                     |   12                                                      |
#      | SELECT a, * FROM 'b'                                                  |    13                                                     |
      | SELECT a, c, d FROM 'b' WHERE a != '1' AND b = '2'                    |                                                         |
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
#      | DELETE FROM 'a' WHERE                                                 | at WHERE: empty WHERE clause                            |
#      | DELETE FROM 'a' WHERE b                                               | at WHERE: condition without operator                    |
#      | DELETE FROM 'a' WHERE b = '1'                                         |       19                                                  |
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
