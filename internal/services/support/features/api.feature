Feature: API service support
  As a user, I want to be able to send a request for support through the API
  and get a list of my previous support requests.

  Scenario: Send support request
    Given I am authenticated with the API
    When I send a POST request to "/support" with the following data:
      | title    | email         | text        | files      |
      | "Help!"  | "user@email"  | "I need help with something"  | "/path/to/file1.txt,/path/to/file2.pdf" |
    Then the response code should be 200
    And the response should contain the following JSON:
      | request_id | title | email         | text        | files      |
      | 1          | "Help!"  | "user@email"  | "I need help with something"  | "/path/to/file1.txt,/path/to/file2.pdf" |

  Scenario: List support requests
    Given I am authenticated with the API
    When I send a GET request to "/support"
    Then the response code should be 200
    And the response should contain a list of previous support requests in the following format:
      | request_id | title | email         | text        | files      |
      | 1          | "Help!"  | "user@email"  | "I need help with something"  | "/path/to/file1.txt,/path/to/file2.pdf" |
      | 2          | "Another issue"  | "user@email"  | "I am having another issue"  | "/path/to/file3.txt" |
