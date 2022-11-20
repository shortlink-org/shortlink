# 3. Observability

Date: 2022-10-28

## Status

Accepted

## Context

We need to be able to monitor the health of our services and be able to collect data on the metrics we are most interested in a standardised way

## Decision

We will add three standard HTTP-endpoints to each service on port 9090.

- `/metrics` will return a prometheus formatted metrics page
- `/healthz` will return a json object with a `status` field that is either `ok` or `error`
- `/ready` will return a json object with a `status` field that is either `ok` or `error`

#### Dockerfile

We will add command `HEALTHCHECK` to the dockerfile for each service that will call the `/healthz` endpoint and check the status.

#### Helm

Also, we add k8s *-probes to each service that will call the `/healthz` endpoint and check the status.

For Helm-chart we use complete template:

```gotemplate
{{- include "shortlink-common.probe" .Values.deploy | indent 6 }}
```

## Consequences

We will be able to monitor the health of our services and collect metrics in a standardised way.
