from rich.panel import Panel

from config.app import console
from services.commands.user import get_all_users
from services.queries.user import register, deposit, withdraw, bet


def run_cli():
    console.print(Panel.fit(
        "[bold magenta]Welcome to User Management CLI![/bold magenta]\n"
        "[bold cyan]Commands available:[/bold cyan]\n"
        "1. register - Register a new user\n"
        "2. deposit - Deposit funds into a user account\n"
        "3. withdraw - Withdraw funds from a user account\n"
        "4. bet - Perform a betting operation (placeholder)\n"
        "5. all - Retrieve all users",
        title="[red]User Management CLI[/red]",
        border_style="red"
    ))

    while True:
        command = input("Enter command (register/deposit/withdraw/bet/all/exit): ").strip().lower()

        if command == "register":
            register()
        elif command == "deposit":
            deposit()
        elif command == "withdraw":
            withdraw()
        elif command == "bet":
            bet()
        elif command == "all":
            get_all_users()
        elif command == "exit":
            console.print("[bold]Exiting CLI. Goodbye![/bold]")
            break
        else:
            console.print("[red]Invalid command. Please enter a valid command.[/red]")
