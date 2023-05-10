class FakeUnitOfWork:
    def __init__(self):
        # self.referrals = FakeRepository()
        self.committed = False

    def commit(self):
        self.committed = True

    def rollback(self):
        pass
