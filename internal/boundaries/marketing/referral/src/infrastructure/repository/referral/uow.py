"""Unit of Work interface."""

import abc
from src.infrastructure.repository.referral import repository


class AbstractUnitOfWork(abc.ABC):
    """Abstract Unit of Work."""

    referral: repository.AbstractRepository

    def __enter__(self) -> "AbstractUnitOfWork":
        """Enter Unit of Work."""
        return self

    def __exit__(self, *args):
        """Exit Unit of Work."""
        self.rollback()

    @abc.abstractmethod
    def commit(self) -> None:
        """Commit transaction."""
        raise NotImplementedError

    @abc.abstractmethod
    def rollback(self) -> None:
        """Rollback transaction."""
        raise NotImplementedError
