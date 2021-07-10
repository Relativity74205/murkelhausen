# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.9-slim

LABEL maintainer="arkadius@schuchhardt.com"

RUN apt-get update && \
    apt-get -y install curl && \
    curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py > get-poetry.py && \
    python get-poetry.py --version 1.1.6
ENV PATH="${PATH}:/root/.poetry/bin"

WORKDIR /usr/app

COPY pyproject.toml poetry.lock src/ /usr/app/

RUN poetry config virtualenvs.create false && \
    poetry install --no-dev --no-interaction --no-ansi

USER 1000

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]