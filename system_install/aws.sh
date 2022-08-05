#!/usr/bin/env bash

sudp apt install unzip
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

sed -i 's/plugins=(\(.*\))/plugins=(\1 aws)/' ~/.zshrc
