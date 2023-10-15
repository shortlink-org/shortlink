# 2. Use oapi-codegen

Date: 2023-10-15

## Status

Accepted

## Context

Post the inception of a BFF service for our web application, there's a need to automate Go boilerplate code generation 
in alignment with [OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md) standards. 
This would further the effectiveness of our BFF service by ensuring a streamlined, error-minimized process while adhering 
to industry standards for API contracts between the web application and microservices.

## Decision

Implement oapi-codegen to automate the generation of Go boilerplate code, enabling us to maintain a coherent API contract 
between servers and clients while focusing more on business logic implementation.

## Consequences

- **Ease of Code Generation**: Simplifies boilerplate code generation, reducing manual work and potential errors.
- **Alignment with OpenAPI Specifications**: Ensures our services are compliant with OpenAPI specifications, facilitating communication between the web application and microservices.
- **Dependency on oapi-codegen**: Introduces a dependency on oapi-codegen, necessitating monitoring for tool updates and potential bugs that could affect the generated code.
