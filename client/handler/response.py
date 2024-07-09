from rich.table import Table

from config.app import console


def print_response(response):
    if response.status_code == 200 and response.json().get("code") == 200:
        console.print("[green]Success[/green]")
        console.print("[green]Response JSON:[/green]", response.json())
    else:
        console.print("[red]Response JSON:[/red]", response.json())


def show_users(users):
    table = Table(title="User List")
    table.add_column("ID", justify="center", style="cyan")
    table.add_column("Balance", justify="center", style="cyan")

    for user in users:
        table.add_row(str(user["id"]), str(user["balance"]))

    console.print(table)
