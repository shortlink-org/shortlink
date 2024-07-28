# 25. Application Configuration Redesign

Date: 2023-05-25

## Status

Accepted

## Context

The present method for configuring our application is through the use of environment variables. 
Although this is a prevalent practice within the industry and generally provides an effective means 
for application configuration, we believe that certain aspects could be improved for better maintainability 
and ease of use.

## Decision

We have decided to follow the "Configuration" principle outlined in the [12-Factor App methodology](https://12factor.net/config). 
This approach emphasizes the importance of storing configuration that varies between deployments in the environment, 
rather than in the codebase.

### Documentation

We can use tools like [shortctl](../../../boundaries/platform/shortctl) to manage and document the configuration changes.

## Consequences

Adopting this change will make it easier to manage and maintain application configurations across different environments, 
reducing potential errors caused by inconsistent configurations.

However, we must also mitigate certain risks associated with this change. Firstly, sensitive data such as credentials 
must be appropriately secured in the environment variables to prevent any security breaches. 
Secondly, all team members need to understand and follow the new method for managing configurations to ensure consistency.

A comprehensive review process for configuration changes will also be necessary to ensure that changes don't inadvertently 
introduce new issues or vulnerabilities.
