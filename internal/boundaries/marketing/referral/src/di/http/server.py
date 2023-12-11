"""HTTP endpoint for the infrastructure."""

from dependency_injector import providers
from quart import Quart

class QuartProvider(providers.Provider):
    """Quart provider."""
    @staticmethod
    def _provide(*args, **kwargs):
        app = Quart(__name__)
        return app
