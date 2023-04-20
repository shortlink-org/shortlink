from redis.backoff import ExponentialBackoff
from redis.retry import Retry
from redis.client import Redis
from redis.exceptions import (
   BusyLoadingError,
   ConnectionError,
   TimeoutError
)

from domain.referral.v1.referral_pb2 import Referral, Referrals
from .repository import AbstractRepository

class Repository(AbstractRepository):
  def __init__(self, host: str):
    # Run 3 retries with exponential backoff strategy
    retry = Retry(ExponentialBackoff(), 3)

    self._redis = Redis(host=host, port=6379, retry=retry, retry_on_error=[BusyLoadingError, ConnectionError, TimeoutError])

    # ping to check if redis is up
    self._redis.ping()

  def add(self, referral: Referral):
    pass

  def get(self, referral_id: str) -> Referral:
    pass

  def update(self, referral: Referral):
    pass

  def delete(self, referral_id: str):
    pass

  def list(self) -> Referrals:
    pass
