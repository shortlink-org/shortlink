import sys

from loguru import logger
from dependency_injector import providers

class LoguruJsonProvider(providers.Provider):
    def _provide(self, *args, **kwargs):
      logger.remove()
      logger_format = (
        "{{'time': '{time}', 'level': '{level}', 'message': '{message}', 'file': '{file}:{line}', 'extra': {extra}}}"
      )
      logger.add(lambda msg: print(msg.rstrip()), format=logger_format,
                 filter=lambda record: self.remove_elapsed_repr(record))
      return logger


    @staticmethod
    def remove_elapsed_repr(record):
      if 'elapsed' in record['extra']:
        del record['extra']
      return True
