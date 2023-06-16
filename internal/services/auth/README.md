## Auth service

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-permissions.md) - Implementing Permissions

### Architecture

#### Use case diagram

The use case diagram shows which functionality of the developed software system is 
available to each group of users.

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml

rectangle Auth {
  usecase (UC-1 Authenticate) as UC1
  usecase (UC-2 Log out) as UC2
  usecase (UC-3 Register) as UC3
  usecase (UC-4 Account recovery) as UC4
  
  url of UC1 is [[https://www.ory.sh/docs/oauth2-oidc/custom-login-consent/flow#sequence-diagram]]
  url of UC2 is [[https://www.ory.sh/docs/oauth2-oidc/oidc-logout#logout-logic-diagram]]
  url of UC3 is [[https://www.ory.sh/docs/kratos/self-service/flows/user-registration#registration-for-server-side-browser-clients]]
  url of UC4 is [[https://www.ory.sh/docs/kratos/self-service/flows/account-recovery-password-reset]]
}

customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
```

**Use cases**:

- [UC-1](https://www.ory.sh/docs/oauth2-oidc/custom-login-consent/flow#sequence-diagram) - Authenticate
- [UC-2](https://www.ory.sh/docs/oauth2-oidc/oidc-logout#logout-logic-diagram) - Log out
- [UC-3](https://www.ory.sh/docs/kratos/self-service/flows/user-registration#registration-for-server-side-browser-clients) - Register
- [UC-4](https://www.ory.sh/docs/kratos/self-service/flows/account-recovery-password-reset) - Account recovery
