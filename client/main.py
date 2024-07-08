"""
This module defines the CLI (Command Line Interface) for the user management system.
It provides commands for user registration, deposit, withdrawal, betting, and retrieving all users.
The CLI uses a loop to continuously prompt the user for commands until the 'exit' command is entered.

Functions:
- run_cli(): Starts the CLI loop and handles user commands.

Imports:
- register, deposit, withdraw, bet, get_all_users from services.queries.user: Functions for user operations.
- console from config.app: Console object for printing messages with style.
"""
from rich.panel import Panel

from config.app import console

if __name__ == "__main__":
    from cli.interface import run_cli

    panel = Panel.fit("[bold cyan]User Management CLI[/bold cyan]", title="User Management", border_style="red")
    console.print(panel)
    run_cli()
