"""Containers module."""

from dependency_injector import containers, providers

from di.logger.logger import LoguruJsonProvider
from di.opentelemetry.opentelemetry import OpenTelemetryProvider
from di.http.server import QuartProvider
from pkg.event_bus import EventBus


class Core(containers.DeclarativeContainer):
  config = providers.Configuration("config")

  logger = LoguruJsonProvider()
  tracer = OpenTelemetryProvider()
  app = QuartProvider()
  eventBus = providers.Singleton(EventBus)
