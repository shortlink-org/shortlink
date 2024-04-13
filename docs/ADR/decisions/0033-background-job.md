# 33. Background Job Processing

Date: 2024-04-13

## Status

Accepted

## Context

The need for efficient background job processing arises from requirements such as handling big data operations, managing long-running tasks, and automating push notifications, report generation, and email/newsletter dispatches. Additionally, caching strategies like invalidation, hot-cache management, and pre-caching require reliable scheduling and execution mechanisms.

## Decision

To address these needs, the following options have been considered for implementing background jobs:

- **Kubernetes Jobs**:
  - **CronJob**: For recurring tasks based on a time-based schedule. [More about Kubernetes CronJobs](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)
  - **Job**: For one-off tasks that need to be run to completion. [More about Kubernetes Jobs](https://kubernetes.io/docs/concepts/workloads/controllers/job/)

- **Temporal**: A workflow orchestration platform ideal for complex workflows and handling retries, failures, and state pass-through between tasks. [Learn more about Temporal](https://temporal.io/)

- **Standard Cron**: Utilizing the operating system's cron for straightforward, time-based job scheduling.

- **For Golang Applications**:
  - **Golang Cron Library**: A library suited for scheduling recurring tasks within Go applications. [Visit Golang Cron Library](https://github.com/robfig/cron)
  - **Go Ticker and Timer**: Utilizing native Go features for fine-grained time-based operations.

- **Message Queueing (MQ)**:
  - **Apache Pulsar**: For decoupling job requests from job execution, supporting event-driven architectures and ensuring scalable message processing. [Explore Apache Pulsar](https://pulsar.apache.org/)

## Consequences

| Technology             | Benefits                                                              | Challenges                                                                         |
|------------------------|-----------------------------------------------------------------------|------------------------------------------------------------------------------------|
| **Standard Cron**      | Simple and reliable for scheduled tasks.                              | Limited to single-node setups, lacks fault tolerance.                              |
| **Kubernetes CronJob** | Scales well with Kubernetes, handles cluster-wide tasks.              | More complex setup, requires Kubernetes management.                                |
| **Temporal**           | Provides robust handling of complex workflows, retries, and failures. | Higher learning curve, more resource-intensive.                                    |
| **Go Ticker/Timer**    | High precision, integrated directly in Go applications.               | Requires custom setup for distributed tasks, more manual management of task state. |
