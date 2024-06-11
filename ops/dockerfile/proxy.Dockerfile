# syntax=docker/dockerfile:1.8

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM node:21.7.3-alpine AS builder

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
RUN corepack enable

# Install dependencies
RUN apk update && \
    apk add --no-cache curl tini

RUN mkdir -p /home/node/.npm/_cacache

WORKDIR /app
COPY ./boundaries/link/proxy /app/

# version for npm: npm ci --cache .npm --prefer-offline --force
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm build

ENTRYPOINT ["/sbin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3020/ready || exit 1

CMD ["pnpm", "prod"]
