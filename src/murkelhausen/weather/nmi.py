"""
https://api.met.no/weatherapi/
https://www.yr.no/en/details/table/2-6553027/Germany/North%20Rhine-Westphalia/D%C3%BCsseldorf%20District/M%C3%BClheim
"""
from typing import Dict, Any
from logging import getLogger

import requests

from murkelhausen.config import WeatherNMI, City

log = getLogger(__name__)


def query_locationforecast(city: City, nmi_settings: WeatherNMI) -> Dict:
    query_params: Dict[str, Any] = {
        "lat": city.gps_lat,
        "lon": city.gps_lon,
    }
    headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64)"}
    r = requests.get(nmi_settings.url_short, params=query_params, headers=headers)
    log.debug(f"Following URL is used for querying NMI API: {r.url}.")

    if r.status_code == 200:
        return_dict: Dict = r.json()
        return return_dict
    else:
        raise RuntimeError(
            f"Query to norwegian meteorological institute one call api returned non 200 status code for city {city.name}: "
            f"status_code: {r.status_code} "
            f"response_text: {r.text}."
        )
