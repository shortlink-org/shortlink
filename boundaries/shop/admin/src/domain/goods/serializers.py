"""This module contains serializers for the Good model.

The serializers convert Good model instances to JSON format and vice versa.
"""

from typing import ClassVar

from rest_framework import serializers

from .models import Good


class GoodSerializer(serializers.ModelSerializer):
    """Serializer for the Good model.

    This serializer converts the Good model instances to JSON format and vice versa.
    """

    class Meta:
        """Metaclass for GoodSerializer.

        Specifies the model to be serialized and the fields to include in the serialization.
        """

        model = Good
        fields: ClassVar[list] = ["id", "name", "price", "description", "created_at", "updated_at"]
