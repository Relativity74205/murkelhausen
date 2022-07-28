#!/usr/bin/env bash
source ./print_color.sh
log=install.log

if ! command -v zsh &> /dev/null
then
  print_color "green" "installing zsh"
  sudo apt-get -y install zsh | sudo tee -a $log > /dev/null
  sudo chsh -s "$(which zsh)" | sudo tee -a $log > /dev/null
  print_color "green" "installing zsh complete"
else
  print_color "yellow" "zsh already installed"
fi


if [ ! -d "$HOME"/.oh-my-zsh ]; then
  print_color "green" "installing oh-my-zsh..."
  sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" | sudo tee -a $log > /dev/null
  print_color "green" "configuring oh-my-zsh"
  cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc
  echo 'export ZSH_AUTOSUGGEST_STRATEGY=(history completion)' >> ~/.zshrc
  sed -i 's/plugins=(\(.*\))/plugins=(\1 pip)/' ~/.zshrc
  git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
  sed -i 's/plugins=(\(.*\))/plugins=(\1 zsh-autosuggestions)/' ~/.zshrc
  print_color "green" "installing oh-my-zsh complete"
else
  print_color "yellow" "oh-my-zsh already installed"
fi

sudo apt-get install fonts-powerline
sudo apt install tmux
echo  "export TERM=xterm-256color" >> ~/.zshrc

#ZSH_THEME="agnoster"
ZSH_THEME="powerlevel10k/powerlevel10k"
DEFAULT_USER="aschuchhardt"
#  plugins=(aws docker docker-compose)
