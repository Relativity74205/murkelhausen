from logging import getLogger

from fastapi import FastAPI

from murkelhausen import __version__, cfg
from murkelhausen.util import logger, backend
from murkelhausen.weather import owm, nmi

app = FastAPI()
log = getLogger(__name__)
logger.setup_logging()


@app.get("/")
def root():
    return {"version": __version__}


@app.get("/health")
def health():
    return {"healthcheck": "OK"}


@app.get("/query_owm/{city_name}")
def query_owm(city_name: str):
    """Queries the API of OpenWeatherMap for the given city name."""
    city = backend.get_city_object(city_name)
    response = owm.query_one_call_api(city, cfg.weather_owm)

    return response


@app.get("/query_nmi/{city_name}")
def query_nmi(city_name: str):
    """Queries the API of the NMI for the given city name."""
    city = backend.get_city_object(city_name)
    response = nmi.query_locationforecast(city, cfg.weather_nmi)

    return response
