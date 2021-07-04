# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.9-slim

# RUN apt update && apt install -y curl
# RUN curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py > get-poetry.py && \
#    python get-poetry.py --version 1.1.6
RUN pip install poetry
ENV PATH="${PATH}:/root/.poetry/bin"

COPY pyproject.toml src/ ./

RUN pip install pydantic && \
    poetry config virtualenvs.create false && \
    poetry install --no-dev --no-interaction --no-ansi

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]