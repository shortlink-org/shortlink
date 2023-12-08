# syntax=docker/dockerfile:1.6

# Defining environment
ARG APP_ENV=development

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

# Install dependencies only when needed
FROM --platform=$BUILDPLATFORM node:21.4-alpine AS development-builder

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN npm config set ignore-scripts false

WORKDIR /app
RUN echo @shortlink-org:registry=https://gitlab.com/api/v4/packages/npm/ >> .npmrc

COPY ./ui/nx-monorepo/ ./

RUN npm ci --cache .npm --prefer-offline --force
RUN npx nx run @shortlink-org/ui-kit:build-storybook

FROM --platform=$BUILDPLATFORM development-builder AS cache

COPY --from=development-builder /app/packages/ui-kit/storybook-static /app/storybook-static

FROM --platform=$BUILDPLATFORM alpine:3.19 AS ci-builder
FROM --platform=$BUILDPLATFORM ${APP_ENV}-builder AS cache

COPY ./ui/nx-monorepo/packages/ui-kit/storybook-static /app/storybook-static

# Production image, copy all the files and run next
FROM --platform=$TARGETPLATFORM ghcr.io/nginxinc/nginx-unprivileged:1.25-alpine

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="ui-kit"
LABEL org.opencontainers.image.description="UI Kit"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

# Use root user to copy dist folder and modify user access to specific folder
USER root

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:8080 || exit 1

# Copy application and custom NGINX configuration
COPY ./ops/dockerfile/conf/ui.local /etc/nginx/conf.d/default.conf
COPY ./ops/docker-compose/gateway/nginx/conf/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/conf/templates /etc/nginx/template
COPY --from=cache /app/storybook-static ./

# Setup unprivileged user 1001
RUN chown -R 1001 /usr/share/nginx/html

# Use user 1001
USER 1001
