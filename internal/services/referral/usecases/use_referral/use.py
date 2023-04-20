"""Use Referral Use Case Module."""

from domain.referral.v1.referral_pb2 import Referral as ReferralModel

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral

class UseReferralService():
  def __init__(self, repository: AbstractRepository) -> None:
    self._repository = repository

  async def use(self, referral_id: int) -> ReferralModel:
    return await self._repository.get(referral_id)

