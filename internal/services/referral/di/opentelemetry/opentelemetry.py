from dependency_injector import providers
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter

class OpenTelemetryProvider(providers.Provider):
    def _provide(self, *args, **kwargs):
        trace.set_tracer_provider(TracerProvider())
        tracer = trace.get_tracer(__name__)

        otlp_exporter = OTLPSpanExporter()
        span_processor = BatchSpanProcessor(otlp_exporter)
        trace.get_tracer_provider().add_span_processor(span_processor)

        return tracer
