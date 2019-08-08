## Aim: 

- To create DevOps tooling to enable the Brave team rapidly deploy Ethereum Proof of Authority (PoA) clusters across different Ethereum Clients for the purpose of performance testing. 
    - While previous set of required involved 5 clients, this scope has been **narrowed down** to **3** clients ; Pantheon, Quorum and POA Network 
- This exercise is to be broken up into 2 phases with a stretch goal:
    - Testing I (Required by: 13/09/19): 
        - Artifacts delivered in this phase will enable Brave to select an Ethereum Client. Brave is interested in exploring Throughput ( Transactions Per Second), and performance under adverse network conditions (e.g. Clock skew, lost or dropped packets, network latency etc.). 
        - **It is important that for each client, all the components are delivered end to end (eg. When building for Pantheon, build all the sevices need (monitoring, chaos testing, etc)). Brave believes that this delivery pathway will enable the Brave development team to operate in an agile fashion.**
    - Testing II (Required by: 12/11/19):
        - After Testing I, the Brave team will proceed to build protoypes which we would like to have tooling that allows End-To-End testing, in addition to DDoS simulation and load testing.
    - Stretch (Flexible):
        - Remove Dependency from Chainhammer
        - Create tooling to manage nodes in the network (i.e. Remove or Add (non-)Validators).
    

## General:

While Brave is agnostic with regards to the technologies used, the DevOps Toolchain should:
- Offer quick deployment and decommissioning
- Support multiple client implementations 
- Be easily configurable via command line inputs or any other method which the vendor may deem fit.
- Provide Monitoring and Alerting Functionality 

## Testing I (Required by: 13/09/19): 
### Nodes:

- A node is an independent computing unit that runs a blockchain client. 

- The DevOps Toolchain  ****MUST****
  - Be capable of deploying an unlimited number of nodes. 
  - The number of nodes should be configurable
  - The whole set up should be deployable with one command. 

- The DevOps Toolchain  **SHOULD**:
	- Be deployable across multiple regions 


### Clients: 

- A blockchain client is a software agent capable of connecting to a blockchain network . 

- The DevOps Toolchain **MUST**
	- Support the following clients:
		- Quorum (https://github.com/jpmorganchase/quorum)
		- Pantheon (https://github.com/PegaSysEng/pantheon)
		- POA Network (https://github.com/poanetwork/parity-ethereum/tree/hbbft)
    
	- Allow configurable client Parameters : 
		- For each client that is supported, the following consensus engines should be configurable:
			- Quorum: IBFT , Raft 
			- Pantheon: IBFT 2.0 
			- POA Network: HBBFT

	- For each client, version number should also be configurable. 
	- It should be possible to specify each off the following from the command line:
		- stepDuration/ Blocktime / Epoch
		- Blockgaslimit 

- Example(s) : 

    - Quorum Cloud: https://github.com/jpmorganchase/quorum-cloud
    - Ether Cluster: https://github.com/ethereum-classic-cooperative/ethercluster
    - Pantheon Kubernetes Deployment: https://github.com/PegaSysEng/pantheon-k8s

         

###  Clock Skew:

- We would like to investigate the effect of clock skew on different clients. LibFakeTime (https://github.com/wolfcw/libfaketime) is the sample library . 

- The DevOps Toolchain **MUST**:
	- Accept command line parameters to set clock skew for each of the nodes / instances of the client (i.e. +1, +2,-1,+5 for a 4 node cluster)

### Network Emulation
- The DevOps Toolchain **MUST**:
		- Allow the user to simulate adverse network conditions. Here is a library that might be under consideration: https://github.com/alexei-led/pumba
			- This should allow us to choose set the duration of jitters, timing , duration and correlation of jitters
			- E.g 

        pumba netem --duration 5m --interface eth1 delay \
        --time 3000 \
        --jitter 30 \
        --correlation 20 

### Transaction Per Second (TPS)
- In order to test TPS , Chainhammer has been selected as the tool of choice. (https://github.com/drandreaskrueger/chainhammer)


### Monitoring and Alerting:

-  The DevOps Toolchain **MUST**  include a monitoring and alerting functionality.
- While Brave is agnostic as to which technology is used for this, we would like the following:
	-  Monitor VM metrics (if VMs are to be used in the final set up)
        - CPU Utilization
	    - Memory Utilization
		- Disk Reads / Writes
		- Network (Egress / Ingress)
			
	- Monitor Ethereum specific metrics:
		- include logs from the client. These should trigger an alert for WARN or ERROR log levels. 
		- The following RPC end points should be polled:
			- Block Number
			- Number of  Connected Peers
			- Transactions Time 
			- Transaction Pool
			- Block Processing 
			- Transaction Propagation 
			- Data Rate
- Example(s):
    - Go-ethereum Prometheus sample : https://github.com/karalabe/geth-prometheus

## Testing II (Required by: 12/11/19): 

- The Brave team will be testing the resilience of our solution to 
    - Client Usage Patterns 
    - Failure Modes 
        - Nodes Failing
        - Layer 7 DDoS Attacks

### Client Usage Patterns 
- These will test the blockchain ability to handle simulated work loads that mimic the actual usage patterns of Brave's clients. 

- The DevOps Toolchain **MUST**:
    - Simulate client usage patterns. 
    - Provide the ability to measure these results. 
- Example(s):
    - WebLOAD: https://www.radview.com/website-load-testing-tools/
    - LoadNinja : https://loadninja.com/
    - SmartMeter.io : https://www.smartmeter.io/
### Failure Modes
- These tests are to understand our installations behaviour during system failures and network attacks. 

#### Node Failure
- This category refers to failures due to either partial or full node failure.

- The DevOps toolchain **MUST**: 
    - Provide the ability to simulate partial or full node failure.
    - Provide the ability to measure the results of partial or full node failure.

#### Distributed Denial of Service (DDoS):

- These will test vulnerabilities from Layer 7 DDoS attacks (i.e. attacks that exploit the application layer / business logic)

- The DevOps Toolchain **MUST**: 
    - Provide a support for simulating DDoS attacks on our network of nodes. 
- Example(s):
    - HULK: https://packetstormsecurity.com/files/112856/HULK-Http-Unbearable-Load-King.html
    - r-u-dead-yet: https://code.google.com/archive/p/r-u-dead-yet/
    - Tor's Hammer: https://code.google.com/archive/p/r-u-dead-yet/
    - Davos Net: https://packetstormsecurity.com/files/123084/DAVOSET-1.1.3.html
    - Golden Eye: https://packetstormsecurity.com/files/120966/GoldenEye-HTTP-Denial-Of-Service-Tool.html


## Stretch (Flexible):

These requirements are not on the project's `critical pathway` (i.e. they are `nice-to-haves`)

### Remove Dependency on Chainhammer
- The Chain Hammer application suite is quite robust and contains a considerable amount of code that we do not require. It would be nice to:
    - Remove parts we do not need
    - Optimize the functionality we require (i.e. sending bulk transactions)

### Adding and Removing Validators
- It would be nice to have the ability to manage the consortium  to add/ remove both validators and non-validators nodes
- This can either be added as a patch to the existing tooling , or as a stand alone script.




