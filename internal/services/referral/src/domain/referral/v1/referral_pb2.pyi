from google.protobuf import field_mask_pb2 as _field_mask_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from src.domain.referral.v1 import commands_pb2 as _commands_pb2
from src.domain.referral.v1 import events_pb2 as _events_pb2
from src.domain.referral.v1 import queries_pb2 as _queries_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

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
