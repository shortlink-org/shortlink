"""Repository interface for referral domain."""

import abc

from src.domain.referral.v1.referral_pb2 import Referral


class AbstractRepository(abc.ABC):
    """Abstract Repository for referral domain."""

    @abc.abstractmethod
    def add(self, referral: Referral) -> Referral:
        """Add referral."""
        raise NotImplementedError

    @abc.abstractmethod
    def get(self, referral_id: str) -> Referral:
        """Get referral."""
        raise NotImplementedError

    @abc.abstractmethod
    def update(self, referral: Referral) -> Referral:
        """Update referral."""
        raise NotImplementedError

    @abc.abstractmethod
    def delete(self, referral_id: str) -> None:
        """Delete referral."""
        raise NotImplementedError

    @abc.abstractmethod
    def list(self) -> list[Referral]:
        """List all referrals."""
        raise NotImplementedError
