#!/usr/bin/env python
"""Django's command-line utility for administrative tasks."""

import logging
import os
import sys

from admin.otel_logging import CustomLogRecord
from opentelemetry.instrumentation.django import DjangoInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor

logging.setLogRecordFactory(CustomLogRecord)


def main():
    """Run administrative tasks."""
    os.environ.setdefault("DJANGO_SETTINGS_MODULE", "admin.settings")

    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line([*sys.argv, "--noreload"])

    DjangoInstrumentor().instrument()
    RequestsInstrumentor().instrument()


if __name__ == "__main__":
    main()
