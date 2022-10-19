import click

from murkelhausen import __version__
from murkelhausen.util import logger
from murkelhausen.mqtt import client


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


# TODO change command name
@cli.command()
def start_mqtt_client():
    client.main()
