## Getting started

### Development mode
To run project locally type:

`bin/run`

### Test
To run tests without watcher:

`bin/run test`

### Test-watch
To run test watcher type:

`bin/run test-watch`

### Build
To build from source:
`bin/run ci`

After success built files will lay within ./dist/{arch}/{binaryName}

To execute apollo binary file:
try `./apollo` to see all commands that are registered
try `./apollo {cmdName} --help` to see help from command

### Quorum execution
to build
`./apollo apply quorum {values.yml}`

### Providing values
See `example/values.yml` that shows how to attach values to apply execution. 
Since any `values-local.yml` file is gitignored
you should copy `example/values.yml` to `values-local.yml` and provide values that you need.

In test mode cli runs with isolated scope with predefined variables and constants.

### Further debugging
After execution of `apply` command certain files will be created on your host:
- `temp.tf` at root of execution dir, which contains whole terraform code that has been executed
- `terraform.tfstate` at root of execution dir, which contains state of execution
- `variables.log` at root of execution dir, which contains last executed variables in recipe
- `.apollo` dir which contains necessary files like ssh key pair to bastion
- `.apollo/$network-name+$timestamp/` is a dir where should end private and public key pair
> Also see CONTRIBUTING.md


## Manually running deploy build

To manually test build please run 

`bin/run`  - or get latest binary release from here: https://github.com/brave-experiments/apollo-devops/releases 

To execute apollo CLI run:

`./apollo apply quorum values.yml`  - with previosly prepared values.yml taken from `examples/` folder in repo

After successful you will find following files in your working directory:
* terraform.tfstate   - current terraform object state
* temp.tf           - full dump of terraform code
* variables.log     - log file with provided vatiables

On successful run on the output you will see following example information:

```
Created output dir:  .apollo/network-name-12345678
[FINAL] Summary execution: [reset][bold][green]
Outputs:

_status = Completed!

Quorum Docker Image         = quorumengineering/quorum:2.2.5
Privacy Engine Docker Image = quorumengineering/tessera:latest
Number of Quorum Nodes      = 3
ECS Task Revision           = 1
CloudWatch Log Group        = /ecs/quorum/network-name-12345678

bastion_host_ip = xxx.xxx.xxx.xxx
bucket_name = us-east-2-ecs-network-name-12345678-b616bc76ee59e4ba
chain_id = 7774
ecs_cluster_name = quorum-network-name-12345678
ethstats_host_url = http://xxx.xxx.xxx.xxx:3000
grafana_host_url = http://xxx.xxx.xxx.xxx:3001
grafana_username = admin
grafana_password = XXXXXXXXX
network_name = network-name
private_key_file = <sensitive>
Wrote summarry output to:  .apollo/quorum-bastion-jkopacze-n3-66790866/output.log
Restoring env variables.
```



## Manually destroying deploy

to destroy run:
`./apollo destroy {values.yml}`

Current success output looks like this ( will be correted in next release ):
```Deploy Name not present
   [FINAL] Summary execution: 
   Wrote summarry output to:  .apollo//output.log
   Deploy Name not present
   Restoring env variables.
```

## Deploy usage


In order to get the current node list please login to bastion using the key:

`ssh -i .apollo/$network-name+$timestamp/id_rsa.pem ec2-user@bastion_host_ip`

On bastion you will find Node{n} scripts. To get nodes IP addresses for each one run:

`cat /usr/local/bin/Node*`

You can then connect to each node from your workstation using the same key like for bastion.

Additional tools you will find under links
* eth-stats: http://bastion_host_ip:3000
* Grafana: http://bastion_host_ip:3001
* Prometheus: http://bastion_host_ip:9090
