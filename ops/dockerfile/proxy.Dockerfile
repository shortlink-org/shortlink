# syntax=docker/dockerfile:1.5

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM node:19.6-alpine as builder

# WARNING: if container limit < MAX_OLD_SPACE_SIZE => Killed
# Docs: https://developer.ibm.com/languages/node-js/articles/nodejs-memory-management-in-container-environments/
ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl

WORKDIR /app
COPY ./internal/services/proxy /app/

RUN npm ci --cache .npm --prefer-offline --force
RUN npm run build

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:3020/ready || exit 1

CMD ["npm", "run", "prod"]
