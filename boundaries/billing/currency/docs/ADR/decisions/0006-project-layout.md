# 6. Project Layout

Date: 2024-09-15

## Status

Accepted

## Context

To maintain a scalable, maintainable, and clear codebase, we need to establish a consistent project layout. 
The layout should facilitate easy navigation, modularity, and adherence to Domain-Driven Design (DDD) principles. 
It should also support growth and complexity as the project evolves, enabling the addition of new features and 
components without creating a monolithic or tangled code structure.

## Decision

We will structure the project into the following main directories, each serving a specific purpose:

### 1. Domain Layer (`src/domain`)

- **Purpose**: Defines the core business logic and rules.
- **Structure**: Organized into submodules representing different domains (e.g., `exchange_rate`, `currency_conversion`).
- **Components**:
    - **Entities**: Core objects representing the state and behavior of the business.
    - **Value Objects**: Immutable objects that describe specific aspects or attributes of entities.

Example:

```
src/
├── domain/
│   ├── exchange_rate/
│   │   ├── mod.rs
│   │   ├── entities.rs
│   │   ├── value_objects.rs
│   ├── currency_conversion/
│   │   ├── mod.rs
│   │   ├── entities.rs
```

### 2. Use Case Layer (`src/usecases`)

- **Purpose**: Contains the application-specific business logic, orchestrating the interaction between the domain layer and other components.
- **Structure**: Each use case is encapsulated in its own submodule.
- **Components**:
    - **Use Cases**: Handle specific business operations, using services and repositories to execute the desired functionality.

Example:

```
src/
├── usecases/
│   ├── mod.rs
│   ├── exchange_rate/
│   │   ├── mod.rs
│   │   ├── fetcher.rs
│   ├── currency_conversion/
│   │   ├── mod.rs
│   │   ├── converter.rs
```

### 3. Repository Layer (`src/repository`)

- **Purpose**: Defines the interfaces for data access and manipulation.
- **Structure**: Organized into submodules corresponding to the domains they support.
- **Components**:
    - **Repositories**: Interfaces for CRUD operations and data retrieval.
    - **Implementations**: In-memory or database-backed implementations of the repositories.

Example:

```
src/
├── repository/
│   ├── mod.rs
│   ├── exchange_rate/
│   │   ├── mod.rs
│   │   ├── repository.rs
│   │   ├── in_memory_repository.rs
```

### 4. Documentation (`./docs/ADR/decisions`)

- **Purpose**: Holds Architecture Decision Records (ADRs) and other project documentation.
- **Structure**: Organized into a folder structure for easy reference.

Example:

```
docs/
├── ADR/
│   ├── decisions/
│   │   ├── 0001-init-project.md
│   │   ├── 0003-c4-model.md
│   │   ├── 0006-project-layout.md
```

## Consequences

- **Modularity**: The project is divided into clear modules, making it easier to manage and extend.
- **Maintainability**: By following DDD principles and organizing the codebase into layers, the code remains clean and maintainable.
- **Scalability**: The structure supports the addition of new domains, use cases, and integrations without introducing complexity.
- **Readability**: Developers can easily navigate the codebase and understand the responsibilities of each module.

## Alternatives Considered

- **Flat Structure**: Placing all modules and files in a single directory. This was rejected as it would lead to a lack of modularity and increased complexity as the project grows.
- **Feature-Based Structure**: Organizing by feature rather than by domain and layer. This was rejected in favor of a structure that better supports DDD principles and separates concerns more effectively.
