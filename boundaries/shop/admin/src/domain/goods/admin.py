"""Define the admin view for the Good model."""

from django.contrib import admin

from .models import Good


class GoodAdmin(admin.ModelAdmin):
    """Define the admin view for the Good model."""

    list_display = ("name", "price", "created_at")
    search_fields = ("name", "price")
    list_filter = ("created_at", "updated_at")
    ordering = ("created_at",)


admin.site.register(Good, GoodAdmin)
