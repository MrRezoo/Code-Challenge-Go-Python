import requests

from config.app import BASE_URL, console
from utils.response import print_response


def get_all_users():
    response = requests.get(f"{BASE_URL}/user/all")
    if response.status_code == 200:
        users = response.json()["result"]
        show_users(users)
    else:
        print_response(response)


def show_users(users):
    table = Table(title="User List")
    table.add_column("ID", justify="center", style="cyan")
    table.add_column("Name", justify="center", style="cyan")
    table.add_column("Balance", justify="center", style="cyan")

    for user in users:
        table.add_row(str(user["id"]), user["name"], str(user["balance"]))

    console.print(table)
