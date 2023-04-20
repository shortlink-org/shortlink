"""Containers module."""

from dependency_injector import containers, providers

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService

from infrastructure.repository.redis import Repository
from di.logger.logger import LoguruJsonProvider


class Core(containers.DeclarativeContainer):
  config = providers.Configuration("config")

  logger = LoguruJsonProvider()


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
    redis=redis,
  )

  use_service = providers.Factory(
    UseReferralService,
    redis=redis,
  )
