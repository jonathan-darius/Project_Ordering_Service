import os
from dotenv import load_dotenv, find_dotenv

load_dotenv(find_dotenv())


class Settings:
    RABBIT_URL: str = os.getenv("RABBIT_URL")
    RABBIT_PORT: int = os.getenv("RABBIT_PORT")
    RABBIT_USERNAME: str = os.getenv("RABBIT_USERNAME")
    RABBIT_PASSWORD: str = os.getenv("RABBIT_PASSWORD")
    RABBIT_QUEUE: str = os.getenv("RABBIT_QUEUE")
    MAIL_NAME: str = os.getenv('MAIL_NAME')
    MAIL_PASSWORD: str = os.getenv("MAIL_PASSWORD")


settings = Settings()
