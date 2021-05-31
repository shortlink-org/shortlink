---
title: "K8s Troubleshooting"
date: 2021-05-31T21:24:14+03:00
draft: true
---

#### Namespace deletion stuck

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
