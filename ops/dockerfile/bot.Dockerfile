# syntax=docker/dockerfile:1.3

FROM maven:3.8.4-jdk-11-slim as builder

ARG CI_COMMIT_TAG
WORKDIR /app

# Load dependencies
COPY internal/services/bot /app
RUN mvn -f /app/pom.xml clean package

FROM openjdk:11-jre-slim
COPY --from=builder /app/target/shortlink-bot-1.0-SNAPSHOT.jar /usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar
ENTRYPOINT ["java", "-jar", "/usr/local/lib/shortlink-bot-1.0-SNAPSHOT.jar"]
