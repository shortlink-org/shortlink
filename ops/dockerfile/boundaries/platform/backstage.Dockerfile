# Global scope
ARG ENVIRONMENT_CONFIG=production

# Stage 1 - Create yarn install skeleton layer
FROM --platform=$BUILDPLATFORM node:22.6.0-bookworm-slim AS packages

ARG ENVIRONMENT_CONFIG

WORKDIR /app

COPY ./boundaries/platform/backstage/packages packages
COPY ./boundaries/platform/backstage/package.json ./boundaries/platform/backstage/yarn.lock ./

RUN find packages \! -name "package.json" -mindepth 2 -maxdepth 2 -print | xargs rm -rf

# Stage 2 - Install dependencies and build packages
FROM --platform=$BUILDPLATFORM node:22.6.0-bookworm-slim AS build

ARG ENVIRONMENT_CONFIG

# Set Python interpreter for `node-gyp` to use
ENV PYTHON /usr/bin/python3

# Install sqlite3 dependencies. You can skip this if you don't use sqlite3 in the image,
# in which case you should also move better-sqlite3 to "devDependencies" in package.json.
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && \
    apt-get install -y --no-install-recommends libsqlite3-dev python3 g++ build-essential && \
    rm -rf /var/lib/apt/lists/*

# enable yarn
RUN corepack enable && yarn set version stable

USER node
WORKDIR /app

COPY --from=packages --chown=node:node /app .

# Stop cypress from downloading it's massive binary.
ENV CYPRESS_INSTALL_BINARY=0
RUN yarn install --immutable

COPY --chown=node:node ./boundaries/platform/backstage .

RUN npx tsc
RUN yarn --cwd packages/backend build

RUN mkdir packages/backend/dist/skeleton packages/backend/dist/bundle \
    && tar xzf packages/backend/dist/skeleton.tar.gz -C packages/backend/dist/skeleton \
    && tar xzf packages/backend/dist/bundle.tar.gz -C packages/backend/dist/bundle
