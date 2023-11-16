# Stage 1 - Create yarn install skeleton layer
FROM --platform=$BUILDPLATFORM node:21.2-bookworm-slim AS packages

WORKDIR /app
COPY ./internal/services/backstage/package.json ./internal/services/backstage/yarn.lock ./

COPY ./internal/services/backstage/packages packages

RUN find packages \! -name "package.json" -mindepth 2 -maxdepth 2 -exec rm -rf {} \+

# Stage 2 - Install dependencies and build packages
FROM --platform=$BUILDPLATFORM node:21.2-bookworm-slim AS build

# install sqlite3 dependencies
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && \
    apt-get install -y --no-install-recommends libsqlite3-dev python3 build-essential && \
    yarn config set python /usr/bin/python3

WORKDIR /app

COPY --from=packages --chown=node:node /app .

# Stop cypress from downloading it's massive binary.
ENV CYPRESS_INSTALL_BINARY=0
RUN --mount=type=cache,target=/home/node/.cache/yarn,sharing=locked,uid=1000,gid=1000 \
    yarn install --frozen-lockfile --network-timeout 600000 --ignore-engines

COPY --chown=node:node ./internal/services/backstage .

RUN yarn tsc
RUN yarn --cwd packages/backend build

RUN mkdir packages/backend/dist/skeleton packages/backend/dist/bundle \
    && tar xzf packages/backend/dist/skeleton.tar.gz -C packages/backend/dist/skeleton \
    && tar xzf packages/backend/dist/bundle.tar.gz -C packages/backend/dist/bundle

# Stage 3 - Build the actual backend image and install production dependencies
FROM --platform=$TARGETPLATFORM node:21.2-bookworm-slim

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="Backstage"
LABEL org.opencontainers.image.description="Backstage"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"


# Install sqlite3 dependencies. You can skip this if you don't use sqlite3 in the image,
# in which case you should also move better-sqlite3 to "devDependencies" in package.json.
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && \
    apt-get install -y --no-install-recommends libsqlite3-dev python3 build-essential curl && \
    yarn config set python /usr/bin/python3 && npm install -g node-gyp

# From here on we use the least-privileged `node` user to run the backend.
USER node

# This should create the app dir AS `node`.
# If it is instead created AS `root` then the `tar` command below will fail: `can't create directory 'packages/': Permission denied`.
# If this occurs, then ensure BuildKit is enabled (`DOCKER_BUILDKIT=1`) so the app dir is correctly created AS `node`.
WORKDIR /app

# Copy the install dependencies from the build stage and context
COPY --from=build --chown=node:node /app/yarn.lock /app/package.json /app/packages/backend/dist/skeleton/ ./

RUN --mount=type=cache,target=/home/node/.cache/yarn,sharing=locked,uid=1000,gid=1000 \
    yarn install --frozen-lockfile --production --network-timeout 600000 --ignore-engines

# Copy the built packages from the build stage
COPY --from=build --chown=node:node /app/packages/backend/dist/bundle/ ./

COPY --chown=node:node \
  ./internal/services/backstage/app-config.yaml \
  ./internal/services/backstage/app-config.production.yaml \
  ./internal/services/backstage/shortlink-org ./

ENV PORT 7007
ENV NODE_ENV production

ENV GITHUB_PRODUCTION_CLIENT_ID ""
ENV GITHUB_PRODUCTION_CLIENT_SECRET ""

ENV GITHUB_DEVELOPMENT_CLIENT_ID ""
ENV GITHUB_DEVELOPMENT_CLIENT_SECRET ""

# For now we need to manually add these configs through environment variables but in the
# future, we should be able to fetch the frontend config from the backend somehow
#ENV APP_CONFIG_app_baseUrl "https://demo.backstage.io"
#ENV APP_CONFIG_backend_baseUrl "https://demo.backstage.io"
#ENV APP_CONFIG_auth_environment "production"
#ENV APP_CONFIG_backend_database_connection_host "localhost"
#ENV APP_CONFIG_backend_database_connection_port "5432"
ENV NODE_OPTIONS "--max-old-space-size=150"

CMD ["node", "packages/backend", "--config", "app-config.yaml", "--config", "app-config.production.yaml"]
