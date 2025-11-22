# OpenAPI Documentation

OpenAPI спецификация для Proxy Service API генерируется из JSDoc комментариев в контроллерах.

## Генерация спецификации

```bash
pnpm docs:generate
```

Это создаст файл `openapi.json` в корне проекта с полной спецификацией API.

## Endpoints

### GET /s/{hash}
Редирект на оригинальный URL по короткой ссылке.

**Параметры:**
- `hash` (path) - хеш короткой ссылки

**Ответы:**
- `301` - Успешный редирект
- `400` - Невалидный формат хеша
- `404` - Ссылка не найдена
- `429` - Превышен лимит запросов
- `500` - Внутренняя ошибка сервера

### GET /ready
Health check endpoint для Kubernetes readiness/liveness probes.

**Ответы:**
- `200` - Сервис готов к работе
- `503` - Сервис не готов (graceful shutdown)

## Использование спецификации

Сгенерированный `openapi.json` можно использовать с:
- Swagger UI (локально или через онлайн редактор)
- Postman (импорт OpenAPI)
- Insomnia (импорт OpenAPI)
- Генерация клиентов (openapi-generator, swagger-codegen)

## Обновление документации

При изменении API endpoints обновите JSDoc комментарии с аннотациями `@swagger` в:
- `src/proxy/interfaces/http/controllers/ProxyController.ts`
- `src/infrastructure/health.ts`

Затем запустите `pnpm docs:generate` для обновления спецификации.

