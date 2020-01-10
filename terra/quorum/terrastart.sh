terraform plan -var "client_name=$CLIENT_NAME" \
               -var "network_name=$NETWORK_NAME" \
               -var "region=$AWS_REGION" \
               -var "number_of_nodes=$NUMBER_OF_NODES" \
               -var "genesis_gas_limit=$GAS_LIMIT" \
               -out=current.plan && \
terraform apply "current.plan" 


# cat terraform.tfstate | jq .modules[].outputs[].value


# chmod 400 quorum-jmeter-test-2.pem

# /tmp/terraform_749961759.sh
# 54.244.164.57

# ssh -i quorum-jmeter-test-2.pem admin@54.244.164.57

# null_resource.bastion_remote_exec (remote-exec):       privacy-address: Kiz6gILCNwnuT8G3m1g+WRPyJsvtxjmCM8BcCEIPJDQ=
# null_resource.bastion_remote_exec (remote-exec):       url: http://52.26.18.131:22000
# null_resource.bastion_remote_exec (remote-exec):       third-party-url: http://52.26.18.131:9080
# null_resource.bastion_remote_exec (remote-exec): #!/bin/bash

# Error: Error applying plan:

# 1 error occurred:
#         * null_resource.bastion_remote_exec: error executing "/tmp/terraform_749961759.sh": Process exited with status 1

