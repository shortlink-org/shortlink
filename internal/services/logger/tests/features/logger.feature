Feature: Print message
  We should be able to output the message

  Scenario Outline: We got the correct payload link URL
    Given Get correct payload with not nil field URL <payload>$
    When get new a event from MQ
    Then print link URL <url>$

    Examples:
      | payload                           | url                |
      | `{"url":"https://google.com"}`    | https://google.com |

  Scenario Outline: We got the incorrect payload link URL
    Given Get random payload <payload>$
    When get new a event from MQ
    Then print error message: `Incorrect format payload` <url>$

    Examples:
      | payload                           | url                      |
      | `{"url":""}`                      | error message            |
      | `123123`                          | error message            |
