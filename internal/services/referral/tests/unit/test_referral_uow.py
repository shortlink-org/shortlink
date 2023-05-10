"""Unit tests for referral unit of work."""

class FakeUnitOfWork:
    """Fake unit of work for testing."""
    def __init__(self):
        """Initialize fake unit of work."""
        # self.referrals = FakeRepository()
        self.committed = False

    def commit(self):
        """Commit transaction."""
        self.committed = True

    def rollback(self):
        """Rollback transaction."""
        pass
