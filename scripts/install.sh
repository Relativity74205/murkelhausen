#!/bin/bash

# atuin (https://github.com/ellie/atuin)
bash <(curl https://raw.githubusercontent.com/ellie/atuin/main/install.sh)
echo 'eval "$(atuin init zsh)"' >> ~/.zshrc
