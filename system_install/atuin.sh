#!/usr/bin/env bash
source ./print_color.sh
log=install.log

# shellcheck disable=SC2016
rc_string='eval "$(atuin init zsh)"'

if ! command -v atuin &> /dev/null
then
  print_color "green" "Installing atuin..."
  curl https://raw.githubusercontent.com/ellie/atuin/main/install.sh | bash | sudo tee -a $log > /dev/null
  print_color "green" "Installing atuin complete."
else
  print_color "yellow" "atuin already installed on system."
fi

# shellcheck disable=SC2016
if ! grep -Fxq "$rc_string" ~/.zshrc
then
  echo "$rc_string" >> ~/.zshrc
  print_color "green" "Modified ~/.zshrc with '$rc_string'"
else
  print_color "yellow" "atuin string ($rc_string) is already in ~/.zshrc"
fi
