from logging import getLogger

import click

from murkelhausen import __version__, cfg
from murkelhausen.util import logger

log = getLogger(__name__)


@click.group(invoke_without_command=True)
@click.option(
    "-c",
    "--cli_config",
    help="Config can be overridden with this option. The config parameters have to be "
         "passed according to the following syntax: "
         "'-c app__loglevel=ERROR'.",
)
@click.option(
    "-v", "--version", is_flag=True, help="Print murkelhausen' version number and exit."
)
@click.pass_context
def cli(ctx, version: bool, cli_config: str):
    """Command line interface for murkelhausen.

    Enter one of the subcommands to execute them, or run their respective --help
    to read about their usage.
    """
    if cli_config:
        cli_loader(cli_config)

    if version:
        click.echo(f"murkelhausen {__version__}")
        exit(0)
    elif ctx.invoked_subcommand is None:
        cli(["--help"])
    logger.setup_logging()


@cli.command()
def test():
    click.echo("TEST")
    log.info("TEST.")
