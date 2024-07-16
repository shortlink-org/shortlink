# Global scope
ARG ENVIRONMENT_CONFIG=production

# Stage 1 - Create yarn install skeleton layer
FROM --platform=$BUILDPLATFORM node:21.7.3-bookworm-slim AS packages

ARG ENVIRONMENT_CONFIG

WORKDIR /app

COPY ./boundaries/platform/backstage/packages packages
COPY ./boundaries/platform/backstage/package.json ./boundaries/platform/backstage/yarn.lock ./

RUN find packages \! -name "package.json" -mindepth 2 -maxdepth 2 -print | xargs rm -rf

# Stage 2 - Install dependencies and build packages
FROM --platform=$BUILDPLATFORM node:21.7.3-bookworm-slim AS build

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

USER node
WORKDIR /app

COPY --from=packages --chown=node:node /app .

# Stop cypress from downloading it's massive binary.
ENV CYPRESS_INSTALL_BINARY=0
RUN --mount=type=cache,target=/home/node/.cache/yarn,sharing=locked,uid=1000,gid=1000 \
    yarn install --immutable

COPY --chown=node:node ./boundaries/platform/backstage .

RUN yarn tsc
RUN yarn --cwd packages/backend build

RUN mkdir packages/backend/dist/skeleton packages/backend/dist/bundle \
    && tar xzf packages/backend/dist/skeleton.tar.gz -C packages/backend/dist/skeleton \
    && tar xzf packages/backend/dist/bundle.tar.gz -C packages/backend/dist/bundle

# Stage 3 - Build the actual backend image and install production dependencies
FROM node:21.7.3-bookworm-slim

ARG ENVIRONMENT_CONFIG

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="Backstage"
LABEL org.opencontainers.image.description="Backstage"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# Set Python interpreter for `node-gyp` to use
ENV PYTHON /usr/bin/python3

# Install sqlite3 dependencies. You can skip this if you don't use sqlite3 in the image,
# in which case you should also move better-sqlite3 to "devDependencies" in package.json.
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && \
    apt-get install -y --no-install-recommends libsqlite3-dev g++ build-essential \
    python3 python3-pip python3-venv \
    curl default-jre graphviz fonts-dejavu fontconfig && \
    rm -rf /var/lib/apt/lists/* && \
    yarn config set python /usr/bin/python3

ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

RUN pip3 install mkdocs-techdocs-core mkdocs-kroki-plugin

RUN curl -o plantuml.jar -L https://github.com/plantuml/plantuml/releases/download/v1.2024.3/plantuml-1.2024.3.jar
RUN echo '#!/bin/sh\n\njava -jar '/opt/plantuml.jar' ${@}' >> /usr/local/bin/plantuml
RUN chmod 755 /usr/local/bin/plantuml

# From here on we use the least-privileged `node` user to run the backend.
USER node

# This should create the app dir AS `node`.
# If it is instead created AS `root` then the `tar` command below will fail: `can't create directory 'packages/': Permission denied`.
# If this occurs, then ensure BuildKit is enabled (`DOCKER_BUILDKIT=1`) so the app dir is correctly created AS `node`.
WORKDIR /app

# Copy the install dependencies from the build stage and context
COPY --from=build --chown=node:node /app/.yarn ./.yarn
COPY --from=build --chown=node:node /app/.yarnrc.yml  ./
COPY --from=build --chown=node:node /app/yarn.lock /app/package.json /app/packages/backend/dist/skeleton/ ./

RUN --mount=type=cache,target=/home/node/.cache/yarn,sharing=locked,uid=1000,gid=1000 \
    yarn workspaces focus --all --production

# Copy the built packages from the build stage
COPY --from=build --chown=node:node /app/packages/backend/dist/bundle/ ./

COPY --chown=node:node \
  ./boundaries/platform/backstage/app-config.yaml \
  ./boundaries/platform/backstage/app-config.production.yaml \
  ./boundaries/platform/backstage/shortlink-org ./

# Heroku will assign the port dynamically; the default value here will be overridden by what Heroku passes in
# For local development the default will be used
ENV PORT 7007
# This switches many Node.js dependencies to production mode.
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

# Sets the max memory size of V8's old memory section
# Also disables node snapshot for Node 20 to work with the Scaffolder
ENV NODE_OPTIONS "--max-old-space-size=1000 --no-node-snapshot"

# Default is 'production', for local testing pass in 'local'
ENV ENVIRONMENT_CONFIG=${ENVIRONMENT_CONFIG}

CMD ["sh", "-c", "node packages/backend --config app-config.yaml --config app-config.${ENVIRONMENT_CONFIG}.yaml"]
