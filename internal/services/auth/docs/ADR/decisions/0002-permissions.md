# 2. Implementing Permissions using ORY Keto

Date: 2023-04-09

## Status

Accepted

## Context

We need a flexible, scalable, and secure way to manage permissions in our application. 
Permissions should be easily manageable and should allow for fine-grained control over 
user actions.

ORY Keto is an open-source policy decision point that uses a flexible attribute-based 
access control (ABAC) model. This allows us to define complex policies 
based on various attributes, making it a suitable choice for our needs.

## Decision

We will use ORY Keto to manage permissions in our application. The following namespaces and 
relationships will be implemented:

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

System_Boundary(keto, "ORY Keto") {
  Container(link, "Link", "Represents links with their associated permissions")
  Container(user_namespace, "User", "Represents the User namespace")
}

System_Boundary(permissions, "Permissions") {
  Container(view, "View", "View permission")
  Container(edit, "Edit", "Edit permission")
  Container(delete, "Delete", "Delete permission")
  Container(share, "Share", "Share permission")
}

Rel(customer, link, "Interacts with")
Rel(customer, user_namespace, "Interacts with")

Rel(link, view, "Uses")
Rel(link, edit, "Uses")
Rel(link, delete, "Uses")
Rel(link, share, "Uses")
```

This structure allows for a flexible and secure permission management system that supports 
fine-grained control over user actions.

## Consequences

Props:

- The system will be more secure, as permissions will be managed centrally by ORY Keto.
- Permissions can be easily managed and updated.
- As our application grows, ORY Keto will be able to scale with it, ensuring that the permission 
  management system remains efficient and performant.

There will be an initial learning curve for developers who are not familiar with ORY Keto. 
However, the benefits of a flexible and secure permission management system outweigh the initial time investment.
