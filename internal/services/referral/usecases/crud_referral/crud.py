"""CRUD Referral use case."""

from google.protobuf.json_format import MessageToJson

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral, Referrals

class CRUDReferralService(AbstractRepository):
  def __init__(self, repository: AbstractRepository) -> None:
    self._repository = repository

  async def add(self, referral: Referral) -> Referral:
    await self._repository.add(referral)

  async def get(self, referral_id: str) -> Referral:
    await self._repository.get(referral_id)

  async def update(self, referral: Referral) -> Referral:
    await self._repository.update(referral)

  async def delete(self, referral_id: str) -> None:
    await self._repository.delete(referral_id)

  async def list(self) -> Referrals:
    await self._repository.list()
