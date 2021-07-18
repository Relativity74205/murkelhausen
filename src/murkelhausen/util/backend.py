from murkelhausen import cfg
from murkelhausen.config import City


def get_city_object(city_name: str) -> City:
    """Retrieves the City object from the config for the given city name.

    Raises ValueError in case the city was not found in the config.
    """
    cities = [city for city in cfg.app.cities if city.name == city_name]
    if len(cities) == 1:
        return cities[0]
    else:
        raise ValueError(f"{city_name=} not found in config.")
