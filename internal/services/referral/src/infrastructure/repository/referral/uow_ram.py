"""Redis Unit of Work."""

from src.infrastructure.repository.referral import uow
from src.infrastructure.repository.referral.repository_ram import Repository

class RamUnitOfWork(uow.AbstractUnitOfWork):
    """Redis Unit of Work."""

    def __init__(self, referral_repo: Repository):
        """Initialize Redis Unit of Work."""
        self.session = referral_repo

    def __enter__(self):
        """Enter Redis Unit of Work."""
        self.referral = self.session
        return super().__enter__()

    def __exit__(self, *args):
        """Exit Redis Unit of Work."""
        super().__exit__(*args)

    def commit(self):
        """Commit transaction."""
        pass

    def rollback(self):
        """Rollback transaction."""
        pass

