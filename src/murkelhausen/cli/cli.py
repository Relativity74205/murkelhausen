from logging import getLogger
import json

import click

from murkelhausen import __version__, cfg
from murkelhausen.config import cli_loader
from murkelhausen.util import logger
from murkelhausen.weather import owm, nmi

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


@cli.command("query-owm")
@click.argument("city_name")
def query_owm(city_name: str):
    cities = [city for city in cfg.cities if city.name == city_name]
    if len(cities) == 1:
        owm_data = owm.query_one_call_api(cities[0], cfg.weather_owm)
    else:
        raise ValueError(f"{city_name=} not found in config.")
    print(json.dumps(owm_data, indent=4))


@cli.command("query-nmi")
@click.argument("city_name")
def query_nmi(city_name: str):
    cities = [city for city in cfg.cities if city.name == city_name]
    if len(cities) == 1:
        nmi_data = nmi.query_locationforecast(cities[0], cfg.weather_nmi)
    else:
        raise ValueError(f"{city_name=} not found in config.")
    print(json.dumps(nmi_data, indent=4))
