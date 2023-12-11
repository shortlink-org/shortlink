# syntax=docker/dockerfile:1.6

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM python:3.13-rc-slim AS builder

WORKDIR /app

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apt-get update && \
    apt-get install -y --no-install-recommends make gcc git libc6-dev g++

COPY internal/boundaries/marketing/referral/requirements.txt .
RUN pip wheel --no-cache-dir --no-deps --wheel-dir /app/wheels -r requirements.txt

# Final image
FROM --platform=$TARGETPLATFORM python:3.13-rc-slim

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-referral"
LABEL org.opencontainers.image.description="shortlink-referral"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

# HTTP API
EXPOSE 8000
# Prometheus metrics
EXPOSE 9090

WORKDIR /app
ENV PYTHONPATH="$PYTHONPATH:$PWD"
ENV PYTHONUNBUFFERED=1

# Install dependencies
RUN \
  apt-get update && \
  apt-get install -y curl tini

ENTRYPOINT ["/usr/bin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:8000/ready || exit 1

COPY --from=builder /app/wheels /wheels
COPY --from=builder /app/requirements.txt .

RUN pip install --no-cache /wheels/*

RUN addgroup --system referall && adduser --system --group referall
USER referall

COPY internal/boundaries/marketing/referral/ .
CMD ["python", "src/__main__.py"]
