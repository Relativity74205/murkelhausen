#!/bin/bash

chmod 600 .ssh/*
eval `ssh-agent`
ssh-add .ssh/github
mkdir -p dev
git clone git@github.com:Relativity74205/murkelhausen.git ~/dev


sudo apt update
sudo apt upgrade -y
sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=/home/arkadius/.local/bin/:$PATH' >> .bashrc
source .bashrc
ansible --version


