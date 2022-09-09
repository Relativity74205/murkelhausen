#!/bin/bash

sudo chown -R arkadius:arkadius ~/.ssh
chmod 600 ~/.ssh/*
eval `ssh-agent`
ssh-add ~/.ssh/github
mkdir -p ~/dev/murkelhausen
git clone git@github.com:Relativity74205/murkelhausen.git ~/dev/murkelhausen


sudo apt update
sudo apt upgrade -y
sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=/home/arkadius/.local/bin/:$PATH' >> ~/.bashrc
source ~/.bashrc
ansible --version


