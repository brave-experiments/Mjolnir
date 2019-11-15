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

    # # Delete Log  Group 
    # aws logs describe-log-groups --region $region | \
    # jq  -r .logGroups[].logGroupName 

    aws logs describe-log-groups --region eu-west-2 | \
        jq  -r .logGroups[].logGroupName | \
            xargs -L 1 -I {}  aws logs delete-log-group --log-group-name {}

    # # Delete ECS

    # aws ecs describe-clusters --cluster default --region eu-central-1

    # # Delete IAM Pol

    # # Delete VPC 

    ## Delete s3 buckets

    for bucket in $(aws s3 ls | awk '{print $3}' | grep us); do  aws s3 rb "s3://${bucket}" --force ; done


    
done