import src.infrastructure.repository.referral.uow as uow
from src.infrastructure.repository.referral.repository_redis import Repository

class RedisUnitOfWork(uow.AbstractUnitOfWork):
    def __init__(self, repository: Repository):
        self.session = repository

    def __enter__(self):
        self.referral = self.session
        self.pipeline = self.session._redis.pipeline(transaction=True)
        return super().__enter__()

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.pipeline.reset()
        super().__exit__(exc_type, exc_val, exc_tb)

    def commit(self):
        self.pipeline.execute()

    def rollback(self):
        self.pipeline.reset()
