# 4. Anti-Corruption Layer (ACL) для внешних сервисов

Date: 2024-12-19

## Status

Accepted

## Context

Proxy Service взаимодействует с внешним Link Service через gRPC/Connect протокол. Внешний сервис использует protobuf типы, которые не соответствуют доменным сущностям Proxy Service. Необходимо изолировать Domain Layer от деталей внешнего API.

Проблемы без ACL:
- Domain Layer зависит от protobuf типов внешнего сервиса
- Изменения в внешнем API требуют изменений в Domain Layer
- Сложно тестировать Domain Layer без реального внешнего сервиса
- Нарушение Dependency Inversion Principle

## Decision

Реализован Anti-Corruption Layer (ACL) для изоляции Domain Layer от внешнего Link Service.

### Структура

```
src/proxy/infrastructure/
├── anti-corruption/
│   └── LinkServiceACL.ts      # Преобразование protobuf <-> Domain
├── adapters/
│   └── LinkServiceConnectAdapter.ts  # Адаптер для Connect/gRPC
└── repositories/
    └── LinkServiceRepository.ts      # Repository реализация
```

### LinkServiceACL

ACL преобразует:
- **protobuf → Domain**: `LinkProto` → `Link` entity
- **Domain → protobuf**: `Hash` → `GetLinkByHashRequest`

ACL инкапсулирует:
- Знание о структуре protobuf сообщений
- Маппинг полей между внешним и доменным представлением
- Обработку edge cases (null, undefined, default values)

### Слои взаимодействия

1. **Domain Layer** - работает только с доменными типами (`Link`, `Hash`)
2. **Repository Interface** (`ILinkRepository`) - определяет контракт в терминах Domain
3. **ACL** (`LinkServiceACL`) - преобразует Domain ↔ protobuf
4. **Adapter** (`LinkServiceConnectAdapter`) - выполняет gRPC вызовы
5. **External Service** - Link Service через gRPC

### Пример использования

```typescript
// Repository использует ACL для преобразования
class LinkServiceRepository {
  async findByHash(hash: Hash): Promise<Link | null> {
    const protoRequest = this.acl.toProtoRequest(hash);
    const protoResponse = await this.adapter.getLinkByHash(protoRequest);
    return this.acl.toDomainEntity(protoResponse);
  }
}
```

## Consequences

### Положительные

- **Изоляция Domain** - Domain Layer не знает о protobuf
- **Независимость от внешнего API** - изменения в Link Service не затрагивают Domain
- **Тестируемость** - можно мокировать ACL и адаптер для тестов
- **Гибкость** - легко заменить внешний сервис или протокол
- **Единая точка преобразования** - вся логика маппинга в одном месте

### Отрицательные

- **Дополнительный слой** - больше кода для поддержки
- **Дублирование** - нужно поддерживать синхронизацию между Domain и protobuf моделями
- **Overhead** - дополнительные преобразования при каждом запросе

### Риски и митигация

**Риск**: Несинхронизированные модели между Domain и protobuf
- **Митигация**: Тесты проверяют преобразования, типизация TypeScript помогает

**Риск**: Потеря данных при преобразовании
- **Митигация**: ACL логирует преобразования, тесты покрывают edge cases

## Implementation Details

### LinkServiceACL

Методы преобразования:
- `toProtoRequest(hash: Hash): GetLinkByHashRequest` - Domain → protobuf request
- `toDomainEntity(proto: LinkProto): Link` - protobuf response → Domain
- Обработка `null` и `undefined` значений
- Валидация данных при преобразовании

### LinkMapper

Отдельный класс для маппинга (может быть частью ACL):
- `fromProto(proto: LinkProto): Link`
- `toProto(link: Link): LinkProto`

## Alternatives Considered

### Альтернатива 1: Использовать protobuf типы напрямую в Domain

**Отклонено** - нарушает изоляцию Domain Layer, создает зависимость от внешнего сервиса

### Альтернатива 2: Преобразование в Repository

**Отклонено** - Repository должен быть простым, преобразование - это отдельная ответственность

### Альтернатива 3: Использовать DTOs вместо ACL

**Отклонено** - DTOs обычно для Application Layer, ACL специфичен для внешних интеграций

## References

- [Anti-Corruption Layer Pattern](https://docs.microsoft.com/en-us/azure/architecture/patterns/anti-corruption-layer)
- `src/proxy/infrastructure/anti-corruption/LinkServiceACL.ts` - реализация
- `src/proxy/infrastructure/repositories/LinkServiceRepository.ts` - использование
