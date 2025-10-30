# syntax=docker/dockerfile:1.19

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM node:25-alpine AS builder

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-proxy"
LABEL org.opencontainers.image.description="shortlink-proxy"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# WARNING: if container limit < MAX_OLD_SPACE_SIZE => Killed
# Docs: https://developer.ibm.com/languages/node-js/articles/nodejs-memory-management-in-container-environments/
ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}
ENV DEBUG=pyroscope

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN apk update && \
    apk add --no-cache curl tini && \
    npm install -g pnpm

RUN mkdir -p /home/node/.npm/_cacache

WORKDIR /app

# ⚡ Copy only package manifests first to improve caching
COPY package.json pnpm-lock.yaml ./

# version for npm: npm ci --cache .npm --prefer-offline --force
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

# 📦 Copy the rest of the application
COPY . .

# 🏗️ Build the application
RUN pnpm build

FROM node:25-alpine AS runtime

RUN apk add --no-cache tini curl \
    && npm install -g pnpm \
    && addgroup -S nodegroup \
    && adduser -S nodeuser -G nodegroup

ENV NODE_ENV=production
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules
COPY prisma ./prisma

USER nodeuser

ENTRYPOINT ["/sbin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3020/ready || exit 1

CMD ["pnpm", "prod"]
