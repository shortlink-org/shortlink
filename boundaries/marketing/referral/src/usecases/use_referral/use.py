"""Use Referral Use Case Module."""

from src.domain.referral.v1.referral_pb2 import Referral
from src.pkg.event_bus import EventBus
from src.infrastructure.repository.referral.uow_redis import RedisUnitOfWork

class UseReferralService:
    """Use Referral Use Case."""
    def __init__(self, uow: RedisUnitOfWork, event_bus: EventBus) -> None:
        """Initialize Use Referral Use Case."""
        self._uow = uow
        self._event_bus = event_bus


    def use(self, referral_id: int) -> Referral:
        """Use Referral."""
        with self._uow:
            return self._uow.referral.get(referral_id)
