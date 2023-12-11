"""Containers module."""

from dependency_injector import containers, providers

from src.di.logger.logger import LoguruJsonProvider
from src.di.observability.opentelemetry import OpenTelemetryProvider
from src.di.observability.prometheus import PrometheusMetricsProvider
from src.di.http.server import QuartProvider
from src.di.config.config import Config
from src.pkg.event_bus import EventBus


class Core(containers.DeclarativeContainer):
    """Core container."""
    config = providers.Singleton(Config)

    logger = LoguruJsonProvider()
    tracer = providers.Singleton(OpenTelemetryProvider)
    app = QuartProvider()
    event_bus = providers.Singleton(EventBus)
    prometheus_metrics = providers.Singleton(PrometheusMetricsProvider)
