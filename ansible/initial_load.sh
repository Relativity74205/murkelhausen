#!/bin/bash

mkdir -p .kube
mkdir -p .aws
mkdir -p .ssh
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/configs/kube_config ~/.kube/config
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/configs/aws_config ~/.aws/config
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/ssh/* ~/.ssh
chmod 600 ~/.ssh/*
eval `ssh-agent`
ssh-add ~/.ssh/github_private
mkdir -p ~/dev/murkelhausen
git clone git@github.com:Relativity74205/murkelhausen.git ~/dev/murkelhausen


sudo apt update && sudo apt upgrade -y && sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=/home/arkadius/.local/bin/:$PATH' >> ~/.bashrc
source ~/.bashrc
ansible --version


