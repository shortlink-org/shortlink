---
title: "Csi"
date: 2020-11-02T03:27:21+03:00
draft: true
categories:
    - Kubernetes
tags:
    - k8s
    - devops
---

### Container Storage Interface (CSI)

![csi image](https://d33wubrfki0l68.cloudfront.net/6f9fa026192623422d7e35e3aa4ba91898e86a35/c299e/images/blog-logging/2018-04-10-container-storage-interface-beta/csi-logo.png)

##### Docs

- [What is The Container Storage Interface (CSI)?](https://www.architecting.it/blog/container-storage-interface/)
- [Official docs](https://kubernetes-csi.github.io/docs/)
- [[RU] Понимаем Container Storage Interface (в Kubernetes и не только)](https://habr.com/ru/company/flant/blog/424211/)

##### Example

- [csi-driver-host-path](https://github.com/kubernetes-csi/csi-driver-host-path) - a sample (non-production) CSI Driver that creates a local directory as a volume on a single node
- [csi-digitalocean](https://github.com/digitalocean/csi-digitalocean) - A Container Storage Interface (CSI) Driver for DigitalOcean Block Storage

**csi-digitalocean** has a more friendly year. 
