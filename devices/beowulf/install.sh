# in addition to ansible (TODO: move to ansible)

# supervisor
sudo apt install supervisor
- add `chown=arkadius` to /etc/supervisor/supervisord.conf
[unix_http_server]
...
chown=arkadius


# access control lists
sudo apt install acl
sudo setfacl -R -m u:arkadius:rwx data

# syncthing
sudo curl -o /usr/share/keyrings/syncthing-archive-keyring.gpg https://syncthing.net/release-key.gpg
echo "deb [signed-by=/usr/share/keyrings/syncthing-archive-keyring.gpg] https://apt.syncthing.net/ syncthing stable" | sudo tee /etc/apt/sources.list.d/syncthing.list
sudo apt-get update
sudo apt-get install syncthing


# prefect
sudo apt update && sudo apt install sqlite3
pipx install prefect
pipx inject prefect prefect-shell

