#!/bin/bash 

set -x

# Sock Puppets!!!


REGION_ARRAY=(eu-west-3 eu-west-2 us-west-2 eu-central-1 us-west-1 us-east-1)

for region in "${REGION_ARRAY[@]}"
do
  echo "Terminating region $region..."
  aws ec2 describe-instances --region $region | \
    jq -r .Reservations[].Instances[].InstanceId | \
      xargs -L 1 -I {} aws ec2 modify-instance-attribute \
        --region $region \
        --no-disable-api-termination \
        --instance-id {}
  aws ec2 describe-instances --region $region | \
    jq -r .Reservations[].Instances[].InstanceId | \
      xargs -L 1 -I {} aws ec2 terminate-instances \
        --region $region \
        --instance-id {}

    
done