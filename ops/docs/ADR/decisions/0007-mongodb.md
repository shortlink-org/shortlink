# 31. MongoDB

Date: 2023-12-24

## Status

Accepted

## Context

We need a database to store data.

## Decision

MongoDB is a document database with the scalability and flexibility that you want with the querying and indexing that you need.

For kubernetes we can use:

- [KubeDB](https://kubedb.com/kubernetes/databases/run-and-manage-mongodb-on-kubernetes/) - KubeDB is a Kubernetes Custom Resource Definition for managing stateful applications. 
KubeDB provides a cloud native, declarative way to manage MongoDB on Kubernetes.
- [Percona Server for MongoDB](https://www.percona.com/doc/kubernetes-operator-for-psmongodb/index.html) - MongoDB operator for Kubernetes provider by Percona.
- [MongoDB Community Kubernetes Operator](https://github.com/mongodb/mongodb-kubernetes-operator) - MongoDB Community Kubernetes Operator provider by MongoDB.

## Consequences

**MongoDB Community Kubernetes Operator** - is the good choice for us. It is a Kubernetes Operator for MongoDB Community Server.
But it is not production ready yet. This operator does not support mongodb v7.

**KubeDB** - interesting solution. It is production ready. But it has a lot of features for a community edition.

So, need watch for **MongoDB Community Kubernetes Operator** and **KubeDB**.
