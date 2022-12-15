#!/bin/bash

# create folders
mkdir -p unifi_config
mkdir -p backup

# install pi-hole (https://github.com/pi-hole/pi-hole/#one-step-automated-install)
curl -sSL https://install.pi-hole.net | bash

# install gravity-sync (https://github.com/vmstan/gravity-sync/wiki/Installing)
curl -sSL https://raw.githubusercontent.com/vmstan/gs-install/main/gs-install.sh | bash
# and run config
gravity-sync config

# rasp2
# install docker (https://docs.docker.com/engine/install/debian/)
sudo apt-get update
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo usermod -aG docker pi

# setup unifi
docker compose /home/pi/docker-compose.yml up -d
