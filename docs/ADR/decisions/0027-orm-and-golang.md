# 27. ORM and Golang

Date: 2023-10-26

## Status

Accepted

## Context

The need to find an effective Object Relational Mapping (ORM) solution for our Golang project has led to the evaluation 
of various ORM tools. I tried, `ent` was utilized, but it presented certain limitations, including the necessity 
for use a custom driver for database interactions and the lack of support for OpenTelemetry which is crucial 
for ensuring observability within our systems.

Currently, our setup incorporates:

1. `pgxpool` for interacting with PostgreSQL, acknowledged for its efficiency and popularity within the community.
2. `squirrel` for handling SQL-like databases, appreciated for its fluent SQL builder.
3. `protoc-gen-go-orm` for generating ORM-like structures based on protobuf definitions.
4. `sqlc` for generating type-safe Go from SQL.

Our `protoc-gen-go-orm` goes a step further in optimizing database interactions by automatically generating filter code from 
Protobuf types. This zero-config/coding approach greatly reduces boilerplate and accelerates development, 
as developers do not need to manually write filter logic for each Protobuf type. 
The plugin intelligently creates filters for both PostgreSQL and MongoDB, ensuring compatibility 
and efficiency across different database systems.

Recent discoveries have introduced [Bun](https://bun.uptrace.dev/) as potential alternatives. 
Bun is commended for its SQL building capabilities and is particularly favored when working with PostgreSQL. 

The aim is to find an ORM solution that not only addresses the limitations encountered with `ent` but also aligns 
with our project's technical needs and long-term sustainability. The decision to explore alternative ORM tools or 
libraries arises from the desire to improve database interaction, enhance observability through OpenTelemetry support, 
and ensure ease of integration with our existing systems.

## Decision

After evaluating the limitations of `ent`, the current setup, the initial results from the custom ORM attempt, 
and reviewing `Bun` and `sqlc`, we have decided to continue the exploration of alternative ORM tools or libraries 
that seamlessly integrate with Golang and address the highlighted issues.

## Comparison Table

| Feature                    | ent        | Bun           | sqlc          | go-orm (Custom) |
|----------------------------|------------|---------------|---------------|-----------------|
| Custom Driver Support      | Yes        | pgx           | No (assumed)  | pgx             |
| OpenTelemetry Support      | No         | Yes (assumed) | Yes (assumed) | No              |
| Ease of Use                | Low        | High          | High          | Medium          |
| Performance                | High       | High          | High          | Medium          |
| Community Support          | Medium     | Medium        | High          | Low             |
| SQL Building Capabilities  | Good       | Excellent     | Excellent     | Good            |
| Database Integration       | Excellent  | Excellent     | Good          | Good            |
| License                    | Apache 2.0 | BSD 2-clause  | MIT           | MIT             |

### Sources and Notes:

- Ent's ease of use and code generation features are appreciated by the community, alongside its capabilities to handle complex relationships and automatic migrations, though it's noted to have a smaller community compared to some other ORMs
- Bun is praised for SQL building, striking a good balance between a full ORM and a query builder, and is particularly favored for use with PostgreSQL
- The 'Ease of Use', 'Performance', 'Community Support', and 'SQL Building Capabilities' columns are subjective and should be validated with further testing and community feedback.

Further research and comparisons among various ORM tools can be found on [ossinsight.io collections](https://ossinsight.io/collections/golang-orm/).

## Consequences

With this change:

1. **Ease of Integration:** Finding an ORM tool that aligns with our requirements will ease the integration process, reducing the time and effort required to adapt to a new tool.
2. **Observability:** OpenTelemetry support is crucial for monitoring and tracing our database operations; hence, a tool that supports it will significantly enhance our observability capabilities.
3. **Performance:** A potential risk could be the performance trade-offs that might come with a new ORM tool, especially if it's less performant than `squirrel`. Mitigating this risk requires thorough testing and benchmarking to ensure the new tool meets our performance standards.
4. **Learning Curve:** There might be a learning curve for the development team if the new ORM tool has a different API or requires new skills to effectively utilize it. This could temporarily affect the development velocity, which needs to be accounted for in our planning.
