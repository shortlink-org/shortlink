"""HTTP endpoint for the infrastructure."""

from dependency_injector import providers
from quart import Quart

class QuartProvider(providers.Provider):
    def _provide(self, *args, **kwargs):
        app = Quart(__name__)
        return app
