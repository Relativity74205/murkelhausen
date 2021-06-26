import logging
from logging.config import dictConfig

from murkelhausen import cfg

log = logging.getLogger(__name__)

ONELINE_FORMATTER = "%(asctime)s [%(levelname)s] %(name)s: %(message)s"


def setup_logging():
    """Set up basic logging to stdout."""
    logging.config.dictConfig(
        {
            "version": 1,
            "disable_existing_loggers": False,
            "formatters": {
                "oneline": {
                    "format": ONELINE_FORMATTER
                },
            },
            "handlers": {
                "console": {
                    "level": cfg.loglevel,
                    "formatter": "oneline",
                    "class": "logging.StreamHandler",
                    "stream": "ext://sys.stdout",
                },
            },
            "loggers": {
                "": {
                    "handlers": ["console"],
                    "level": cfg.loglevel,
                    "propagate": True,
                },
            },
        }
    )
    log.info("Logger configuration set up successfully.")
