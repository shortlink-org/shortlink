# syntax=docker/dockerfile:1.9

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM maven:3-eclipse-temurin-21-alpine AS builder

ARG CI_COMMIT_TAG
WORKDIR /app

# Load dependencies
COPY boundaries/notification/bot /app
RUN mvn -f /app/pom.xml clean package

FROM openjdk:22-ea-25-jdk-slim-bullseye

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-bot"
LABEL org.opencontainers.image.description="shortlink-bot"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# Install dependencies
RUN \
  apt update && \
  apt install -y curl tini

ENTRYPOINT ["/usr/bin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

WORKDIR /usr/local/lib/

COPY --from=builder /app/target/shortlink-bot-1.0-SNAPSHOT.jar ./shortlink-bot-1.0-SNAPSHOT.jar

CMD ["java", "-jar", "/usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar"]
