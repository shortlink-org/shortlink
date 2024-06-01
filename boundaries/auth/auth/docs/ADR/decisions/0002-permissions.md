# 2. Implementing Permissions using SpiceDB

Date: 2023-05-20

## Status

Accepted

## Context

We need a flexible, scalable, and secure way to manage permissions in our application. 
Permissions should be easily manageable and should allow for fine-grained control over user actions. 
To meet these requirements, we have evaluated two open-source systems: ORY Keto and SpiceDB.

ORY Keto is an open-source policy decision point that uses a flexible attribute-based access control (ABAC) model. 
This allows us to define complex policies based on various attributes. 
However, we identified certain limitations in terms of scalability and ease of management for our growing application.

SpiceDB, on the other hand, is a Google Zanzibar-inspired database system for managing application permissions. 
It's designed for high scalability and offers a flexible model for defining and evaluating access control lists (ACLs), 
allowing us to define complex policies based on various attributes. 

After careful evaluation, we found that SpiceDB aligns better with our needs.

## Decision

We will use SpiceDB to manage permissions in our application. The following namespaces and
relationships will be implemented:

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

System_Boundary(spicedb, "SpiceDB") {
  Container(link, "Link", "Represents links with their associated permissions")
  Container(user_namespace, "User", "Represents the User namespace")
}

System_Boundary(permissions, "Permissions") {
  Container(reader, "View", "View permission")
  Container(writer, "Edit/Delete", "Edit/Delete permission")
  Container(share, "Share", "Share permission")
}

Rel(customer, spicedb, "Interacts with")

Rel(link, writer, "Uses")
Rel(link, reader, "Uses")
Rel(link, share, "Uses")
```

This structure allows for a flexible and secure permission management system that supports 
fine-grained control over user actions.

### Architecture SpiceDB

> Docs:
> - [Architecture](https://authzed.com/blog/spicedb-architecture)
> - [ABAC on SpiceDB: Enabling Netflix’s Complex Identity Types](https://authzed.com/blog/abac-on-spicedb-enabling-netflix-complex-identity-types)

## Consequences

**Props**:

- The system will be more secure, as permissions will be managed centrally by SpiceDB.
- Permissions can be easily managed and updated.
- As our application grows, SpiceDB will be able to scale with it, ensuring that the permission
  management system remains efficient and performant.

### Alternatives

We considered using ORY Keto for managing permissions in our application. 
While ORY Keto offered an attribute-based access control (ABAC) model, 
which initially seemed appealing, we found that SpiceDB's model 
for defining and evaluating access control lists (ACLs) was more suitable 
for our needs and offered better scalability for our growing application.

### Proof of Concept

**Example Permission**:

- [ORY/Keto](./proof/ADR-0002/permissions/permissions-v1.ts)
- [SpiceDB](./proof/ADR-0002/permissions/permissions-v1.zed)
