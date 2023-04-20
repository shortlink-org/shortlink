"""Use Referral Use Case Module."""

from domain.referral.v1.referral_pb2 import Referral as ReferralModel

import redis

class UseReferralService:
  def __init__(self, redis: redis) -> None:
    self._redis = redis

  async def use(self, referral_id: int) -> ReferralModel:
    referral = ReferralModel()
    referral.ParseFromString(await self._redis.get(referral_id))
    await self._redis.set(referral.id, referral.SerializeToString())
    return referral

