FROM ubuntu:latest

ENV TZ=Europe/Berlin

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install tzdata \
    && apt-get install -y sudo lsb-core

RUN adduser --disabled-password --gecos '' docker \
    && adduser docker sudo \
    && echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER docker

WORKDIR /home/docker

COPY --chown=docker:docker python-dependencies.sh apt-installs.sh print_color.sh ./
RUN sudo ./apt-installs.sh
RUN sudo ./python-dependencies.sh

COPY --chown=docker:docker zsh.sh ./
RUN ./zsh.sh
COPY --chown=docker:docker pyenv.sh ./
RUN ./pyenv.sh
COPY --chown=docker:docker atuin.sh ./
RUN ./atuin.sh