## Ops

### CI/CD

| [GitLab CI](./ops/gitlab/README.md) | [GitHub CI](./.github/DOCS.md) | [ArgoCD](./ops/argocd/README.md) |
|-------------------------------------|--------------------------------|----------------------------------|


### MQ

| [Kafka](https://kafka.apache.org/) | [RabbitMQ](https://www.rabbitmq.com/) | [NATS](https://nats.io/) |
|------------------------------------|---------------------------------------|--------------------------|

### Cloud-Native stack

+ Development
  + [Skaffold](https://skaffold.dev/)
  + Telepresence
+ Security
  + [SOPS](https://github.com/mozilla/sops)
  + [Teleport](https://goteleport.com/) - Identity-Native Infrastructure Access
  + cert-manager
    + cloudfare
    + spiffe
  + [kubescape](https://github.com/kubescape/kubescape) - Kubernetes security auditing
+ HealthCheck
+ Support K8S
  + Helm Chart
    + [pingcap/chaos-meshh](https://github.com/pingcap/chaos-mesh)
  + Minikube
  + Backup/Restore [(Velero)](https://velero.io/)
  + Custom CSI driver (fork [csi-driver-host-pat](https://github.com/kubernetes-csi/csi-driver-host-path))
+ MetalLB
+ [kyverno](https://kyverno.io/) - Kubernetes Native Policy Management
+ Storage
  + [rook-ceph](https://rook.io/)
    + ceph cluster (3 node)
    + grafana dashboard
    + prometheus metrics
+ Ingress (Gateway)
  + Istio
    + [kiali](https://kiali.io/) - The Console for Istio Service Mesh
  + Nginx
  + Traefik
