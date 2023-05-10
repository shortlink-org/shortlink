"""CRUD Referral use case."""

from src.domain.referral.v1.referral_pb2 import Referral
from src.domain.referral.v1.events_pb2 import REFERRAL_EVENT_ADD, REFERRAL_EVENT_UPDATE, REFERRAL_EVENT_DELETE, ReferralEvent
from src.pkg.event_bus import EventBus
from src.infrastructure.repository.referral.uow_redis import RedisUnitOfWork

class CRUDReferralService():
    def __init__(self, uow: RedisUnitOfWork, event_bus: EventBus) -> None:
        self._uow = uow
        self._event_bus = event_bus

    def get(self, referral_id: str) -> Referral:
        with self._uow:
            return self._uow.referral.get(referral_id)

    def add(self, referral: Referral) -> Referral:
        with self._uow as uow:
            result = uow.referral.add(referral)
            uow.commit()
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_ADD), referral)
        return result

    def update(self, referral: Referral) -> Referral:
        with self._uow as uow:
            result = uow.referral.update(referral)
            uow.commit()
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_UPDATE), referral)
        return result

    def delete(self, referral_id: str) -> None:
        with self._uow as uow:
            uow.referral.delete(referral_id)
            uow.commit()
        self._event_bus.publish(ReferralEvent.Name(REFERRAL_EVENT_DELETE), referral_id)

    def list(self) -> list[Referral]:
        with self._uow:
            return self._uow.referral.list()
