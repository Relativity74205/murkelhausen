# https://stackoverflow.com/questions/53835198/integrating-python-poetry-with-docker
FROM python:3.9-slim

# COPY requirements.txt requirements.txt
COPY dist/* dist/

# RUN pip install -r requirements.txt --no-deps
RUN pip install asgiref==3.4.1 --no-deps
RUN pip install certifi==2021.5.30 --no-deps
RUN pip install chardet==4.0.0 --no-deps
RUN pip install click==8.0.1 --no-deps
RUN pip install fastapi==0.65.3 --no-deps
RUN pip install h11==0.12.0 --no-deps
RUN pip install idna==2.10 --no-deps
RUN pip install pydantic==1.8.2 --no-deps
RUN pip install requests==2.25.1 --no-deps
RUN pip install starlette==0.14.2 --no-deps
RUN pip install toml==0.10.2 --no-deps
RUN pip install typing-extensions==3.10.0.0 --no-deps
RUN pip install urllib3==1.26.6 --no-deps
RUN pip install uvicorn==0.14.0 --no-deps
RUN pip install dist/murkelhausen-0.1.0-py3-none-any.whl --no-deps

EXPOSE 5000

ENTRYPOINT [ "murkelhausen", "serve", "-h", "0.0.0.0" ]