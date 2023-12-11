"""Redis Unit of Work."""

from src.infrastructure.repository.referral import uow
from src.infrastructure.repository.referral.repository_redis import Repository

class RedisUnitOfWork(uow.AbstractUnitOfWork):
    """Redis Unit of Work."""

    def __init__(self, repository: Repository):
        """Initialize Redis Unit of Work."""
        self.session = repository

    def __enter__(self):
        """Enter Redis Unit of Work."""
        self.referral = self.session
        self.pipeline = self.session._redis.pipeline(transaction=True)
        return super().__enter__()

    def __exit__(self, exc_type, exc_val, exc_tb):
        """Exit Redis Unit of Work."""
        self.pipeline.reset()
        super().__exit__(exc_type, exc_val, exc_tb)

    def commit(self):
        """Commit transaction."""
        self.pipeline.execute()

    def rollback(self):
        """Rollback transaction."""
        self.pipeline.reset()
