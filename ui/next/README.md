## NEXT-UI Application

This UI for shortlink application

#### Feature/Require

- NextJS (SSR/Static generate content)
- Auth (Kratos)
- Monitoring (Sentry)
- Pretty UI
  - TailwindCSS
  - Material-UI
- Pretty code base
  - Typescript
  - ESLint/Prettier
- Storybook (for UI)
- Cypress (for E2E)

### Getting start

```bash
npm ci
npm start

// Ready on http://127.0.0.1:3000/next/auth/login
```

### ENV

Use `.env` file in `ui/next` directories for setting your UI

| Name            | Value                   | Description                             |
| --------------- | ----------------------- | --------------------------------------- |
| `API_URI`       | `http://localhost:7070` | API port                                |
| `PROXY_URI`     | `http://localhost:3030` | Proxy service address                   |
| `SENTRY_ENABLE` | `false`                 | Init Sentry                             |
| `NODE_ENV`      |                         | Select: production, development, etc... |
| `SENTRY_DSN`    |                         | Your sentry DSN                         |

### Build docker image

```bash
docker buildx build -t next-ui -f ops/dockerfile/ui-next.Dockerfile .
```

### UI Screenshot

<details>

| Describe   | Screenshot                               |
| ---------- | ---------------------------------------- |
| Link Table | ![link table](docs/next-js-ui.png) |

</details>

## Storybook

```bash
npm run storybook
```
