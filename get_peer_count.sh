#!/bin/bash 

# set -x

#  get peer count
IPs=(52.15.75.118 3.15.25.194 18.223.213.139 3.16.207.40)
for i in "${IPs[@]}" 
do :
    curl --data '{"method":"net_peerCount","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST $i:8545
done
