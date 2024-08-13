from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import GoodViewSet

router = DefaultRouter()
router.register(r'', GoodViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
