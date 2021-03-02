FROM node:15.8-alpine as builder

WORKDIR /app
COPY ./ui/next /app/

RUN npm config set ignore-scripts false && \
  npm i && \
  npm run generate

FROM nginx:1.19-alpine

RUN apk add --no-cache curl

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/out ./
COPY ./ops/docker-compose/gateway/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/dockerfile/conf/ui-next.local /etc/nginx/conf.d/ui-next.local
COPY ./ops/docker-compose/gateway/nginx/templates /etc/nginx/template
