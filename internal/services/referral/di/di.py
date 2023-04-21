"""Containers module."""

from dependency_injector import containers, providers

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService

from infrastructure.repository.referral.redis import Repository
from di.logger.logger import LoguruJsonProvider
from di.opentelemetry.opentelemetry import OpenTelemetryProvider
from di.http.server import QuartProvider


class Core(containers.DeclarativeContainer):
  config = providers.Configuration("config")

  logger = LoguruJsonProvider()
  tracer = OpenTelemetryProvider()
  app = QuartProvider()


class Application(containers.DeclarativeContainer):
  config = providers.Configuration()

  redis = providers.Singleton(
    Repository,
    host="localhost",
  )

  core = providers.Container(
    Core,
    config=config.core,
  )

  referral_service = providers.Factory(
    CRUDReferralService,
    repository=redis,
  )

  use_service = providers.Factory(
    UseReferralService,
    repository=redis,
  )
