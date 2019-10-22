
# Mjölnir 
## ...the hammer of Thor. 

### *N.B.* 
- Some key features may not be compatible in Windows environment
- Only Amazon Web Services (AWS) is supported now. We are however open to PRs for other cloud providers!

```
|"  \    /"  |     |"  |  /    " \  |"  |    (\"   \|"  \  |" \    /"      \  
 \   \  //   |     ||  | // ____  \ ||  |    |.\\   \    | ||  |  |:        | 
 /\\  \/.    |     |:  |/  /    ) :)|:  |    |: \.   \\  | |:  |  |_____/   ) 
|: \.        |  ___|  /(: (____/ //  \  |___ |.  \    \. | |.  |   //      /  
|.  \    /:  | /  :|_/ )\        /  ( \_|:  \|    \    \ | /\  |\ |:  __   \  
|___|\__/|___|(_______/  \"_____/    \_______)\___|\____\)(__\_|_)|__|  \___)
```
Mjolnir is a tool for rapidly deploying and testing Ethereum Clients. Tool currently allows users to test the through put of Ethereum Clients both on its own, and under adverse network conditions (i.e. Clock Skew , dropped tcp packets, jitters, etc.)

At this moment, Mjolnir supports the following clients:

- Quorum
- Patheon (now Hyperledger Besu)
- Parity (Honey Badger / POSDAO version)


## Getting started
./dist/{cli-version}/{arch}/{binaryName}

### Terminology

- {cli-version} : Semantic version of binary 
- {arch}: The OS architecture. Currently supported are 
   - `osx` 
   - `unix`
- {binaryName}: `apollo`
- {client}: The Ethereum client been tested. Currently supported are:
   - `quorum`
   - `pantheon`
   - `parity`
- {cmdName} : Binary's sub command. 


- Step 1 : Deploy Infrastructure
   - Clone this repo
   - Enter `bin/run` to run locally. This will create a local docker container and ssh into it. 
   - Create a copy of the configuration files in `examples/values.yaml` to  `examples/values-local.yaml`
   - Update `examples/values-local.yaml`
   - enter `./dist/{cli-version}/{arch}/apollo apply {client} examples/values-local.yml `. This will deploy the requsite infrastrucure on your AWS account.

- Step 2: Fire Transactions 
   - Once this is complete, enter `./dist/{cli-version}/{arch}/apollo bastion` to tunnel into the bastion host. It is from here, we are able to access chainhamer for sending transactions to the clients. 
   - Move in the `chainhammer` directory by entering `cd chainhammer`
   - To send transactions, ` CH_TXS=40000 CH_THREADING="threaded2 300" ./run.sh "{TESTNAME}"`; Where `CH_TXS` is the number of transactions to be send, `CH_THREADING` is the threading algorithm, and `{TESTNAME}` is the name that the run will be save under.
   - If all goes well, 


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

After success built files will lay within `./dist/{cli-version}/{arch}/{binaryName}`

To execute apollo binary file:
try `./dist/{cli-version}/{arch}/apollo` to see all commands that are registered
try `./dist/{cli-version}/{arch}/apollo {cmdName} --help` to see help from command

### Example : Quorum execution
to build
`../dist/{cli-version}/{arch}/apollo apply quorum {values.yml}`

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
`./dist/{cli-version}/{arch}/apollo destroy {values-local.yml}`

Current success output looks like this ( will be correted in next release ):
```Deploy Name not present
   [FINAL] Summary execution: 
   Wrote summarry output to:  .apollo//output.log
   Deploy Name not present
   Restoring env variables.
```

## Deploy usage


In order to get the current node list please login to bastion using the key:

`ssh -i .apollo/$network-name+$timestamp/id_rsa admin@bastion_host_ip`

On bastion you will find Node{n} scripts. To get nodes IP addresses for each one run:

`cat /usr/local/bin/Node*`

You can then connect to each node from your workstation using the same key like for bastion.

Additional tools you will find under links
* eth-stats: http://bastion_host_ip:3000
* Grafana: http://bastion_host_ip:3001
* Prometheus: http://bastion_host_ip:9090

## Monitoring and logs

### Grafana logs

For infrastructure monitoring and incident response, you no longer need to switch to other tools to debug what went wrong. Explore allows you to dig deeper into your metrics and logs to find the cause. Grafana’s new logging data source, Loki is tightly integrated into Explore and allows you to correlate metrics and logs by viewing them side-by-side
More info: https://grafana.com/docs/features/explore/

### Dashboard JSON

A dashboard in Grafana is represented by a JSON object, which stores metadata of its dashboard. Dashboard metadata includes dashboard properties, metadata from panels, template variables, panel queries, etc.
More info: https://grafana.com/docs/reference/dashboard/

### Error handling
When you are running command through CLI it should end with exit code status. Statuses are present in:
`commands.go`

## Built with

- Chainhammer : https://github.com/drandreaskrueger/chainhammer

- Quorum Cloud: https://github.com/jpmorganchase/quorum-cloud

- Terraform: https://www.terraform.io/

## Contributing 

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning 

We use [SemVer](https://semver.org/) for versioning.


## License 

This project is licensed under the MIT License - see the LICENSE.md file for details

## Acknowledgements

We would like to thank the following teams for their contributions to the project:

- [binarapps](https://binarapps.com) for their DevOps expertise
- Dr Andreas Krueger for [chainhammer](https://github.com/drandreaskrueger/chainhammer/tree/master/hammer). The name is inspired by his project, and much of the code for firing the transactions is mostly his.
- The JP Morgan Team for [quorum-cloud](https://github.com/jpmorganchase/quorum-cloud). This formed the boiler plate for the deployments of other clients. 
