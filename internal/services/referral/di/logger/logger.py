from loguru import logger

async def logger() -> logger:
  return logger.add(serialize=True)
