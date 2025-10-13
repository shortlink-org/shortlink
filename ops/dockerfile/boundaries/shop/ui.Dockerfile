# syntax=docker/dockerfile:1.19

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

# Install dependencies only when needed
FROM --platform=$BUILDPLATFORM node:23.11.1-alpine AS development-builder

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-shop-ui"
LABEL org.opencontainers.image.description="shortlink-shop-ui"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# Defining environment
ARG API_URI
ARG CI_COMMIT_TAG
ARG CI_COMMIT_REF_NAME
ARG CI_PIPELINE_ID
ARG CI_PIPELINE_URL

# Set environment variables
ENV CI_COMMIT_TAG=$CI_COMMIT_TAG
ENV CI_COMMIT_REF_NAME=$CI_COMMIT_REF_NAME
ENV CI_PIPELINE_ID=$CI_PIPELINE_ID
ENV CI_PIPELINE_URL=$CI_PIPELINE_URL
ENV NEXT_TELEMETRY_DISABLED=1

# Enable corepack and set pnpm home
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

# Create the necessary directory for corepack
RUN mkdir -p /home/node/.cache/node/corepack/v1 \
    && chown -R node:node /home/node/.cache

RUN npm i -g next@15.0.0-rc.0

RUN corepack enable

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat

WORKDIR /app
RUN echo @shortlink-org:registry=https://gitlab.com/api/v4/packages/npm/ >> .npmrc

COPY ./boundaries/shop/ui ./

# version for npm: npm ci --cache .npm --prefer-offline --force
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

RUN pnpm build

RUN chown -R node:node /app

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3000 || exit 1

EXPOSE 3000

# server.js is created by next build from the standalone output
CMD ["next", "start", "-H", "0.0.0.0"]
