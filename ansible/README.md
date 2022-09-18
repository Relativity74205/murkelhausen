# WSL 


## Microsoft WSL configs

https://docs.microsoft.com/de-de/windows/wsl/wsl-config


## WSL links

- https://github.com/bostrot/wsl2-distro-manager
- https://github.com/DDoSolitary/LxRunOffline
- https://github.com/microsoft/wslg/wiki/Diagnosing-%22cannot-open-display%22-type-issues-with-WSLg#still-having-a-problem
- https://hjoelr.medium.com/wsl2-gui-app-shortcuts-in-windows-with-wslg-fcc66d3134e7



# TODO

- golang
- gauxcli
- taskwarrior config + remote server


# Setup

Setup passwd vault:
```bash
ansible-vault create passwd.yaml
```

with following entry
```
ansible_become_pass: <<my_pass>>
```

For local
```bash
ansible-playbook local.yaml
```

For Beowulf
```bash
ansible-playbook -i inventory.yaml beowulf.yaml
```

Interesting options:
- `--check`
- `--verbose` or `-v`

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

Verwendung: wsl.exe [Argument] [Optionen...] [Befehlszeile]

Argumente zum Ausführen von Linux-Binärdateien:

    Wenn keine Befehlszeile angegeben wird, startet wsl.exe die Standardshell.

    --exec,-e <Befehlszeile>
        Führen Sie den angegebenen Befehl aus, ohne die Standard-Linux-Shell zu verwenden.

    --
        Übergeben Sie die restliche Befehlszeile ohne Änderung.

Optionen:
    --cd <Verzeichnis>
        Legt das angegebene Verzeichnis als aktuelles Arbeitsverzeichnis fest.
        Bei Angabe von „~“ wird der Startpfad des Linux-Benutzers verwendet. Wenn der Pfad mit einem
        „/“-Zeichen beginnt, wird er als absoluter Linux-Pfad interpretiert.
        Andernfalls muss der Wert ein absoluter Windows-Pfad sein.

    --distribution,-d <Distribution>
        Führt die angegebene Distribution aus.

    --user,-u <Benutzername>
        Verwendet für die Ausführung den angegebenen Benutzer.

    --System
        Startet eine Shell für die Systemverteilung.



- Copy SSH keys with powershell
```powershell
Copy-Item -Path C:\Users\arkad\OneDrive\Documents\wsl_dev\.ssh -Destination \\wsl.localhost\Ubuntu\home\arkadius\.ssh -Recurse
Copy-Item -Path C:\Users\arkad\OneDrive\Documents\wsl_dev\initial_load.sh -Destination \\wsl.localhost\Ubuntu\home\arkadius\initial_load.sh
```

- copy ssh keys and initial loda scripts
```bash
cp -r /mnt/c/Users/arkad/OneDrive/Documents/wsl_dev/.ssh ~
cp /mnt/c/Users/arkad/OneDrive/Documents/wsl_dev/initial_load.sh ~
```

- change SSH key permissions of ssh keys to rw for current user
```bash
mkdir -p .kube
mkdir -p .aws
mkdir -p .ssh
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/configs/kube_config ~/.kube/config
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/configs/aws_config ~/.aws/config
cp /mnt/c/Users/ArkadiusSchuchhardt/OneDrive\ -\ auxmoney\ GmbH/Dokumente/ssh/* ~/.ssh
chmod 600 ~/.ssh/*
eval `ssh-agent`
ssh-add ~/.ssh/github_private
mkdir -p ~/dev/murkelhausen
git clone git@github.com:Relativity74205/murkelhausen.git ~/dev/murkelhausen
```


- Install ansible
```bash
sudo apt update && sudo apt upgrade -y && sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=/home/arkadius/.local/bin/:$PATH' >> ~/.bashrc
source ~/.bashrc
ansible --version
```

- Install fonts for Windows Terminal on host:
https://github.com/romkatv/powerlevel10k#manual-font-installation
