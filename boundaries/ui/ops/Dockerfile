# syntax=docker/dockerfile:1.20

###############################################################################
# Global BuildKit options (optional SBOM attestation―leave if you need them)
# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
###############################################################################
ARG BUILDKIT_SBOM_SCAN_STAGE=true
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true


###############################################################################
# Build-stage arguments
###############################################################################
ARG NODE_VERSION=24.13.0-alpine

########################
# 1️⃣  Builder stage
########################
FROM --platform=$BUILDPLATFORM node:${NODE_VERSION} AS builder

# CI metadata (build-time only)
ARG CI_COMMIT_SHA
ARG CI_COMMIT_TAG
ARG CI_PIPELINE_URL
ARG BUILD_DATE

# production build args
ARG NEXT_PUBLIC_CLOUDFLARE_SITE_KEY
ENV NEXT_PUBLIC_CLOUDFLARE_SITE_KEY=${NEXT_PUBLIC_CLOUDFLARE_SITE_KEY}

ARG NEXT_PUBLIC_GOOGLE_ANALYTICS_ID
ENV NEXT_PUBLIC_GOOGLE_ANALYTICS_ID=${NEXT_PUBLIC_GOOGLE_ANALYTICS_ID}

# pnpm only needed here
ENV PNPM_HOME="/pnpm"
ENV PATH="${PNPM_HOME}:${PATH}"
# Disable npm warnings about unknown env configs
ENV NPM_CONFIG_PROGRESS=false
ENV NPM_CONFIG_AUDIT=false
ENV NPM_CONFIG_FUND=false

RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

COPY . .

# Set development environment for build (stage can be overridden at runtime)
ENV NODE_ENV=development

# Create npmrc with proper configuration
RUN echo '@shortlink-org:registry=https://gitlab.com/api/v4/packages/npm/' > .npmrc \
 && echo 'registry=https://registry.npmjs.org/' >> .npmrc \
 && echo 'auto-install-peers=true' >> .npmrc \
 && echo 'strict-peer-dependencies=false' >> .npmrc \
 && echo 'prefer-offline=true' >> .npmrc

# Install & build with cache for pnpm store
RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store \
    npm install -g patch-package \
 && rm -rf node_modules package-lock.json \
 && pnpm install --config.frozen-lockfile=true --include=dev \
 && pnpm run build

########################
# 2️⃣  Runtime stage
########################
FROM ghcr.io/nginx/nginx-unprivileged:1.29-alpine-otel AS runtime

###############################################################################
# Build-time args (re-declared so we can inject them into labels)
###############################################################################
ARG CI_COMMIT_SHA
ARG CI_COMMIT_TAG
ARG BUILD_DATE

# Set production environment
ENV NODE_ENV=production

###############################################################################
# OCI labels (single instruction = single layer)
###############################################################################
LABEL \
  org.opencontainers.image.title="link-ui" \
  org.opencontainers.image.description="Shortlink UI" \
  org.opencontainers.image.url="https://shortlink.best/next" \
  org.opencontainers.image.source="https://github.com/shortlink-org/shortlink" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.vendor="Shortlink" \
  org.opencontainers.image.authors="Viktor Login <batazor111@gmail.com>" \
  org.opencontainers.image.version="${CI_COMMIT_TAG:-${CI_COMMIT_SHA}}" \
  org.opencontainers.image.revision="${CI_COMMIT_SHA}" \
  org.opencontainers.image.created="${BUILD_DATE}"

###############################################################################
# Layer that needs root (package install + removing default config)
###############################################################################
USER root
RUN apk add --no-cache curl \
 && rm /etc/nginx/conf.d/default.conf

###############################################################################
# NGINX configuration & static assets
###############################################################################
COPY ./ops/conf/ui.local   /etc/nginx/conf.d/default.conf
COPY ./ops/conf/nginx.conf /etc/nginx/nginx.conf

# Next.js export build from the builder stage
COPY --from=builder /app/out /usr/share/nginx/html

# Ensure correct permissions for the runtime user
RUN chown -R 101:0 /etc/nginx /usr/share/nginx/html

###############################################################################
# Health-check, port & runtime user
###############################################################################
# Use the image's built-in non-root user (UID 101)
USER 101

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
  CMD curl -f http://localhost:8080/ || exit 1

EXPOSE 8080
