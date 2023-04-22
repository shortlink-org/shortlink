"""Containers module."""

from dependency_injector import containers, providers

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService
from infrastructure.repository.referral.redis import Repository
from di.core import Core

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
