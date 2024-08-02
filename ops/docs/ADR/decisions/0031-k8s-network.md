# 31. k8s network

Date: 2024-07-25

## Status

Accepted

## Context

To enhance the networking capabilities within our Kubernetes cluster, we need to integrate advanced networking 
and service mesh solutions. Cilium will provide improved networking and security, while Istio will manage traffic 
and enhance observability, security, and reliability of microservices.

## Decision

We will integrate Cilium for networking and security policies, and Istio for service mesh capabilities within our Kubernetes cluster. 
These tools will be used to manage network traffic, security policies, and service-to-service communication.

## Consequences

+ Enhanced network security and observability.
+ Improved traffic management and service reliability.

### Schema for Kubernetes Network

```plantuml
@startuml k8s_nodes_network

' Define KubernetesPuml for URL inclusion
!define KubernetesPuml https://raw.githubusercontent.com/dcasati/kubernetes-PlantUML/master/dist

' Include common Kubernetes components
!includeurl KubernetesPuml/kubernetes_Common.puml
!includeurl KubernetesPuml/kubernetes_Context.puml

' Include specific components for nodes and Cilium
!includeurl KubernetesPuml/OSS/KubernetesNode.puml
!includeurl KubernetesPuml/OSS/KubernetesPod.puml
!includeurl KubernetesPuml/OSS/KubernetesSvc.puml

' Include C4 model
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

skinparam backgroundColor #FEFEFE
skinparam handwritten false

left to right direction

' Define simple blocks for Kubernetes Ingress and Gateway API
rectangle "NGINX Ingress" as nginx_ingress #98FB98
rectangle "Gateway API" as gateway_api #FF69B4

' Define C4 Context
LAYOUT_TOP_DOWN()

System_Boundary(k8s_cluster, "K8s Cluster") {
    ' Define Nodes
    KubernetesNode(talos_bxi_ihe, "talos-bxi-ihe\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ffcccc")
    KubernetesNode(talos_coj_emd, "talos-coj-emd\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ccffcc")
    KubernetesNode(talos_dks_th8, "talos-dks-th8\nStatus: Ready\nRoles: <none>\nAge: 78d\nVersion: v1.30.3", "#ccccff")
    KubernetesNode(talos_t8v_b87, "talos-t8v-b87\nStatus: Ready\nRoles: control-plane\nAge: 78d\nVersion: v1.30.3", "#ffcc00")
    
    ' Define Cilium pods
    KubernetesPod(cilium_pod1, "cilium-pod1", "#ffd700")
    KubernetesPod(cilium_pod2, "cilium-pod2", "#ffd700")
    KubernetesPod(cilium_pod3, "cilium-pod3", "#ffd700")
    KubernetesPod(cilium_pod4, "cilium-pod4", "#ffd700")
    
    ' Define Istio
    KubernetesPod(istio_pod, "Istio Pod", "#87CEEB")
    
    ' Define Networking
    Rel(talos_bxi_ihe, cilium_pod1, "Network with Cilium")
    Rel(talos_coj_emd, cilium_pod2, "Network with Cilium")
    Rel(talos_dks_th8, cilium_pod3, "Network with Cilium")
    Rel(talos_t8v_b87, cilium_pod4, "Network with Cilium")

    ' Network connections between nodes
    Rel_Left(talos_bxi_ihe, talos_coj_emd, "Inter-node Network")
    Rel_Left(talos_coj_emd, talos_dks_th8, "Inter-node Network")
    Rel_Left(talos_dks_th8, talos_t8v_b87, "Inter-node Network")
    
    ' Define Relationships for Istio, NGINX Ingress, and Gateway API
    Rel(istio_pod, nginx_ingress, "Ingress Traffic")
    Rel(istio_pod, gateway_api, "Gateway Traffic")
}

@enduml
```
