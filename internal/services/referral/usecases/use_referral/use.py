"""Use Referral Use Case Module."""

from domain.referral.v1.referral_pb2 import Referral as ReferralModel

class UseReferralService:
  def __init__(self) -> None: ...

  def use(self, referral_id: int) -> None: ...
