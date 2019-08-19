Iterations:

* * *

  

1. Initial sprint 60h of work. @Blazej @Jerzy 2 weeks 

    

 

Goal:

 

Cli that can run GETH and Quorum with parametrization

 

Definition of done:

 

As authenticated user of infrastructure

I can manage infrastructure

With additional parameters

So I can monit infrastructure

 

Business requirements filled

 

- The number of nodes should be configurable 
- The whole set up should be deployable with one command.  
- Support the following clients: (B. =&gt; All clients can be accessible within one docker container, or within pod. Decision is ours) 

        ○ Quorum (https://github.com/jpmorganchase/quorum)

        ○ Go-Ethereum ([https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)) 

 

- For each client that is supported, the following consensus engines should be configurable: 

            § Quorum: IBFT , Raft, Clique

            § Go-Ethereum: Clique

 

- Monitor VM metrics 

            § CPU Utilization

            § Memory Utilization

            § Disk Reads / Writes

            § Network (Egress / Ingress)

 

- Basic client metrics from relevant exporter 

- Client container logs available though AWS CloudWatch 

 

Knowledge to gather:

 

Tools to use:

  
  

2. Sprint Blockchain Monitoring 120h of work 2 weeks @Jerzy @Grzegorz @Błażej @[ADDITIONAL] @[ADDITIONAL]

  
  

 

Goal:

 

We can have deployable infrastructure and log whole results

 

Definition of done:

 

As authenticated user of infrastructure

I can monitor blockchain infrastructure standard circumstances 

 

Business requirements filled

 

- Be deployable across multiple regions 2 regions 

 

- Support the following clients: 

    § Parity: Aura , Clique

      § Pantheon: IBFT 2.0, Clique

- Monitor Ethereum specific metrics: 

  § Should include logs from the client. These should trigger 

    and alert for WARN or ERROR log levels.

  § The following RPC endpoints should be polled:

                □ Block Number

                □ Number of  Connected Peers

                □ Transactions Time

                □ Transaction Pool

                □ Block Processing

                □ Transaction Propagation

                □ Data Rate

 

Knowledge to gather:

 

Tools to use:

  
  
  

 

3. What have left:

~~ 80 - 120 H

  
  

 

Goal:

 

We can perform chaos testing and log results

 

Definition of done:

 

As authenticated user of infrastructure

I can spread chaotic circumstances

So I can monitor blockchain infrastructure within chaotic circumstances 

 

Business requirements filled

 

- Allow the user to simulate adverse network conditions. Here is a library that might be under consideration: https://github.com/alexei-led/pumba 

  § This should allow us to choose set the duration of

    jitters, timing , duration and correlation of jitters

  § E.g

  □ pumba netem --duration 5m --interface eth1 delay \

     --time 3000 \

     --jitter 30 \

     --correlation 20

-  It should be possible to specify each on the following from the command line: 

        ○ stepDuration/ Blocktime / Epoch 

        ○ Blockgaslimit 

 

Knowledge to gather:

 

Tools to use:

  
  
  

SUM MAX: 300H =&gt; 6 Weeks
