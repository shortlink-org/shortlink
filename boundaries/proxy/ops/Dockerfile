# syntax=docker/dockerfile:1.20

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

# 🧱 Using Debian-based image to support glibc binaries (Datadog, pprof, etc.)
FROM node:25.2.1-bookworm-slim AS builder

LABEL maintainer="batazor111@gmail.com"
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

# 📦 install dependencies
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends curl tini python3 make g++ && \
    npm install -g pnpm && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p /home/node/.npm/_cacache

WORKDIR /app

# ⚡ Copy only package manifests first to improve caching
COPY package.json pnpm-lock.yaml ./

# version for npm: npm ci --cache .npm --prefer-offline --force
RUN --mount=type=cache,id=pnpm,target=/pnpm/store \
    --mount=type=cache,target=/root/.npm \
    pnpm install --frozen-lockfile

# 📦 Copy app files
COPY . .

# 🏗️ Build the application
RUN pnpm build

# -----------------------------
# Runtime image
# -----------------------------
FROM node:25.2.1-bookworm-slim AS runtime

# Minimal system deps
RUN apt-get update -y && \
    apt-get install -y --no-install-recommends tini curl && \
    npm install -g pnpm@9 && \
    useradd -m -u 1001 nodeuser && \
    rm -rf /var/lib/apt/lists/*

ENV NODE_ENV=production
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/node_modules ./node_modules
USER nodeuser

ENTRYPOINT ["/usr/bin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3020/ready || exit 1

# Using Permissions API for Zero-Trust security
# Restricting access to filesystem and network
# Environment variables are controlled at the Kubernetes/Docker level
# Allow /etc read for node-gyp-build to detect system type (Alpine/Debian)
CMD ["node", "--permission", "--allow-fs-read=/app", "--allow-fs-read=/etc", "--allow-net", "./dist/main.js"]
