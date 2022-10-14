## NEXT-UI Application

This UI for shortlink application

#### Feature/Require

- NextJS (SSR/Static generate content)
- Auth (Kratos)
- Monitoring (Sentry)
- Pretty UI
  - tailwind CSS
  - Material-UI
- Pretty code base
  - Typescript
  - ESLint/Prettier
- Storybook (for UI)

### Getting start

```bash
$> npm ci
$> npm start
$>
$> Ready on http://127.0.0.1:3000/next/auth/login
```

#### ENV

| Name            | Value                   | Description           |
| --------------- | ----------------------- | --------------------- |
| `API_URI`       | `http://localhost:7070` | API port              |
| `PROXY_URI`     | `http://localhost:3030` | Proxy service address |
| `SENTRY_ENABLE` | `false`                 | Init Sentry           |

### Storybook

```bash
$> npm run storybook
```

### Build docker image

```bash
$> docker buildx build -t next-ui -f ops/dockerfile/ui-next.Dockerfile .
```
