# 3. Select REST Framework for Public API

Date: 2024-08-13

## Status

Accepted

## Context

We need a robust and flexible framework to build and manage our public API. 
The framework should support easy integration with our existing Django project, 
provide comprehensive documentation capabilities, and ensure security and scalability.

## Decision

We have decided to use Django REST framework (DRF) for our public API. DRF is a powerful and flexible toolkit for building Web APIs in Django. 
It provides a wide range of features, including:

- Serialization that supports both ORM and non-ORM data sources.
- Authentication and permissions.
- Viewsets and routers for easy URL routing.
- Pagination, filtering, and ordering.
- Browsable API for easy testing and debugging.

Additionally, we will use `drf-spectacular` to generate and serve our API documentation. This includes:

- **Swagger UI**: A user-friendly interface for exploring and testing the API.
- **ReDoc UI**: A responsive and customizable API documentation interface.

### Swagger and ReDoc URLs

- **Swagger UI**: `/api/schema/swagger-ui/`
- **ReDoc UI**: `/api/schema/redoc/`

## Consequences

By selecting Django REST framework, we ensure that our public API is built on a well-supported and widely-used framework. 
This decision will make it easier to maintain and extend the API in the future. 
The integration with `drf-spectacular` will provide comprehensive and interactive API documentation, 
improving the developer experience and facilitating easier integration with third-party services.
