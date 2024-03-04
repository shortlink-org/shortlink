# ADR 4: How We Approach Testing

Date: 2024-03-03

Status: Accepted

## Context

This ADR outlines our chosen strategy for testing to ensure consistent code quality and 
maintainability across the project.

## Decision

We have selected Pytest and Tox as our primary testing frameworks. 
Configuration details are provided in [pyproject.toml](../../../pyproject.toml).

## Rules

- **Testing Structure**: Unit tests must mirror the structure of the codebase. 
- This rule ensures coherence and facilitates easier maintenance.

## Consequences

Adopting these frameworks and rules not only standardizes our testing practices but also enhances our ability 
to maintain and navigate our test suite efficiently, improving overall code quality and development processes.
