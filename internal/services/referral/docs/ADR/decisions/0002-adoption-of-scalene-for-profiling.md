# 2. Adoption of Scalene for Profiling

Date: 2023-09-12

## Status

Accepted

## Context

Our project requires a performance profiling tool to measure and communicate the efficiency of our codebase. 
Scalene has been considered as a potential solution.

## Decision

[Scalene](https://github.com/plasma-umass/scalene) is a high-performance, high-precision CPU, GPU, and memory profiler for Python.

| 👍 Pros                                                         | 👎 Cons                                                       |
|-----------------------------------------------------------------|---------------------------------------------------------------|
| Scalene provides high-precision CPU, GPU, and memory profiling. | Potential complications when run within virtual environments. |
| Minimal overhead.                                               |                                                               |
| Line-by-line performance metrics.                               |                                                               |
| Can run without modifications on existing programs.             |                                                               |


Given the detailed and accurate profiling benefits, we've decided to integrate Scalene into our development workflow.

## Consequences

Scalene will be integrated into our development workflow.
