# 1. Connect Interceptors для cross-cutting concerns

Date: 2024-11-13

## Status

Accepted

## Context

После миграции на `@connectrpc/connect` через transport API возникла необходимость в централизованной обработке cross-cutting concerns для всех Connect вызовов:

- **Логирование**: необходимость логировать все Connect запросы и ответы для отладки
- **Retry logic**: автоматические повторные попытки при временных ошибках (сетевые проблемы, таймауты, rate limiting)
- **Метрики**: автоматическая запись метрик для мониторинга (success/error rates, latency)
- **Трейсинг**: автоматическое создание OpenTelemetry spans для distributed tracing

Ранее эта логика была размазана по коду адаптера (`LinkServiceConnectAdapter`), что приводило к:
- Дублированию кода
- Сложности поддержки
- Невозможности переиспользования логики для других адаптеров
- Нарушению Single Responsibility Principle

## Decision

Реализовать Connect Interceptors для централизованной обработки cross-cutting concerns:

1. **LoggingInterceptor** - логирование всех Connect запросов/ответов
2. **RetryInterceptor** - автоматические повторные попытки с exponential backoff
3. **MetricsInterceptor** - автоматическая запись метрик
4. **TracingInterceptor** - автоматическое создание OpenTelemetry spans

Interceptors применяются при создании Connect transport и автоматически обрабатывают все вызовы через `transport.unary()`.

### Структура

```
src/proxy/infrastructure/adapters/connect/interceptors/
├── LoggingInterceptor.ts    # Логирование запросов/ответов
├── RetryInterceptor.ts      # Retry logic с exponential backoff
├── MetricsInterceptor.ts    # Запись метрик
├── TracingInterceptor.ts    # OpenTelemetry трейсинг
└── index.ts                 # Экспорты
```

### Порядок выполнения

Порядок interceptors важен - они выполняются в обратном порядке (как middleware):

1. **TracingInterceptor** (внешний) - создает span в начале, завершает в конце
2. **MetricsInterceptor** - записывает метрики
3. **RetryInterceptor** - выполняет повторные попытки при ошибках
4. **LoggingInterceptor** (внутренний) - логирует каждый запрос/ответ

### Конфигурация

Retry конфигурация берется из `ExternalServicesConfig`:

```typescript
retryCount: ConfigReader.number("EXTERNAL_SERVICE_RETRY_COUNT", 3)
```

Можно переопределить через переменные окружения:
- `EXTERNAL_SERVICE_RETRY_COUNT` - количество повторных попыток (по умолчанию: 3)

### Retry логика

RetryInterceptor повторяет запросы при следующих ошибках:

- **HTTP статус коды**: 429 (Too Many Requests), 500-599 (Server Errors)
- **gRPC коды**: UNAVAILABLE (14), DEADLINE_EXCEEDED (4), RESOURCE_EXHAUSTED (8)
- **Сетевые ошибки**: ECONNREFUSED, NetworkError

Использует exponential backoff:
- Попытка 1 → Попытка 2: 100ms
- Попытка 2 → Попытка 3: 200ms
- Попытка 3 → Попытка 4: 400ms
- Максимум: 5000ms

## Consequences

### Положительные

- **Централизованная обработка** - вся cross-cutting логика в одном месте
- **Переиспользование** - interceptors можно использовать для любых Connect адаптеров
- **Упрощение кода адаптера** - код `LinkServiceConnectAdapter` сократился с ~150 строк до ~75 строк
- **Автоматизация** - логирование, метрики, трейсинг и retry работают автоматически
- **Настраиваемость** - конфигурация через переменные окружения
- **Тестируемость** - interceptors можно тестировать отдельно
- **Соответствие принципам** - Single Responsibility, Separation of Concerns

### Отрицательные

- **Дополнительная сложность** - нужно понимать порядок выполнения interceptors
- **Отладка** - ошибки могут быть скрыты retry логикой (но логируются)
- **Производительность** - interceptors добавляют небольшие overhead (минимальный)

### Риски и митигация

**Риск**: Retry может маскировать реальные проблемы
- **Митигация**: Все retry логируются, метрики записываются для каждой попытки

**Риск**: Неправильный порядок interceptors может привести к неожиданному поведению
- **Митигация**: Порядок документирован и зафиксирован в коде

**Риск**: Exponential backoff может увеличить latency при множественных ошибках
- **Митигация**: Максимальная задержка ограничена (5000ms), можно настроить через конфигурацию

## Implementation Details

### LoggingInterceptor

Логирует:
- Начало запроса: service, method, url
- Успешный ответ: service, method, duration, status
- Ошибка: service, method, duration, errorCode, errorMessage

### RetryInterceptor

Конфигурация:
```typescript
interface RetryConfig {
  maxAttempts: number;           // По умолчанию: 3
  initialDelayMs: number;       // По умолчанию: 100ms
  maxDelayMs: number;           // По умолчанию: 5000ms
  backoffMultiplier: number;    // По умолчанию: 2
  retryableStatusCodes?: number[]; // По умолчанию: [429, 500-599]
}
```

### MetricsInterceptor

Записывает метрики через `IGrpcMetrics`:
- `recordRequest(method, status)` - счетчик запросов
- `recordDuration(method, durationMs)` - длительность выполнения
- `recordError(method, errorCode)` - счетчик ошибок

### TracingInterceptor

Создает OpenTelemetry spans через `IGrpcTracing`:
- Атрибуты: `rpc.service`, `rpc.method`, `rpc.system`, `rpc.grpc.status_code`
- Завершает span с успешным результатом или ошибкой

## Alternatives Considered

### Альтернатива 1: Оставить логику в адаптере

**Отклонено** - приводит к дублированию кода и нарушению SRP

### Альтернатива 2: Использовать декораторы

**Отклонено** - TypeScript декораторы не подходят для async функций и не интегрируются с Connect API

### Альтернатива 3: Middleware pattern

**Отклонено** - Connect уже предоставляет interceptors, нет смысла изобретать велосипед

## References

- [Connect Interceptors Documentation](https://connectrpc.com/docs/node/interceptors)
- [OpenTelemetry RPC Semantic Conventions](https://opentelemetry.io/docs/specs/semconv/rpc/)
- `src/proxy/infrastructure/adapters/connect/interceptors/` - реализация
- `src/proxy/infrastructure/adapters/LinkServiceConnectAdapter.ts` - использование
