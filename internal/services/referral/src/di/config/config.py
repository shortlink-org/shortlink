"""Config module."""

from dotenv import load_dotenv

class Config:
    """Config class."""
    def _provide(self, *args, **kwargs):
        return load_dotenv()
