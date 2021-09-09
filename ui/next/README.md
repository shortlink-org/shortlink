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

### Getting start

```
$> npm i
$> npm start
$>
$> Ready on http://localhost:3000/next/auth/login
```

#### ENV

| Name            | Value                   | Description                |
| --------------- | ----------------------- | -------------------------- |
| `API_URI`       | `http://localhost:7070` | API port                   |
| `PROXY_URI`     | `http://localhost:3030` | Proxy service address      |
| `SENTRY_ENABLE` | `false`                 | Init Sentry                |
