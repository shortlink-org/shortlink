# C4 Architecture Diagrams

C4-модель для визуализации архитектуры Proxy Service на разных уровнях детализации.

## Диаграммы

### Context Diagram (`context.puml`)

Высокоуровневое представление системы и её взаимодействия с внешними системами:
- **User** - пользователь, который переходит по короткой ссылке
- **Proxy Service** - основной сервис для редиректа
- **Link Service** - внешний gRPC сервис для управления ссылками
- **Message Bus** - AMQP брокер для публикации событий
- **Observability Stack** - OpenTelemetry, Prometheus, Grafana

### Container Diagram (`container.puml`)

Структура Proxy Service на уровне контейнеров:
- **Web Application** - Express.js HTTP сервер
- **Application Layer** - Use Cases, Application Services, Pipeline
- **Domain Layer** - Entities, Value Objects, Domain Services
- **Infrastructure Layer** - Repositories, Adapters, ACL, Messaging

### Component Diagram (`component.puml`)

Детальная архитектура Application и Infrastructure слоев:
- Компоненты Application Layer (Use Cases, Services, Pipeline, Interceptors)
- Компоненты Domain Layer (Entities, Value Objects, Events)
- Компоненты Infrastructure Layer (Repositories, Adapters, ACL, Messaging, Observability)

## Генерация диаграмм

### Требования

- [PlantUML](https://plantuml.com/) установлен локально или используется онлайн редактор
- Или используйте VS Code расширение "PlantUML"

### Локальная генерация

```bash
# Установить PlantUML (macOS)
brew install plantuml

# Генерация PNG
plantuml -tpng docs/c4/context.puml
plantuml -tpng docs/c4/container.puml
plantuml -tpng docs/c4/component.puml

# Генерация SVG
plantuml -tsvg docs/c4/*.puml
```

### Онлайн генерация

1. Откройте файл `.puml` в [PlantUML Online Editor](http://www.plantuml.com/plantuml/uml/)
2. Или используйте VS Code расширение "PlantUML" для предпросмотра

## Использование

Диаграммы используются для:
- **Документации архитектуры** - визуальное представление структуры системы
- **Онбординга новых разработчиков** - быстрое понимание архитектуры
- **Обсуждения архитектурных решений** - визуализация для обсуждений
- **ADR (Architecture Decision Records)** - иллюстрация архитектурных решений

## Обновление диаграмм

При изменении архитектуры обновите соответствующие диаграммы:
- **Context Diagram** - при добавлении/удалении внешних систем
- **Container Diagram** - при изменении структуры слоев
- **Component Diagram** - при добавлении/удалении компонентов

## Ссылки

- [C4 Model](https://c4model.com/)
- [C4-PlantUML](https://github.com/plantuml-stdlib/C4-PlantUML)
- [PlantUML Documentation](https://plantuml.com/)
