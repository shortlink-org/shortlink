## Ops

### ADR

- [README.md](./docs/ADR/README.md) - Architecture Decision Records

### MQ

| [Kafka](https://kafka.apache.org/) | [RabbitMQ](https://www.rabbitmq.com/) | [NATS](https://nats.io/) |
|------------------------------------|---------------------------------------|--------------------------|

### Third-party Service

| Boundary      | Service       | Description                                                             | Language/Framework | Docs                                                | Status                                                                                                                                                  |
|---------------|---------------|-------------------------------------------------------------------------|--------------------|-----------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Observability | grafana       | Grafana is the open source analytics & monitoring solution for          | More               | [docs](https://grafana.com/docs/)                   | [![App Status](https://argo.shortlink.best/api/badge?name=grafana&revision=true)](https://argo.shortlink.best/applications/grafana)               |       
| Platform      | cert-manager  | Automatically provision and manage TLS certificates in Kubernetes       | Go                 | [docs](https://cert-manager.io/docs/)               | [![App Status](https://argo.shortlink.best/api/badge?name=cert-manager&revision=true)](https://argo.shortlink.best/applications/cert-manager)     |  
| Network       | nginx-ingress | Ingress controller for Kubernetes using NGINX                           | Go                 | [docs](https://kubernetes.github.io/ingress-nginx/) | [![App Status](https://argo.shortlink.best/api/badge?name=nginx-ingress&revision=true)](https://argo.shortlink.best/applications/nginx-ingress)   | 
| Network       | istio         | Istio is an open platform to connect, manage, and secure microservices. | Go                 | [docs](https://istio.io/latest/docs/) |
| MQ            | kafka         | Kafka is used as a message broker                                       | Java               | [docs](https://kafka.apache.org/)                   | [![App Status](https://argo.shortlink.best/api/badge?name=kafka&revision=true)](https://argo.shortlink.best/applications/kafka)                   |         
| Security      | keycloak      | Keycloak is an open source identity and access management solution      | Java               | [docs](https://www.keycloak.org/documentation.html) | [![App Status](https://argo.shortlink.best/api/badge?name=keycloak&revision=true)](https://argo.shortlink.best/applications/keycloak)             | 

#### Draft

| Boundary | Service       | Description                                                        | Language/Framework | Docs                                                | Status                                                                                                                                                  |
|----------|---------------|--------------------------------------------------------------------|--------------------|-----------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Platform | keycloak      | Keycloak is an open source identity and access management solution | Java               | [docs](https://www.keycloak.org/documentation.html) | [![App Status](https://argo.shortlink.best/api/badge?name=keycloak&revision=true)](https://argo.shortlink.best/applications/keycloak)             | 

### Cloud-Native stack

+ Development
  + [Skaffold](https://skaffold.dev/)
  + [Telepresence](https://www.telepresence.io/)
+ Security
  + [SOPS](https://github.com/mozilla/sops)
  + [Teleport](https://goteleport.com/) - Identity-Native Infrastructure Access
  + cert-manager
    + cloudfare
    + spiffe
  + [kubescape](https://github.com/kubescape/kubescape) - Kubernetes security auditing
+ Support K8S (support version 1.29+)
  + Helm Chart
    + [pingcap/chaos-meshh](https://github.com/pingcap/chaos-mesh)
  + Minikube
  + Backup/Restore [(Velero)](https://velero.io/)
  + Custom CSI driver (fork [csi-driver-host-pat](https://github.com/kubernetes-csi/csi-driver-host-path))
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
