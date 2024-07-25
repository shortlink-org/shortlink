# 6. Kubernetes tips

Date: 2022-12-10

## Status

Accepted

## Context

`CODEOWNERS` file is used to define individuals or teams that are responsible for code in a repository.

## Decision

### Example from GitLab

![gitlab-codeowner](images/ADR-0006/codeowner.png)

- [Best practices to set up organizational hierarchies that scale](https://about.gitlab.com/blog/2024/07/22/best-practices-to-set-up-organizational-hierarchies-that-scale/)

### Structure

```
- groupCompany
  - project-workspace-1
    - project-1
    - project-2
    - project-3
  - project-workspace-N
    - project-1
    - project-N
  - maintainers
    - frontend
    - backend
    - database
    - etc...
```

## Consequences

- We can easily add people to teams and teams to projects.
- We assign teams to projects, not people.
