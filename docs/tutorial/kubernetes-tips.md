## Kubernetes tips

#### Namespace deletion stuck

Script: https://github.com/shortlink-org/shortlink/tree/main/ops/Helm/tooling

```
$> kubectl proxy
Starting to serve on 127.0.0.1:8001
```

```
$> kubectl get namespace istio-system -o json > tmp.json
# remove "spec.kubernetes"
$> vim tmp.json
$> curl -k -H "Content-Type: application/json" -X PUT --data-binary @tmp.json http://127.0.0.1:8001/api/v1/namespaces/istio-system/finalize
```

#### Helm upgraid-faild

```
$> kubectl -n app-namespace patch secret release-name.v123 --type=merge -p '{"metadata":{"labels":{"status":"deployed"}}}'
```

## Tool

- [krew](https://github.com/kubernetes-sigs/krew) manager plugin for kubectl
