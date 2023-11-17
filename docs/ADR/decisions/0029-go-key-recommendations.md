# 29. Go: Key Recommendations

Date: 2023-11-17

## Status

Accepted

## Context

In our quest to enhance the quality and reliability of our Go codebase, we identified the need for robust tools and guidelines. 
Effective linting, handling of nil pointers, and adherence to a consistent style guide are essential for maintaining code quality.

## Decision

1. **Adopt `golangci-lint`**: This comprehensive linter will be our primary tool for identifying and fixing issues in our Go code.

2. **Implement `nilaway` by Uber**: To specifically address nil pointer errors, 
we will utilize [Uber's `nilaway`](https://github.com/uber-go/nilaway).

3. **Follow the Uber Go Style Guide**: Our coding style will align with the [Uber Go Style Guide](https://github.com/uber-go/guide), 
ensuring consistency and adherence to industry best practices.

## Consequences

- **Easier Code Maintenance**: With `golangci-lint`, identifying and fixing stylistic and logical errors becomes more streamlined.
- **Reduced Nil Pointer Errors**: `nilaway` will significantly lower the risk of nil pointer dereferences, one of the common pitfalls in Go.
- **Consistent Coding Style**: By adhering to the Uber Go Style Guide, our codebase will be more uniform and easier to read, 
which is beneficial for both current and future team members.
