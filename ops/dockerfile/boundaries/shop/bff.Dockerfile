# syntax=docker/dockerfile:1.20

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM node:24.11.1-alpine

# WARNING: if container limit < MAX_OLD_SPACE_SIZE => Killed
# Docs: https://developer.ibm.com/languages/node-js/articles/nodejs-memory-management-in-container-environments/
ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}

# This is the public node url of the wundergraph node you want to include in the generated client
ARG wg_public_node_url

# Defining environment =================================================================================================

# We set the public node url as an environment variable so the generated client points to the correct url
# See for node options a https://docs.wundergraph.com/docs/wundergraph-config-ts-reference/configure-wundernode-options and
# for server options https://docs.wundergraph.com/docs/wundergraph-server-ts-reference/configure-wundergraph-server-options
ENV WG_PUBLIC_NODE_URL=${wg_public_node_url}
# Listen to all interfaces, 127.0.0.1 might produce errors with ipv6 dual stack
ENV WG_NODE_URL=http://127.0.0.1:9991
ENV WG_NODE_INTERNAL_URL=http://127.0.0.1:9993
ENV WG_NODE_HOST=0.0.0.0
ENV WG_NODE_PORT=9991
ENV WG_NODE_INTERNAL_PORT=9993
ENV WG_SERVER_URL=http://127.0.0.1:9992
ENV WG_SERVER_HOST=127.0.0.1
ENV WG_SERVER_PORT=9992

# Enable corepack and set pnpm home
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

COPY ./boundaries/shop/bff ./
# We place the binary in /usr/bin/wunderctl so we can find it without a relative path
ENV CI=true WG_COPY_BIN_PATH=/usr/bin/wunderctl
# Ensure you lock file is up to date otherwise the build will fail
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

RUN pnpm build

# Expose only the node, server is private
EXPOSE 9991

CMD npx wunderctl start --wundergraph-dir=.wundergraph
