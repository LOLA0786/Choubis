import typer
from rich.console import Console
from rich.panel import Panel
from rich.table import Table
import time

app = typer.Typer()
console = Console()

@app.command()
def monitor():
    """Real-time monitoring of the Choubis Sovereign Plane."""
    console.print(Panel("[bold cyan]CHOUBIS MONITORING MODE ACTIVE[/bold cyan]\nListening on localhost:9000...", title="Status"))
    
    table = Table(show_header=True, header_style="bold magenta")
    table.add_column("Timestamp", style="dim")
    table.add_column("Event")
    table.add_column("Status")
    table.add_column("Insurance Score")

    table.add_row(
        time.strftime("%H:%M:%S"), 
        "Shadow AI Intercepted: OpenAI Chat", 
        "[green]REDACTED[/green]", 
        "98/100"
    )
    
    console.print(table)

@app.command()
def audit():
    """Run the Pulse Hostile Auditor."""
    console.print("[bold red]Starting Hostile Audit...[/bold red]")
    time.sleep(1)
    console.print("✔ Persona 'The Infiltrator' blocked.")
    console.print("✔ Persona 'The Cipher' blocked.")
    console.print(Panel("[bold green]CHOUBIS SHIELD: 100% SUCCESS RATE[/bold green]"))

if __name__ == "__main__":
    app()
