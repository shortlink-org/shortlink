# syntax=docker/dockerfile:1.7

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

# Install dependencies only when needed
FROM --platform=$BUILDPLATFORM public.ecr.aws/g6m3a0y9/landscape2:latest as builder

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN

WORKDIR /app

COPY ./boundaries/platform/landscape ./

USER root

RUN landscape2 build \
          --data-file data.yml \
          --settings-file settings.yml \
          --guide-file guide.yml \
          --logos-path logos \
          --output-dir build

# Production image, copy all the files and run next
FROM --platform=$TARGETPLATFORM ghcr.io/nginxinc/nginx-unprivileged:1.27-alpine

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-landscape"
LABEL org.opencontainers.image.description="shortlink-landscape"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

EXPOSE 8080

# Delete default config
RUN rm /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/html

# Use root user to copy dist folder and modify user access to specific folder
USER root

# Install dependencies
RUN \
  apk update && \
  apk add --no-cache curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:8080 || exit 1

# Copy application and custom NGINX configuration
COPY ./ops/dockerfile/boundary_platform/landscape/landscape.local /etc/nginx/conf.d/default.conf
COPY ./ops/docker-compose/gateway/nginx/conf/nginx.conf /etc/nginx/nginx.conf
COPY ./ops/docker-compose/gateway/nginx/conf/templates /etc/nginx/template
COPY --from=builder /app/build ./

# Setup unprivileged user 1001
RUN chown -R 1001 /usr/share/nginx/html

# Use user 1001
USER 1001
