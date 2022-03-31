#!/usr/bin/env bash
source ./print_color.sh
log=install.log

# shellcheck disable=SC2016
pyenv_config='
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init --path)"'

if [ -d "$HOME"/.pyenv ] && ! command -v pyenv &> /dev/null; then
  print_color "red" ".pyenv directory in home found, however pyenv command not in path."
  exit 1
elif [ -d "$HOME"/.pyenv ] && command -v pyenv &> /dev/null; then
  print_color "yellow" "pyenv already installed."
else
  print_color "green" "installing pyenv"
  curl -L https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | bash | sudo tee -a $log > /dev/null

  # https://github.com/pyenv/pyenv/#installation
  git clone https://github.com/momo-lab/xxenv-latest.git "$(pyenv root)"/plugins/xxenv-latest >> $log
fi

echo 'eval "$(pyenv init -)"' >> ~/.zshrc
echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.zshrc
if ! grep -xq "plugins=\(.*pyenv.*\)" ~/.zshrc; then
  sed -i 's/plugins=(\(.*\))/plugins=(\1 pyenv)/' ~/.zshrc
  print_color "green" "modified .zshrc::plugins for pyenv."
else
  print_color "yellow" ".zshrc::plugins for pyenv already modified."
fi

if ! grep -xq "$pyenv_config" ~/.profile; then
  echo "$pyenv_config" >> ~/.profile
  print_color "green" "Modified .profile for pyenv."
else
  print_color "yellow" ".profile already modified for pyenv."
fi

if [ ! -f ~/.zprofile ]; then
  touch ~/.zprofile
fi

if ! grep -xq "$pyenv_config" ~/.zprofile; then
  echo "$pyenv_config" >> ~/.zprofile
  print_color "green" "Modified .zprofile for pyenv."
else
  print_color "yellow" ".zprofile already modified for pyenv."
fi

sudo apt install python3.8-venv
curl -sSL https://install.python-poetry.org | python3 -
echo 'export PATH="/home/docker/.local/bin:$PATH"' >> ~/.zshrc
