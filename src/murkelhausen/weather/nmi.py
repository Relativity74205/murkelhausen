"""
https://api.met.no/weatherapi/
https://www.yr.no/en/details/table/2-6553027/Germany/North%20Rhine-Westphalia/D%C3%BCsseldorf%20District/M%C3%BClheim
"""
from typing import Dict

import requests

from murkelhausen.config import WeatherNMI, City


def query_locationforecast(city: City, nmi_settings: WeatherNMI) -> Dict:
    query_params = {
        "lat": city.gps_lat,
        "lon": city.gps_lon,
    }
    r = requests.get(nmi_settings.url, params=query_params)
    print(r.url)

    if r.status_code == 200:
        return r.json()
    else:
        raise RuntimeError(
            f"Query to norwegian meteorological institute one call api returned non 200 status code for city {city.name}:"
            f"status_code: {r.status_code}"
            f"response_text: {r.text}"
        )
