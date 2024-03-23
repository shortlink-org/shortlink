from django.http import HttpResponse
from opentelemetry import trace

tracer = trace.get_tracer(__name__)

def hello(request):
    # Create a custom span
    with tracer.start_as_current_span("hello"):
        return HttpResponse("Hello, World!")
