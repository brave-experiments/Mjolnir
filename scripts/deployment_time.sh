#!/bin/bash 

# @dev : This is the script measuring deployment and destroy times.
set -x

for ((i=4;i<=50;i++)); 
do 
    mkdir -p results/deploy
    mkdir -p results/destroy
    sed -i.bak "s/ number_of_nodes:.*/ number_of_nodes: \"${i}\"/" examples/values-local.yml 
    { time  make quorum ; } 2> results/deploy/deploy-${i}.txt
    { time  make destroy ; } 2> results/destroy/destroy-${i}.txt
    echo "${i} done!!!"

done