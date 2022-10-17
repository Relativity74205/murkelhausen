"""Settings and functions related to our logging setup."""

import logging
from functools import wraps
from logging.config import dictConfig
from typing import Any, Callable, TypeVar, cast


from murkelhausen.config import config

log = logging.getLogger(__name__)

#: **Function-type identifier** - By asserting to mypy that a decorator only ever consumes
#: functions, the wrapped function will keep its original signature.
F = TypeVar("F", bound=Callable[..., Any])
ONELINE_FORMATTER_STRING = "%(asctime)s [%(levelname)s] %(name)s: %(message)s"


def run_once(func: F) -> F:
    """Decorator which ensures that a function can only be run a single time.

    The value that the original function returns is cached internally, and will
    be returned on consecutive calls without running the original function's code.

    Notes:
        By introspecting the wrapper's internals, the run_once logic can be
        circumvented. Every wrapped function has a dunder attribute
        `__wrapped__`, which is used to persist whether the internal
        function was already run or not. Hence, manually setting its attribute
        `has_run` to False will allow additional executions.

    """

    @wraps(func)
    def wrapper(*args, **kwargs):
        if not func.has_run:
            func.result = func(*args, **kwargs)
            func.has_run = True
        return func.result

    func.has_run = False  # type: ignore
    # casts have no influence on runtime behavior and only serve as assertion for mypy
    return cast(F, wrapper)


def at_least(loglevel, min_log_level="CRITICAL"):
    """Helper function to enforce minimum log levels."""
    level_order = list(logging._nameToLevel)[::-1]
    min_log_level = max([level_order.index(loglevel), level_order.index(min_log_level)])
    return level_order[min_log_level]


@run_once
def setup_logging():
    """Set up basic logging to stdout."""
    logging.config.dictConfig(
        {
            "version": 1,
            "disable_existing_loggers": False,
            "formatters": {
                "oneline": {"format": ONELINE_FORMATTER_STRING},
            },
            "handlers": {
                "console": {
                    "level": config.app.loglevel,
                    "formatter": "oneline",
                    "class": "logging.StreamHandler",
                    "stream": "ext://sys.stdout",
                },
            },
            "loggers": {
                "murkelhausen": {
                    "handlers": ["console"],
                    "level": config.app.loglevel,
                    "propagate": True,
                },
                # the __main__ logger is needed when certain files like train/main.py are run standalone
                "__main__": {
                    "handlers": ["console"],
                    "level": config.app.loglevel,
                    "propagate": True,
                },
                "paromiko": {
                    "handlers": ["console"],
                    "level": at_least(config.app.loglevel, "INFO"),
                    "propagate": True,
                },
            },
        }
    )
    log.info("Logger configuration set up successfully.")
