# syntax=docker/dockerfile:1.3

FROM node:17.6-alpine as deps

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
RUN npm config set ignore-scripts false

WORKDIR /app
COPY ./ui/next ./
RUN npm install && npm build

# Setup unprivileged user 101
RUN chown -R 101 /app

# Use user 101
USER 101

CMD ["npm", "run", "start:prod"]
