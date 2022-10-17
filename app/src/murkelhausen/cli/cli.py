import click

from murkelhausen import __version__
from murkelhausen.util import logger


@click.group(invoke_without_command=True)
@click.pass_context
# fmt: off
@click.option(
    "-v",
    "--version",
    is_flag=True,
    help="Print murkelhausen' version number and exit."
)
# fmt: on
def cli(ctx: click.Context, version: bool):
    """Command line interface for murkelhausen development utilities.

    Enter one of the subcommands to execute them, or run their respective --help
    to read about their usage."""

    if version:
        click.echo(f"murkelhausen {__version__}")
        exit(0)
    elif ctx.invoked_subcommand is None:
        cli(["--help"])
    logger.setup_logging()


@cli.command()
@click.option("-c", "--count", default=1, help="Number of greetings.")
@click.option("-n", "--name", prompt="Your name", help="The person to greet.")
def hello(count: int, name: str):  # pragma: no cover
    for x in range(count):
        click.echo(f"Hello {name}!")
