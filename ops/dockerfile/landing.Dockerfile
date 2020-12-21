FROM node:14.15-alpine as builder

RUN apk add --update alpine-sdk

WORKDIR /app
COPY ./pkg/ui/landing /app/

RUN npm i && \
  npm run generate

FROM nginx:1.19-alpine

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/dockerfile/conf/landing.local /etc/nginx/conf.d/landing.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
