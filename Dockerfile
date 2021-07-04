# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.9-slim

COPY requirements.txt requirements.txt
COPY dist/* dist/

RUN pip install -r requirements.txt --no-deps && \
    pip install dist/murkelhausen-0.1.0-py3-none-any.whl --no-deps

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]