FROM node:13.14 as builder

WORKDIR /app
ADD ./pkg/ui/nuxt /app/

RUN npm i && \
  npm run generate

FROM nginx:1.18-alpine

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/dist ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/conf.d/ui-nuxt.local /etc/nginx/conf.d/ui-nuxt.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
