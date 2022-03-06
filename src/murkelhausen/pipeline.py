from prefect import Flow

from murkelhausen.weather import nmi, owm
from murkelhausen.util import backend
from murkelhausen import cfg


with Flow("Test") as flow:
    city = backend.get_city_object("Mülheim")
    nmi_data = nmi.query_compact(city, cfg.weather_nmi)
    backend.save_json("nmi_compact", nmi_data)
    owm_data = owm.query_weather(city, cfg.weather_owm)
    backend.save_json("owm_weather", owm_data)

flow.register(project_name="murkelhausen")
# flow.run()
