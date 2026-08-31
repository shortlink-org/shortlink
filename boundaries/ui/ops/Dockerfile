# syntax=docker/dockerfile:1.26

ARG NODE_VERSION=24.15.0-alpine

# ---------------------------------------------------------------------------
# Build
# ---------------------------------------------------------------------------
FROM --platform=$BUILDPLATFORM node:${NODE_VERSION} AS builder

ARG CI_COMMIT_SHA
ARG CI_COMMIT_TAG
ARG BUILD_DATE

ARG NEXT_PUBLIC_CLOUDFLARE_SITE_KEY
ENV NEXT_PUBLIC_CLOUDFLARE_SITE_KEY=${NEXT_PUBLIC_CLOUDFLARE_SITE_KEY}

ARG NEXT_PUBLIC_GOOGLE_ANALYTICS_ID
ENV NEXT_PUBLIC_GOOGLE_ANALYTICS_ID=${NEXT_PUBLIC_GOOGLE_ANALYTICS_ID}

ENV PNPM_HOME=/pnpm \
    PATH="/pnpm:$PATH" \
    NPM_CONFIG_PROGRESS=false \
    NPM_CONFIG_AUDIT=false \
    NPM_CONFIG_FUND=false \
    NEXT_TELEMETRY_DISABLED=1

RUN corepack enable

WORKDIR /app

RUN echo '@shortlink-org:registry=https://gitlab.com/api/v4/packages/npm/' > .npmrc \
 && echo 'registry=https://registry.npmjs.org/' >> .npmrc

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store \
    pnpm install --frozen-lockfile

COPY . .

# `public/` is not tracked in the repo and is only materialised by the PWA plugin
# when it is enabled; the runtime stage copies it unconditionally, so guarantee it exists.
RUN mkdir -p public

RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store \
    pnpm run build

# ---------------------------------------------------------------------------
# Runtime — Next.js standalone
# ---------------------------------------------------------------------------
FROM node:${NODE_VERSION} AS runtime

ARG CI_COMMIT_SHA
ARG CI_COMMIT_TAG
ARG BUILD_DATE

ENV NODE_ENV=production \
    PORT=8080 \
    HOSTNAME=0.0.0.0 \
    NEXT_TELEMETRY_DISABLED=1

LABEL org.opencontainers.image.title="link-ui" \
      org.opencontainers.image.description="Shortlink UI (Next.js standalone)" \
      org.opencontainers.image.url="https://shortlink.best/next" \
      org.opencontainers.image.source="https://github.com/shortlink-org/shortlink" \
      org.opencontainers.image.licenses="MIT" \
      org.opencontainers.image.vendor="Shortlink" \
      org.opencontainers.image.authors="Viktor Login <batazor111@gmail.com>" \
      org.opencontainers.image.version="${CI_COMMIT_TAG:-${CI_COMMIT_SHA}}" \
      org.opencontainers.image.revision="${CI_COMMIT_SHA}" \
      org.opencontainers.image.created="${BUILD_DATE}"

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder --chown=node:node /app/public ./public
COPY --from=builder --chown=node:node /app/.next/standalone ./
COPY --from=builder --chown=node:node /app/.next/static ./.next/static

USER node

HEALTHCHECK --interval=10s --timeout=5s --retries=3 \
  CMD curl -fsS http://127.0.0.1:8080/next >/dev/null || exit 1

EXPOSE 8080

CMD ["node", "server.js"]
