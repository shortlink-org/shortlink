## Deploy K8S

1. Run GitLab-Agent on your Kubernetes cluster

**Link**: https://gitlab.com/<ORG_NAME>/<PROJECT_NAME>/-/clusters

```bash
helm repo add gitlab https://charts.gitlab.io
helm repo update
helm upgrade --install contabo gitlab/gitlab-agent \
  --namespace <namespace> \
  --create-namespace \
  --set image.tag=<version> \
  --set config.token=<token> \
  --set config.kasAddress=wss://kas.gitlab.com
```

2. Set Pod Security Policy

```
kubectl label --overwrite ns <namespace> \
  pod-security.kubernetes.io/enforce=privileged
```
