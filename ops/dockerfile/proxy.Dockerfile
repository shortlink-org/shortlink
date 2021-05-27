FROM node:16.2-alpine as builder

ARG MAX_OLD_SPACE_SIZE=8192
ENV NODE_OPTIONS=--max_old_space_size=${MAX_OLD_SPACE_SIZE}

WORKDIR /app
COPY ./internal/services/proxy /app/

RUN npm i

CMD ["npm", "start"]
