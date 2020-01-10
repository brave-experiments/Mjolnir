terraform plan -var "client_name=$CLIENT_NAME" \
               -var "network_name=$NETWORK_NAME" \
               -var "region=$AWS_REGION" \
               -var 'profile=horizon-admin' \
               -var "number_of_nodes=$NUMBER_OF_NODES" \
               -var "genesis_gas_limit=$GAS_LIMIT" \
               -out=current.plan && \
terraform detroy "current.plan"