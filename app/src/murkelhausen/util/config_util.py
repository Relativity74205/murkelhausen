import importlib.resources
from logging import getLogger
from pathlib import Path
from typing import Any, MutableMapping

import toml

log = getLogger(__name__)


def default_toml_loader(*_) -> MutableMapping[str, Any]:
    """Loads default variables from src/murkelhausen/default.toml to config"""
    try:
        with importlib.resources.path(
                "murkelhausen", "default.toml"
        ) as default_config:
            with default_config.open() as f:
                return toml.load(f)
    except FileNotFoundError:
        log.error(f"Default config expected at src/murkelhausen/default.toml, but not found.")
        raise RuntimeError("Config not found, aborting!")


def user_toml_loader(*_) -> MutableMapping[str, Any]:
    """Loads user-specified variables from user.toml to config.

    user.toml has to be placed in the top project path and follows the syntax
    of the default.toml (located in src/murkelhausen/default.toml).
    user.toml must not be committed to git and is included in .gitignore.
    """
    # project_root must be adapted in case this file is moved in the project structure
    project_root = Path(__file__).parent.parent.parent.parent
    try:
        with open(project_root / "user.toml") as f:
            log.info("User.toml found, loading.")
            return toml.load(f)
    except FileNotFoundError:
        log.warning(f"User config not found. It is expected at {str(project_root / 'user.toml')}.")
        return {}
