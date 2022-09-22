#!/usr/bin/bash

sudo apt install -y cmake wget build-essential tar
wget https://github.com/GothenburgBitFactory/taskwarrior/releases/download/v2.6.2/task-2.6.2.tar.gz
tar xzvf task-2.6.2.tar.gz
cd task-2.6.2
cmake -DCMAKE_BUILD_TYPE=release .
make
sudo make install
sed -i 's/plugins=(\(.*\))/plugins=(\1 taskwarrior)/' ~/.zshrc