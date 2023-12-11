"""Config module."""

from dotenv import load_dotenv

class Config:
    """Config class."""
    @staticmethod
    def _provide(*args, **kwargs):
        return load_dotenv()
