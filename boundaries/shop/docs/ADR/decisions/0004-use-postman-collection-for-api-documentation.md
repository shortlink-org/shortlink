# 4. Use Postman Collection for API Documentation

Date: 2024-09-03

## Status

Accepted

## Context

Postman collections offer an interactive environment where users can directly experiment with GraphQL queries, mutations, 
and subscriptions while viewing the documentation. This approach also allows for the inclusion of sample queries, variables, 
and responses, making it easier for developers to understand how to interact with the API.

## Decision

We will use a Postman collection to document our GraphQL API examples. This collection will include detailed 
examples of queries, mutations, and subscriptions, along with explanations of the GraphQL schema, variables, 
and expected responses. The collection will be maintained alongside the development of the GraphQL API 
to ensure accuracy and relevance.

The Postman collection can be accessed [here](https://raw.githubusercontent.com/shortlink-org/shortlink/main/boundaries/shop/docs/API/shortlink-shop.postman_collection.json).

## Consequences

Using Postman collections for documenting GraphQL API examples will provide a more interactive and hands-on experience for developers. 
This will allow them to easily experiment with queries, mutations, and subscriptions, leading to a better understanding 
of how to interact with the API. 
