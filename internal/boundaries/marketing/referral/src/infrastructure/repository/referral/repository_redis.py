"""Implementation of Referral Repository using RAM as storage."""

import json
import os
from urllib.parse import urlparse
from redis.backoff import ExponentialBackoff
from redis.retry import Retry
from redis.client import Redis
from redis.exceptions import (
   BusyLoadingError,
   ConnectionError,
   TimeoutError
)
from google.protobuf.json_format import MessageToJson, ParseDict

from src.domain.referral.v1.referral_pb2 import Referral
from .repository import AbstractRepository
from src.domain.referral.v1.exception import ReferralNotFoundError

class Repository(AbstractRepository):
    """Repository implementation for referral domain."""

    def __init__(self):
        """Initialize Redis connection."""
        # Run 3 retries with exponential backoff strategy
        retry = Retry(ExponentialBackoff(), 3)

        # parse DATABASE_URI as host, port, db
        uri = os.environ.get("DATABASE_URI")
        parsed_uri = urlparse(uri)

        self._redis = Redis(
            host=parsed_uri.hostname,
            port=parsed_uri.port,
            retry=retry,
            retry_on_error=[BusyLoadingError, ConnectionError, TimeoutError])

        # ping to check if redis is up
        self._redis.ping()

    def get(self, referral_id: str) -> Referral:
        """Get referral."""
        referral_data = self._redis.get(referral_id)

        # If the referral_id doesn't exist in the Redis database, raise a ReferralNotFoundError exception
        if referral_data is None:
            raise ReferralNotFoundError(f"Referral with id {referral_id} not found")

        # If the referral_id exists, load the json data and return the Referral object
        payload = json.loads(referral_data)
        referral = Referral()
        ParseDict(payload, referral)
        return referral

    def add(self, referral: Referral) -> Referral:
        """Add referral."""
        self._redis.set(referral.id, MessageToJson(referral))
        return referral

    def update(self, referral: Referral) -> Referral:
        """Update referral."""
        self._redis.set(referral.id, MessageToJson(referral))
        return referral

    def delete(self, referral_id: str) -> None:
        """Delete referral."""
        self._redis.delete(referral_id)

    def list(self) -> list[Referral]:
        """List all referrals."""
        referrals = []
        for referral_id in self._redis.scan_iter('*'):
            referral = Referral()
            payload = json.loads(self._redis.get(referral_id))
            ParseDict(payload, referral)
            referrals.append(referral)

        return referrals
