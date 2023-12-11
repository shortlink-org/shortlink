from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from typing import ClassVar as _ClassVar

DESCRIPTOR: _descriptor.FileDescriptor

class ReferralEvent(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    REFERRAL_EVENT_UNSPECIFIED: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_ADD: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_GET: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_LIST: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_UPDATE: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_DELETE: _ClassVar[ReferralEvent]
REFERRAL_EVENT_UNSPECIFIED: ReferralEvent
REFERRAL_EVENT_ADD: ReferralEvent
REFERRAL_EVENT_GET: ReferralEvent
REFERRAL_EVENT_LIST: ReferralEvent
REFERRAL_EVENT_UPDATE: ReferralEvent
REFERRAL_EVENT_DELETE: ReferralEvent
