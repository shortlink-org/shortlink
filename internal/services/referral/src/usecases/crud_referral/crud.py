"""CRUD Referral use case."""

from src.infrastructure.repository.referral.redis import AbstractRepository
from src.domain.referral.v1.referral_pb2 import Referral
from src.domain.referral.v1.events_pb2 import REFERRAL_EVENT_ADD, REFERRAL_EVENT_UPDATE, REFERRAL_EVENT_DELETE, ReferralEvent
from src.pkg.event_bus import EventBus

class CRUDReferralService(AbstractRepository):
    def __init__(self, repository: AbstractRepository, event_bus: EventBus) -> None:
        self._repository = repository
        self._event_bus = event_bus

    def get(self, referral_id: str) -> Referral:
        return self._repository.get(referral_id)

    def add(self, referral: Referral) -> Referral:
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_ADD), referral)
        return self._repository.add(referral)

    def update(self, referral: Referral) -> Referral:
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_UPDATE), referral)
        return self._repository.update(referral)

    def delete(self, referral_id: str) -> None:
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_DELETE), referral_id)
        return self._repository.delete(referral_id)

    def list(self) -> list[Referral]:
        return self._repository.list()
