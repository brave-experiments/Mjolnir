# Iterations v2.0

1. Initial sprint 180h of work. @Blazej @Grzegorz @Jerzy ​3 weeks ( from 13.08 to 01.09 )

## Goal

 - Cli that can run complete Quorum cluster with parametrization and side services

## Definition of done:

- As an authenticated user of infrastructure
using a CLI command with additional parameters
I can deploy modify and destroy a quorum client cluster with side services ( monitoring, logs, chaos testing, etc )
So I have a complete testing environment available.

## Business requirements filled

- [x] The number of nodes should be configurable
- [] Be deployable across multiple regions
- [x] The whole set up should be deployable with one command.
- [x] Support the following clients:
   [x] - Quorum (https://github.com/jpmorganchase/quorum)
- [] For each client that is supported, the following consensus engines should be configurable:
    - [] Quorum: IBFT , Raft, Clique
- [] Monitor VM metrics
    - [] CPU Utilization
    - [] Memory Utilization
    - [] Disk Reads / Writes
    - [] Network (Egress / Ingress)

- [] For each client, version number should also be configurable.
- [] It should be possible to specify each off the following from the
command line:
    - [x] stepDuration/ Blocktime / Epoch
    - [x] Blockgaslimit
- [] Basic client metrics from relevant exporter
- [] Client container logs available though AWS CloudWatch
- [] Basic client metrics from relevant exporter
- [] Client container logs available though Logging tool
- [] Monitor Ethereum specific metrics:
    - [] Should include logs from the client. These should trigger and alert for WARN or ERROR log levels.
    - [] The following RPC endpoints should be polled: 
        - [] Block Number
        - [] Number of Connected Peers 
        - [] Transactions Time
        - [] Transaction Pool
        - [] Block Processing
        - [] Transaction Propagation 
        - [] Data Rate
- [] Clock Skew
- [] Chaos Testing (Dropped packets , network latency, etc)


## Knowledge to gather:


## Tools to use: 

## CheckList
I can:
- [x] deploy
- [] modify
- [] destroy

a quorum client cluster 
with side services: 
- [x] monitoring
- [] logs
- [] chaos testing

2. Sprint ​Blockchain Monitoring​ 120h of work 2 weeks @Jerzy @Grzegorz @Błażej ​( from 02.08 to 12.09 )

## Goal 

- To extend CLI by adding new clients: Pantheon and POA Network

## Definition of done:
- As authenticated user of infrastructure
using a CLI command with additional parameters
I deploy all three clients: Quorum, Pantheon and POA Network with all necessary side services
so I can extend my testing to next clients.

## Business requirements filled

- Support the following clients:
    - Pantheon (​ https://github.com/PegaSysEng/pantheon​ ) 
    -  POA Network
(​ https://github.com/poanetwork/parity-ethereum/tree/hbbft​ )
-  For each client that is supported, the following consensus engines should be configurable:
    - Pantheon: IBFT 2.0 
    -  POA Network: HBBFT