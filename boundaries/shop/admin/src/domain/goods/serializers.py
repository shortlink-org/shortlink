from rest_framework import serializers
from .models import Good

class GoodSerializer(serializers.ModelSerializer):
    """
    Serializer for the Good model.

    This serializer converts the Good model instances to JSON format and vice versa.
    """

    class Meta:
        """
        Metaclass for GoodSerializer.

        Specifies the model to be serialized and the fields to include in the serialization.
        """
        model = Good
        fields = ['id', 'name', 'price', 'description', 'created_at', 'updated_at']
