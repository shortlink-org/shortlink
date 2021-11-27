Feature: SQL Parser
  We should be able to parse the SQL expression

  Scenario Outline: We got an error on an empty query
    Given Get empty query
    When we want to parse an SQL expression
    Then we get an error message

    Examples:
      | query                        | response                    |
      | ""                           | query type cannot be empty   |

  Scenario Outline: We got an error on select empty table name
    Given Get empty table name
    When we want to parse an SQL expression
    Then we get an error message

    Examples:
      | query                             | response                   |
      | SELECT without FROM fails         | table name cannot be empty |
