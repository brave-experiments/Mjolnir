#!/bin/bash 

set -x

sudo apt-get -y upgrade && 
sudo apt-get -y update &&

sudo apt -y install docker.io git &&
sudo systemctl start docker &&
sudo systemctl enable docker &&

sudo usermod -aG docker ${USER} &&

sudo reboot 


enode://12d663d23c88bd3bea539bef3b4513b0ba4476f84f11148a76e15df6d0b89b7093ed053cafbe3efbbddcc6199d97ff33dc095094d6f436f40154546c210cec2f@172.17.0.2:30303
enode://72964e7d25f883db3c59cc9fd0c421a93304ba6cbc4495116ca43b95c2cda8429bbacae6397ac7da4941d8f6eb220bec7fe1a65ef9eea39be0eb4495e4d0b414@172.17.0.2:30303
enode://c9ad99c778c7f031e73f2bc4f7b1770e3cd82a0e7b99f07dce72b90ba6f8b9407adabbf9b9f977b77ae78c868c808be09e3518b58b74fd49a03e33dca482201c@172.17.0.2:30303
enode://335e199e70db3bcee345aa33aa4e167b65fb8938690e27c33f3684c67520257da2a644835930a69fff4988441a277dda40c33e07e8b94b18e7901bc4a7b5ce13@172.17.0.2:30303


# Docker Run command

docker run -it \
    -p 30303:30303 -p 8545:8545 \
    -v node:~/ \
    brave/honey-badger
ps aux
docker run --rm -ti -p  30303:30303/tcp \
    -p  30303:30303/udp  -p 8545:8545/tcp \
    -v /home/ubuntu/node:/node --name honey-badger \
    brave/honey-badger:dave --config node.toml --unsafe-expose \
    --tx-queue-per-sender 25000 --tx-queue-size 25000 \
    --jsonrpc-server-threads 32 --jsonrpc-threads=0

docker run --rm -d -e FAKETIME=+15d \
    -e LD_PRELOAD=/usr/local/lib/faketime/libfaketime.so.1 \
    -e FAKETIME_NO_CACHE=1 --name honey-badger \
    brave/honey-badger:dave_faketime --config node.toml --unsafe-expose \
    --tx-queue-per-sender 25000 --tx-queue-size 25000 \
    --jsonrpc-server-threads 32 --jsonrpc-threads=0

docker run --rm -d -e FAKETIME=+0 \
    -e LD_PRELOAD=/usr/local/lib/faketime/libfaketime.so.1 \
    -e FAKETIME_NO_CACHE=1 --name honey-badger \
    brave/honey-badger:dave_faketime --config node.toml --unsafe-expose \
    --tx-queue-per-sender 25000 --tx-queue-size 25000 \
    --jsonrpc-server-threads 32 --jsonrpc-threads=0

p 172.17.0.1:$PORT:$PORT/tcp -p 172.17.0.1:$PORT:$PORT/udp


docker build -t 
# IPs
# Node 1
ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@52.15.75.118
# Node 2
ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@3.15.25.194
# Node 3
ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@18.223.213.139
#  Node 4
ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@3.16.207.40

# chain hammer

ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem"  admin@18.188.117.77

geth attach http://18.223.213.139


# Copying Parity Binary

scp -o "StrictHostKeyChecking  no" \
    -o "IdentitiesOnly=yes" \
    -i "~/Documents/Code/AWS/SSH/Launcher.pem" -r \
    ubuntu@18.221.106.0:/home/ubuntu/parity-ethereum/target/release/parity . 

scp -o "StrictHostKeyChecking  no" \
    -o "IdentitiesOnly=yes" \
    -i "~/Documents/Code/AWS/SSH/Launcher.pem" -r \
    node1 ubuntu@52.15.75.118:~/node && \

scp -o "StrictHostKeyChecking  no" \
    -o "IdentitiesOnly=yes" \
    -i "~/Documents/Code/AWS/SSH/Launcher.pem" -r \
    node2 ubuntu@3.15.25.194:~/node && \

scp -o "StrictHostKeyChecking  no" \
    -o "IdentitiesOnly=yes" \
    -i "~/Documents/Code/AWS/SSH/Launcher.pem" -r \
    node3 ubuntu@18.223.213.139:~/node && \

scp -o "StrictHostKeyChecking  no" \
    -o "IdentitiesOnly=yes" \
    -i "~/Documents/Code/AWS/SSH/Launcher.pem" -r \
    node4 ubuntu@3.16.207.40:~/node


    # Running Chain hammer

CH_TXS=25000 CH_THREADING="threaded2 300" ./run.sh "HBBFT"
CH_TXS=25000 CH_THREADING="threaded2 300" ./run.sh "IBFT"

ps aux | grep python
sudo kill 9 25005
# Build Failure

Remeber to remove go interactive install from terra 



rm -rf  node/data/chains/DPoSChain/

geth attach http://3.16.207.40:8545

eth.getBalance(eth.accounts[0]) 


# get peer count
IPs=(52.15.75.118 3.15.25.194 18.223.213.139 3.16.207.40)
for i in IPS[@] do;
    curl --data '{"method":"net_peerCount","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST $i:8545

# Modify clienttools to take out parity's unlock


var sender = eth.accounts[0];

var receiver = "0x8f2e3b0b69894d95a16ed91e3c9b8d7a63f6a504";

var amount = web3.toWei(0.000001, "ether")

eth.sendTransaction({from:sender, to:receiver, value: amount})
176.24.45.227:8541
176.24.45.227:8542

# Changes to Chainhammer to make HBBFT work:

- send.py : make it parity line 581


# Parity modificatins

https://github.com/paritytech/parity-ethereum/issues/10382#issuecomment-466373932


# Idea for building parity

- Mofify tool to produce single variable like quorum. 


./dist/v0.1.0-alpha/osx/mjolnir destroy  examples/values-local.yml \
 &&  ./dist/v0.1.0-alpha/osx/mjolnir destroy  examples/values-local.yml  