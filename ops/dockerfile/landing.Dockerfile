# syntax=docker/dockerfile:1.3

# Install dependencies only when needed
FROM node:17.5-alpine as builder

ENV PYTHONUNBUFFERED=1

RUN apk add --no-cache alpine-sdk python3 libsass \
  && ln -sf python3 /usr/bin/python

WORKDIR /app
COPY ./ui/landing /app/

RUN npm i --force && \
  npm rebuild node-sass && \
  npm run generate

# Production image, copy all the files
FROM nginxinc/nginx-unprivileged:1.21-alpine

WORKDIR /usr/share/nginx/html

# Use root user to copy dist folder and modify user access to specific folder
USER root

COPY --from=builder /app/dist ./

# Setup unprivileged user 101
RUN chown -R 101 /usr/share/nginx/html

# Use user 101
USER 101
