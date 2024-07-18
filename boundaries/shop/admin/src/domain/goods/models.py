"""Define the Good model."""

from django.db import models
from django_prometheus.models import ExportModelOperationsMixin


class Good(ExportModelOperationsMixin("goods"), models.Model):
    """Define the Good model."""

    name = models.CharField(max_length=255)
    price = models.DecimalField(max_digits=5, decimal_places=2)
    description = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def __str__(self):
        """Return the name of the good."""
        return self.name
