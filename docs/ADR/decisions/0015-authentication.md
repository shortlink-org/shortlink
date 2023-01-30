# 15. authentication

Date: 2023-01-30

## Status

Accepted

## Context

We need to decide how we will authenticate users. We have a few options:

* OAuth2
* JWT
* 2FA

## Decision

We want use third-party authentication, because it is more secure and we don't support this service.

We research next third-party authentication:
+ [Keycloak](https://www.keycloak.org/)
+ [ory/hydra](https://www.ory.sh/hydra/)
+ [dexidp/dex](https://dexidp.io/)

## Consequences

### Step 1. Try to use Keycloak

1. Install Keycloak
2. Try to use Keycloak
3. Made decision about use Keycloak and update this ADR
