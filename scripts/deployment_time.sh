#!/bin/bash 

# @dev : This is the script measuring deployment and destroy times.


set -x

for ((i=4;i<=50;i++)); 
do 
   # your-unix-command-here
    sed -i.bak "s/ number_of_nodes:.*/ number_of_nodes: \"${i}\"/" examples/values-local.yml 
    { time  make quorum ; } 2> results/deploy${i}.txt
    { time  make destroy ; } 2> results/deploy${i}.txt
    echo "${i} done!!!"

done