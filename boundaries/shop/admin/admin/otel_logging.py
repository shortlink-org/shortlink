import logging
from opentelemetry import trace

class CustomLogRecord(logging.LogRecord):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        current_span = trace.get_current_span()
        if current_span is not None:
            self.otelTraceID = current_span.get_span_context().trace_id
            self.otelSpanID = current_span.get_span_context().span_id
        else:
            self.otelTraceID = None
            self.otelSpanID = None
