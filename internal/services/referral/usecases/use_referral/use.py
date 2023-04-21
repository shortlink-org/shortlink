"""Use Referral Use Case Module."""

from dependency_injector.wiring import inject, Provide
from opentelemetry.trace import Tracer

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral

class UseReferralService():
    def __init__(self, repository: AbstractRepository) -> None:
        self._repository = repository

    @inject
    def use(self, referral_id: str, tracer: Tracer = Provide["Core.tracer"]) -> Referral:
        with tracer.start_as_current_span("use_referral") as span:
            if span.is_recording():
                span.set_attribute("referral_id", referral_id)
            return self._repository.get(referral_id)
