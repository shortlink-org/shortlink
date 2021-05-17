FROM node:16.1-alpine as builder

WORKDIR /app
COPY ./internal/services/proxy /app/

RUN npm i

CMD ["npm", "start"]
