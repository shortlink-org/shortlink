## Ops

### ADR

- **CI/CD**
  - [ADR-0001](./docs/ADR/decisions/0001-ci-cd.md) - Use CI/CD
  - [ADR-0016](./docs/ADR/decisions/0016-lifecycle-deploy.md) - Lifecycle deploy
- **Kubernetes**
  - [ADR-0022](./docs/ADR/decisions/0022-kubernetes.md) - Kubernetes
  - [ADR-0027](./docs/ADR/decisions/0027-local-kubernetes-development-tools.md) - Local Kubernetes Development Tools
- **DataBase**
  - [ADR-0002](./docs/ADR/decisions/0002-cache-system.md) - Cache System
  - **PostgreSQL**
  - [ADR-0005](./docs/ADR/decisions/0005-postgres-optimization.md) - Postgres optimization
  - [ADR-0020](./docs/ADR/decisions/0020-postgresql-operator.md) - Research Summary and Decision for PostgreSQL Operators: Choosing Crunchy Postgres Operator

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
+ Support K8S (support version 1.24+)
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
