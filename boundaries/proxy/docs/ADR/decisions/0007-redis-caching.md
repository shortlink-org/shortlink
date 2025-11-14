# 7. Redis кэширование для оптимизации производительности

Date: 2025-01-XX

## Status

Accepted

## Context

Proxy Service обрабатывает большое количество запросов на редирект ссылок. Каждый запрос требует обращения к внешнему Link Service через gRPC, что создает:

- **Задержку**: каждый запрос требует сетевого вызова к Link Service
- **Нагрузку**: Link Service обрабатывает множество запросов на одни и те же ссылки
- **Зависимость**: доступность Proxy Service зависит от доступности Link Service

Проблемы без кэширования:

- Высокая латентность для популярных ссылок
- Избыточная нагрузка на Link Service
- Нет защиты от временной недоступности Link Service
- Высокие затраты на сетевые вызовы

## Decision

Использовать Redis для кэширования результатов поиска ссылок с поддержкой положительного и отрицательного кэширования.

### Архитектура

**Слои кэширования:**

1. **RedisLinkCache** (`infrastructure/cache/`) - реализация кэша на базе ioredis
2. **ILinkCache** - интерфейс для абстракции от конкретной реализации
3. **Интеграция в LinkServiceRepository** - кэш проверяется перед обращением к адаптеру

**Стратегия кэширования:**

- **Положительный кэш**: кэшируются найденные ссылки с TTL (по умолчанию 1 час)
- **Отрицательный кэш**: кэшируются отсутствующие ссылки с меньшим TTL (по умолчанию 5 минут)
- **Graceful degradation**: при недоступности Redis приложение продолжает работать без кэша

### Формат ключей

```
${CACHE_KEY_PREFIX}:hash:${hash}
```

Пример: `shortlink:proxy:hash:abc123`

### Конфигурация

Переменные окружения:

- `CACHE_ENABLED` (default: false) - включить/выключить кэш
- `REDIS_URL` (default: redis://localhost:6379) - URL подключения к Redis
- `CACHE_TTL_POSITIVE` (default: 3600) - TTL для положительных результатов (секунды)
- `CACHE_TTL_NEGATIVE` (default: 300) - TTL для отрицательных результатов (секунды)
- `CACHE_KEY_PREFIX` (default: shortlink:proxy) - префикс для ключей кэша

### Реализация

**Поток запроса:**

1. `LinkServiceRepository.findByHash()` вызывается
2. Проверяется кэш через `cache.get()`
3. Если cache hit (positive) - возвращается ссылка из кэша
4. Если cache hit (negative) - возвращается null без обращения к адаптеру
5. Если cache miss - запрос к адаптеру
6. Результат сохраняется в кэш (positive или negative)

**Обработка ошибок:**

- Ошибки Redis логируются, но не прерывают работу приложения
- При недоступности Redis все запросы идут напрямую к адаптеру
- Поврежденные записи кэша автоматически очищаются

### Метрики

Собираются Prometheus метрики через OpenTelemetry Metrics API:

**cache_requests_total** (Counter)
- Описание: Общее количество запросов к кэшу
- Labels:
  - `operation`: тип операции (get, set, clear)
  - `type`: результат запроса (hit, miss, error)
  - `result`: для hit - тип результата (positive, negative)
  - `reason`: для miss - причина (unavailable)

**cache_duration_ms** (Histogram)
- Описание: Время выполнения операций кэша в миллисекундах
- Labels:
  - `operation`: тип операции (get, set, clear)
  - `type`: для set - тип кэша (positive, negative)

**cache_errors_total** (Counter)
- Описание: Количество ошибок кэша
- Labels:
  - `operation`: тип операции (get, set, clear)
  - `type`: для set - тип кэша (positive, negative)
  - `error_type`: тип ошибки (redis_error, parse_error)

**Примеры метрик:**

```
# Cache hit rate
rate(cache_requests_total{operation="get",type="hit"}[5m]) / rate(cache_requests_total{operation="get"}[5m])

# Cache error rate
rate(cache_errors_total[5m])

# Average cache operation duration
histogram_quantile(0.95, rate(cache_duration_ms_bucket[5m]))
```

## Consequences

### Положительные

- **Снижение латентности** - популярные ссылки возвращаются из кэша (<1ms вместо сетевого вызова)
- **Снижение нагрузки** - Link Service обрабатывает меньше запросов
- **Устойчивость** - graceful degradation при недоступности Redis
- **Масштабируемость** - Redis может использоваться несколькими инстансами Proxy Service
- **Мониторинг** - метрики позволяют отслеживать эффективность кэша

### Отрицательные

- **Дополнительная инфраструктура** - требуется Redis
- **Консистентность** - возможны устаревшие данные в кэше (TTL решает проблему)
- **Сложность** - дополнительный компонент для управления
- **Память** - Redis требует памяти для хранения кэша

### Риски и митигация

**Риск**: Redis недоступен, все запросы идут к Link Service

- **Митигация**: Graceful degradation - приложение продолжает работать без кэша

**Риск**: Устаревшие данные в кэше (ссылка удалена, но кэш еще валиден)

- **Митигация**: TTL ограничивает время жизни кэша, можно очистить через `clear()`

**Риск**: Redis переполняется памятью

- **Митигация**: TTL автоматически удаляет старые записи, можно настроить maxmemory policy в Redis

**Риск**: Неправильная конфигурация TTL приводит к неоптимальной производительности

- **Митигация**: Метрики позволяют отслеживать hit rate и настраивать TTL

## Implementation Details

### Файлы

- `src/infrastructure/config/CacheConfig.ts` - конфигурация кэша
- `src/proxy/infrastructure/cache/RedisLinkCache.ts` - реализация Redis кэша
- `src/proxy/infrastructure/cache/ILinkCache.ts` - интерфейс кэша (экспортируется из RedisLinkCache.ts)
- `src/proxy/infrastructure/repositories/LinkServiceRepository.ts` - интеграция кэша
- `src/inversify.config.ts` - DI bindings для кэша

### Зависимости

- `ioredis` - клиент Redis для Node.js
- `@opentelemetry/api` - для метрик

### Тесты

- Unit тесты: `src/proxy/infrastructure/cache/__tests__/RedisLinkCache.test.ts`
- Integration тесты: `src/proxy/infrastructure/cache/__tests__/RedisLinkCache.integration.test.ts`
- Обновлены тесты репозитория для проверки кэширования

## Alternatives Considered

### Альтернатива 1: In-memory кэш (Node.js Map)

**Отклонено** - не масштабируется между инстансами, теряется при перезапуске, ограничен памятью процесса

### Альтернатива 2: Memcached

**Отклонено** - Redis более функционален, уже используется в инфраструктуре, лучшая поддержка в Node.js

### Альтернатива 3: Кэширование на уровне HTTP (CDN)

**Отклонено** - не подходит для динамических редиректов, требует дополнительной инфраструктуры

### Альтернатива 4: Без кэширования

**Отклонено** - высокая латентность и нагрузка на Link Service, плохая производительность

## References

- [ioredis documentation](https://github.com/redis/ioredis)
- [Redis best practices](https://redis.io/docs/manual/patterns/)
- [OpenTelemetry Metrics API](https://opentelemetry.io/docs/specs/otel/metrics/api/)
- `src/proxy/infrastructure/cache/RedisLinkCache.ts` - реализация
- `src/infrastructure/config/CacheConfig.ts` - конфигурация

