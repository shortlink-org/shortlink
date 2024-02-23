from collections.abc import Iterable as _Iterable
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar

from google.protobuf import descriptor as _descriptor
from google.protobuf import field_mask_pb2 as _field_mask_pb2
from google.protobuf import message as _message
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers

DESCRIPTOR: _descriptor.FileDescriptor

class Referral(_message.Message):
    __slots__ = ("id", "name", "user_id", "created_at", "updated_at", "field_mask")
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
    def __init__(self, id: str | None = ..., name: str | None = ..., user_id: str | None = ..., created_at: _timestamp_pb2.Timestamp | _Mapping | None = ..., updated_at: _timestamp_pb2.Timestamp | _Mapping | None = ..., field_mask: _field_mask_pb2.FieldMask | _Mapping | None = ...) -> None: ...

class Referrals(_message.Message):
    __slots__ = ("referrals",)
    REFERRALS_FIELD_NUMBER: _ClassVar[int]
    referrals: _containers.RepeatedCompositeFieldContainer[Referral]
    def __init__(self, referrals: _Iterable[Referral | _Mapping] | None = ...) -> None: ...
