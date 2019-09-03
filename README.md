## Getting started

To run project locally type:

`bin/run`

To run tests without watcher:

`bin/run test`

To run test watcher type:

`bin/run test-watch`

In test mode cli runs with isolated scope with predefined variables and constants.

> Also see CONTRIBUTING.md


## Manually running build

To manually test build please run 

`bin/run`  - or get latest binary release from here: https://github.com/brave-experiments/apollo-devops/releases 

To execute apollo CLI run:

`./apollo apply quorum values.yml`  - with previosly prepared values.yml taken from `examples/` folder in repo

After successful you will find following files in your working directory:
* default.tfstate   - current terraform object state
* temp.tf           - full dump of terraform code
* variables.log     - log file with provided vatiables


As we don't have a working output function yet ( depends of Blazej availability  should be ready  tomorrow 4.09.2019 )
you need to run following command on `default.tfstate` file to extract deployment details.

`cat default.tfstate | jq '.modules[0].outputs'`

This will show you a json formated list of outputs including:
* _status           - deployment details
* bastion_host_ip   - bastion public IP
* bucket_name       - S3 bucket name with deployment meta files
* chain_id          - deploy chain ID
* ecs_cluster_name  - full cluster name combined of provided vars
* network_name      - provided network_name
* private_key_file  - SSH key for bastion and cluster nodes

In order to get the current node list please copy a file from `private_key_file` var into local dir and chmod it to 0400.

eg. 
`cp /tmp/.terranova141687627/network_name-deploy.pem .`
`chmod 0400 network_name-deploy.pem`

then login to bastion using the key:

`ssh -i ./network_name-deploy.pem ec2-user@bastion_host_ip`

On bastion you will find Node{n} scripts. To get nodes IP addresses for each one run:

`cat /usr/local/bin/Node*`

You can then connect to each node from your workstation using the same key like for bastion.

We will update this file as soon as CLI `output` is fixed. 