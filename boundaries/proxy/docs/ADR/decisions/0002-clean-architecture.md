# 2. Clean Architecture с разделением на слои

Date: 2024-12-19

## Status

Accepted

## Context

Proxy Service должен быть легко тестируемым, поддерживаемым и независимым от деталей реализации (БД, внешние сервисы, фреймворки). Необходимо четкое разделение ответственности между бизнес-логикой и инфраструктурой.

## Decision

Применена Clean Architecture с явным разделением на слои:

1. **Domain Layer** - ядро приложения, бизнес-логика без зависимостей
   - Entities (Hash, Link)
   - Value Objects
   - Domain Services (LinkDomainService)
   - Domain Events
   - Repository Interfaces (ILinkRepository)

2. **Application Layer** - use cases и оркестрация
   - Use Cases (GetLinkByHashUseCase, PublishEventUseCase)
   - Application Services (LinkApplicationService)
   - DTOs (Application DTOs)
   - Pipeline & Interceptors (UseCasePipeline, LoggingInterceptor, MetricsInterceptor, AuthorizationInterceptor)
   - Event Handlers (EventDispatcher, IEventHandler)

3. **Infrastructure Layer** - детали реализации
   - Repository Implementations (LinkServiceRepository)
   - Adapters (LinkServiceConnectAdapter)
   - Anti-Corruption Layer (LinkServiceACL)
   - Messaging (AMQPEventPublisher)
   - Logging (WinstonLogger)
   - Metrics & Tracing (OpenTelemetryGrpcMetrics, OpenTelemetryGrpcTracing)
   - Configuration (AppConfig, ExternalServicesConfig)

4. **Interfaces Layer** - точки входа
   - HTTP Controllers (ProxyController)
   - HTTP Middleware (validationMiddleware, errorHandler)
   - Swagger/OpenAPI

### Структура директорий

```
src/
├── proxy/
│   ├── domain/              # Domain Layer
│   │   ├── entities/
│   │   ├── value-objects/
│   │   ├── services/
│   │   ├── events/
│   │   └── repositories/
│   ├── application/         # Application Layer
│   │   ├── use-cases/
│   │   ├── services/
│   │   ├── pipeline/
│   │   └── handlers/
│   └── infrastructure/      # Infrastructure Layer
│       ├── repositories/
│       ├── adapters/
│       ├── anti-corruption/
│       ├── messaging/
│       ├── metrics/
│       └── tracing/
├── infrastructure/          # Shared Infrastructure
│   ├── logging/
│   ├── config/
│   ├── telemetry.ts
│   └── health.ts
└── interfaces/              # Interfaces Layer
    └── http/
```

### Правила зависимостей

- **Domain** не зависит ни от чего
- **Application** зависит только от Domain
- **Infrastructure** зависит от Domain и Application
- **Interfaces** зависит от Application и Infrastructure

## Consequences

### Положительные

- **Независимость от фреймворков** - Domain и Application слои не знают о Express, Inversify, Winston
- **Тестируемость** - Domain и Application легко тестировать без инфраструктуры
- **Гибкость** - можно заменить любую реализацию в Infrastructure без изменения бизнес-логики
- **Явные зависимости** - через Dependency Injection (Inversify) все зависимости явны
- **Соответствие SOLID** - Single Responsibility, Dependency Inversion

### Отрицательные

- **Больше кода** - требуется больше абстракций и интерфейсов
- **Сложность навигации** - нужно понимать структуру слоев
- **Overhead** - дополнительные уровни абстракции

### Риски и митигация

**Риск**: Избыточная абстракция для простых случаев
- **Митигация**: Используем только там, где есть реальная необходимость в гибкости

**Риск**: Нарушение правил зависимостей
- **Митигация**: ESLint правила для проверки зависимостей между слоями (планируется)

## Implementation Details

### Dependency Injection

Используется Inversify для управления зависимостями:
- Все зависимости регистрируются в `inversify.config.ts`
- Символы типов в `types.ts` для избежания строковых зависимостей
- Контейнер создается один раз при старте приложения

### Domain Layer

- Entities содержат только бизнес-логику
- Value Objects инкапсулируют валидацию
- Repository Interfaces определяют контракты без деталей реализации

### Application Layer

- Use Cases представляют одну бизнес-операцию
- Application Services оркестрируют несколько Use Cases
- Pipeline с Interceptors для cross-cutting concerns
- Event Handlers для асинхронной обработки событий

## Alternatives Considered

### Альтернатива 1: Монолитная структура

**Отклонено** - сложно тестировать, тесная связанность

### Альтернатива 2: MVC без слоев

**Отклонено** - бизнес-логика смешивается с инфраструктурой

### Альтернатива 3: Hexagonal Architecture

**Рассмотрено** - очень похоже, но Clean Architecture более явно разделяет слои

## References

- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- `src/proxy/domain/` - Domain Layer
- `src/proxy/application/` - Application Layer
- `src/proxy/infrastructure/` - Infrastructure Layer
- `src/interfaces/` - Interfaces Layer
