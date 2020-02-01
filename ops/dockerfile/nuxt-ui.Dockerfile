FROM node:13.7 as builder

WORKDIR /app
ADD ./pkg/ui/nuxt /app/

RUN npm i && \
  npm run generate

FROM nginx:1.17-alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/conf.d/nuxt-ui.local /etc/nginx/conf.d/nuxt-ui.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
