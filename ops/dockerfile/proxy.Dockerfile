# syntax=docker/dockerfile:1.3

FROM node:16.7-alpine as builder

# WARNING: if container limit < MAX_OLD_SPACE_SIZE => Killed
# Docs: https://developer.ibm.com/languages/node-js/articles/nodejs-memory-management-in-container-environments/
ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}

WORKDIR /app
COPY ./internal/services/proxy /app/

RUN npm i

CMD ["npm", "run", "prod"]
