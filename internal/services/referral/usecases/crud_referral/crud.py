"""CRUD Referral use case."""

from typing import List
from dependency_injector.wiring import inject, Provide
from opentelemetry.trace import Tracer

from infrastructure.repository.referral.repository import AbstractRepository
from domain.referral.v1.referral_pb2 import Referral, Referrals

class CRUDReferralService(AbstractRepository):
    def __init__(self, repository: AbstractRepository) -> None:
        self._repository = repository

    @inject
    def get(self, referral_id: str, tracer: Tracer = Provide["Core.tracer"]) -> Referral:
        with tracer.start_as_current_span("get_referral") as span:
            if span.is_recording():
                span.set_attribute("referral_id", referral_id)
            return self._repository.get(referral_id)

    @inject
    def add(self, referral: Referral, tracer: Tracer = Provide["Core.tracer"]) -> Referral:
        with tracer.start_as_current_span("add_referral") as span:
            if span.is_recording():
                span.set_attribute("referral_id", referral.id)
            return self._repository.add(referral)

    @inject
    def update(self, referral: Referral, tracer: Tracer = Provide["Core.tracer"]) -> Referral:
        with tracer.start_as_current_span("update_referral") as span:
            if span.is_recording():
                span.set_attribute("referral_id", referral.id)
            return self._repository.update(referral)

    @inject
    def delete(self, referral_id: str, tracer: Tracer = Provide["Core.tracer"]) -> None:
        with tracer.start_as_current_span("delete_referral") as span:
            if span.is_recording():
                span.set_attribute("referral_id", referral_id)
            return self._repository.delete(referral_id)

    @inject
    def list(self, tracer: Tracer = Provide["Core.tracer"]) -> list[Referral]:
        with tracer.start_as_current_span("list_referrals") as span:
            if span.is_recording():
                span.set_attribute("referral_id", "all")
            return self._repository.list()
