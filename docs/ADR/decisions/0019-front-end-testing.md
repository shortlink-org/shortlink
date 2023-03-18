# 19. front-end testing

Date: 2023-03-18

## Status

Accepted

## Context

As our application grows in complexity, we need to establish a reliable front-end testing strategy. 
Additionally, we need a consistent and effective way to test our APIs.

## Decision

After careful consideration, we have decided to use [Cypress](https://www.cypress.io/) for front-end testing,
and we will utilize [Cypress Testing Library](https://testing-library.com/docs/cypress-testing-library/intro/) 
for writing tests. We have chosen Cypress because it offers a comprehensive set of tools for testing, debugging, 
and CI/CD integration. Furthermore, its ability to simulate user interactions and automate repetitive tasks makes 
it an ideal choice for front-end testing. 

Cypress Testing Library offers a simple and intuitive API for testing React components, which will help us
write more reliable and maintainable tests.

For API testing, we will use Postman and Newman. Postman will allow us to test our APIs manually, and we can use Newman
for automated API testing as part of our CI/CD pipeline.

## Consequences

By adopting Cypress and Cypress Testing Library for front-end testing, we can expect to have a more reliable and 
efficient testing process. Additionally, using Postman and Newman for API testing will allow us to catch issues early in 
the development process and ensure that our APIs are working as expected. Overall, this decision will help us achieve 
higher quality software and faster delivery times.
