"""
https://openweathermap.org/api/one-call-api
"""
from typing import Dict

import requests

from murkelhausen.config import WeatherOWM


def query_one_call_api(owm_settings: WeatherOWM) -> Dict:
    query_params = {
        "lat": owm_settings.gps_lat,
        "lon": owm_settings.gps_lon,
        "appid": owm_settings.api_key,
        "units": owm_settings.units,

    }
    r = requests.get(owm_settings.url, params=query_params)

    if r.status_code == 200:
        return r.json()
    else:
        raise RuntimeError(
            f"Query to openweatherapi one call api returned non 200 status code: "
            f"status_code: {r.status_code}"
            f"response_text: {r.text}"
        )


def get_weather_map(layer: str, owm_settings: WeatherOWM):
    """
    https://openweathermap.org/api/weathermaps
    https://github.com/google/maps-for-work-samples/blob/master/samples/maps/OpenWeatherMapLayer/OpenWeatherMapLayer.pdf
    """
    pass


def query_air_pollution():
    """
    https://openweathermap.org/api/air-pollution
    """
    pass