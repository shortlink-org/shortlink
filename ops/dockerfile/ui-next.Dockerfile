FROM node:15.6-alpine as builder

WORKDIR /app
COPY ./pkg/ui/next /app/

RUN npm config set ignore-scripts false && \
  npm i && \
  npm run generate

FROM nginx:1.19-alpine

RUN apk add --update curl

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/out ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/dockerfile/conf/ui-nuxt.local /etc/nginx/conf.d/ui-nuxt.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
