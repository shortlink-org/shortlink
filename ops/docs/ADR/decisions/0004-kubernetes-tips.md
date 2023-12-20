# 4. Kubernetes tips [Cookbook]

Date: 2023-12-20

## Status

Accepted

## Cookbook

### Recipe 1: Approving Certificate Signing Requests (CSRs)

When working with Kubernetes, you might encounter situations where you need to approve Certificate Signing Requests (CSRs). This is a common task when adding new nodes to a cluster or issuing certificates for user authentication. To efficiently manage CSRs, you can use the following Kubernetes commands.

> [!TIP]
> **Approve Certificate Signing Requests**
>
> To list Certificate Signing Requests sorted by their creation timestamp, use:
> ```bash
> kubectl get csr --sort-by=.metadata.creationTimestamp
> ```
> To approve a specific CSR, for example, `csr-fm7hw`, run:
> ```bash
> kubectl certificate approve csr-fm7hw
> ```

This process helps maintain the security and integrity of your Kubernetes cluster by ensuring that only valid and verified nodes or users can interact with the cluster.
