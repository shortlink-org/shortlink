"""This module defines URL patterns for the goods domain.

It includes routes for the GoodViewSet, which handles CRUD operations for the Good model.
"""

from django.urls import include, path
from rest_framework.routers import DefaultRouter

from .views import GoodViewSet

router = DefaultRouter()
router.register(r"", GoodViewSet)

urlpatterns = [
    path("", include(router.urls)),
]
