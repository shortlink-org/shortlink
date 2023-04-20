from redis.backoff import ExponentialBackoff
from redis.retry import Retry
from redis.client import Redis
from redis.exceptions import (
   BusyLoadingError,
   ConnectionError,
   TimeoutError
)

class Repository:
  def __init__(self, host: str):
    # Run 3 retries with exponential backoff strategy
    retry = Retry(ExponentialBackoff(), 3)

    self._redis = Redis(host=host, port=6379, retry=retry, retry_on_error=[BusyLoadingError, ConnectionError, TimeoutError])

    # ping to check if redis is up
    self._redis.ping()

