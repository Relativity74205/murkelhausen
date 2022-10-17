#--build-arg REGISTRY="gitlab.office.auxmoney.com:4567"

ARG REGISTRY
FROM ${REGISTRY}/auxmoney/docker-images/python-dev:3.10-latest as dev

WORKDIR /usr/app

COPY src src
COPY tests tests
COPY poetry.lock ./pyproject.toml ./

RUN poetry config virtualenvs.in-project true && \
    poetry install && \
    poetry export -f requirements.txt --without-hashes -o requirements.txt && \
    poetry build -f wheel && \
    poetry run pip install --upgrade pip && \
    poetry run pip wheel -w wheels -r requirements.txt && \
	mv dist/* wheels


FROM ${REGISTRY}/auxmoney/docker-images/python:3.10-latest AS prod

WORKDIR /usr/app

COPY --from=dev /usr/app/wheels ./wheels/

RUN pip install wheels/* --no-deps --no-index

USER 1000
