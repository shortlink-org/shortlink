# 5. Event-Driven Architecture с Notification Pattern

Date: 2024-12-19

## Status

Accepted

## Context

После выполнения редиректа необходимо асинхронно обработать события:
- Публикация события `LinkRedirected` в message bus (AMQP)
- Сбор статистики (через eBPF, но событие нужно для других систем)
- Возможные будущие обработчики (аналитика, уведомления)

Необходимо:
- Разделить синхронную обработку редиректа и асинхронную обработку событий
- Обеспечить возможность добавления новых обработчиков без изменения существующего кода
- Использовать Notification Pattern для реакции на события

## Decision

Реализована Event-Driven Architecture с использованием Notification Pattern через EventDispatcher.

### Структура

```
src/proxy/
├── domain/
│   └── events/
│       ├── LinkRedirectedEvent.ts    # Доменное событие
│       └── index.ts
└── application/
    ├── use-cases/
    │   └── PublishEventUseCase.ts    # Use case для публикации событий
    └── handlers/
        ├── EventDispatcher.ts        # Диспетчер событий
        ├── IEventHandler.ts          # Интерфейс обработчика
        └── LinkRedirectedEventHandler.ts  # Обработчик события
```

### Notification Pattern

EventDispatcher реализует Notification Pattern:
- **Регистрация обработчиков** - несколько обработчиков для одного события
- **Reaction Chain** - все обработчики выполняются асинхронно через `Promise.all`
- **Неблокирующее выполнение** - ошибки в одном обработчике не останавливают другие

### Поток событий

1. **Domain Event** создается в Domain Layer (`LinkRedirectedEvent`)
2. **Use Case** публикует событие через `PublishEventUseCase`
3. **EventDispatcher** получает событие и находит все зарегистрированные обработчики
4. **Event Handlers** выполняются параллельно:
   - `LinkRedirectedEventHandler` - публикует в AMQP
   - Будущие обработчики могут быть добавлены без изменения существующего кода

### Event Handlers

Каждый обработчик:
- Реализует `IEventHandler<TEvent>`
- Определяет `canHandle(event)` - может ли обработать событие
- Реализует `handle(event)` - логика обработки

## Consequences

### Положительные

- **Разделение ответственности** - синхронная и асинхронная логика разделены
- **Расширяемость** - легко добавить новые обработчики
- **Независимость** - обработчики не знают друг о друге
- **Тестируемость** - каждый обработчик можно тестировать отдельно
- **Отказоустойчивость** - ошибка в одном обработчике не влияет на другие
- **Соответствие принципам** - Open/Closed Principle, Single Responsibility

### Отрицательные

- **Сложность отладки** - асинхронное выполнение может усложнить отладку
- **Нет гарантии порядка** - обработчики выполняются параллельно
- **Нет транзакционности** - если нужна атомарность, нужны дополнительные механизмы

### Риски и митигация

**Риск**: Ошибка в обработчике может быть незамеченной
- **Митигация**: Все ошибки логируются, можно добавить метрики для мониторинга

**Риск**: Порядок выполнения обработчиков важен
- **Митигация**: Если порядок важен, использовать последовательное выполнение или приоритеты

## Implementation Details

### LinkRedirectedEvent

Доменное событие содержит:
- `link` - ссылка, на которую произошел редирект
- `timestamp` - время события
- `type` - тип события для диспетчеризации

### EventDispatcher

Методы:
- `register<TEvent>(eventType: string, handler: IEventHandler<TEvent>)` - регистрация обработчика
- `dispatch<TEvent>(event: TEvent)` - диспетчеризация события

### LinkRedirectedEventHandler

Обработчик:
- Публикует событие в AMQP через `IEventPublisher`
- Логирует публикацию
- Обрабатывает ошибки публикации

### PublishEventUseCase

Use case для публикации событий:
- Получает событие и EventDispatcher
- Вызывает `dispatch()` для диспетчеризации
- Не блокирует выполнение основного flow

## Alternatives Considered

### Альтернатива 1: Прямой вызов обработчиков в Use Case

**Отклонено** - нарушает Open/Closed Principle, сложно расширять

### Альтернатива 2: Использовать готовую библиотеку (EventEmitter)

**Отклонено** - избыточная сложность, нужен простой механизм

### Альтернатива 3: Message Queue напрямую из Use Case

**Отклонено** - смешивает синхронную и асинхронную логику, сложно тестировать

## References

- [Notification Pattern](https://martinfowler.com/eaaDev/Notification.html)
- [Domain Events Pattern](https://martinfowler.com/eaaDev/DomainEvent.html)
- `src/proxy/domain/events/` - доменные события
- `src/proxy/application/handlers/EventDispatcher.ts` - диспетчер
- `src/proxy/application/handlers/LinkRedirectedEventHandler.ts` - обработчик
