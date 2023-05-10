import src.infrastructure.repository.referral.uow as uow
from src.infrastructure.repository.referral.repository_ram import Repository

class RamUnitOfWork(uow.AbstractUnitOfWork):
    def __init__(self, referral_repo: Repository):
        self.session = referral_repo

    def __enter__(self):
        self.referral = self.session
        return super().__enter__()

    def __exit__(self, *args):
        super().__exit__(*args)

    def commit(self):
        pass

    def rollback(self):
        pass

