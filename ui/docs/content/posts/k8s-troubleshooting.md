---
title: "K8s Troubleshooting"
date: 2021-05-31T21:24:14+03:00
draft: false
categories:
    - Kubernetes
tags:
    - k8s
    - devops
    - troubleshooting
---

#### Namespace deletion stuck

Script: https://github.com/batazor/shortlink/tree/main/ops/Helm/tooling

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
