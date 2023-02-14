# 17. Profile-guided optimization

Date: 2023-02-15

## Status

Accepted

## Context

Profile-guided optimization (PGO), also known as feedback-directed optimization (FDO), is a compiler optimization 
technique that feeds information (a profile) from representative runs of the application back into to the compiler 
for the next build of the application, which uses that information to make more informed optimization decisions. 
For example, the compiler may decide to more aggressively inline functions which the profile indicates 
are called frequently.

In Go, the compiler uses CPU pprof profiles as the input profile, such as from runtime/pprof or net/http/pprof.

As of Go 1.20, benchmarks for a representative set of Go programs show that building with PGO improves performance 
by around 2-4%. We expect performance gains to generally increase over time as additional optimizations take 
advantage of PGO in future versions of Go.

## Decision

We will use PGO for our application.

### Pipeline

1. Build docker image withour PGO or with old PGO profile
2. Push to registry
3. Deploy application
4. Run job for get and save PGO profile
5. Merge PGO profile for application
6. Build docker image with new PGO profile
7. Push to registry
8. Deploy application

![Pipeline](./images/profile-guided-optimization.png)

## Consequences

We will get a little bit faster application.
