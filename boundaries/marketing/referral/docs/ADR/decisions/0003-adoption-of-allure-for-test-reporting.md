# 3. Adoption of Allure for Test Reporting

Date: 2023-09-12

## Status

Accepted

## Context

Our project needs a comprehensive test reporting tool to convey the reliability of our codebase. 
Allure has been evaluated for this purpose.

## Decision


| 👍 Pros                                               | 👎 Cons                                                               |
|-------------------------------------------------------|-----------------------------------------------------------------------|
| Allure provides visually appealing reports.           | Requires additional setup and possibly infrastructure for hosting.    |
| Compatibility with various testing frameworks.        | Team might need training to maximize the use of the detailed reports. |
| Supports historical data tracking for trend analysis. |                                                                       |

Recognizing the need for a comprehensive test reporting tool offering both detailed insights and visual appeal, 
we've chosen to integrate Allure into our testing workflow.

## Consequences

* Post-test run, Allure reports will be generated and should be reviewed for test results insights and potential anomalies.
* Depending on the hosting choice for Allure reports, infrastructure decisions might be necessary.
