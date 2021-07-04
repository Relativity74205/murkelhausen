# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.8-slim

COPY wheelhouse/* wheelhouse/

RUN pip install wheelhouse/* --no-deps --no-index

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]