# syntax=docker/dockerfile:1.9

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM python:3.12-slim

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-referral"
LABEL org.opencontainers.image.description="shortlink-referral"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

ENV PYTHONPATH="$PYTHONPATH:$PWD"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV VIRTUAL_ENV=/usr/local

# HTTP API
EXPOSE 8000
# Prometheus metrics
EXPOSE 9090

RUN apt-get update && \
    apt-get install -y --no-install-recommends make gcc git libc6-dev g++ curl tini libffi-dev

WORKDIR /app

COPY boundaries/shop/admin/pyproject.toml .

# Install dependency manager
# https://github.com/astral-sh/uv
RUN pip install uv
# Create a virtual environment at .venv
RUN uv venv

RUN uv pip install -r pyproject.toml --no-deps

ENTRYPOINT ["/usr/bin/tini", "--"]

HEALTHCHECK \
  --interval=5s \
  --timeout=5s \
  --retries=3 \
  CMD curl -f localhost:8000/healthz/ready || exit 1

RUN addgroup --system referall && adduser --system --group referall
USER referall

COPY boundaries/marketing/referral/ .
CMD ["python", "src/__main__.py"]
