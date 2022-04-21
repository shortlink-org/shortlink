# syntax=docker/dockerfile:1.3

# Install dependencies only when needed
FROM node:18.0-alpine as builder

ENV PYTHONUNBUFFERED=1

RUN apk add --no-cache alpine-sdk python3 libsass \
  && ln -sf python3 /usr/bin/python

WORKDIR /app
COPY ./ui/landing /app/

RUN npm ci --force && \
  npm rebuild node-sass && \
  npm run generate

# Production image, copy all the files
FROM nginxinc/nginx-unprivileged:1.21-alpine

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

# Use root user to copy dist folder and modify user access to specific folder
USER root

# Copy application and custom NGINX configuration
COPY --from=builder /app/dist ./
COPY ./ops/dockerfile/conf/ui.local /etc/nginx/conf.d/ui.local
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template

# Setup unprivileged user 101
RUN chown -R nginx /usr/share/nginx/html

# Use user 101
USER nginx
