import os
from pydantic_settings import BaseSettings if 'pydantic_settings' in globals() else object

class Settings:
    def __init__(self):
        self.app_env: str = os.getenv("APP_ENV", "local")
        self.database_url: str = os.getenv("DATABASE_URL", "")
        self.openai_api_key: str = os.getenv("OPENAI_API_KEY", "")
        self.google_vision_key: str = os.getenv("GOOGLE_VISION_KEY", "")

settings = Settings()
