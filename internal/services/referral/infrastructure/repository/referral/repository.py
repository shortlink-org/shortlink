import abc
from typing import List

from domain.referral.v1.referral_pb2 import Referral


class AbstractRepository(abc.ABC):
    @abc.abstractmethod
    def add(self, referral: Referral):
        raise NotImplementedError

    @abc.abstractmethod
    def get(self, referral_id: str) -> Referral:
        raise NotImplementedError

    @abc.abstractmethod
    def update(self, referral: Referral):
        raise NotImplementedError

    @abc.abstractmethod
    def delete(self, referral_id: str) -> None:
        raise NotImplementedError

    @abc.abstractmethod
    def list(self) -> List[Referral]:
        raise NotImplementedError
