
# Mjölnir 
## ...[the hammer of Thor](https://en.wikipedia.org/wiki/Mj%C3%B6lnir). 

[![Build Status](https://travis-ci.com/brave-experiments/Mjolnir.svg?token=KboonuECWJN5n1otaykj&branch=master)](https://travis-ci.com/brave-experiments/Mjolnir)
<!--[![Coverage Status](https://coveralls.io/repos/github/<github username>/<repo name>/badge.svg?branch=master)](https://coveralls.io/github/<github username>/<repo name>?branch=master)-->

         
                  |"  \    /"  |     |"  |  /    " \  |"  |    (\"   \|"  \  |" \    /"      \  
                  \   \  //   |     ||  | // ____  \ ||  |    |.\\   \    | ||  |  |:        | 
                  /\\  \/.    |     |:  |/  /    ) :)|:  |    |: \.   \\  | |:  |  |_____/   ) 
                  |: \.        |  ___|  /(: (____/ //  \  |___ |.  \    \. | |.  |   //      /  
                  |.  \    /:  | /  :|_/ )\        /  ( \_|:  \|    \    \ | /\  |\ |:  __   \  
                  |___|\__/|___|(_______/  \"_____/    \_______)\___|\____\)(__\_|_)|__|  \___)

                  T                                    \`.    T
                        |    T     .--------------.___________) \   |    T
                        !    |     |//////////////|___________[ ]   !  T |
                              !     `--------------'           ) (      | !
                                                         mn  '-'      !
         
We needed DevOps tooling to enable the Brave team rapidly deploy Ethereum Proof of Authority (PoA) clusters across different Ethereum Clients for benchmarking.

As there was no tool out there that fulfilled this requirement,this gap gave birth to Mjolnir ...  a tool for rapidly deploying and testing Ethereum Clients. 

The tool currently allows users to test the through put of Ethereum Clients both on its own, and under adverse network conditions (i.e. Clock Skew , dropped tcp packets, jitters, etc.)

At this moment, Mjolnir supports the following clients:

- [Quorum](https://github.com/jpmorganchase/quorum)
- [Patheon](https://github.com/hyperledger/besu) (now Hyperledger Besu)
- [Parity](https://github.com/poanetwork/hbbft) (Honey Badger / POSDAO version)



## Table of Contents
- [Architecture](tbd)
- [Terminology](https://github.com/brave-experiments/Mjolnir#terminology)
- [Requirements](https://github.com/brave-experiments/Mjolnir#requirements)
- [Getting-Started](https://github.com/brave-experiments/Mjolnir#requirements)
   - [Development-Mode](https://github.com/brave-experiments/Mjolnir#development-mode)
   - [Test](https://github.com/brave-experiments/Mjolnir#test)
   - [Test-Watch](https://github.com/brave-experiments/Mjolnir#test-watch)
   - [Build](https://github.com/brave-experiments/Mjolnir#build)
   - [Providing-Values](https://github.com/brave-experiments/Mjolnir#providing-values)
   - [Debugging](https://github.com/brave-experiments/Mjolnir#debugging)
- [Build-Details](https://github.com/brave-experiments/Mjolnir#build-details)
- [Subcommands](https://github.com/brave-experiments/Mjolnir#subcommands)
- [Cleaning-Up](https://github.com/brave-experiments/Mjolnir#cleaning-up)
- [Monitoring-and-Logs](https://github.com/brave-experiments/Mjolnir#monitoring-and-logs)
   - [Grafana-Logs](https://github.com/brave-experiments/Mjolnir#grafana-logs)
   - [Dashboard-JSON](https://github.com/brave-experiments/Mjolnir#dashboard-json)
- [Error-Handling](https://github.com/brave-experiments/Mjolnir#error-handling)
- [Limitations](https://github.com/brave-experiments/Mjolnir#limitations)
-[Built-With](https://github.com/brave-experiments/Mjolnir#built-with)
- [Conributing](https://github.com/brave-experiments/Mjolnir#contributing)
- [License](https://github.com/brave-experiments/Mjolnir#license)

- [Acknowledgements](https://github.com/brave-experiments/Mjolnir#license)



## Architecture

<a href="https://ibb.co/1L5LbVc"><img src="https://i.ibb.co/56y62DZ/M.png" alt="M" border="0"></a>

## Terminology

- **{cli-version}**: Semantic version of binary 
- **{arch}**: The OS architecture. Currently supported are 
   - `osx` 
   - `unix`
- **{binaryName}**: `apollo`
- **{client}**: The Ethereum client been tested. Currently supported are:
   - `quorum`
   - `pantheon`
   - `parity`
- **{cmdName}**: Binary's sub command. 

## Requirements
- A UNIX based machine.
- Go (>= v 1.12.7)
- Docker Engine Community (>= v 19.03.1)
- Terraform (>= v0.12.5)


## Getting started
- **Step 1: Deploy Infrastructure**
   - Clone this repo - ` git clone git@github.com:brave-experiments/Mjolnir.git`
   - Enter `bin/run` to run locally. This will create a local docker container and ssh into it. 
   - Create a copy of the configuration files in `examples/values.yaml` to  `examples/values-local.yaml`
   - Update `examples/values-local.yaml`
   - enter `./dist/{cli-version}/{arch}/apollo apply {client} examples/values-local.yml `. This will deploy the requsite infrastrucure on your AWS account.

- **Step 2: Fire Transactions**
   - Once this is complete, enter `./dist/{cli-version}/{arch}/apollo bastion` to tunnel into the bastion host. It is from here, we are able to access chainhamer for sending transactions to the clients. 
   - Move in the `chainhammer` directory by entering `cd chainhammer`
   - Run `scripts/install-initialize.sh` to intialize chainhammer. 
   - To send transactions, ` CH_TXS=40000 CH_THREADING="threaded2 300" ./run.sh "{TESTNAME}"`; Where `CH_TXS` is the number of transactions to be send, `CH_THREADING` is the threading algorithm, and `{TESTNAME}` is the name that the run will be save under.
   - If all goes well, files will be saved under:
      - ../results/runs/{client}_{date}-{time}_{no_of_transactions_sent}.md
      - ../results/runs/{client}_{date}-{time}_{no_of_transactions_sent}.html
      - ..reader/img/{TESTNAME}-{date}-{time}_blks.pgn


### Development Mode
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

### Providing values
See `example/values.yml` that shows how to attach values to apply execution. 
Since any `values-local.yml` file is gitignored
you should copy `example/values.yml` to `values-local.yml` and provide values that you need.

In test mode cli runs with isolated scope with predefined variables and constants.

### Debugging
After execution of `apply` command certain files will be created on your host:
- `temp.tf` at root of execution dir, which contains whole terraform code that has been executed
- `terraform.tfstate` at root of execution dir, which contains state of execution
- `variables.log` at root of execution dir, which contains last executed variables in recipe
- `.apollo` dir which contains necessary files like ssh key pair to bastion
- `.apollo/$network-name+$timestamp/` is a dir where should end private and public key pair



## Build Details

To manually test build run 

`bin/run`  - or get latest binary release from here: https://github.com/brave-experiments/Mjolnir/releases 

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
## Subcommands

- SSH into the bastion:
   
   `./dist/{cli-version}/{arch}/apollo bastion`

- SSH into an Ethereum node:

   `./dist/{cli-version}/{arch}/apollo node n`

    where `n` is the node number. 


- To attach an interactive geth console to any node: 

   `./dist/{cli-version}/{arch}/apollo geth n` 

   where `n` is the node number. 


- Get information on currently deployed nodes

   `./dist/{cli-version}/{arch}/apollo node-info` 

## Cleaning Up

to destroy run:
`./dist/{cli-version}/{arch}/apollo destroy {values-local.yml}`

Current success output looks like this ( will be correted in next release ):
```Deploy Name not present
   [FINAL] Summary execution: 
   Wrote summarry output to:  .apollo//output.log
   Deploy Name not present
   Restoring env variables.
```


## Monitoring and logs

### Grafana logs

For infrastructure monitoring and incident response, you no longer need to switch to other tools to debug what went wrong. Explore allows you to dig deeper into your metrics and logs to find the cause. Grafana’s new logging data source, Loki is tightly integrated into Explore and allows you to correlate metrics and logs by viewing them side-by-side
More info: https://grafana.com/docs/features/explore/


To access tools to visualise your cluster's performance visit:

* eth-stats: http://bastion_host_ip:3000
* Grafana: http://bastion_host_ip:3001
* Prometheus: http://bastion_host_ip:9090

### Dashboard JSON

A dashboard in Grafana is represented by a JSON object, which stores metadata of its dashboard. Dashboard metadata includes dashboard properties, metadata from panels, template variables, panel queries, etc.
More info: https://grafana.com/docs/reference/dashboard/

## Error handling
When you are running command through CLI it should end with exit code status. Statuses are present in:
`commands.go`

##  Limitations

- **This tool is meant for benchmarking alone and should not be used to deploy production instances.** 
- Some features may not be compatible in Windows environment.
- Only Amazon Web Services (AWS) is supported now. We are however open to PRs for other cloud providers!


## Built with

- Chainhammer: https://github.com/drandreaskrueger/chainhammer
- Quorum Cloud: https://github.com/jpmorganchase/quorum-cloud
- Terraform: https://www.terraform.io/

## Contributing 

Refer to [CONTRIBUTING.md](https://github.com/brave-experiments/Mjolnir/blob/master/docs/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning 

We use [SemVer](https://semver.org/) for versioning.

## Code of Conduct

We subscribe to a strict code of conduct. For more information visit our [CODE_OF_CONDUCT](https://github.com/brave-experiments/Mjolnir/blob/master/CODE_OF_CONDUCT.md) page


## License 

This project is licensed under the Mozilla Public License 2.0- see the [LICENSE](https://github.com/brave-experiments/Mjolnir/blob/master/LICENSE) file for details

## Frequently Asked Questions

- **Why did my deployment fail?:** Although we have made efforts to provide descriptive error codes, there are times when a deployment can fail due to random faults (i.e. connectivity issues). Should this happen to you, clean up the resources and start the process again.

- **When I try to log in to the bastion, it get rejected stating that there have been too many attempts:** Mjolnir registers an identity everytime you deploy new infrastructure. Unfortunately, this is not deleted after the instance is destroyed. As more as more instances are brought up, this leads to too many registers identities. Eventually when you try to log in, it fails as it iterates through all the old ones and times out. The solution is:
   - run `ssh-add -l` to list all the identities
   - Flush them with `ssh-add -D`
   - After doing this, you should be able to log into the bastion. 






## Acknowledgements

We would like to thank the following teams for their contributions to the project:

- [binarapps](https://binarapps.com) for their ability to dive into both the infrastructure and software and deliver on our requirements. 
- Dr Andreas Krueger for [chainhammer](https://github.com/drandreaskrueger/chainhammer/tree/master/hammer). Mjolnir inspired by his project, and much of the code for firing the transactions is mostly his.
- The JP Morgan Team for [quorum-cloud](https://github.com/jpmorganchase/quorum-cloud). This was the boiler plate for the deployments of other clients. 
