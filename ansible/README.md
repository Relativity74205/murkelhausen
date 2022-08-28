```bash
ansible-playbook -i inventory.yaml beowulf.yaml -K
```

Interesting options:
- `--check`
- `--verbose`

## Links

- Docker with ansible: https://www.digitalocean.com/community/tutorials/how-to-use-ansible-to-install-and-set-up-docker-on-ubuntu-20-04
- Ansible Roles: https://www.digitalocean.com/community/tutorials/how-to-use-ansible-roles-to-abstract-your-infrastructure-environment
- Docker container module: https://docs.ansible.com/ansible/2.9/modules/docker_container_module.html



## Ansible Galaxy Roles:

- https://galaxy.ansible.com/geerlingguy/docker
- https://galaxy.ansible.com/gantsign/oh-my-zsh


## multiple ubuntu Installations on WSL

- https://cloudbytes.dev/snippets/how-to-install-multiple-instances-of-ubuntu-in-wsl2


Get wsl installations:
```bash
wsl -l -v
```

Install custom WSL OS version:
- download image, e.g. ubuntu 22.04:
```shell
curl (("https://cloud-images.ubuntu.com", "releases/jammy/release", "ubuntu-22.04-server-cloudimg-amd64-wsl.rootfs.tar.gz") -join "/")
```

- Install the second instance of Ubuntu in WSL2
```shell
wsl --import <Distribution Name> <Installation Folder> <Ubuntu WSL2 Image Tarball path>
```
with 
- Replace the <Distribution Name> with the name you want to give, e.g. ubuntu-2,
- Replace <Installation Folder> with the folder where you want to install the second instance of Ubuntu
- and finally replace <Ubuntu Tarball path> with the path of the Ubuntu WSL2 image tarball you downloaded earlier.

Run the following in new WSL:
```bash

NEW_USER=arkadius
useradd -m -G sudo -s /bin/bash "$NEW_USER"
passwd "$NEW_USER"

tee /etc/wsl.conf <<_EOF
[user]
default=${NEW_USER}
_EOF
```

- Copy SSH keys (change permission of ssh keys to rw for current user)
```bash
chmod 600 .ssh/*
eval `ssh-agent`
ssh-add .ssh/github
git clone git@github.com:Relativity74205/murkelhausen.git
```


- Install ansible
```bash
sudo apt update
sudo apt upgrade
sudo apt install python3-pip
python3 -m pip install --user ansible
export PATH=/home/arkadius/.local/bin/:$PATH
source .bashrc
ansible --version
```