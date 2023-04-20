"""CRUD Referral use case."""

import redis
from google.protobuf.json_format import MessageToJson

from domain.referral.v1.referral_pb2 import Referral as ReferralModel, Referrals as ReferralsModel

class CRUDReferralService:
  def __init__(self, redis: redis) -> None:
    self._redis = redis

  async def create(self, referral: ReferralModel) -> ReferralModel:
    await self._redis.set(referral.id, MessageToJson(referral))

  async def get(self, referral_id: str) -> ReferralModel:
    referral = ReferralModel()
    referral.ParseFromString(await self._redis.get(referral_id))
    return referral

  async def update(self, referral: ReferralModel) -> ReferralModel:
    await self._redis.set(referral.id, MessageToJson(referral))

  async def delete(self, referral_id: str) -> None:
    await self._redis.delete(referral_id)

  async def list(self, limit: int, offset: int) -> ReferralsModel:
    referrals = ReferralsModel()
    for referral_id in await self._redis.keys():
      referral = ReferralModel()
      referral.ParseFromString(await self._redis.get(referral_id))
      referrals.referrals.append(referral)
    return referrals

