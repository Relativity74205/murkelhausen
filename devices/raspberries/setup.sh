#!/bin/bash

# create folders
mkdir -p unifi_config
mkdir -p nas
mkdir -p pihole

# setup nas connection
echo "//192.168.1.19/arkadiusbackup /home/pi/nas cifs credentials=/home/pi/.smbcredentials,vers=2.0,uid=1000,gid=1000 0 0" | sudo tee -a /etc/fstab
echo -e "username=arkadius\npassword=$PASSWORD" >> .smbcredentials
sudo mount -a

# setup git (incl. adding git ssh key to ssh-agent)
sudo apt install git
# TODO: ssh keys


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


# install pi-hole (https://github.com/pi-hole/pi-hole/#one-step-automated-install)
curl -sSL https://install.pi-hole.net | bash
