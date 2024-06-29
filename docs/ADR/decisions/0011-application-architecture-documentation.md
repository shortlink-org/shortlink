# 11. Application architecture documentation

Date: 2023-01-07

## Status

Accepted

## Context

We want to have a standard documentation for easy understanding of 
the logic of the platform and its architecture.

We research models:

  - [4C model](https://c4model.com/)
  - 4+1 model
  - ARC42
  - Archimate
  - Diagrams.net

We research tools:

  - archimate model
  - archi -> neo4j -> grafana
  - PlantUML
    - PlantUML-C4
  - OpenAPI specification
  - Gherkin
  - Miro
  - Notion
  - [diagrams](https://diagrams.mingrammer.com/)

## Decision

+ We use the [4C model](https://c4model.com/) for the documentation of the application architecture.
+ We use [PlantUML](https://plantuml.com/) for the documentation of the application architecture.
+ We use [C4-PlantUML](https://github.com/plantuml-stdlib/C4-PlantUML) for the documentation of the application architecture.

### Architecture documentation

- [C4](./docs/c4)
- [Docs in Notion](https://shortlink-org.notion.site/Low-level-f61e3d5ab4ad484784cada86de569eba)
- [Low-level schema](https://miro.com/app/board/o9J_laImQpo=/)
- [Auth](https://miro.com/app/board/o9J_lA5Wmhg=/)
- [Event Sourcing](https://miro.com/app/board/o9J_l-6o1U0=/)

### References

- [Example C4 shop-service](https://gitlab.com/microarch-ru/microservices/system-design)

### We analyze next tools and documentation models:

- [archi editor](https://www.archimatetool.com/)
  + [C4 Model, Architecture Viewpoint and Archi 4.7](https://www.archimatetool.com/blog/2020/04/18/c4-model-architecture-viewpoint-and-archi-4-7/)
  + References:
    + [RecruitmentCompany](https://github.com/rogerkohler/RecruitmentCompany) - Archimate All layers
    + [Cloud Native Architectural Models using Archi](https://github.com/alberttwong/archi_cloudnative)

## Consequences

+ We have a standard documentation for the application architecture.
+ We use the [4C model](https://c4model.com/) for the documentation of the application architecture.
