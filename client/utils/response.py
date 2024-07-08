def print_response(response):
    if response.status_code == 200:
        console.print("[green]Success[/green]")
    else:
        console.print(f"[red]Error: {response.status_code}[/red]")
        console.print(response.text)
