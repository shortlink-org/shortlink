from dotenv import load_dotenv

class Config:
    def _provide(self, *args, **kwargs):
        return load_dotenv()
