# syntax=docker/dockerfile:1.8

# Link: https://github.com/moby/buildkit/blob/master/docs/attestations/sbom.md
# enable scanning for the intermediate build stage
ARG BUILDKIT_SBOM_SCAN_STAGE=true
# scan the build context only if the build is run to completion
ARG BUILDKIT_SBOM_SCAN_CONTEXT=true

FROM --platform=$BUILDPLATFORM python:3.12-slim

LABEL maintainer=batazor111@gmail.com
LABEL org.opencontainers.image.title="shortlink-shop-admin"
LABEL org.opencontainers.image.description="shortlink-shop-admin"
LABEL org.opencontainers.image.authors="Login Viktor @batazor"
LABEL org.opencontainers.image.vendor="Login Viktor @batazor"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.url="http://shortlink.best/"
LABEL org.opencontainers.image.source="https://github.com/shortlink-org/shortlink"

ENV PYTHONPATH="$PYTHONPATH:$PWD"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV VIRTUAL_ENV=/usr/local

# Kratos ENV
ENV ORY_SDK_URL="http://host.docker.internal:4433"
ENV ORY_UI_URL="http://host.docker.internal:3000/next/auth"

# HTTP API && Prometheus metrics
EXPOSE 8000

RUN apt-get update && \
    apt-get install -y --no-install-recommends curl tini sqlite3

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
  CMD curl -f localhost:8000/healthz || exit 1

RUN addgroup --system shop && adduser --system --group shop

COPY boundaries/shop/admin/ .
RUN chown -R shop:shop /app/src

USER shop

CMD ["python", "src/manage.py", "runserver", "0.0.0.0:8000"]
