#!/bin/bash

DEV_DIRECTORY=dev  # folder, to which the bi-wsl repo is cloned

# (optional) copies SSH keys from Windows Folder; if needed set COPY_SSH_KEY to true and edit GITLAB_SSH_KEY and WINDOWS_SSH_FOLDER
COPY_SSH_KEY=false
GITLAB_SSH_KEY=gitlab
WINDOWS_SSH_FOLDER="/mnt/c/Users/ArkadiusSchuchhardt/OneDrive - auxmoney GmbH/Dokumente/ssh/"
# (optional) copies aws config; if needed set COPY_AWS_CONFIG to true and edit AWS_CONFIG_PATH
COPY_AWS_CONFIG=false
AWS_CONFIG_PATH="/mnt/c/Users/ArkadiusSchuchhardt/OneDrive - auxmoney GmbH/Dokumente/configs/aws_config"
# (optional) copies kube config; if needed set COPY_KUBE_CONFIG to true and edit KUBE_CONFIG_PATH
COPY_KUBE_CONFIG=false
KUBE_CONFIG_PATH="/mnt/c/Users/ArkadiusSchuchhardt/OneDrive - auxmoney GmbH/Dokumente/configs/aws_config"


# (optional) copy ssh key and add to ssh-agent
if [ $COPY_SSH_KEY = true ]
then
  mkdir -p .ssh
  cp "$WINDOWS_SSH_FOLDER"/* ~/.ssh
  chmod 600 ~/.ssh/*
  eval `ssh-agent`  # needed?
  ssh-add ~/.ssh/$GITLAB_SSH_KEY  # needed?
fi
# (optional) copies aws and kube config (if needed)
if [ $COPY_AWS_CONFIG = true ]
then
  mkdir -p ~/.aws
  cp "$AWS_CONFIG_PATH" ~/.aws/config
fi
if [ $COPY_KUBE_CONFIG = true ]
then
  mkdir -p ~/.kube
  cp "$KUBE_CONFIG_PATH" ~/.kube/config
fi

# clones ansible scripts
mkdir -p $DEV_DIRECTORY
cd $DEV_DIRECTORY
git clone git@gitlab.office.auxmoney.com:data-infra/bi-wsl.git


# installs ansible
sudo apt update
sudo apt upgrade -y
sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=~/.local/bin/:$PATH' >> ~/.bashrc
