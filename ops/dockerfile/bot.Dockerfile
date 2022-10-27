# syntax=docker/dockerfile:1.4

FROM maven:3.8.6-jdk-11-slim as builder

ARG CI_COMMIT_TAG
WORKDIR /app

# Load dependencies
COPY internal/services/bot /app
RUN mvn -f /app/pom.xml clean package

FROM openjdk:11.0.16-jre-slim

# Install dependencies
RUN \
  apt update && \
  apt install -y curl

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:9090/ready || exit 1

COPY --from=builder /app/target/shortlink-bot-1.0-SNAPSHOT.jar /usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar

ENTRYPOINT ["java", "-jar", "/usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar"]
