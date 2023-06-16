"""Containers module."""

from dependency_injector import containers, providers

from src.usecases.crud_referral.crud import CRUDReferralService
from src.usecases.use_referral.use import UseReferralService
from src.infrastructure.repository.referral.repository_redis import Repository
from src.di.core import Core
from src.infrastructure.repository.referral.uow_redis import RedisUnitOfWork

class Application(containers.DeclarativeContainer):
    """Application container."""

    referral_repository = providers.Singleton(
        Repository,
    )

    referral_uow = providers.Factory(
        RedisUnitOfWork,
        repository=referral_repository,
    )

    core = providers.Container(
        Core,
    )

    referral_service = providers.Factory(
        CRUDReferralService,
        uow=referral_uow,
        event_bus=core.event_bus,
    )

    use_service = providers.Factory(
        UseReferralService,
        uow=referral_uow,
        event_bus=core.event_bus,
    )
