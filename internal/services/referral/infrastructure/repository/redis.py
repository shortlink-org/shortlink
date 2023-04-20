import redis as Redis

class Repository:
  def __init__(self, host: str):
    pool = Redis.ConnectionPool(host=host, port=6379, db=0)
    self._redis = Redis.Redis(connection_pool=pool)

