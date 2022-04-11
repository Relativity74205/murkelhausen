#!/usr/bin/env bash

source ./print_color.sh
log=install.log

print_color "green" "apt update"
sudo apt-get update | sudo tee -a $log > /dev/null
#print_color "green" "Installing apt-utils..."
#sudo apt-get -y install apt-utils | sudo tee -a $log > /dev/null
#print_color "green" "Installing apt-utils complete."
print_color "green" "Installing wget curl git nano htop..."
sudo apt-get -y install wget curl git nano htop | sudo tee -a $log > /dev/null
print_color "green" "Installing wget curl git nano htop complete."

print_color "green" "Installing pavucontrol..."
sudo apt install pavucontrol
print_color "green" "Installing pavucontrol complete."
