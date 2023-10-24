"""OpenTelemetry provider."""

from dependency_injector import providers

import pyroscope
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.sdk.resources import Resource
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter


class OpenTelemetryProvider(providers.Provider):
    """OpenTelemetry provider."""
    @staticmethod
    def _provide(*args, **kwargs):
        resource = Resource.create(attributes={
            "service.name": "referral-service",
        })

        trace.set_tracer_provider(TracerProvider(resource=resource))
        tracer = trace.get_tracer(__name__)

        otlp_exporter = OTLPSpanExporter(endpoint="localhost:5050", insecure=True)
        span_processor = BatchSpanProcessor(otlp_exporter)
        trace.get_tracer_provider().add_span_processor(span_processor)

        # pyroscope
        pyroscope.configure(
            application_name="referral-service",
            server_address="http://pyroscope.pyroscope:4040",
        )

        return tracer
