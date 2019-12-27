#!/bin/bash 
set -x 

## The IPs of nodes running HBBFT 
IP_FT="52.15.75.118"
IPs_NORM=(3.15.25.194 18.223.213.139 3.16.207.40)
IPs_ALL=(52.15.75.118 3.15.25.194 18.223.213.139 3.16.207.40)
## Very bad - do not commit this 
KEY_FILE="~/Documents/Code/AWS/SSH/Launcher.pem"
# Command to start containers 
START_FT_NODE="docker run --rm -d -e FAKETIME=+0s -e LD_PRELOAD=/usr/local/lib/faketime/libfaketime.so.1 -e FAKETIME_NO_CACHE=1 --name honey-badger brave/honey-badger:dave_faketime --config node.toml --unsafe-expose --tx-queue-per-sender 25000 --tx-queue-size 25000 --jsonrpc-server-threads 32 --jsonrpc-threads=0"
START_NORM_NODE="docker run --rm -d -p 30303:30303/tcp -p  30303:30303/udp  -p 8545:8545/tcp -v /home/ubuntu/node:/node --name honey-badger brave/honey-badger:dave --config node.toml --unsafe-expose --tx-queue-per-sender 25000 --tx-queue-size 25000 --jsonrpc-server-threads 32 --jsonrpc-threads=0"
# Command to stop containers
STOP="docker container stop honey-badger"

#  VM with Fake Time
ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@$IP_FT $START_FT_NODE
echo "Started Parity in Fake VM"

# ## Start Nodes Normal Nodes 
for i in "${IPs_NORM[@]}" 
do :
    ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@$i $START_NORM_NODE
    echo "Started Parity in $i"
done

# Fire Transactions 


# Echo commands to create a script that we can run in parallel . 
for i in "${IPs_ALL[@]}" 
do :
    #  Setting duration to 5 for now CHANGE for main tests 
    echo 
    ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@$i sudo tcpdump -G 600 -W 1 -w $i-dump.pcap  -i eth0 \'port 30303\'
    scp -o "StrictHostKeyChecking  no" -o "IdentitiesOnly=yes" -i "~/Documents/Code/AWS/SSH/Launcher.pem"  ubuntu@52.15.75.118:~/$i-dump.pcap 
done

# Take a tcp dump of p2p comms for the node and store it in a file ?? Do this for the container or VM?
for i in "${IPs_ALL[@]}" 
do :
    #  Setting duration to 5 for now CHANGE for main tests 
    ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@$i sudo tcpdump -G 600 -W 1 -w $i-dump.pcap  -i eth0 \'port 30303\'
    scp -o "StrictHostKeyChecking  no" -o "IdentitiesOnly=yes" -i "~/Documents/Code/AWS/SSH/Launcher.pem"  ubuntu@52.15.75.118:~/$i-dump.pcap 
done

#Copy Results 



# Stop Nodes
for i in "${IPs_ALL[@]}" 
do :
    ssh -i "~/Documents/Code/AWS/SSH/Launcher.pem" ubuntu@$i $STOP
    echo "Stopped Parity in $i"
done



