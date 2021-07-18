# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.9-slim as dev

# ggf gcc und make statt build-essential
RUN apt-get update && \
    apt-get -y install curl build-essential && \
    curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py > get-poetry.py && \
    python get-poetry.py --version 1.1.6
ENV PATH="${PATH}:/root/.poetry/bin"

WORKDIR /usr/app

COPY . .

RUN pip install --upgrade pip && \
    poetry install
RUN poetry build -f wheel && \
    poetry export -f requirements.txt --without-hashes -o requirements.txt && \
    poetry run pip wheel -w wheels -r requirements.txt && \
    mv dist/* wheels

FROM python:3.9-slim AS prod

WORKDIR /usr/app

COPY --from=dev /usr/app/wheels ./wheels/

RUN pip install --upgrade pip && \
    pip install wheels/* --no-deps --no-index

USER 1000

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]