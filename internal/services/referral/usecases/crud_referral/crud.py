"""CRUD Referral use case."""
from typing import List

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral, Referrals

class CRUDReferralService(AbstractRepository):
  def __init__(self, repository: AbstractRepository) -> None:
    self._repository = repository

  def get(self, referral_id: str) -> Referral:
    return self._repository.get(referral_id)

  def add(self, referral: Referral) -> Referral:
    return self._repository.add(referral)

  def update(self, referral: Referral) -> Referral:
    return self._repository.update(referral)

  def delete(self, referral_id: str) -> None:
    return self._repository.delete(referral_id)

  def list(self) -> list[Referral]:
    return self._repository.list()
