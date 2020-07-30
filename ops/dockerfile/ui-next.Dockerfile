FROM node:14.7-alpine as builder

WORKDIR /app
ADD ./pkg/ui/next /app/

RUN npm i && \
  npm run generate

FROM nginx:1.19-alpine

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/out ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/dockerfile/conf/ui-nuxt.local /etc/nginx/conf.d/ui-nuxt.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
