"""Logger provider."""

from loguru import logger
from dependency_injector import providers
from opentelemetry.trace import Span
from opentelemetry.context import Context

class LoguruJsonProvider(providers.Provider):
    """Logger provider."""

    def _provide(self, *args, **kwargs):
        logger.remove()
        logger_format = (
          "{{'time': '{time}', 'level': '{level}', 'message': '{message}', 'file': '{file}:{line}', 'trace_id': '{"
         "trace_id}', 'extra': {extra}}}"
        )
        logger.add(lambda msg: print(msg.rstrip()), format=logger_format,
                 filter=lambda record: self.process_record(record))
        return logger


    @staticmethod
    def process_record(record):
        """Process log record."""
        # Remove elapsed
        if 'elapsed' in record['extra']:
            del record['extra']['elapsed']

        # Add trace_id
        current_span = Span.get_span_context(Context)
        if current_span:
            record['trace_id'] = current_span.trace_id
        else:
            record['trace_id'] = 'None'

        return True

