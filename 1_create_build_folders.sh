#!/bin/bash -e

set -x

# Sock Puppets!!!


# Constants 

CLI_VERSION='v0.1.0-alpha'
APOLLO=./dist/$CLI_VERSION/unix/apollo

# CLIENT=(quorum pantheon parity)
CLIENT=(quorum)
VM_FAMILY_ARRAY=(t2.2xlarge t3a.xlarge t3a.2xlarge a1.4xlarge a1.metal m4.xlarge m4.2xlarge m5.4xlarge  m5.8xlarge m5.12xlarge m5.16xlarge m5.24xlarge m5.metal)
GAS_LIMIT_ARRAY=(8388608 33554432 134217728 536870912 2147483648 8589934592 34359738368 137438953472 549755813888 2199023255552 8796093022208 35184372088832 140737488355328 562949953421312 2251799813685250 9007199254740990)
BLOCKTIME_ARRAY=(1 2 4 8 15 30 60)

# Duplicating devops folder to prevent race condtions 

# for client in "${CLIENT[@]}"
# do
#     CLIENT_NAME=$(echo "${client::3}")
#     for vm in "${VM_FAMILY_ARRAY[@]}"
#     do 
#         for gas in "${GAS_LIMIT_ARRAY[@]}"
#         do 
#             GAS_LIMIT_NAME=$(echo "${gas::4}")
#             NETWORK_NAME="${GAS_LIMIT_NAME}${CLIENT_NAME}"

#             # Creating directory for builds and results
#             sudo mkdir -p ../build/"${NETWORK_NAME}"
#             sudo mkdir -p ../results/"${client}"/gas/"${vm}"/snip
#             sudo mkdir -p ../results/"${client}"/gas/"${vm}"/markdown
#             sudo cp -R ../Mjolnir/*  ../build/"${NETWORK_NAME}"

#             for file in /build/"${NETWORK_NAME}"
#             do
#                 # cd ../build/"${NETWORK_NAME}"
#                 sudo cp ../Mjolnir/examples/values.yml  ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
#                 sudo sed -i "s/ network_name:.*/ network_name: \"${NETWORK_NAME}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
#                 sudo sed -i "s/ # genesis_gas_limit:.*/ genesis_gas_limit: \"${gas}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
#                 sudo sed -i "s/ # asg_instance_type:.*/ asg_instance_type: \"${vm}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
#                 sudo sed -i "s/ # tf_log:.*/ tf_log: \"y\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml  
#             done
#         done
#     done
# done



# # ## Very hacky for setting the region   
# # ## Each region can only have a maximum of 5 VPCs 
# # sudo chmod -R 777 ../build
# cd ../build

# for f in *

# do 
#     # echo f
#     # case $(echo "${f::1}")
#     if 
#         [ "$(echo "${f::1}")" == "1"  ]; then
#         echo "$f/examples/*.yaml" 
#         sed -i "s/ region:.*/ region: \"us-east-1\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region
#     fi
#     if 
#         [ "$(echo "${f::1}")" == "2"  ]; then
#         echo "$f"
#         # $f/examples/values-local-${gas}-${vm}.yaml
#         sudo sed -i "s/ region:.*/ region: \"us-west-1\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region 
#     fi
#     if 
#         [ "$(echo "${f::1}")" == "3"  ]; then
#         echo "$f" 
#         # $f/examples/values-local-${gas}-${vm}.yaml
#         sudo sed -i "s/ region:.*/ region: \"eu-central-1\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region
#     fi
#     if 
#         [ "$(echo "${f::1}")" == "5"  ]; then
#         echo "$f" 
#         # $f/examples/values-local-${gas}-${vm}.yaml
#         sudo sed -i "s/ region:.*/ region: \"us-west-2\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region
#     fi
#     if 
#         [ "$(echo "${f::1}")" == "8"  ]; then
#         echo "$f" 
#         # $f/examples/values-local-${gas}-${vm}.yaml
#         sudo sed -i "s/ region:.*/ region: \"eu-west-2\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region
#     fi
#     if 
#         [ "$(echo "${f::1}")" == "9"  ]; then
#         echo "$f" 
#         # $f/examples/values-local-${gas}-${vm}.yaml
#         sudo sed -i "s/ region:.*/ region: \"ca-central-1\"/" $f/examples/*.yaml
#         sudo cat $f/examples/*.yaml | grep region
#     fi
    
# done
cd ../build
CLIENT=quorum
# BANG=$(echo */"$APOLLO" apply "$CLIENT" */examples/*.yaml &)
CLI_VERSION='v0.1.0-alpha'
APOLLO=./dist/$CLI_VERSION/unix/apollo

pids=() 

for f in *
do 
    file=$f
    # cd $file/example/
    for yaml in $file/examples/*.yaml
    do  
       
        $file/$APOLLO apply $CLIENT $yaml &
        echo $BANG
        pids+=($!)
    done
done 

for pid in "${pids[@]}"; do
  wait "$pid"
done

# for f in *
# do 
#     file=$f
#     # cd $file/example/
#     for yaml in $file/examples/*.yaml
#     do  
#         BASTION_IP=$(cat $file/.apollo/*/output.log | grep bastion_host_ip | cut -f2 -d"=" |  sed -e 's/^[ \t]*//')
#         KEY=$file/.apollo/*/id_rsa && ssh -o "StrictHostKeyChecking no" -o "IdentitiesOnly=yes"  -i "$file/$KEY" admin@"$BASTION_IP" "bash -s" < ./remote_script.sh 
#         pids+=($!)
#     done
# done 

# for pid in "${pids[@]}"; do
#   wait "$pid"
# done






# # ./2_send_parallel_clientobs.sh 

