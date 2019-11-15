#!/bin/bash 

set -x

# Sock Puppets!!!
 REGIONS=`aws ec2 describe-regions --region us-east-1 --output text --query Regions[*].[RegionName]`
  for REGION in $REGIONS
  do
    echo -e "\nInstances in '$REGION'..";
    aws ec2 describe-instances --region $REGION | \
      jq '.Reservations[].Instances[] | "EC2: \(.InstanceId): \(.State.Name)"'
  done