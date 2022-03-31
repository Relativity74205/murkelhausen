#!/usr/bin/env bash
source ./print_color.sh
log=install.log

print_color "green" "apt update"
sudo apt-get update | sudo tee -a $log > /dev/null
#print_color "green" "Installing apt-utils..."
#sudo apt-get -y install apt-utils | sudo tee -a $log > /dev/null
#print_color "green" "Installing apt-utils complete."
print_color "green" "Installing wget curl git nano..."
sudo apt-get -y install wget curl git nano | sudo tee -a $log > /dev/null
print_color "green" "Installing wget curl git nano complete."