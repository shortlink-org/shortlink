"""Implementation of Referral Repository using RAM as storage."""

from google.protobuf.json_format import MessageToJson, ParseDict

from src.domain.referral.v1.referral_pb2 import Referral
from .repository import AbstractRepository
from src.domain.referral.v1.exception import ReferralNotFoundError

class Repository(AbstractRepository):
    """Repository implementation for referral domain."""

    def __init__(self):
        """Initialize RAM."""
        self._referrals = {}

    def get(self, referral_id: str) -> Referral:
        """Get referral."""
        payload = self._referrals.get(referral_id)

        if payload is None:
            raise ReferralNotFoundError

        referral = Referral()
        ParseDict(payload, referral)
        return referral

    def add(self, referral: Referral) -> Referral:
        """Add referral."""
        self._referrals[referral.id] = MessageToJson(referral)
        return referral

    def update(self, referral: Referral) -> Referral:
        """Update referral."""
        self._referrals[referral.id] = MessageToJson(referral)
        return referral

    def delete(self, referral_id: str) -> None:
        """Delete referral."""
        self._referrals.pop(referral_id, None)

    def list(self) -> list[Referral]:
        """List all referrals."""
        referrals = []
        for payload in self._referrals.values():
            referral = Referral()
            ParseDict(payload, referral)
            referrals.append(referral)

        return referrals
