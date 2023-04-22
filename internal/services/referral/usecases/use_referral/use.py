"""Use Referral Use Case Module."""

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral
from pkg.event_bus import EventBus

class UseReferralService():
  def __init__(self, repository: AbstractRepository, event_bus: EventBus) -> None:
    self._repository = repository
    self._event_bus = event_bus

  def use(self, referral_id: int) -> Referral:
    return self._repository.get(referral_id)
