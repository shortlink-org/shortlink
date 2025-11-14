# 3. Use Case Pipeline с Interceptors

Date: 2024-12-19

## Status

Accepted

## Context

Необходимо централизованно обрабатывать cross-cutting concerns для всех Use Cases:
- Логирование выполнения use cases
- Метрики (latency, success/error rates)
- Авторизация и проверка прав доступа
- Обработка ошибок

Ранее эта логика была размазана по коду, что приводило к дублированию и сложности поддержки.

## Decision

Реализован Use Case Pipeline с системой Interceptors для централизованной обработки cross-cutting concerns.

### Структура

```
src/proxy/application/pipeline/
├── UseCasePipeline.ts           # Основной pipeline
├── IUseCaseInterceptor.ts       # Интерфейс interceptor
├── LoggingInterceptor.ts        # Логирование use cases
├── MetricsInterceptor.ts        # Метрики use cases
├── AuthorizationInterceptor.ts  # Авторизация use cases
└── index.ts                     # Экспорты
```

### Порядок выполнения Interceptors

Interceptors выполняются в следующем порядке:

1. **AuthorizationInterceptor** (первый) - проверка прав доступа
2. **MetricsInterceptor** - запись метрик до выполнения
3. **LoggingInterceptor** - логирование начала выполнения
4. **Use Case Execution** - выполнение самого use case
5. **LoggingInterceptor** - логирование результата
6. **MetricsInterceptor** - запись метрик после выполнения

### UseCasePipeline

Pipeline оборачивает выполнение use case и применяет все зарегистрированные interceptors:

```typescript
const result = await pipeline.execute(useCase, request, context);
```

### Interceptor Lifecycle

Каждый interceptor имеет три точки входа:

1. **beforeExecute** - перед выполнением use case
2. **afterExecute** - после успешного выполнения
3. **onError** - при ошибке выполнения

### Контекст выполнения

UseCaseExecutionContext содержит:
- Request данные
- Metadata (для передачи данных между interceptors)
- Timestamp начала выполнения
- Use Case информация

## Consequences

### Положительные

- **Централизованная обработка** - вся cross-cutting логика в одном месте
- **Переиспользование** - interceptors применяются ко всем use cases автоматически
- **Гибкость** - легко добавить новые interceptors
- **Тестируемость** - interceptors можно тестировать отдельно
- **Соответствие принципам** - Single Responsibility, Separation of Concerns
- **Автоматизация** - логирование, метрики, авторизация работают автоматически

### Отрицательные

- **Дополнительная сложность** - нужно понимать порядок выполнения
- **Overhead** - interceptors добавляют небольшие накладные расходы
- **Отладка** - может быть сложнее отследить выполнение через несколько слоев

### Риски и митигация

**Риск**: Неправильный порядок interceptors может привести к неожиданному поведению
- **Митигация**: Порядок документирован и зафиксирован в коде

**Риск**: Interceptors могут скрывать ошибки
- **Митигация**: Все ошибки логируются, можно отключить interceptors для отладки

## Implementation Details

### LoggingInterceptor

Логирует:
- Начало выполнения use case: useCaseName, request
- Успешное завершение: useCaseName, duration, response
- Ошибка: useCaseName, duration, error

### MetricsInterceptor

Записывает метрики через OpenTelemetry:
- `use_case_requests_total` - счетчик запросов
- `use_case_duration_ms` - длительность выполнения
- `use_case_errors_total` - счетчик ошибок

### AuthorizationInterceptor

Проверяет права доступа через `IAuthorizationChecker`:
- Вызывает `checkAuthorization(useCase, request)`
- Если проверка не пройдена - выбрасывает `AuthorizationError`
- Можно отключить для определенных use cases через конфигурацию

## Alternatives Considered

### Альтернатива 1: Декораторы TypeScript

**Отклонено** - декораторы не подходят для async функций и сложнее тестировать

### Альтернатива 2: Middleware в контроллерах

**Отклонено** - дублирование кода, не применяется к use cases напрямую

### Альтернатива 3: AOP (Aspect-Oriented Programming)

**Отклонено** - избыточная сложность для текущих требований

## References

- `src/proxy/application/pipeline/` - реализация
- `src/proxy/application/services/LinkApplicationService.ts` - использование
