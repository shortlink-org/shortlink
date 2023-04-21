"""Use Referral Use Case Module."""

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral

class UseReferralService():
  def __init__(self, repository: AbstractRepository) -> None:
    self._repository = repository

  def use(self, referral_id: int) -> Referral:
    return self._repository.get(referral_id)
