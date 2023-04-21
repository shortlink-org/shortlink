import json
from typing import List
from redis.backoff import ExponentialBackoff
from redis.retry import Retry
from redis.client import Redis
from redis.exceptions import (
   BusyLoadingError,
   ConnectionError,
   TimeoutError
)
from google.protobuf.json_format import MessageToJson, ParseDict

from domain.referral.v1.referral_pb2 import Referral
from .repository import AbstractRepository
from domain.referral.v1.exception import ReferralNotFound

class Repository(AbstractRepository):
    def __init__(self, host: str):
        # Run 3 retries with exponential backoff strategy
        retry = Retry(ExponentialBackoff(), 3)

        self._redis = Redis(host=host, port=6379, retry=retry, retry_on_error=[BusyLoadingError, ConnectionError, TimeoutError])

        # ping to check if redis is up
        self._redis.ping()

    def get(self, referral_id: str) -> Referral:
        payload = json.loads(self._redis.get(referral_id))

        if payload is None:
            raise ReferralNotFound

        referral = Referral()
        ParseDict(payload, referral)
        return referral

    def add(self, referral: Referral) -> Referral:
        self._redis.set(referral.id, MessageToJson(referral))
        return referral

    def update(self, referral: Referral) -> Referral:
            self._redis.set(referral.id, MessageToJson(referral))
            return referral

    def delete(self, referral_id: str) -> None:
            self._redis.delete(referral_id)
            return

    def list(self) -> List[Referral]:
        referrals = []
        for referral_id in self._redis.keys():
            referral = Referral()
            payload = json.loads(self._redis.get(referral_id))
            ParseDict(payload, referral)
            referrals.append(referral)

        return referrals
