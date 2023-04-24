from google.protobuf import field_mask_pb2 as _field_mask_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReferralEvent(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    REFERRAL_EVENT_UNSPECIFIED: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_ADD: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_GET: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_LIST: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_UPDATE: _ClassVar[ReferralEvent]
    REFERRAL_EVENT_DELETE: _ClassVar[ReferralEvent]

class ReferralCommand(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    REFERRAL_COMMAND_UNSPECIFIED: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_ADD: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_UPDATE: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_DELETE: _ClassVar[ReferralCommand]

class ReferralQuery(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    REFERRAL_QUERY_UNSPECIFIED: _ClassVar[ReferralQuery]
    REFERRAL_QUERY_GET: _ClassVar[ReferralQuery]
    REFERRAL_QUERY_LIST: _ClassVar[ReferralQuery]
REFERRAL_EVENT_UNSPECIFIED: ReferralEvent
REFERRAL_EVENT_ADD: ReferralEvent
REFERRAL_EVENT_GET: ReferralEvent
REFERRAL_EVENT_LIST: ReferralEvent
REFERRAL_EVENT_UPDATE: ReferralEvent
REFERRAL_EVENT_DELETE: ReferralEvent
REFERRAL_COMMAND_UNSPECIFIED: ReferralCommand
REFERRAL_COMMAND_ADD: ReferralCommand
REFERRAL_COMMAND_UPDATE: ReferralCommand
REFERRAL_COMMAND_DELETE: ReferralCommand
REFERRAL_QUERY_UNSPECIFIED: ReferralQuery
REFERRAL_QUERY_GET: ReferralQuery
REFERRAL_QUERY_LIST: ReferralQuery

class ReferralAddCommand(_message.Message):
    __slots__ = ["name", "user_id"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    user_id: str
    def __init__(self, name: _Optional[str] = ..., user_id: _Optional[str] = ...) -> None: ...

class ReferralUpdateCommand(_message.Message):
    __slots__ = ["id", "name", "user_id", "field_mask"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    FIELD_MASK_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    user_id: str
    field_mask: _field_mask_pb2.FieldMask
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., user_id: _Optional[str] = ..., field_mask: _Optional[_Union[_field_mask_pb2.FieldMask, _Mapping]] = ...) -> None: ...

class ReferralDeleteCommand(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ReferralGetQuery(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ReferralListQuery(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class Referral(_message.Message):
    __slots__ = ["id", "name", "user_id", "created_at", "updated_at", "field_mask"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    FIELD_MASK_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    user_id: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    field_mask: _field_mask_pb2.FieldMask
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., user_id: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., field_mask: _Optional[_Union[_field_mask_pb2.FieldMask, _Mapping]] = ...) -> None: ...

class Referrals(_message.Message):
    __slots__ = ["referrals"]
    REFERRALS_FIELD_NUMBER: _ClassVar[int]
    referrals: _containers.RepeatedCompositeFieldContainer[Referral]
    def __init__(self, referrals: _Optional[_Iterable[_Union[Referral, _Mapping]]] = ...) -> None: ...
