# syntax=docker/dockerfile:1.3

FROM node:17.3-alpine as builder

ENV PYTHONUNBUFFERED=1

RUN apk add --no-cache alpine-sdk python3 libsass \
  && ln -sf python3 /usr/bin/python

WORKDIR /app
COPY ./ui/landing /app/

RUN npm i --force && \
  npm rebuild node-sass && \
  npm run generate

FROM nginx:1.21-alpine

RUN apk add --no-cache curl

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist ./
COPY ./ops/dockerfile/conf/landing.local /etc/nginx/conf.d/landing.local
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
