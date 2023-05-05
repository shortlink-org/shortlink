"""Containers module."""

from dependency_injector import containers, providers

from src.di.logger.logger import LoguruJsonProvider
from src.di.observability.opentelemetry import OpenTelemetryProvider
from src.di.observability.prometheus import PrometheusMetricsProvider
from src.di.http.server import QuartProvider
from src.pkg.event_bus import EventBus


class Core(containers.DeclarativeContainer):
  config = providers.Configuration("config")

  logger = LoguruJsonProvider()
  tracer = providers.Singleton(OpenTelemetryProvider)
  app = QuartProvider()
  eventBus = providers.Singleton(EventBus)
  prometheus_metrics = PrometheusMetricsProvider()
