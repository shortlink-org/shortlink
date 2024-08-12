# 3. Auth mode

Date: 2024-03-24

## Status

Accepted

## Context

We need a robust and secure authentication system for our Django application. 
Ory Kratos is an open-source identity system that fits our requirements.

## Decision

We have decided to use `django_ory_auth`, a Django middleware that simplifies the integration with Ory Kratos. 
This middleware handles the communication between our Django application and the Ory Kratos service, 
providing a seamless authentication experience.

## Consequences

By using `django_ory_auth`, we can leverage the features of Ory Kratos without having to implement the integration ourselves. 
This decision speeds up the development process and ensures that we are using a tested and secure authentication system.

### Admin roles

For set up the admin roles, need open Postgres and set next values:

- `is_staff` - True
- `is_superuser` - True
