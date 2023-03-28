# syntax=docker/dockerfile:1.5

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM maven:3.8.6-jdk-11-slim as builder

ARG CI_COMMIT_TAG
WORKDIR /app

# Load dependencies
COPY internal/services/bot /app
RUN mvn -f /app/pom.xml clean package

FROM openjdk:21-ea-13-jdk-slim

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

ENTRYPOINT ["/sbin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

COPY --from=builder /app/target/shortlink-bot-1.0-SNAPSHOT.jar /usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar

CMD ["java", "-jar", "/usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar"]
