from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar

from google.protobuf import descriptor as _descriptor
from google.protobuf import field_mask_pb2 as _field_mask_pb2
from google.protobuf import message as _message
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper

DESCRIPTOR: _descriptor.FileDescriptor

class ReferralCommand(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    REFERRAL_COMMAND_UNSPECIFIED: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_ADD: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_UPDATE: _ClassVar[ReferralCommand]
    REFERRAL_COMMAND_DELETE: _ClassVar[ReferralCommand]
REFERRAL_COMMAND_UNSPECIFIED: ReferralCommand
REFERRAL_COMMAND_ADD: ReferralCommand
REFERRAL_COMMAND_UPDATE: ReferralCommand
REFERRAL_COMMAND_DELETE: ReferralCommand

class ReferralAddCommand(_message.Message):
    __slots__ = ("name", "user_id")
    NAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    user_id: str
    def __init__(self, name: str | None = ..., user_id: str | None = ...) -> None: ...

class ReferralUpdateCommand(_message.Message):
    __slots__ = ("id", "name", "user_id", "field_mask")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    FIELD_MASK_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    user_id: str
    field_mask: _field_mask_pb2.FieldMask
    def __init__(self, id: str | None = ..., name: str | None = ..., user_id: str | None = ..., field_mask: _field_mask_pb2.FieldMask | _Mapping | None = ...) -> None: ...

class ReferralDeleteCommand(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: str | None = ...) -> None: ...
