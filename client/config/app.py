import os

from rich.console import Console

APP_ENV = os.getenv('APP_ENV', 'development')
if APP_ENV == "development":
    BASE_URL = "http://localhost:5005"
if APP_ENV == "docker":
    BASE_URL = "http://server:5005"

console = Console()
