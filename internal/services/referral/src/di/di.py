"""Containers module."""

from dependency_injector import containers, providers

from src.usecases.crud_referral.crud import CRUDReferralService
from src.usecases.use_referral.use import UseReferralService
from src.infrastructure.repository.referral.repository_redis import Repository
from src.di.core import Core

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
        event_bus=core.eventBus,
    )

    use_service = providers.Factory(
        UseReferralService,
        repository=redis,
        event_bus=core.eventBus,
    )
