class FakeUnitOfWork:
    def __init__(self):
        self.referrals = FakeRepository()

    def commit(self):
        self.committed = True

    def rollback(self):
        pass
