from redis.backoff import ExponentialBackoff
from redis.retry import Retry
from redis.client import Redis
from redis.exceptions import (
   BusyLoadingError,
   ConnectionError,
   TimeoutError
)
from google.protobuf.json_format import MessageToJson

from domain.referral.v1.referral_pb2 import Referral, Referrals
from .repository import AbstractRepository

class Repository(AbstractRepository):
  def __init__(self, host: str):
    # Run 3 retries with exponential backoff strategy
    retry = Retry(ExponentialBackoff(), 3)

    self._redis = Redis(host=host, port=6379, retry=retry, retry_on_error=[BusyLoadingError, ConnectionError, TimeoutError])

    # ping to check if redis is up
    self._redis.ping()

  async def add(self, referral: Referral):
    await self._redis.set(referral.id, MessageToJson(referral))

  async def get(self, referral_id: str) -> Referral:
    referral = Referral()
    referral.ParseFromString(await self._redis.get(referral_id))
    return referral

  async def update(self, referral: Referral):
    await self._redis.set(referral.id, MessageToJson(referral))

  async def delete(self, referral_id: str):
    await self._redis.delete(referral_id)

  async def list(self) -> Referrals:
    referrals = Referrals()
    for referral_id in await self._redis.keys():
      referral = Referral()
      referral.ParseFromString(await self._redis.get(referral_id))
      referrals.referrals.append(referral)
    return referrals
