# 1. Init project

Date: 2024-09-13

## Status

Accepted

## Context

We are building an **A/B testing platform**, now named `ab-platform`, to support the creation, execution, and analysis 
of experiments that measure the effectiveness of different variations in a controlled environment. 
The platform will allow teams to make data-driven decisions by comparing variations across features, designs, or workflows.

Initially, the platform will follow a **monolithic architecture**, which will simplify development and deployment in 
the early stages. As the system evolves and scales, we may later transition to a more modular or service-oriented architecture.

The first module will implement **CRUD operations** for managing test cases, which are the core elements of each experiment.

## Decision

We will initiate the project with the following steps:

### 1. **Monolithic Architecture**:
- The platform will be designed as a single deployable unit to simplify development and testing.
- All components for managing experiments, user assignment, metrics collection, and reporting will be part of the same codebase.

### 2. **First Module - CRUD for Test Cases**:
- The first focus will be on creating a module for **test case management**.
- This will include creating, reading, updating, and deleting (CRUD) test cases, which define the various aspects of experiments.
- Each test case will store information like:
  - Test name
  - Variations
  - Target user groups
  - Start and end dates
  - Associated metrics for evaluation

## Consequences

- **Monolithic Architecture**: Simplifies development and initial deployment but might require refactoring as the platform scales.
- **CRUD for Test Cases**: Provides the foundation for managing experiments and their configurations, 
which will enable further features such as user assignment and metrics tracking.
