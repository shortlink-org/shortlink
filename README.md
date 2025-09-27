<div align="center">

# shortlink 

ShortLink is an open-source educational project that provides a pretty user interface and respects GDPR. 

The goal of the project is to demonstrate the practical application of microservices architecture.

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/shortlink)](https://artifacthub.io/packages/search?repo=shortlink)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/shortlink-org/shortlink)](https://pkg.go.dev/mod/github.com/shortlink-org/shortlink)
[![codecov](https://codecov.io/gh/shortlink-org/shortlink/branch/main/graph/badge.svg?token=Wxz5bI4QzF)](https://codecov.io/gh/shortlink-org/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/shortlink-org/shortlink)](https://goreportcard.com/report/github.com/shortlink-org/shortlink)
[![Releases](https://img.shields.io/github/release-pre/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/blob/main/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)
[![StackShare](http://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/shortlink-org/shortlink)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=shield)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_shield)
[![DeepSource](https://app.deepsource.com/gh/shortlink-org/shortlink.svg/?label=active+issues&show_trend=true&token=DL-zlqtnyx6CvlHCroG0Jdx5)](https://app.deepsource.com/gh/shortlink-org/shortlink/)

<hr />

<div style="align-items: center; display: flex;">
  <a href="https://www.producthunt.com/posts/shortlink-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-shortlink&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=374140&theme=light" alt="ShortLink - Get&#0032;ready&#0032;to&#0032;share&#0032;your&#0032;links&#0032;with&#0032;ease&#0033; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>
  <img height="100px" src="https://slsa.dev/images/SLSA-Badge-full-level1.svg" alt="SLSA">
</div>

</div>
<hr />

> #### Contributing
>
> - [Getting Started](./CONTRIBUTING.md#getting-started)

### Architecture decision records (ADR)

> [!IMPORTANT]
> An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.
>
>+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)
>
> **Decisions:**
>  + [main decisions](./docs/ADR/README.md)
>  + [ops decisions](./ops/docs/ADR/README.md)
>
> Also, each boundary context and service has its own ADR. You can find them in the relevant sections.

### Core Packages

The ShortLink project provides a comprehensive set of reusable packages that implement modern software engineering practices and clean architecture principles.

#### Observability Package (`pkg/observability`)

Advanced observability capabilities with comprehensive tracing, metrics, and logging support:

- **Enhanced Tracing with FlightRecorder**: Implements Go 1.25 `trace.FlightRecorder` for "perfect tracing"
  - **Domain Layer**: Core business logic and interfaces (`pkg/observability/traicing/domain/`)
    - `recorder.go` - Core recorder interfaces and value objects
    - `errors.go` - Domain-specific error definitions
  - **Application Layer**: Use cases and service orchestration (`pkg/observability/traicing/application/`)
    - `service.go` - High-level recorder operations and business workflows
  - **Infrastructure Layer**: External integrations and implementations (`pkg/observability/traicing/infrastructure/`)
    - `recorder.go` - Go 1.25 FlightRecorder adapter
    - `repository.go` - File system and storage implementations
    - `events.go` - Logging and monitoring event handlers
  - **Factory Pattern**: Dependency injection and component wiring (`pkg/observability/traicing/factory.go`)

**Key Features:**
- **Clean Architecture**: Follows hexagonal architecture with clear separation of concerns
- **Perfect Tracing**: Continuous, low-overhead tracing with configurable rolling buffer (default: 1 minute, 3MB)
- **Professional Error Handling**: Comprehensive error types and validation
- **Thread-Safe Operations**: Mutex-protected operations with proper resource management
- **Multiple Storage Backends**: Filesystem and in-memory storage implementations
- **Event-Driven Architecture**: Composite event handling for logging, metrics, and monitoring
- **Dependency Injection**: Factory pattern for clean component assembly

**Configuration Options:**
```bash
FLIGHT_RECORDER_ENABLED=true     # Enable/disable flight recorder
FLIGHT_RECORDER_MIN_AGE=1m       # Minimum trace data retention
FLIGHT_RECORDER_MAX_BYTES=3145728 # Maximum buffer size (3MB)
```

**Usage Example:**
```go
// Create factory with configuration
config := traicing.DefaultFactoryConfig(logger)
factory, err := traicing.NewFactory(config)

// Create service with all dependencies wired
service, err := factory.CreateRecorderService()

// Start recording
err = service.StartRecording(ctx)

// Capture trace on significant events
traceID, err := service.CaptureTrace(ctx, "error_occurred")
```

For detailed implementation documentation, see [`FLIGHT_RECORDER_IMPLEMENTATION.md`](./FLIGHT_RECORDER_IMPLEMENTATION.md).

#### Additional Packages

- **Database (`pkg/db`)**: Multi-database support with connection pooling and health checks
- **Dependency Injection (`pkg/di`)**: Structured dependency injection with lifecycle management
- **HTTP Utilities (`pkg/http`)**: Middleware, routing, and server utilities
- **Concurrency (`pkg/concurrency`)**: Advanced concurrency patterns and utilities
- **Caching (`pkg/cache`)**: Multi-backend caching with TTL and eviction policies
- **Message Queue (`pkg/mq`)**: Multi-provider message queue abstractions
- **Pattern Implementations (`pkg/pattern`)**: Common software patterns (CQRS, Event Sourcing, etc.)

Each package follows clean architecture principles with proper separation of concerns, comprehensive testing, and professional documentation.

### License

> [!WARNING]
> 
> This project includes dependencies licensed under the GNU Lesser General Public License (LGPL). 
> Users must comply with LGPL terms when using or modifying these dependencies. 
> For detailed information on each LGPL library used in this project, please refer to the respective license documentation 
> included with each library. For comprehensive license compliance information, including dependencies and their licenses, 
> you can read more details in our FOSSA report.

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=large)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_large)

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/shortlink-org/shortlink&style=flat
