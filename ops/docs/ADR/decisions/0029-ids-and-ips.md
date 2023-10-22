# 29. Intrusion Detection System (IDS) and Intrusion Prevention System (IPS) for Kubernetes (k8s)

Date: 2023-10-22

## Status

Accepted

## Context

In the ever-evolving landscape of cyber threats, it's crucial for Kubernetes (k8s) environments to have robust security 
measures in place. Kubernetes, being an open-source container orchestration platform, is often targeted by malicious 
entities for unauthorized access and malicious activities. Two essential components for enhancing the security posture of 
Kubernetes clusters are the Intrusion Detection System (IDS) and Intrusion Prevention System (IPS). 
There is a need to evaluate and select the most suitable IDS/IPS solution for our Kubernetes environment.

## Decision

After evaluating multiple options, including Tigera Calico Cloud and Falco, we have decided to implement **Falco** as 
our IDS/IPS solution for Kubernetes. Our decision is based on the following considerations:

1. **Integration Capabilities**: Falco's ability to seamlessly integrate with Argo events and Argo workflows makes it 
    a versatile tool for our Kubernetes environment.
2. **Self-hosted Advantage**: Falco, being a self-hosted solution, offers us more control over our security infrastructure.
3. **Active Community Support**: Falco's active open-source community ensures that we benefit from regular updates, 
    bug fixes, and new feature introductions.
4. **Runtime Security**: Falco specializes in providing runtime security, enabling us to detect and alert on any abnormal 
    behavior in applications running within containers.

## Consequences

**Benefits**:
1. **Enhanced Security**: With Falco, our Kubernetes environment will be better protected against unauthorized access and malicious activities.
2. **Automated Workflows**: Integration with Argo events and workflows allows us to automate responses to specific types 
    of alerts, enhancing our incident response capabilities.

**Challenges**:
1. **Maintenance**: Being an open-source tool, Falco requires regular maintenance to ensure it is configured correctly 
    and updated to the latest version.
2. **Technical Expertise**: Effective use of Falco necessitates a certain level of technical expertise for setup, 
    configuration, and rule creation.
3. **Potential False Positives**: Like any IDS/IPS, Falco might generate false positive alerts, which would require 
    tuning and refining of rules.

**Risks**:
1. **Dependency on Community**: While Falco has an active community, any potential lapses in updates or bug fixes might affect our security posture.
2. **Complexity**: Integration with Argo events and workflows, although beneficial, introduces an additional layer of complexity to our Kubernetes setup.

### How it works

```
┌─────────────┐           ┌─────────┐          ┌────────────────┐
│             │  detect   │         │  push    │                │
│  pwned pod  ├───────────►  falco  ├──────────► falcosidekick  ├────┐
│             │           │         │          │                │    │
└──────▲──────┘           └─────────┘          └────────────────┘    │ notify
       │                                                             │
       │                                                             │
delete │   ┌──────────────┐          ┌───────────────┐        ┌──────▼──────┐
       │   │              │          │               │        │             │
       └───┤ deletion pod ◄──────────┤ argo workflow │        │ argo events │
           │              │  create  │               │        │             │
           └──────────────┘          └────────────▲──┘        └─┬───────────┘
                                                  │             │
                                          trigger │             │ push
                                                  │             │
                                                ┌─┴─────────────▼──┐
                                                │       bus        │
                                                └──────────────────┘
```

## References

1. [Falco](https://falco.org/)
2. [Falco rules](https://falcosecurity.github.io/rules/)
3. [Kubernetes Response Engine, Part 5: Falcosidekick + Argo](https://falco.org/blog/falcosidekick-response-engine-part-5-argo/)
