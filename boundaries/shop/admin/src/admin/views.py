"""Define the views for the admin app."""

from django.http import HttpResponse
from opentelemetry import trace

tracer = trace.get_tracer(__name__)


def hello(request):
    """Return a simple hello world response."""
    # Create a custom span
    with tracer.start_as_current_span("hello"):
        return HttpResponse("Hello, World!")
