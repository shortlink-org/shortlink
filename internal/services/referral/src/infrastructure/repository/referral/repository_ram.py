from typing import List

from google.protobuf.json_format import MessageToJson, ParseDict

from src.domain.referral.v1.referral_pb2 import Referral
from .repository import AbstractRepository
from src.domain.referral.v1.exception import ReferralNotFound

class Repository(AbstractRepository):
    def __init__(self):
        self._referrals = {}

    def get(self, referral_id: str) -> Referral:
        payload = self._referrals.get(referral_id)

        if payload is None:
            raise ReferralNotFound

        referral = Referral()
        ParseDict(payload, referral)
        return referral

    def add(self, referral: Referral) -> Referral:
        self._referrals[referral.id] = MessageToJson(referral)
        return referral

    def update(self, referral: Referral) -> Referral:
        self._referrals[referral.id] = MessageToJson(referral)
        return referral

    def delete(self, referral_id: str) -> None:
        self._referrals.pop(referral_id, None)
        return

    def list(self) -> List[Referral]:
        referrals = []
        for payload in self._referrals.values():
            referral = Referral()
            ParseDict(payload, referral)
            referrals.append(referral)

        return referrals
