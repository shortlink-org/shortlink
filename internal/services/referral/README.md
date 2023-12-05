## Referral Service

> [!NOTE]
> This service is responsible for the referral program.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-adoption-of-scalene-for-profiling.md) - Adoption of Scalene for Profiling
- [ADR-0003](./docs/ADR/decisions/0003-adoption-of-allure-for-test-reporting.md) - Adoption of Allure for Test Reporting

**Functionality**:

  * CRUD operations for referral

### Architecture

We use the C4 model to describe architecture.

#### Container diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

!include boundaries/marketing.puml

customer --> marketingBoundary : uses
```

#### Use case diagram

The use case diagram shows which functionality of the developed software system is
available to each group of users.

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml
!include actors/manager.puml

rectangle Referral {
  usecase (UC-1 CRUD referral) as UC1
  usecase (UC-2 Use referral) as UC2
}

customer --> UC2
manager --> UC1
```

**Use cases**:

- [UC-1](src/usecases/crud_referral/README.md) - CRUD referral
- [UC-2](src/usecases/use_referral/README.md) - Use referral

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

#### Config

| Name           | Description   | Default value |
|----------------|---------------|---------------|
| DATABASE_URI   | Database URI  | -             |
