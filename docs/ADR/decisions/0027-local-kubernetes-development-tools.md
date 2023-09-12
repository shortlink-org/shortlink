# 27. 🛠️ Local Kubernetes Development Tools

Date: 2023-09-12

## 📌 Status

Accepted

## 🌐 Context

The Kubernetes development landscape is continuously evolving, with numerous tools aiming to simplify the process. 
Each tool has its distinct methodology, and the selection often hinges on individual project needs, 
the team's proficiency, and the favored workflow. A category termed “Local K8S Development tools” has surfaced, 
aiming to amplify the Kubernetes development experience by linking locally running components to the Kubernetes cluster. 
This connection allows swift testing of fresh code in cloud conditions, bypassing the conventional Dockerization, CI, 
and deployment cycle.

## 🚀 Decision

To provide a comprehensive understanding of the local Kubernetes development tools, 
we will delve into three prominent solutions: Telepresence, Gefyra, and mirrord.

| Tool         | Description                                                                                                                 | Link                                         |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|----------------------------------------------|
| Telepresence | Uses a VPN to connect the user's machine to the cluster's network. Requires a local daemon and a Traffic Manager component. | [Telepresence](https://www.telepresence.io/) |
| Gefyra       | Uses a VPN to connect Docker containers to the cluster. Focuses on network traffic.                                         | [Gefyra](https://gefyra.dev/)                |
| mirrord      | Injects itself into the local binary and proxies to an agent in the cluster. Supports multiple local processes.             | [mirrord](https://mirrord.dev/)              |

Given the context and the features of the tools, I'd prefer to try [mirrord](https://mirrord.dev/).

## ⚖️ Consequences

- **Easier**: Streamlining the Kubernetes development cycle, testing new code rapidly in cloud conditions, 
  and bypassing traditional development cycles.
- **Difficult**: Choosing the right tool based on project requirements, team's expertise, and preferred workflow. 
  Each tool has its learning curve and setup complexities.
- **Risks**: Potential integration issues with existing workflows, and the need to stay updated with 
  the evolving Kubernetes tool ecosystem to ensure compatibility and security.
