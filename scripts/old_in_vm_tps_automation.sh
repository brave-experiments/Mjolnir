#!/bin/bash 

set -x

# Sock Puppets!!!


# Constants 

CLI_VERSION='v0.1.0-alpha'
MJOLNIR=../../dist/$CLI_VERSION/unix/mjolnir

# CLIENT=(quorum pantheon parity)
CLIENT=(quorum)
# VM_FAMILY_ARRAY=(t2.2xlarge t3a.xlarge t3a.2xlarge m4.xlarge m4.2xlarge m5.4xlarge  m5.8xlarge m5.12xlarge m5.16xlarge m5.24xlarge m5.metal)
VM_FAMILY_ARRAY=(t2.2xlarge t3a.2xlarge m4.2xlarge m5.4xlarge  m5.8xlarge)
GAS_LIMIT_ARRAY=(8388608 33554432 134217728 536870912 2147483648 8589934592 34359738368 137438953472 549755813888 2199023255552 8796093022208 35184372088832 140737488355328 562949953421312 2251799813685250 9007199254740990)
BLOCKTIME_ARRAY=(1 2 4 8 15 30 60)

# Duplicating devops folder to prevent race condtions 

for client in "${CLIENT[@]}"
do
    CLIENT_NAME=$(echo "${client::3}")
    for vm in "${VM_FAMILY_ARRAY[@]}"
    do 
        for gas in "${GAS_LIMIT_ARRAY[@]}"
        do 
            GAS_LIMIT_NAME=$(echo "${gas::4}")
            NETWORK_NAME="${CLIENT_NAME}${GAS_LIMIT_NAME}"

            # Creating directory for builds and results
            sudo mkdir -p ../build/"${NETWORK_NAME}"/examples/
            sudo mkdir -p ../build/"${NETWORK_NAME}"/dist/v0.1.0-alpha/unix/
            sudo chmod -R 777 ../build/"${NETWORK_NAME}"
            sudo mkdir -p ../results/"${client}"/"${gas}"/"${vm}"/snip
            sudo mkdir -p ../results/"${client}"/"${gas}"/"${vm}"/markdown
            sudo mkdir -p ../results/"${client}"/"${gas}"/"${vm}"/img
            sudo chmod -R 777 ../results/
            sudo cp -R ../Mjolnir/examples/*  ../build/"${NETWORK_NAME}"/examples/
            sudo cp  ../Mjolnir/dist/v0.1.0-alpha/unix/mjolnir ../build/"${NETWORK_NAME}"/dist/v0.1.0-alpha/unix/
            sudo cp  ../Mjolnir/remote_script.sh ../build/"${NETWORK_NAME}"/
            for file in /build/"${NETWORK_NAME}"
            do
                mkdir -p ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}/
                sudo cp ../Mjolnir/examples/values.yml  ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
                sudo sed -i "s/ network_name:.*/ network_name: \"${NETWORK_NAME}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
                sudo sed -i "s/ # genesis_gas_limit:.*/ genesis_gas_limit: \"${gas}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
                sudo sed -i "s/ # asg_instance_type:.*/ asg_instance_type: \"${vm}\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
                sudo sed -i "s/ # tf_log:.*/ tf_log: \"y\"/" ../build/"${NETWORK_NAME}"/examples/values-local-${gas}-${vm}.yaml
                sudo rm -rf ../build/"${NETWORK_NAME}"/examples/values-local.yml
                sudo rm -rf ../build/"${NETWORK_NAME}"/examples/values.yml
                rm -rf ../build/"${NETWORK_NAME}"/examples/values-local.yaml
                rm -rf ../build/"${NETWORK_NAME}"/examples/values-local.yaml.bak
                mkdir -p ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}/
                echo \#!/bin/bash -e >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo set -x >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo $MJOLNIR apply $CLIENT ../../examples/values-local-${gas}-${vm}.yaml >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo "BASTION_IP=\$(cat .mjolnir/*/output.log | grep bastion_host_ip | cut -f2 -d"=" |  sed -e 's/^[ \t]*//')"  >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo KEY=.mjolnir/*/id_rsa >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo ssh -o \"StrictHostKeyChecking no\" -o \"IdentitiesOnly=yes\"  -i \$KEY admin@\$BASTION_IP" "bash -s" < ../../remote_script.sh  "$NETWORK_NAME"" >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo scp -o \"StrictHostKeyChecking no\" -o \"IdentitiesOnly=yes\"  -i \$KEY  admin@\$BASTION_IP:chainhammer/results-*.txt ../../../../results/"${client}"/"${gas}"/"${vm}"/snip >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo scp -o \"StrictHostKeyChecking no\" -o \"IdentitiesOnly=yes\"  -i \$KEY  admin@\$BASTION_IP:chainhammer/results/runs/*.md ../../../../results/"${client}"/"${gas}"/"${vm}"/markdown >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                echo scp -o \"StrictHostKeyChecking no\" -o \"IdentitiesOnly=yes\"  -i \$KEY  admin@\$BASTION_IP:chainhammer/reader/img/* ../../../../results/"${client}"/"${gas}"/"${vm}"/img >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo $MJOLNIR destroy ../../examples/values-local-${gas}-${vm}.yaml >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                sudo echo wait >> ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
                chmod +x ../build/"${NETWORK_NAME}"/run/"${NETWORK_NAME}"-${gas}-${vm}/"${NETWORK_NAME}"-${gas}-${vm}.sh
            done
        done
    done
done



# # ## Very hacky for setting the region   
# # #`a# Each region can only have a maximum of 5 VPCs 

cd ../build


for f in *

do 
    if 
        [ "$(echo "${f:3:1}")" == "1"  ]; then
        sed -i "s/ region:.*/ region: \"us-east-1\"/" $f/examples/*.yaml
    fi
    if 
        [ "$(echo "${f:3:1}")"  == "2"  ]; then
        sudo sed -i "s/ region:.*/ region: \"us-west-1\"/" $f/examples/*.yaml  
    fi
    if 
        [ "$(echo "${f:3:1}")" == "3"  ]; then   
        sudo sed -i "s/ region:.*/ region: \"eu-central-1\"/" $f/examples/*.yaml
    fi
    if 
        [ "$(echo "${f:3:1}")" == "5"  ]; then
        sudo sed -i "s/ region:.*/ region: \"us-west-2\"/" $f/examples/*.yaml
       
    fi
    if 
        [ "$(echo "${f:3:1}")"  == "8"  ]; then
        sudo sed -i "s/ region:.*/ region: \"eu-west-2\"/" $f/examples/*.yaml
       
    fi
    if 
        [ "$(echo "${f:3:1}")" == "9"  ]; then
        sudo sed -i "s/ region:.*/ region: \"eu-west-3\"/" $f/examples/*.yaml
    fi
    
done



# Running all the jobs concurrently 

function run_folder {
   local  dir=$1 f=
   cd $dir/run
   # sequential execution
   for f in */*.sh ; do
       # Execute each test in it's folder.
       (cd ${f%/*} && bash ${f#*/} -H)
   done
}

# parallel execution 
for j in * ; do
   run_folder $j &
done
wait