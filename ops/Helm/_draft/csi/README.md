### CSI driver

#### Install

1. Deploy the CSI plugin and sidecars

```
kubectl apply -n default --prune --applyset=csi-driver -f ./templates
```

#### Info

| Name     | Value        |
|----------|--------------|
| name     | csi.shrts.ru |
