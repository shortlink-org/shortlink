# syntax=docker/dockerfile:1.3

# Install dependencies only when needed
FROM node:16.12-alpine as deps

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
RUN npm config set ignore-scripts false

WORKDIR /app
COPY ./ui/next/package.json ./ui/next/package-lock.json ./

RUN npm i

# Rebuild the source code only when needed
FROM node:16.12-alpine as builder

WORKDIR /app
COPY ./ui/next /app/
COPY --from=deps /app/node_modules ./node_modules

RUN npm run generate

# Production image, copy all the files and run next
FROM nginx:1.21-alpine

RUN apk add --no-cache curl

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/out ./
COPY ./ops/dockerfile/conf/ui-next.local /etc/nginx/conf.d/ui-next.local
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
