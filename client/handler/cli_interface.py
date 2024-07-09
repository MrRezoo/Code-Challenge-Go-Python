from rich.panel import Panel

from config.app import console
from services.commands.user import get_all_users
from services.queries.user import register, deposit, withdraw, bet


def run_cli():
    console.print(Panel.fit(
        "[bold magenta]Welcome to the Challenge![/bold magenta]\n"
        "[bold cyan]Commands available:[/bold cyan]\n"
        "[magenta]1.[/magenta] register - Register a new user\n"
        "[magenta]2.[/magenta] deposit - Deposit funds into a user account\n"
        "[magenta]3.[/magenta] withdraw - Withdraw funds from a user account\n"
        "[magenta]4.[/magenta] bet - Perform a betting operation (placeholder)\n"
        "[magenta]5.[/magenta] all - Retrieve all users",
        title="[red] Challenge CLI [/red]",
        border_style="red",
    ))

    while True:
        command = input("Enter command (register/deposit/withdraw/bet/all/exit): ").strip().lower()

        if command in ["1", "register"]:
            register()
        elif command in ["2", "deposit"]:
            deposit()
        elif command in ["3", "withdraw"]:
            withdraw()
        elif command in ["4", "bet"]:
            bet()
        elif command in ["5", "all"]:
            get_all_users()
        elif command == "exit":
            console.print("[bold]Exiting CLI. Goodbye![/bold]")
            break
        else:
            console.print("[red]Invalid command. Please enter a valid command.[/red]")
