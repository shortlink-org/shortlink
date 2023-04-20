"""Containers module."""

from dependency_injector import containers, providers

from usecases.crud_referral.crud import CRUDReferralService
from usecases.use_referral.use import UseReferralService

class Container(containers.DeclarativeContainer):

  # Services ==================================================================

  referral_service = providers.Factory(
    CRUDReferralService,
  )

  use_service = providers.Factory(
    UseReferralService,
  )
