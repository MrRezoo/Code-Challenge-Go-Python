from rich.table import Table

from config.app import console


def print_response(response):
    if response.status_code == 200:
        console.print("[green]Success[/green]")
    else:
        console.print(f"[red]Error: {response.status_code}[/red]")
        console.print(response.text)


def show_users(users):
    table = Table(title="User List")
    table.add_column("ID", justify="center", style="cyan")
    table.add_column("Name", justify="center", style="cyan")
    table.add_column("Balance", justify="center", style="cyan")

    for user in users:
        table.add_row(str(user["id"]), user["name"], str(user["balance"]))

    console.print(table)
