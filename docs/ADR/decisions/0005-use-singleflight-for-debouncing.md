# 5. Use singleflight for debounce

Date: 2023-10-15

## Status

Accepted

## Context

When our system experiences a surge in duplicate requests, especially during cache misses, 
it leads to an unnecessary load, redundant computations, and potential bottlenecks. 
This not only impacts system performance but can also degrade the user experience due to increased latency.

## Decision

We've chosen to implement the `singleflight` package as it provides a proven solution to debounce redundant requests, 
ensuring that for any given function with a particular key, only one execution happens. 
This choice was made after evaluating other potential solutions, with `singleflight` emerging as the most effective and 
widely adopted in the industry.

Implementation involves:

1. Pinpointing areas with frequent duplicate requests.
2. Applying `singleflight.Group` structures therein.
3. Ensuring waiting requests get results from a single computation, optimizing system load.

### Available Packages for Different Languages

- **Go:** [link](golang.org/x/sync/singleflight) - A reliable and efficient package tailored for Go's concurrency model.
- **Rust:** [crate](https://docs.rs/singleflight/latest/singleflight/) - Offers similar debounce capabilities, optimized for Rust's memory safety guarantees.
- **Python:** [PyPI](https://pypi.org/project/singleflight/) - A lightweight Pythonic approach to handling redundant requests.

## Consequences

**Pros:**

- **Reduced Redundancy:** Minimizes resource wastage.
- **Improved Performance:** Enhances user experience with reduced latency.
- **Prevent Overload:** Shields the system from surges of duplicate requests.

**Cons:**

- **Potential Stale Data:** Risk of outdated data in long computations. 
  - **Mitigation:** Monitor computation durations, optimize where possible.
- **Locking Overhead:** Overhead introduced by locks and channels. 
  - **Mitigation:** Regularly profile and monitor system performance.

---

Post-implementation, we'll monitor the system's performance, specifically focusing on areas where singleflight was implemented. 
Any anomalies will trigger an analysis, and if necessary, a rollback plan is in place to revert the changes.
