# 12. Adoption of Talos OS for Kubernetes Infrastructure

Date: 2024-11-01

## Status

Accepted

## Context

As our Kubernetes infrastructure grows, we require an operating system that is secure, minimal, and purpose-built for container orchestration. The OS should reduce maintenance overhead, enhance security, and integrate seamlessly with Kubernetes.

We considered alternatives like Flatcar Container Linux, which has recently entered the CNCF Incubator ([Flatcar brings Container Linux to the CNCF Incubator](https://www.cncf.io/blog/2024/10/29/flatcar-brings-container-linux-to-the-cncf-incubator/)). While Flatcar offers a container-optimized OS, we seek a solution that provides an immutable and API-driven operating system. Additionally, [Talos Factory](https://factory.talos.dev/) offers advanced customization and plugin support that aligns with our operational requirements.

## Decision

We have decided to adopt [Talos OS](https://www.talos.dev/) as the operating system for our Kubernetes nodes. As part of this adoption, we have enabled the `intel-ucode` plugin to ensure up-to-date microcode updates for enhanced hardware security and stability.

## Alternative

### Flatcar Container Linux

Flatcar Container Linux is a leading alternative that provides a container-optimized, immutable operating system. It recently entered the CNCF Incubator, indicating strong community and industry support. Flatcar offers automatic updates, robust security features, and compatibility with existing container orchestration tools.

**Pros:**
- **Mature Ecosystem**: Established user base and extensive documentation.
- **CNCF Incubator**: Backed by the Cloud Native Computing Foundation, ensuring ongoing support and integration.
- **Automatic Updates**: Seamless OS updates without downtime.

**Cons:**
- **Customization Limitations**: Less flexibility in deep OS customization compared to Talos OS.
- **API Management**: Relies more on traditional management tools rather than a fully API-driven approach.
- **Community Size**: While growing, it may not match the specific features and support provided by Talos OS for certain use cases.

### Other Alternatives

- **Ubuntu Core**: Offers a secure, minimal OS with snap package management. However, it may require more maintenance and lacks some of the Kubernetes-specific optimizations of Talos OS.
- **CoreOS**: The predecessor to Flatcar, now largely succeeded by Flatcar, but still considered for legacy systems.

After evaluating these alternatives, Talos OS emerged as the best fit due to its immutability, API-driven management, and specialized features tailored for Kubernetes environments.

## Consequences

- **Benefits**:
  - **Security**: Talos OS is immutable and has no SSH access, significantly reducing the attack surface. Enabling the `intel-ucode` plugin ensures that our hardware is protected against known vulnerabilities.
  - **Simplicity**: Managed entirely via APIs, it simplifies operations and automation.
  - **Consistency**: Immutable infrastructure ensures predictable environments across all nodes.
  - **Customization**: Utilizing [Talos Factory](https://factory.talos.dev/) allows for tailored configurations to meet our specific needs.

- **Drawbacks**:
  - **Learning Curve**: Team members will need to learn Talos OS concepts and management.
  - **Tooling Adjustments**: Existing tools relying on traditional OS access may require modification.
  - **Community Support**: As a newer OS, it may have less extensive community support compared to established alternatives.

To address these drawbacks, we will provide team training and update our tooling and processes to align with Talos OS.

## References

- [Talos OS](https://www.talos.dev/)
- [Talos Factory](https://factory.talos.dev/)
- [Flatcar Container Linux CNCF Blog](https://www.cncf.io/blog/2024/10/29/flatcar-brings-container-linux-to-the-cncf-incubator/)
