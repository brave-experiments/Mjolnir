#!/bin/bash 

set -x

# Sock Puppets!!!

#TODO move all constants to .env file 

# Constants 

MJOLNIR=./mjolnir
KEY=.mjolnir/*/id_rsa 
BASTION_IP=$(cat .mjolnir/*/output.log | grep bastion_host_ip | cut -f2 -d"=" |  sed -e 's/^[ \t]*//')
client="quorum"
NETWORK_NAME=$client


mkdir -p results
ssh -o "StrictHostKeyChecking no" -o "IdentitiesOnly=yes"  -i $KEY admin@$BASTION_IP "bash -s" < ./scripts/remote_script.sh  "$NETWORK_NAME" &
BACK_PID=$!
wait $BACK_PID
scp -o "StrictHostKeyChecking no" -o "IdentitiesOnly=yes"  -i $KEY admin@$BASTION_IP:chainhammer/results/runs/Quorum_* results/ 

scp -o "StrictHostKeyChecking no" -o "IdentitiesOnly=yes"  -i $KEY admin@$BASTION_IP:chainhammer/reader/img/*  results/ 


