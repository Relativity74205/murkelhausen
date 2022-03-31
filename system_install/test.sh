#!/usr/bin/env bash
source ./print_color.sh

if ! grep -xq "plugins=\(.*pyenv.*\)" ~/.zshrc
then
  print_color "red" "foo"
else
  print_color "green" "bar"
fi