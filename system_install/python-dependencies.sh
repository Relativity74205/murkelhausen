#!/usr/bin/env bash
source ./print_color.sh
log=install.log

print_color "green" "apt update"
sudo apt-get update | sudo tee -a $log > /dev/null
print_color "green" "Installing python dependencies..."
sudo apt-get -y install make build-essential libssl-dev zlib1g-dev \
libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm \
libncursesw5-dev xz-utils tk-dev libxml2-dev libxmlsec1-dev libffi-dev liblzma-dev | sudo tee -a $log > /dev/null
print_color "green" "Installing python dependencies complete."

