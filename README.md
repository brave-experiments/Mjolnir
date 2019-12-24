
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

As there was no tool out there that fulfilled this requirement, this vacuum  gave birth to Mjolnir ...  a tool for rapidly deploying and testing Ethereum Clients. 

Mjolnir allows users to test the through put of Ethereum Clients both on its own, and under adverse network conditions (i.e. Clock Skew , dropped tcp packets, jitters, etc.)

Mjolnir supports the following clients:

- [Quorum](https://github.com/jpmorganchase/quorum)
- [Patheon](https://github.com/hyperledger/besu) (now Hyperledger Besu)
- [Parity](https://github.com/poanetwork/hbbft) (Honey Badger / POSDAO version)



## Table of Contents
- [Architecture](https://github.com/brave-experiments/Mjolnir#architecture)
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
- [Built-With](https://github.com/brave-experiments/Mjolnir#built-with)
- [Contributing](https://github.com/brave-experiments/Mjolnir#contributing)
- [License](https://github.com/brave-experiments/Mjolnir#license)
- [Acknowledgements](https://github.com/brave-experiments/Mjolnir#license)



## Architecture

<a href="https://ibb.co/1L5LbVc"><img src="https://i.ibb.co/56y62DZ/M.png" alt="M" border="0"></a>

## Terminology

- **{binaryName}**: `mjolnir`
- **{client}**: The Ethereum client been tested. Currently supported are:
   - `quorum`
   - `parity`
   -  Support for Hyperlegder Besu is WIP. 
- **{cmdName}**: Binary's sub command. 

## Requirements
- A UNIX based machine.
- Go (>= v 1.12.7)
- Docker Engine Community (>= v 19.03.1)
- Terraform (>= v0.12.5)


## Getting started

**Before proceeding, it is important to ensure that your AWS user has the right permissions for creating IAM roles / policies, EC2 , ECS, s3 and Cloudwatch**. 

- **Step 1: Deploy Infrastructure**
   - Clone this repo - ` git clone git@github.com:brave-experiments/Mjolnir.git`
   - Run `make build` from in root folder. This will build compile the terraform and go modules, creating a `mjolnir` binary in the root folder.  
   - Create a copy of the configuration files in `examples/values.yaml` to  `examples/values-local.yaml` i.e.

      `cp examples/values.yaml examples/values-local.yaml` . This is important as this file is added to the .gitignore file and protects the user from accidentally commiting secrets to their online repository.

   - Update `examples/values-local.yaml` with your desired configuration. 
   - The user can either deploy their cluster by either of the following commands

      `make {client}` 

      OR

      `./mjolnir apply examples/values-local.yml`.

       This will deploy the requisite infrastrucure in the user's AWS account.
      For a 4 node cluster, this takes about 15 minutes. 


- **Step 2: Fire Transactions**
   - Once this is complete, enter `mjolnir bastion`
    to tunnel into the bastion host. It is from here, we are able to access chainhamer for sending transactions to the clients. 
   - Move in the `chainhammer` directory by entering `cd chainhammer`
   - Run `scripts/install-initialize.sh` to intialize chainhammer. 
   - To send transactions, ` CH_TXS=25000 CH_THREADING="threaded2 300" ./run.sh "{TESTNAME}"`; Where `CH_TXS` is the number of transactions to be send, `CH_THREADING` is the threading algorithm, and `{TESTNAME}` is the name that the run will be save under.
   - If all goes well, files will be saved under:
      - ../results/runs/{client}_{date}-{time}_{no_of_transactions_sent}.md
      - ../results/runs/{client}_{date}-{time}_{no_of_transactions_sent}.html
      - ..reader/img/{TESTNAME}-{date}-{time}_blks.pgn


### Build
To build from source:
`make build`

### Test
To run tests without watcher:

`make test`

### Test-watch
To run test watcher type:

`make test-watch`


After a successful build , the binary `mjolnir` will be in the root folder. 

To execute mjolnir binary file:
try `./mjolnir` to see all commands that are registered
try `./mjolnir {cmdName} --help` to see help from command

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
- `.mjolnir` dir which contains necessary files like ssh key pair to bastion
- `.mjolnir/$network-name+$timestamp/` is a dir where should end private and public key pair



## Build Details

To manually test build run 

`bin/run ci`  - or get latest binary release from here: https://github.com/brave-experiments/Mjolnir/releases 

To execute mjolnir CLI run:

`./mjolnir apply quorum values.yml`  - with previosly prepared values.yml taken from `examples/` folder in repo

After successful you will find following files in your working directory:
* terraform.tfstate   - current terraform object state
* temp.tf           - full dump of terraform code
* variables.log     - log file with provided vatiables

On successful run on the output you will see following example information:

```
Created output dir:  .mjolnir/network-name-12345678
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
Wrote summarry output to:  .mjolnir/quorum-bastion-jkopacze-n3-66790866/output.log
Restoring env variables.
```
## Subcommands

- SSH into the bastion:
   
   `./mjolnir bastion`

- SSH into an Ethereum node:

   `./mjolnir node n`

    where `n` is the node number. 


- To attach an interactive geth console to any node: 

   `./mjolnir geth n` 

   where `n` is the node number. 


- Get information on currently deployed nodes

   `./mjolnir node-info` 

## Cleaning Up

to destroy run:

`./mjolnir destroy {values-local.yml}`

or

`make destroy`

Current success output looks like this ( will be correted in next release ):
```Deploy Name not present
   [FINAL] Summary execution: 
   Wrote summarry output to:  .mjolnir//output.log
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

### GIT FLOW

- Create branch with issue you are working with, 4eg:

`feature/01/CLI-initialize-project`

- If it is `fix` for current code that already on master:

`fix/01/CLI-initialize-project-output`

- Create a Pull Pequest to branch `master`

- Master is the developer branch

- Releases will lay on certain locked branches, it will occur here after we will be ready with CI/CD pattern.

### Build, deploy, run

- All commands that must be run to fill CI/CD process are described in `Makefile`
- All commands/scripts that are run from host to set up environment are run via `bin/run`

This logic should be sustained to clarify where code should be executed

### IDE Goland
To forward go dependencies from container to your host write `go mod vendor` within container

### Terraform Code:
- Never use "`" (&#96;) sign in terraform code.
- We pass it to build as static asset, so this sign will be removed from whole string
- All variables that are strictly used like ```"${var.something}"``` must be declared as variable in TF recipe. Otherwise validator will fail

For example:
```
variable "something" {
    description = "variable description"
    default = "default value"
}

```

## Versioning 

We use [SemVer](https://semver.org/) for versioning.


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
