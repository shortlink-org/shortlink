### Dashboard

##### Auth

Get token:

```
export NAMESPACE=dashboard
kubectl -n ${NAMESPACE} describe secret $(kubectl -n ${NAMESPACE} get secret | grep admin-user | awk '{print $1}')
```
