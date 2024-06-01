## Common Expression Language

### Goal of the POC

The goal of this...

### Example

#### Rules

You can see more examples [here](./rules).

```
# Validate that the three fields defining replicas are ordered appropriately
self.minReplicas <= self.replicas && self.replicas <= self.maxReplicas

# Validate that an object's name matches a specific value (making it a singleton)
self.metadata.name == 'singleton'
```

#### Curl

```shell
curl -X POST http://localhost:8080/evaluate \
  -H "Content-Type: application/json" \
  -d '{
    "rule": "jwt_audience.cel",
    "claims": {
      "exp": 1234567890,
      "aud": "my-audience"
    },
    "now": 1234567891
  }'
```

**response**:

```json
{"jwt_audience.cel":true,"jwt_exp.cel":true}
```

### Docs

- [docs](https://cel.dev/)
- [Language Definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md)
- [cel-lab](https://codelabs.developers.google.com/codelabs/cel-go#1)
- [Common Expression Language in Kubernetes](https://kubernetes.io/docs/reference/using-api/cel/)
