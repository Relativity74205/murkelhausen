# installs ansible
sudo apt update
sudo apt upgrade -y
sudo apt install -y python3-pip
python3 -m pip install --user ansible
echo 'export PATH=~/.local/bin/:$PATH' >> ~/.bashrc
