"""Config parser module.

This module cascades through the different ways in which configuration
parameters for the application can be set. They follow the
standard priority of

 - default.toml (default parameters, src/murkelhausen/default.toml)
 - user.toml (top project path)
 - environment variable parameters

Parameters set through a method lower down the priority chain will
overwrite settings from higher up.
That is settings from the user.toml will overwrite the settings from default.toml.

Environment variables need to start with murkelhausen
following with the section they belong to and have the key name follow
after a double-underscore, e.g. murkelhausen_APP__LOGLEVEL.
"""

from logging import getLogger
from typing import Literal

from pydantic import BaseModel, BaseSettings

from murkelhausen.util.config_util import (
    default_toml_loader,
    user_toml_loader
)

log = getLogger(__name__)

loglevels = Literal["DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"]


class App(BaseModel, validate_assignment=True):
    loglevel: loglevels


class Settings(BaseSettings):
    app: App

    class Config:
        validate_assignment = True

        @classmethod
        def customise_sources(cls, init_settings, env_settings, file_secret_settings):
            return init_settings, env_settings, user_toml_loader, default_toml_loader

        env_prefix = "MURKELHAUSEN_"
        env_nested_delimiter = "__"


config = Settings()
