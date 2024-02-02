from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ReferralQuery(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    REFERRAL_QUERY_UNSPECIFIED: _ClassVar[ReferralQuery]
    REFERRAL_QUERY_GET: _ClassVar[ReferralQuery]
    REFERRAL_QUERY_LIST: _ClassVar[ReferralQuery]
REFERRAL_QUERY_UNSPECIFIED: ReferralQuery
REFERRAL_QUERY_GET: ReferralQuery
REFERRAL_QUERY_LIST: ReferralQuery

class ReferralGetQuery(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ReferralListQuery(_message.Message):
    __slots__ = ("user_id",)
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...
