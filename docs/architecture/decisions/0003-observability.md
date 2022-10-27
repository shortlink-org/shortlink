# 4. Observability

Date: 2022-10-28

## Status

Accepted

## Context

We need to be able to monitor the health of our services and be able to collect data on the metrics we are most interested in a standardised way

## Decision

We will add two standard endpoints to each service on port 9090.

- `/metrics` will return a prometheus formatted metrics page
- `/health` will return a json object with a `status` field that is either `ok` or `error`

+ Also, we will add command `HEALTHCHECK` to the dockerfile for each service that will call the `/health` endpoint and check the status.
+ Also, we add k8s liveness and readiness probes to each service that will call the `/health` endpoint and check the status.

## Consequences

We will be able to monitor the health of our services and collect metrics in a standardised way.
