import requests

from config.app import BASE_URL
from handler.response import print_response, show_users


def get_all_users():
    response = requests.get(f"{BASE_URL}/user/all")
    if response.status_code == 200:
        users = response.json()["result"]
        show_users(users)
    else:
        print_response(response)
