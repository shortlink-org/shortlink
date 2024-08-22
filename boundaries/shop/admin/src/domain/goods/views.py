"""This module contains viewsets for the goods domain.

It includes the GoodViewSet, which handles viewing and editing Good instances.
"""

from rest_framework import viewsets

from .models import Good
from .serializers import GoodSerializer


class GoodViewSet(viewsets.ModelViewSet):
    """A viewset for viewing and editing Good instances."""

    serializer_class = GoodSerializer
    queryset = Good.objects.all()
