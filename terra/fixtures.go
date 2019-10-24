package terra

const (
	InvalidOutputFixture = `{"some": "value", "modules": {}}`
	ProperOutputFixture  = `{
    "version": 3,
    "terraform_version": "0.11.13",
    "serial": 3,
    "lineage": "56a6b1ae-aa91-ccb9-bb23-cfd3ab19176f",
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {
                "_status": {
                    "sensitive": false,
                    "type": "string",
                    "value": "Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n"
                },
                "bastion_host_dns": {
                    "sensitive": false,
                    "type": "string",
                    "value": ""
                },
                "bastion_host_ip": {
                    "sensitive": false,
                    "type": "string",
                    "value": "invalid.ip.666"
                },
                "bucket_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1"
                },
                "chain_id": {
                    "sensitive": false,
                    "type": "string",
                    "value": "4856"
                },
                "ecs_cluster_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "quorum-network-cocroaches-attack"
                },
                "grafana_host_url": {
                    "sensitive": false,
                    "type": "string",
                    "value": "http://invalid.ip.666:3001"
                },
                "grafana_password": {
                    "sensitive": false,
                    "type": "string",
                    "value": "-\u003cND{!FA)tLFDoGB"
                },
                "network_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "cocroaches-attack"
                },
                "private_key_file": {
                    "sensitive": false,
                    "type": "string",
                    "value": "/tmp/.terranova273240257/quorum-cocroaches-attack.pem"
                }
            },
            "resources": {
                "aws_autoscaling_group.asg": {
                    "type": "aws_autoscaling_group",
                    "depends_on": [
                        "aws_launch_configuration.lc",
                        "aws_subnet.public.*",
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "quorum-network-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:autoscaling:us-east-2:051582052996:autoScalingGroup:30919fe0-83af-4d9f-b570-8399e4c432eb:autoScalingGroupName/quorum-network-cocroaches-attack",
                            "availability_zones.#": "1",
                            "availability_zones.4293815384": "us-east-2a",
                            "default_cooldown": "30",
                            "desired_capacity": "0",
                            "enabled_metrics.#": "0",
                            "force_delete": "false",
                            "health_check_grace_period": "300",
                            "health_check_type": "EC2",
                            "id": "quorum-network-cocroaches-attack",
                            "launch_configuration": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "launch_template.#": "0",
                            "load_balancers.#": "0",
                            "max_size": "0",
                            "metrics_granularity": "1Minute",
                            "min_size": "0",
                            "name": "quorum-network-cocroaches-attack",
                            "placement_group": "",
                            "protect_from_scale_in": "false",
                            "service_linked_role_arn": "arn:aws:iam::051582052996:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling",
                            "suspended_processes.#": "0",
                            "tags.#": "2",
                            "tags.0.%": "3",
                            "tags.0.key": "ecs_cluster",
                            "tags.0.propagate_at_launch": "1",
                            "tags.0.value": "quorum-network-cocroaches-attack",
                            "tags.1.%": "3",
                            "tags.1.key": "created_by",
                            "tags.1.propagate_at_launch": "1",
                            "tags.1.value": "terraform",
                            "target_group_arns.#": "0",
                            "termination_policies.#": "1",
                            "termination_policies.0": "Default",
                            "vpc_zone_identifier.#": "1",
                            "vpc_zone_identifier.2389729966": "subnet-035251cd0096068f1",
                            "wait_for_capacity_timeout": "10m"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "delete": 600000000000
                            }
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_cloudwatch_log_group.quorum": {
                    "type": "aws_cloudwatch_log_group",
                    "depends_on": [
                        "local.common_tags"
                    ],
                    "primary": {
                        "id": "/ecs/quorum/cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:logs:us-east-2:051582052996:log-group:/ecs/quorum/cocroaches-attack:*",
                            "id": "/ecs/quorum/cocroaches-attack",
                            "kms_key_id": "",
                            "name": "/ecs/quorum/cocroaches-attack",
                            "retention_in_days": "7",
                            "tags.%": "4",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_ecs_cluster.quorum": {
                    "type": "aws_ecs_cluster",
                    "depends_on": [
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                            "id": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                            "name": "quorum-network-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_ecs_task_definition.quorum": {
                    "type": "aws_ecs_task_definition",
                    "depends_on": [
                        "aws_iam_role.ecs_task",
                        "local.container_definitions",
                        "local.shared_volume_name"
                    ],
                    "primary": {
                        "id": "quorum-istanbul-tessera-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:ecs:us-east-2:051582052996:task-definition/quorum-istanbul-tessera-cocroaches-attack:2",
                            "container_definitions": "[{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"mkdir -p /qdata/dd/geth\\necho \\\"\\\" \u003e /qdata/passwords.txt\\nbootnode -genkey /qdata/dd/geth/nodekey\\nexport NODE_ID=$(bootnode -nodekey /qdata/dd/geth/nodekey -writeaddress)\\necho Creating an account for this node\\ngeth --datadir /qdata/dd account new --password /qdata/passwords.txt\\nexport KEYSTORE_FILE=$(ls /qdata/dd/keystore/ | head -n1)\\nexport ACCOUNT_ADDRESS=$(cat /qdata/dd/keystore/$KEYSTORE_FILE | sed 's/^.*\\\"address\\\":\\\"\\\\([^\\\"]*\\\\)\\\".*$/\\\\1/g')\\necho Writing account address $ACCOUNT_ADDRESS to /qdata/first_account_address\\necho $ACCOUNT_ADDRESS \u003e /qdata/first_account_address\\necho Writing Node Id [$NODE_ID] to /qdata/node_id\\necho $NODE_ID \u003e /qdata/node_id\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -f /qdata/node_id ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/quorum:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"node-key-bootstrap\",\"portMappings\":[],\"volumesFrom\":[]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until Node Key is ready ...\\nwhile [ ! -f \\\"/qdata/node_id\\\" ]; do sleep 1; done\\napk update\\napk add curl jq\\nexport TASK_REVISION=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq '.Revision' -r)\\necho \\\"Task Revision: $TASK_REVISION\\\"\\necho $TASK_REVISION \u003e /qdata/task_revision\\nexport HOST_IP=$(/usr/bin/curl http://169.254.169.254/latest/meta-data/public-ipv4)\\necho \\\"Host IP: $HOST_IP\\\"\\necho $HOST_IP \u003e /qdata/host_ip\\nexport TASK_ARN=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')\\nexport REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')\\naws ecs describe-tasks --region $REGION --cluster quorum-network-cocroaches-attack --tasks $TASK_ARN | jq -r '.tasks[0] | .group' \u003e /qdata/service\\nmkdir -p /qdata/hosts\\nmkdir -p /qdata/nodeids\\nmkdir -p /qdata/accounts\\nmkdir -p /qdata/lib\\naws s3 cp s3://us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/libs/libfaketime.so /qdata/lib/libfaketime.so\\naws s3 cp /qdata/node_id s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/nodeids/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\naws s3 cp /qdata/host_ip s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/hosts/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\naws s3 cp /qdata/first_account_address s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/accounts/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/hosts | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/hosts /qdata/hosts \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other containers to report their IPs ... $count/0\\\"; sleep 1; done\\necho \\\"All containers have reported their IPs\\\"\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/accounts | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/accounts /qdata/accounts \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other nodes to report their accounts ... $count/0\\\"; sleep 1; done\\necho \\\"All nodes have registered accounts\\\"\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/nodeids | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/nodeids /qdata/nodeids \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other nodes to report their IDs ... $count/0\\\"; sleep 1; done\\necho \\\"All nodes have registered their IDs\\\"\\nalloc=\\\"\\\"; for f in $(ls /qdata/accounts); do address=$(cat /qdata/accounts/$f); alloc=\\\"$alloc,\\\\\\\"$address\\\\\\\": { \\\"balance\\\": \\\"\\\\\\\"1000000000000000000000000000\\\\\\\"\\\"}\\\"; done\\nalloc=\\\"{${alloc:1}}\\\"\\nextraData=\\\"\\\\\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\\\\\"\\\"\\napk add --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/community go=1.9.4-r0\\napk add git gcc musl-dev linux-headers\\ngit clone https://github.com/getamis/istanbul-tools /istanbul-tools/src/github.com/getamis/istanbul-tools\\nexport GOPATH=/istanbul-tools\\nexport GOROOT=/usr/lib/go\\necho 'package main\\n\\nimport (\\n\\t\\\"encoding/hex\\\"\\n\\t\\\"fmt\\\"\\n\\t\\\"os\\\"\\n\\n\\t\\\"github.com/ethereum/go-ethereum/crypto\\\"\\n\\t\\\"github.com/ethereum/go-ethereum/p2p/discover\\\"\\n)\\n\\nfunc main() {\\n\\tif len(os.Args) \u003c 2 {\\n\\t\\tfmt.Println(\\\"missing enode value\\\")\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tenode := os.Args[1]\\n\\tnodeId, err := discover.HexID(enode)\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tpub, err := nodeId.Pubkey()\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tfmt.Printf(\\\"0x%s\\\\n\\\", hex.EncodeToString(crypto.PubkeyToAddress(*pub).Bytes()))\\n}\\n' \u003e /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go\\nall=\\\"\\\"; for f in $(ls /qdata/nodeids); do address=$(cat /qdata/nodeids/$f); all=\\\"$all,$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go $address)\\\"; done\\nall=\\\"${all:1}\\\"\\necho Validator Addresses: $all\\nextraData=\\\"\\\\\\\"$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/cmd/istanbul/main.go extra encode --validators $all | awk -F: '{print $2}' | tr -d ' ')\\\\\\\"\\\"\\nmixHash=\\\"\\\\\\\"0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365\\\\\\\"\\\"\\ndifficulty=\\\"\\\\\\\"0x01\\\\\\\"\\\"\\necho '{\\\"alloc\\\":{},\\\"coinbase\\\":\\\"0x0000000000000000000000000000000000000000\\\",\\\"config\\\":{\\\"byzantiumBlock\\\":1,\\\"chainId\\\":4856,\\\"eip150Block\\\":1,\\\"eip150Hash\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"eip155Block\\\":0,\\\"eip158Block\\\":1,\\\"homesteadBlock\\\":0,\\\"isQuorum\\\":true},\\\"difficulty\\\":\\\"0x0\\\",\\\"extraData\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"gasLimit\\\":\\\"0xE0000000\\\",\\\"mixHash\\\":\\\"0x00000000000000000000000000000000000000647572616c65787365646c6578\\\",\\\"nonce\\\":\\\"0x0\\\",\\\"parentHash\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"timestamp\\\":\\\"0x00\\\"}' | jq \\\". + { alloc : $alloc, extraData: $extraData, mixHash: $mixHash, difficulty: $difficulty} | .config=.config + {istanbul: {epoch: 30000, policy: 0} }\\\" \u003e /qdata/genesis.json\\ncat /qdata/genesis.json\\necho \\\"Done!\\\" \u003e /qdata/metadata_bootstrap_container_status\\necho Wait until privacy engine initialized ...\\nwhile [ ! -f \\\"/qdata/.pub\\\" ]; do sleep 1; done\\naws s3 cp /qdata/.pub s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/privacyaddresses/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -f /qdata/metadata_bootstrap_container_status ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"senseyeio/alpine-aws-cli:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"metamain-bootstrap\",\"portMappings\":[],\"volumesFrom\":[{\"sourceContainer\":\"node-key-bootstrap\"}]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until metadata bootstrap completed ...\\nwhile [ ! -f \\\"/qdata/metadata_bootstrap_container_status\\\" ]; do sleep 1; done\\necho Wait until tessera is ready ...\\nwhile [ ! -S \\\"/qdata/tm.ipc\\\" ]; do sleep 1; done\\nmkdir -p /qdata/dd/geth\\necho \\\"\\\" \u003e /qdata/passwords.txt\\necho \\\"Creating /qdata/dd/static-nodes.json and /qdata/dd/permissioned-nodes.json\\\"\\nall=\\\"\\\"; for f in $(ls /qdata/nodeids); do nodeid=$(cat /qdata/nodeids/$f); ip=$(cat /qdata/hosts/$f); all=\\\"$all,\\\\\\\"enode://$nodeid@$ip:21000?discport=0\u0026\\\\\\\"\\\"; done; all=${all:1}\\necho \\\"[$all]\\\" \u003e /qdata/dd/static-nodes.json\\necho \\\"[$all]\\\" \u003e /qdata/dd/permissioned-nodes.json\\necho Permissioned Nodes: $(cat /qdata/dd/permissioned-nodes.json)\\ngeth --datadir /qdata/dd init /qdata/genesis.json\\nexport IDENTITY=$(cat /qdata/service | awk -F: '{print $2}')\\necho 'Running geth with: --datadir /qdata/dd --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --rpcport 22000 --port 21000 --unlock 0 --password /qdata/passwords.txt --nodiscover --networkid 4856 --verbosity 5 --debug --identity $IDENTITY --ethstats \\\"$IDENTITY:b4402044e9f7b5cde81aa4f70a0552a8@10.0.0.239:3000\\\" --istanbul.blockperiod 1 --emitcheckpoints --syncmode full --mine --minerthreads 1'\\ngeth --datadir /qdata/dd --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --rpcport 22000 --port 21000 --unlock 0 --password /qdata/passwords.txt --nodiscover --networkid 4856 --verbosity 5 --debug --identity $IDENTITY --ethstats \\\"$IDENTITY:b4402044e9f7b5cde81aa4f70a0552a8@10.0.0.239:3000\\\" --istanbul.blockperiod 1 --emitcheckpoints --syncmode full --mine --minerthreads 1\"],\"environment\":[{\"name\":\"LD_PRELOAD\",\"value\":\"/qdata/lib/libfaketime.so\"},{\"name\":\"PRIVATE_CONFIG\",\"value\":\"/qdata/tm.ipc\"}],\"essential\":true,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -S /qdata/dd/geth.ipc ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/quorum:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"quorum-run\",\"portMappings\":[{\"containerPort\":22000,\"hostPort\":22000,\"protocol\":\"tcp\"},{\"containerPort\":21000,\"hostPort\":21000,\"protocol\":\"tcp\"},{\"containerPort\":50400,\"hostPort\":50400,\"protocol\":\"tcp\"}],\"volumesFrom\":[{\"sourceContainer\":\"metamain-bootstrap\"},{\"sourceContainer\":\"tessera-run\"}]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until metadata bootstrap completed ...\\nwhile [ ! -f \\\"/qdata/metadata_bootstrap_container_status\\\" ]; do sleep 1; done\\napk update\\napk add jq\\ncd /qdata; echo \\\"\\n\\\" | java -jar /tessera/tessera-app.jar -keygen /qdata/\\nexport HOST_IP=$(cat /qdata/host_ip)\\nexport TM_PUB=$(cat /qdata/.pub)\\nexport TM_KEY=$(cat /qdata/.key)\\necho \\\"\\nHost IP: $HOST_IP\\\"\\necho \\\"Public Key: $TM_PUB\\\"\\nall=\\\"\\\"; for f in $(ls /qdata/hosts | grep -v ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g')); do ip=$(cat /qdata/hosts/$f); all=\\\"$all,{ \\\\\\\"url\\\\\\\": \\\\\\\"http://$ip:9000/\\\\\\\" }\\\"; done\\nall=\\\"[{ \\\\\\\"url\\\\\\\": \\\\\\\"http://$HOST_IP:9000/\\\\\\\" }$all]\\\"\\nexport TESSERA_VERSION=latest\\nexport V=$(echo -e \\\"0.8\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\necho \\\"Creating /qdata/tessera.cfg\\\"\\nDDIR=/qdata/dd\\nunzip -p /tessera/tessera-app.jar META-INF/MANIFEST.MF | grep Tessera-Version | cut -d: -f2 | xargs\\necho \\\"Tessera Version: $TESSERA_VERSION\\\"\\nV08=$(echo -e \\\"0.8\\\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\nV09=$(echo -e \\\"0.9\\\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\ncase \\\"$TESSERA_VERSION\\\" in\\n    \\\"$V09\\\"|latest)\\n    # use new config\\n    cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n  \\\"useWhiteList\\\": false,\\n  \\\"jdbc\\\": {\\n    \\\"username\\\": \\\"sa\\\",\\n    \\\"password\\\": \\\"\\\",\\n    \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n    \\\"autoCreateTables\\\": true\\n  },\\n  \\\"serverConfigs\\\":[\\n  {\\n    \\\"app\\\":\\\"ThirdParty\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"http://$HOST_IP:9080\\\",\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"Q2T\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"unix:/qdata/tm.ipc\\\",\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"P2P\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"http://$HOST_IP:9000\\\",\\n    \\\"sslConfig\\\": {\\n      \\\"tls\\\": \\\"OFF\\\",\\n      \\\"generateKeyStoreIfNotExisted\\\": true,\\n      \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n      \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n      \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n      \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n      \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n      \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  }\\n  ],\\n  \\\"peer\\\": $all,\\n  \\\"keys\\\": {\\n    \\\"passwords\\\": [],\\n    \\\"keyData\\\": [\\n      {\\n        \\\"config\\\": $TM_KEY,\\n        \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n      }\\n    ]\\n  },\\n  \\\"alwaysSendTo\\\": []\\n}\\nEOF    \\n      ;;\\n    \\\"$V08\\\")\\n      # use enhanced config\\n      cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n  \\\"useWhiteList\\\": false,\\n  \\\"jdbc\\\": {\\n    \\\"username\\\": \\\"sa\\\",\\n    \\\"password\\\": \\\"\\\",\\n    \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n    \\\"autoCreateTables\\\": true\\n  },\\n  \\\"serverConfigs\\\":[\\n  {\\n    \\\"app\\\":\\\"ThirdParty\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"INET\\\",\\n      \\\"port\\\": 9080,\\n      \\\"hostName\\\": \\\"http://$HOST_IP\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"Q2T\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"UNIX\\\",\\n      \\\"path\\\":\\\"/qdata/tm.ipc\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"UNIX_SOCKET\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"P2P\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"INET\\\",\\n      \\\"port\\\": 9000,\\n      \\\"hostName\\\": \\\"http://$HOST_IP\\\"\\n    },\\n    \\\"sslConfig\\\": {\\n      \\\"tls\\\": \\\"OFF\\\",\\n      \\\"generateKeyStoreIfNotExisted\\\": true,\\n      \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n      \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n      \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n      \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n      \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n      \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  }\\n  ],\\n  \\\"peer\\\": $all,\\n  \\\"keys\\\": {\\n    \\\"passwords\\\": [],\\n    \\\"keyData\\\": [\\n      {\\n        \\\"config\\\": $TM_KEY,\\n        \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n      }\\n    ]\\n  },\\n  \\\"alwaysSendTo\\\": []\\n}\\nEOF\\n      ;;\\n    *)\\n    # use old config\\n    cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n    \\\"useWhiteList\\\": false,\\n    \\\"jdbc\\\": {\\n        \\\"username\\\": \\\"sa\\\",\\n        \\\"password\\\": \\\"\\\",\\n        \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n        \\\"autoCreateTables\\\": true\\n    },\\n    \\\"server\\\": {\\n        \\\"port\\\": 9000,\\n        \\\"hostName\\\": \\\"http://$HOST_IP\\\",\\n        \\\"sslConfig\\\": {\\n            \\\"tls\\\": \\\"OFF\\\",\\n            \\\"generateKeyStoreIfNotExisted\\\": true,\\n            \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n            \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n            \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n            \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n            \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n            \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n            \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n            \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n            \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n            \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n            \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n            \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n        }\\n    },\\n    \\\"peer\\\": $all,\\n    \\\"keys\\\": {\\n        \\\"passwords\\\": [],\\n        \\\"keyData\\\": [\\n            {\\n                \\\"config\\\": $TM_KEY,\\n                \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n            }\\n        ]\\n    },\\n    \\\"alwaysSendTo\\\": [],\\n    \\\"unixSocketFile\\\": \\\"/qdata/tm.ipc\\\"\\n}\\nEOF\\n      ;;\\nesac\\ncat /qdata/tessera.cfg\\n\\njava -jar /tessera/tessera-app.jar -configfile /qdata/tessera.cfg\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -S /qdata/tm.ipc ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/tessera:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"tessera-run\",\"portMappings\":[{\"containerPort\":9000,\"hostPort\":9000,\"protocol\":\"tcp\"},{\"containerPort\":9080,\"hostPort\":9080,\"protocol\":\"tcp\"}],\"volumesFrom\":[{\"sourceContainer\":\"metamain-bootstrap\"}]}]",
                            "cpu": "4096",
                            "execution_role_arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "family": "quorum-istanbul-tessera-cocroaches-attack",
                            "id": "quorum-istanbul-tessera-cocroaches-attack",
                            "memory": "8192",
                            "network_mode": "bridge",
                            "placement_constraints.#": "0",
                            "requires_compatibilities.#": "1",
                            "requires_compatibilities.2737437151": "EC2",
                            "revision": "2",
                            "task_role_arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "volume.#": "1",
                            "volume.4175909684.host_path": "",
                            "volume.4175909684.name": "quorum_shared_volume"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_instance_profile.bastion": {
                    "type": "aws_iam_instance_profile",
                    "depends_on": [
                        "aws_iam_role.bastion",
                        "local.default_bastion_resource_name"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:instance-profile/quorum-bastion-cocroaches-attack",
                            "create_date": "2019-09-05T10:13:33Z",
                            "id": "quorum-bastion-cocroaches-attack",
                            "name": "quorum-bastion-cocroaches-attack",
                            "path": "/",
                            "role": "quorum-bastion-cocroaches-attack",
                            "roles.#": "0",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_instance_profile.ecsInstanceProfile": {
                    "type": "aws_iam_instance_profile",
                    "depends_on": [
                        "aws_iam_role.ecsInstanceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceProfile-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:instance-profile/ecsInstanceProfile-dccfbe87",
                            "create_date": "2019-09-05T10:13:33Z",
                            "id": "ecsInstanceProfile-dccfbe87",
                            "name": "ecsInstanceProfile-dccfbe87",
                            "path": "/",
                            "role": "ecsInstanceRole-dccfbe87",
                            "roles.#": "0",
                            "unique_id": ""
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_policy.bastion": {
                    "type": "aws_iam_policy",
                    "depends_on": [
                        "data.aws_iam_policy_document.bastion"
                    ],
                    "primary": {
                        "id": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "description": "This policy allows task to access S3 bucket and ECS",
                            "id": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "name": "quorum-bastion-policy-cocroaches-attack",
                            "path": "/",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:*\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowEC2\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ec2:*\",\n      \"Resource\": \"*\"\n    }\n  ]\n}"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_policy.ecs_task": {
                    "type": "aws_iam_policy",
                    "depends_on": [
                        "data.aws_iam_policy_document.ecs_task"
                    ],
                    "primary": {
                        "id": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "description": "This policy allows task to access S3 bucket",
                            "id": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "name": "quorum-ecs-task-policy-cocroaches-attack",
                            "path": "/",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3Access\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:DescribeTasks\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    }\n  ]\n}"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.bastion": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "local.default_bastion_resource_name"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/quorum-bastion-cocroaches-attack",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "quorum-bastion-cocroaches-attack",
                            "max_session_duration": "3600",
                            "name": "quorum-bastion-cocroaches-attack",
                            "path": "/",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecsInstanceRole": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceRole-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecsInstanceRole-dccfbe87",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"\",\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "ecsInstanceRole-dccfbe87",
                            "max_session_duration": "3600",
                            "name": "ecsInstanceRole-dccfbe87",
                            "path": "/",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecsServiceRole": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsServiceRole-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecsServiceRole-dccfbe87",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"\",\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ecs.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "ecsServiceRole-dccfbe87",
                            "max_session_duration": "3600",
                            "name": "ecsServiceRole-dccfbe87",
                            "path": "/",
                            "unique_id": ""
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecs_task": {
                    "type": "aws_iam_role",
                    "depends_on": [],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":[\"ecs-tasks.amazonaws.com\",\"ecs.amazonaws.com\"]},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "quorum-ecs-task-cocroaches-attack",
                            "max_session_duration": "3600",
                            "name": "quorum-ecs-task-cocroaches-attack",
                            "path": "/ecs/",
                            "unique_id": ""
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy.ecsInstanceRolePolicy": {
                    "type": "aws_iam_role_policy",
                    "depends_on": [
                        "aws_iam_role.ecsInstanceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceRole-dccfbe87:ecsInstanceRolePolicy-dccfbe87",
                        "attributes": {
                            "id": "ecsInstanceRole-dccfbe87:ecsInstanceRolePolicy-dccfbe87",
                            "name": "ecsInstanceRolePolicy-dccfbe87",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ecs:CreateCluster\",\n        \"ecs:DeregisterContainerInstance\",\n        \"ecs:DiscoverPollEndpoint\",\n        \"ecs:Poll\",\n        \"ecs:RegisterContainerInstance\",\n        \"ecs:StartTelemetrySession\",\n        \"ecs:Submit*\",\n        \"ecr:GetAuthorizationToken\",\n        \"ecr:BatchCheckLayerAvailability\",\n        \"ecr:GetDownloadUrlForLayer\",\n        \"ecr:BatchGetImage\",\n        \"logs:CreateLogStream\",\n        \"logs:PutLogEvents\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\n",
                            "role": "ecsInstanceRole-dccfbe87"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy.ecsServiceRolePolicy": {
                    "type": "aws_iam_role_policy",
                    "depends_on": [
                        "aws_iam_role.ecsServiceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsServiceRole-dccfbe87:ecsServiceRolePolicy-dccfbe87",
                        "attributes": {
                            "id": "ecsServiceRole-dccfbe87:ecsServiceRolePolicy-dccfbe87",
                            "name": "ecsServiceRolePolicy-dccfbe87",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ec2:AuthorizeSecurityGroupIngress\",\n        \"ec2:Describe*\",\n        \"elasticloadbalancing:DeregisterInstancesFromLoadBalancer\",\n        \"elasticloadbalancing:DeregisterTargets\",\n        \"elasticloadbalancing:Describe*\",\n        \"elasticloadbalancing:RegisterInstancesWithLoadBalancer\",\n        \"elasticloadbalancing:RegisterTargets\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\n",
                            "role": "ecsServiceRole-dccfbe87"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.bastion": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_policy.bastion",
                        "aws_iam_role.bastion"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack-20190905101338541400000004",
                        "attributes": {
                            "id": "quorum-bastion-cocroaches-attack-20190905101338541400000004",
                            "policy_arn": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "role": "quorum-bastion-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_cloudwatch": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101333138700000002",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101333138700000002",
                            "policy_arn": "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_execution": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101333138100000001",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101333138100000001",
                            "policy_arn": "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_s3": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_policy.ecs_task",
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101338522200000003",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101338522200000003",
                            "policy_arn": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_instance.bastion": {
                    "type": "aws_instance",
                    "depends_on": [
                        "aws_iam_instance_profile.bastion",
                        "aws_key_pair.ssh",
                        "aws_security_group.bastion-ethstats",
                        "aws_security_group.bastion-ssh",
                        "aws_security_group.quorum",
                        "aws_subnet.public",
                        "data.aws_ami.this",
                        "local.bastion_bucket",
                        "local.common_tags",
                        "local.default_bastion_resource_name",
                        "local.ethstats_docker_image",
                        "local.ethstats_port",
                        "local.quorum_docker_image",
                        "random_id.ethstat_secret",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "i-092a0879925d45d82",
                        "attributes": {
                            "ami": "ami-00c03f7f7f2ec15c3",
                            "associate_public_ip_address": "true",
                            "availability_zone": "us-east-2a",
                            "cpu_core_count": "2",
                            "cpu_threads_per_core": "1",
                            "credit_specification.#": "1",
                            "credit_specification.0.cpu_credits": "standard",
                            "disable_api_termination": "false",
                            "ebs_block_device.#": "0",
                            "ebs_optimized": "false",
                            "ephemeral_block_device.#": "0",
                            "get_password_data": "false",
                            "iam_instance_profile": "quorum-bastion-cocroaches-attack",
                            "id": "i-092a0879925d45d82",
                            "instance_state": "running",
                            "instance_type": "t2.large",
                            "ipv6_addresses.#": "0",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "monitoring": "false",
                            "network_interface.#": "0",
                            "network_interface_id": "eni-0e00d78db53cf56c9",
                            "password_data": "",
                            "placement_group": "",
                            "primary_network_interface_id": "eni-0e00d78db53cf56c9",
                            "private_dns": "ip-10-0-0-239.us-east-2.compute.internal",
                            "private_ip": "10.0.0.239",
                            "public_dns": "",
                            "public_ip": "invalid.ip.666",
                            "root_block_device.#": "1",
                            "root_block_device.0.delete_on_termination": "true",
                            "root_block_device.0.iops": "100",
                            "root_block_device.0.volume_id": "vol-0efe095602110db6f",
                            "root_block_device.0.volume_size": "8",
                            "root_block_device.0.volume_type": "gp2",
                            "security_groups.#": "0",
                            "source_dest_check": "true",
                            "subnet_id": "subnet-035251cd0096068f1",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "tenancy": "default",
                            "user_data": "c388a9205eb078013748e9431a2c713189dc038c",
                            "volume_tags.%": "0",
                            "vpc_security_group_ids.#": "3",
                            "vpc_security_group_ids.1114044214": "sg-02d4a1108481c7db4",
                            "vpc_security_group_ids.1424535282": "sg-047de8aa5ccdc10f5",
                            "vpc_security_group_ids.1485421960": "sg-04e6635eb212ff555"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 1200000000000,
                                "update": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_internet_gateway.this": {
                    "type": "aws_internet_gateway",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "igw-0d383894b8f6b5e94",
                        "attributes": {
                            "id": "igw-0d383894b8f6b5e94",
                            "tags.%": "1",
                            "tags.Name": "",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_key_pair.ssh": {
                    "type": "aws_key_pair",
                    "depends_on": [
                        "local.default_bastion_resource_name",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "id": "quorum-bastion-cocroaches-attack",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC775eGcJ2GZBfUfpiQnXqW/WfhQD8C7mnI1qz1COnR3yft7fpwVNP1wJ9lC8DB6HUdsKl0WISydrMrkwk5aIh4Hr5WJh2hcZq73FFUvQhEwnKvjce6rGTjl5pIRFwLgA9KB4hXfwCii6J+ul5CZs9zIteqaKovQbXcWh/6y0u0yvppP16Y2NuoDWqJjNNDo1QeiS27Ft88jbxVqX/B35cIzkjBiWvoLeoG6wD+K6r/kX0OMXomwyu4Ofc088WON4cXIiGHudDDMJPZD/h0hKAseDyflwpsY4qCG4bJpij3IcruvrMB1G+z0cpcERf3MtxFs+ob+LPZCD6v4QI2Eszv"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_kms_key.bucket": {
                    "type": "aws_kms_key",
                    "depends_on": [
                        "data.aws_iam_policy_document.kms_policy",
                        "local.common_tags"
                    ],
                    "primary": {
                        "id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                        "attributes": {
                            "arn": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "deletion_window_in_days": "7",
                            "description": "Used to encrypt/decrypt objects stored inside bucket created for this deployment",
                            "enable_key_rotation": "false",
                            "id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "is_enabled": "true",
                            "key_id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "key_usage": "ENCRYPT_DECRYPT",
                            "policy": "{\"Statement\":[{\"Action\":\"kms:*\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::051582052996:root\"},\"Resource\":\"*\",\"Sid\":\"AllowAccess\"}],\"Version\":\"2012-10-17\"}",
                            "tags.%": "4",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_launch_configuration.lc": {
                    "type": "aws_launch_configuration",
                    "depends_on": [
                        "aws_iam_instance_profile.ecsInstanceProfile",
                        "aws_key_pair.ssh",
                        "aws_security_group.quorum",
                        "data.aws_ami.ecs_ami",
                        "data.template_file.user_data",
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "quorum-network-cocroaches-attack20190905101403318200000005",
                        "attributes": {
                            "associate_public_ip_address": "false",
                            "ebs_block_device.#": "0",
                            "ebs_optimized": "false",
                            "enable_monitoring": "true",
                            "ephemeral_block_device.#": "0",
                            "iam_instance_profile": "ecsInstanceProfile-dccfbe87",
                            "id": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "image_id": "ami-035a1bdaf0e4bf265",
                            "instance_type": "t2.xlarge",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "name": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "name_prefix": "quorum-network-cocroaches-attack",
                            "root_block_device.#": "1",
                            "root_block_device.0.delete_on_termination": "true",
                            "root_block_device.0.iops": "0",
                            "root_block_device.0.volume_size": "16",
                            "root_block_device.0.volume_type": "",
                            "security_groups.#": "1",
                            "security_groups.1114044214": "sg-02d4a1108481c7db4",
                            "spot_price": "",
                            "user_data": "08fe58a67d425d4ad60356e369bf6e3c353f403b",
                            "vpc_classic_link_id": "",
                            "vpc_classic_link_security_groups.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route.public_internet_gateway": {
                    "type": "aws_route",
                    "depends_on": [
                        "aws_internet_gateway.this",
                        "aws_route_table.public"
                    ],
                    "primary": {
                        "id": "r-rtb-04cfca31968dae4771080289494",
                        "attributes": {
                            "destination_cidr_block": "0.0.0.0/0",
                            "destination_prefix_list_id": "",
                            "egress_only_gateway_id": "",
                            "gateway_id": "igw-0d383894b8f6b5e94",
                            "id": "r-rtb-04cfca31968dae4771080289494",
                            "instance_id": "",
                            "instance_owner_id": "",
                            "nat_gateway_id": "",
                            "network_interface_id": "",
                            "origin": "CreateRoute",
                            "route_table_id": "rtb-04cfca31968dae477",
                            "state": "active",
                            "vpc_peering_connection_id": ""
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 300000000000,
                                "delete": 300000000000
                            }
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.0": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-0539e90b9bba7657f",
                        "attributes": {
                            "id": "rtb-0539e90b9bba7657f",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.1": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-03b7807afd2862da6",
                        "attributes": {
                            "id": "rtb-03b7807afd2862da6",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2b",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.2": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-0d69c07019210736f",
                        "attributes": {
                            "id": "rtb-0d69c07019210736f",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2c",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.public": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-04cfca31968dae477",
                        "attributes": {
                            "id": "rtb-04cfca31968dae477",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-public",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.0": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-0b1ffd059fd0a8a08",
                        "attributes": {
                            "id": "rtbassoc-0b1ffd059fd0a8a08",
                            "route_table_id": "rtb-0539e90b9bba7657f",
                            "subnet_id": "subnet-08c762a35be7b0024"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.1": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-00f88936a60c4cce2",
                        "attributes": {
                            "id": "rtbassoc-00f88936a60c4cce2",
                            "route_table_id": "rtb-03b7807afd2862da6",
                            "subnet_id": "subnet-0360554e5daf7b5d4"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.2": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-05da0e76079fcfc35",
                        "attributes": {
                            "id": "rtbassoc-05da0e76079fcfc35",
                            "route_table_id": "rtb-0d69c07019210736f",
                            "subnet_id": "subnet-0dfb4a947cdf1ae22"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.public": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.public",
                        "aws_subnet.public.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-010c308d38c127217",
                        "attributes": {
                            "id": "rtbassoc-010c308d38c127217",
                            "route_table_id": "rtb-04cfca31968dae477",
                            "subnet_id": "subnet-035251cd0096068f1"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_s3_bucket.bastion": {
                    "type": "aws_s3_bucket",
                    "depends_on": [
                        "local.bastion_bucket"
                    ],
                    "primary": {
                        "id": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                        "attributes": {
                            "acceleration_status": "",
                            "acl": "private",
                            "arn": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket_domain_name": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1.s3.amazonaws.com",
                            "bucket_regional_domain_name": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1.s3.us-east-2.amazonaws.com",
                            "cors_rule.#": "0",
                            "force_destroy": "true",
                            "hosted_zone_id": "Z2O1EMRO9K5GLX",
                            "id": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "logging.#": "0",
                            "region": "us-east-2",
                            "replication_configuration.#": "0",
                            "request_payer": "BucketOwner",
                            "server_side_encryption_configuration.#": "0",
                            "tags.%": "0",
                            "versioning.#": "1",
                            "versioning.0.enabled": "true",
                            "versioning.0.mfa_delete": "false",
                            "website.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_s3_bucket.quorum": {
                    "type": "aws_s3_bucket",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "data.aws_iam_policy_document.bucket_policy",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                        "attributes": {
                            "acceleration_status": "",
                            "acl": "private",
                            "arn": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket_domain_name": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1.s3.amazonaws.com",
                            "bucket_regional_domain_name": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1.s3.us-east-2.amazonaws.com",
                            "cors_rule.#": "0",
                            "force_destroy": "true",
                            "hosted_zone_id": "Z2O1EMRO9K5GLX",
                            "id": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "logging.#": "0",
                            "policy": "{\"Statement\":[{\"Action\":\"s3:*\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::051582052996:root\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"AllowAccess\"},{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"DenyAccess1\"},{\"Action\":\"s3:PutObject\",\"Condition\":{\"StringNotEquals\":{\"s3:x-amz-server-side-encryption\":\"aws:kms\"}},\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"DenyAccess2\"}],\"Version\":\"2012-10-17\"}",
                            "region": "us-east-2",
                            "replication_configuration.#": "0",
                            "request_payer": "BucketOwner",
                            "server_side_encryption_configuration.#": "1",
                            "server_side_encryption_configuration.0.rule.#": "1",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.#": "1",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.0.kms_master_key_id": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.0.sse_algorithm": "aws:kms",
                            "tags.%": "0",
                            "versioning.#": "1",
                            "versioning.0.enabled": "true",
                            "versioning.0.mfa_delete": "false",
                            "website.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-ethstats": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "aws_subnet.public",
                        "local.common_tags",
                        "local.quorum_rpc_port",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-047de8aa5ccdc10f5",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-047de8aa5ccdc10f5",
                            "description": "Security group used by external to access ethstats for Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-047de8aa5ccdc10f5",
                            "ingress.#": "1",
                            "ingress.2872103538.cidr_blocks.#": "1",
                            "ingress.2872103538.cidr_blocks.0": "10.0.0.0/24",
                            "ingress.2872103538.description": "Allow geth console",
                            "ingress.2872103538.from_port": "22000",
                            "ingress.2872103538.ipv6_cidr_blocks.#": "0",
                            "ingress.2872103538.protocol": "tcp",
                            "ingress.2872103538.security_groups.#": "0",
                            "ingress.2872103538.self": "false",
                            "ingress.2872103538.to_port": "22000",
                            "name": "quorum-bastion-ethstats-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "client-bastion-geth-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-geth": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-05f12b4adeda993e5",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-05f12b4adeda993e5",
                            "description": "Security group used by external to access geth for quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-05f12b4adeda993e5",
                            "ingress.#": "1",
                            "ingress.338493142.cidr_blocks.#": "1",
                            "ingress.338493142.cidr_blocks.0": "0.0.0.0/0",
                            "ingress.338493142.description": "Allow ethstats",
                            "ingress.338493142.from_port": "3000",
                            "ingress.338493142.ipv6_cidr_blocks.#": "0",
                            "ingress.338493142.protocol": "tcp",
                            "ingress.338493142.security_groups.#": "0",
                            "ingress.338493142.self": "false",
                            "ingress.338493142.to_port": "3000",
                            "name": "quorum-bastion-get-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-ethstats-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-ssh": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-04e6635eb212ff555",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-04e6635eb212ff555",
                            "description": "Security group used by Bastion node to access Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-04e6635eb212ff555",
                            "ingress.#": "1",
                            "ingress.1248900448.cidr_blocks.#": "1",
                            "ingress.1248900448.cidr_blocks.0": "0.0.0.0/0",
                            "ingress.1248900448.description": "Allow SSH",
                            "ingress.1248900448.from_port": "22",
                            "ingress.1248900448.ipv6_cidr_blocks.#": "0",
                            "ingress.1248900448.protocol": "tcp",
                            "ingress.1248900448.security_groups.#": "0",
                            "ingress.1248900448.self": "false",
                            "ingress.1248900448.to_port": "22",
                            "name": "quorum-bastion-ssh-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-ssh-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.quorum": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-02d4a1108481c7db4",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-02d4a1108481c7db4",
                            "description": "Security group used in Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-02d4a1108481c7db4",
                            "ingress.#": "0",
                            "name": "quorum-sg-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-sg-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.ethstats": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.ethstats_port"
                    ],
                    "primary": {
                        "id": "sgrule-2136424293",
                        "attributes": {
                            "description": "ethstats traffic",
                            "from_port": "3000",
                            "id": "sgrule-2136424293",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "3000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.ethstats-bastion": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.bastion-ethstats",
                        "local.ethstats_port"
                    ],
                    "primary": {
                        "id": "sgrule-435918442",
                        "attributes": {
                            "description": "ethstats traffic",
                            "from_port": "3000",
                            "id": "sgrule-435918442",
                            "protocol": "tcp",
                            "security_group_id": "sg-047de8aa5ccdc10f5",
                            "self": "true",
                            "to_port": "3000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.geth_admin_rpc": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.quorum_rpc_port"
                    ],
                    "primary": {
                        "id": "sgrule-3810161799",
                        "attributes": {
                            "description": "Geth Admin RPC traffic",
                            "from_port": "22000",
                            "id": "sgrule-3810161799",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "22000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.geth_p2p": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.quorum_p2p_port"
                    ],
                    "primary": {
                        "id": "sgrule-256513769",
                        "attributes": {
                            "description": "Geth P2P traffic",
                            "from_port": "21000",
                            "id": "sgrule-256513769",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "21000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.open-all-ingress-research": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum"
                    ],
                    "primary": {
                        "id": "sgrule-2623329713",
                        "attributes": {
                            "cidr_blocks.#": "1",
                            "cidr_blocks.0": "0.0.0.0/0",
                            "description": "Open all ports",
                            "from_port": "0",
                            "id": "sgrule-2623329713",
                            "protocol": "-1",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "false",
                            "to_port": "0",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.tessera": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.tessera_port"
                    ],
                    "primary": {
                        "id": "sgrule-637454889",
                        "attributes": {
                            "description": "Tessera API traffic",
                            "from_port": "9000",
                            "id": "sgrule-637454889",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "9000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.tessera_thirdparty": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.tessera_thirdparty_port"
                    ],
                    "primary": {
                        "id": "sgrule-4039332832",
                        "attributes": {
                            "description": "Tessera Thirdparty API traffic",
                            "from_port": "9080",
                            "id": "sgrule-4039332832",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "9080",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.0": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-08c762a35be7b0024",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2a",
                            "cidr_block": "10.0.1.0/24",
                            "id": "subnet-08c762a35be7b0024",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.1": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-0360554e5daf7b5d4",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2b",
                            "cidr_block": "10.0.2.0/24",
                            "id": "subnet-0360554e5daf7b5d4",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2b",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.2": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-0dfb4a947cdf1ae22",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2c",
                            "cidr_block": "10.0.3.0/24",
                            "id": "subnet-0dfb4a947cdf1ae22",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2c",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.public": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-035251cd0096068f1",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2a",
                            "cidr_block": "10.0.0.0/24",
                            "id": "subnet-035251cd0096068f1",
                            "map_public_ip_on_launch": "true",
                            "tags.%": "1",
                            "tags.Name": "-public-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_vpc.this": {
                    "type": "aws_vpc",
                    "depends_on": [],
                    "primary": {
                        "id": "vpc-010e95f77f8c7f7ee",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:vpc/vpc-010e95f77f8c7f7ee",
                            "assign_generated_ipv6_cidr_block": "false",
                            "cidr_block": "10.0.0.0/16",
                            "default_network_acl_id": "acl-0f9342c12ab343237",
                            "default_route_table_id": "rtb-08cc5b75ea27736b1",
                            "default_security_group_id": "sg-0825371596c4eebbb",
                            "dhcp_options_id": "dopt-d710e5bc",
                            "enable_dns_hostnames": "false",
                            "enable_dns_support": "true",
                            "id": "vpc-010e95f77f8c7f7ee",
                            "instance_tenancy": "default",
                            "main_route_table_id": "rtb-08cc5b75ea27736b1",
                            "tags.%": "1",
                            "tags.Name": ""
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_ami.ecs_ami": {
                    "type": "aws_ami",
                    "depends_on": [],
                    "primary": {
                        "id": "ami-035a1bdaf0e4bf265",
                        "attributes": {
                            "architecture": "x86_64",
                            "block_device_mappings.#": "2",
                            "block_device_mappings.2538004115.device_name": "/dev/xvdcz",
                            "block_device_mappings.2538004115.ebs.%": "5",
                            "block_device_mappings.2538004115.ebs.delete_on_termination": "true",
                            "block_device_mappings.2538004115.ebs.encrypted": "false",
                            "block_device_mappings.2538004115.ebs.iops": "0",
                            "block_device_mappings.2538004115.ebs.volume_size": "22",
                            "block_device_mappings.2538004115.ebs.volume_type": "gp2",
                            "block_device_mappings.2538004115.no_device": "",
                            "block_device_mappings.2538004115.virtual_name": "",
                            "block_device_mappings.340275815.device_name": "/dev/xvda",
                            "block_device_mappings.340275815.ebs.%": "6",
                            "block_device_mappings.340275815.ebs.delete_on_termination": "true",
                            "block_device_mappings.340275815.ebs.encrypted": "false",
                            "block_device_mappings.340275815.ebs.iops": "0",
                            "block_device_mappings.340275815.ebs.snapshot_id": "snap-09a5a92a34c4cf8c3",
                            "block_device_mappings.340275815.ebs.volume_size": "8",
                            "block_device_mappings.340275815.ebs.volume_type": "gp2",
                            "block_device_mappings.340275815.no_device": "",
                            "block_device_mappings.340275815.virtual_name": "",
                            "creation_date": "2019-08-16T22:34:38.000Z",
                            "description": "Amazon Linux AMI 2018.03.w x86_64 ECS HVM GP2",
                            "filter.#": "1",
                            "filter.3350713981.name": "name",
                            "filter.3350713981.values.#": "1",
                            "filter.3350713981.values.0": "amzn-ami-*-amazon-ecs-optimized",
                            "hypervisor": "xen",
                            "id": "ami-035a1bdaf0e4bf265",
                            "image_id": "ami-035a1bdaf0e4bf265",
                            "image_location": "amazon/amzn-ami-2018.03.w-amazon-ecs-optimized",
                            "image_owner_alias": "amazon",
                            "image_type": "machine",
                            "most_recent": "true",
                            "name": "amzn-ami-2018.03.w-amazon-ecs-optimized",
                            "owner_id": "591542846629",
                            "owners.#": "1",
                            "owners.0": "amazon",
                            "product_codes.#": "0",
                            "public": "true",
                            "root_device_name": "/dev/xvda",
                            "root_device_type": "ebs",
                            "root_snapshot_id": "snap-09a5a92a34c4cf8c3",
                            "sriov_net_support": "simple",
                            "state": "available",
                            "state_reason.%": "2",
                            "state_reason.code": "UNSET",
                            "state_reason.message": "UNSET",
                            "tags.%": "0",
                            "virtualization_type": "hvm"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_ami.this": {
                    "type": "aws_ami",
                    "depends_on": [],
                    "primary": {
                        "id": "ami-00c03f7f7f2ec15c3",
                        "attributes": {
                            "architecture": "x86_64",
                            "block_device_mappings.#": "1",
                            "block_device_mappings.340275815.device_name": "/dev/xvda",
                            "block_device_mappings.340275815.ebs.%": "6",
                            "block_device_mappings.340275815.ebs.delete_on_termination": "true",
                            "block_device_mappings.340275815.ebs.encrypted": "false",
                            "block_device_mappings.340275815.ebs.iops": "0",
                            "block_device_mappings.340275815.ebs.snapshot_id": "snap-0e43f2c4403c9a35b",
                            "block_device_mappings.340275815.ebs.volume_size": "8",
                            "block_device_mappings.340275815.ebs.volume_type": "gp2",
                            "block_device_mappings.340275815.no_device": "",
                            "block_device_mappings.340275815.virtual_name": "",
                            "creation_date": "2019-08-30T07:06:00.000Z",
                            "description": "Amazon Linux 2 AMI 2.0.20190823.1 x86_64 HVM gp2",
                            "filter.#": "3",
                            "filter.2026626658.name": "name",
                            "filter.2026626658.values.#": "1",
                            "filter.2026626658.values.0": "amzn2-ami-hvm-*",
                            "filter.3386043752.name": "architecture",
                            "filter.3386043752.values.#": "1",
                            "filter.3386043752.values.0": "x86_64",
                            "filter.490168357.name": "virtualization-type",
                            "filter.490168357.values.#": "1",
                            "filter.490168357.values.0": "hvm",
                            "hypervisor": "xen",
                            "id": "ami-00c03f7f7f2ec15c3",
                            "image_id": "ami-00c03f7f7f2ec15c3",
                            "image_location": "amazon/amzn2-ami-hvm-2.0.20190823.1-x86_64-gp2",
                            "image_owner_alias": "amazon",
                            "image_type": "machine",
                            "most_recent": "true",
                            "name": "amzn2-ami-hvm-2.0.20190823.1-x86_64-gp2",
                            "owner_id": "137112412989",
                            "owners.#": "1",
                            "owners.0": "137112412989",
                            "product_codes.#": "0",
                            "public": "true",
                            "root_device_name": "/dev/xvda",
                            "root_device_type": "ebs",
                            "root_snapshot_id": "snap-0e43f2c4403c9a35b",
                            "sriov_net_support": "simple",
                            "state": "available",
                            "state_reason.%": "2",
                            "state_reason.code": "UNSET",
                            "state_reason.message": "UNSET",
                            "tags.%": "0",
                            "virtualization_type": "hvm"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_caller_identity.this": {
                    "type": "aws_caller_identity",
                    "depends_on": [],
                    "primary": {
                        "id": "2019-09-05 10:13:23.699585902 +0000 UTC",
                        "attributes": {
                            
                            "arn": "arn:aws:iam::051582052996:user/bkrzakala",
                            "id": "2019-09-05 10:13:23.699585902 +0000 UTC",
                            "user_id": "***REMOVED***"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.bastion": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "local.bastion_bucket",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "3317039311",
                        "attributes": {
                            "id": "3317039311",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:*\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowEC2\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ec2:*\",\n      \"Resource\": \"*\"\n    }\n  ]\n}",
                            "statement.#": "5",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "0",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowS3",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.1834123015": "s3:*",
                            "statement.1.condition.#": "0",
                            "statement.1.effect": "Allow",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "0",
                            "statement.1.resources.#": "2",
                            "statement.1.resources.2997558646": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.1.resources.94925287": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.1.sid": "AllowS3Bastion",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.1011197950": "kms:*",
                            "statement.2.condition.#": "0",
                            "statement.2.effect": "Allow",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "0",
                            "statement.2.resources.#": "1",
                            "statement.2.resources.3285651814": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "statement.2.sid": "AllowKMSAccess",
                            "statement.3.actions.#": "1",
                            "statement.3.actions.3112138991": "ecs:*",
                            "statement.3.condition.#": "0",
                            "statement.3.effect": "Allow",
                            "statement.3.not_actions.#": "0",
                            "statement.3.not_principals.#": "0",
                            "statement.3.not_resources.#": "0",
                            "statement.3.principals.#": "0",
                            "statement.3.resources.#": "1",
                            "statement.3.resources.2679715827": "*",
                            "statement.3.sid": "AllowECS",
                            "statement.4.actions.#": "1",
                            "statement.4.actions.2597799863": "ec2:*",
                            "statement.4.condition.#": "0",
                            "statement.4.effect": "Allow",
                            "statement.4.not_actions.#": "0",
                            "statement.4.not_principals.#": "0",
                            "statement.4.not_resources.#": "0",
                            "statement.4.principals.#": "0",
                            "statement.4.resources.#": "1",
                            "statement.4.resources.2679715827": "*",
                            "statement.4.sid": "AllowEC2"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.bucket_policy": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "data.aws_caller_identity.this",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "2088926263",
                        "attributes": {
                            "id": "2088926263",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"arn:aws:iam::051582052996:root\"\n      }\n    },\n    {\n      \"Sid\": \"DenyAccess1\",\n      \"Effect\": \"Deny\",\n      \"Action\": \"s3:PutObject\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"*\"\n      },\n      \"Condition\": {\n        \"Null\": {\n          \"s3:x-amz-server-side-encryption\": \"true\"\n        }\n      }\n    },\n    {\n      \"Sid\": \"DenyAccess2\",\n      \"Effect\": \"Deny\",\n      \"Action\": \"s3:PutObject\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"*\"\n      },\n      \"Condition\": {\n        \"StringNotEquals\": {\n          \"s3:x-amz-server-side-encryption\": \"aws:kms\"\n        }\n      }\n    }\n  ]\n}",
                            "statement.#": "3",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "1",
                            "statement.0.principals.3691871349.identifiers.#": "1",
                            "statement.0.principals.3691871349.identifiers.2401438501": "arn:aws:iam::051582052996:root",
                            "statement.0.principals.3691871349.type": "AWS",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowAccess",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.315547055": "s3:PutObject",
                            "statement.1.condition.#": "1",
                            "statement.1.condition.3673734150.test": "Null",
                            "statement.1.condition.3673734150.values.#": "1",
                            "statement.1.condition.3673734150.values.4043113848": "true",
                            "statement.1.condition.3673734150.variable": "s3:x-amz-server-side-encryption",
                            "statement.1.effect": "Deny",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "1",
                            "statement.1.principals.636693895.identifiers.#": "1",
                            "statement.1.principals.636693895.identifiers.2679715827": "*",
                            "statement.1.principals.636693895.type": "AWS",
                            "statement.1.resources.#": "2",
                            "statement.1.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.1.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.1.sid": "DenyAccess1",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.315547055": "s3:PutObject",
                            "statement.2.condition.#": "1",
                            "statement.2.condition.3059326814.test": "StringNotEquals",
                            "statement.2.condition.3059326814.values.#": "1",
                            "statement.2.condition.3059326814.values.800761281": "aws:kms",
                            "statement.2.condition.3059326814.variable": "s3:x-amz-server-side-encryption",
                            "statement.2.effect": "Deny",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "1",
                            "statement.2.principals.636693895.identifiers.#": "1",
                            "statement.2.principals.636693895.identifiers.2679715827": "*",
                            "statement.2.principals.636693895.type": "AWS",
                            "statement.2.resources.#": "2",
                            "statement.2.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.2.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.2.sid": "DenyAccess2"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.ecs_task": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "local.bastion_bucket",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "2507052418",
                        "attributes": {
                            "id": "2507052418",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3Access\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:DescribeTasks\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    }\n  ]\n}",
                            "statement.#": "4",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "0",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowS3Access",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.1011197950": "kms:*",
                            "statement.1.condition.#": "0",
                            "statement.1.effect": "Allow",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "0",
                            "statement.1.resources.#": "1",
                            "statement.1.resources.3285651814": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "statement.1.sid": "AllowKMSAccess",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.974674342": "ecs:DescribeTasks",
                            "statement.2.condition.#": "0",
                            "statement.2.effect": "Allow",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "0",
                            "statement.2.resources.#": "1",
                            "statement.2.resources.2679715827": "*",
                            "statement.2.sid": "AllowECS",
                            "statement.3.actions.#": "1",
                            "statement.3.actions.1834123015": "s3:*",
                            "statement.3.condition.#": "0",
                            "statement.3.effect": "Allow",
                            "statement.3.not_actions.#": "0",
                            "statement.3.not_principals.#": "0",
                            "statement.3.not_resources.#": "0",
                            "statement.3.principals.#": "0",
                            "statement.3.resources.#": "2",
                            "statement.3.resources.2997558646": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.3.resources.94925287": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.3.sid": "AllowS3Bastion"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.kms_policy": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "data.aws_caller_identity.this"
                    ],
                    "primary": {
                        "id": "1405806754",
                        "attributes": {
                            "id": "1405806754",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"*\",\n      \"Principal\": {\n        \"AWS\": \"arn:aws:iam::051582052996:root\"\n      }\n    }\n  ]\n}",
                            "statement.#": "1",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1011197950": "kms:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "1",
                            "statement.0.principals.3691871349.identifiers.#": "1",
                            "statement.0.principals.3691871349.identifiers.2401438501": "arn:aws:iam::051582052996:root",
                            "statement.0.principals.3691871349.type": "AWS",
                            "statement.0.resources.#": "1",
                            "statement.0.resources.2679715827": "*",
                            "statement.0.sid": "AllowAccess"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_security_group.default": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "aws_vpc.this"
                    ],
                    "primary": {
                        "id": "sg-0825371596c4eebbb",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-0825371596c4eebbb",
                            "description": "default VPC security group",
                            "id": "sg-0825371596c4eebbb",
                            "name": "default",
                            "tags.%": "0",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.template_file.user_data": {
                    "type": "template_file",
                    "depends_on": [
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "c08a5a7b7fca738108790bf3d0b997c3ccffebc143541469ec3785f8122811c8",
                        "attributes": {
                            "id": "c08a5a7b7fca738108790bf3d0b997c3ccffebc143541469ec3785f8122811c8",
                            "rendered": "    #!/bin/bash\n    echo ECS_CLUSTER=quorum-network-cocroaches-attack \u003e\u003e /etc/ecs/ecs.config\n\n    # node_exporter part\n    set -e\n    cd /tmp\n    curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz\n    tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz\n    mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/\n    useradd -rs /bin/false node_exporter\n\n\n    tee -a /etc/init.d/node_exporter \u003c\u003c END\n#!/bin/bash\n\n### BEGIN INIT INFO\n# processname:       node_exporter\n# Short-Description: Exporter for machine metrics.\n# Description:       Prometheus exporter for machine metrics,\n#                    written in Go with pluggable metric collectors.\n#\n# chkconfig: 2345 80 80\n# pidfile: /var/run/node_exporter/node_exporter.pid\n#\n#\n### END INIT INFO\n\n# Source function library.\n. /etc/init.d/functions\n\nNAME=node_exporter\nDESC=\"Exporter for machine metrics\"\nDAEMON=/usr/local/bin/node_exporter\nUSER=node_exporter\nCONFIG=\nPID=/var/run/node_exporter/\\$NAME.pid\nLOG=/var/log/node_exporter/\\$NAME.log\n\nDAEMON_OPTS=\nRETVAL=0\n\n# Check if DAEMON binary exist\n[ -f \\$DAEMON ] || exit 0\n\n[ -f /etc/default/node_exporter ]  \u0026\u0026  . /etc/default/node_exporter\n\nservice_checks() {\n  # Prepare directories\n  mkdir -p /var/run/node_exporter /var/log/node_exporter\n  chown -R \\$USER /var/run/node_exporter /var/log/node_exporter\n\n  # Check if PID exists\n  if [ -f \"\\$PID\" ]; then\n    PID_NUMBER=\\$(cat \\$PID)\n    if [ -z \"\\$(ps axf | grep \\$PID_NUMBER | grep -v grep)\" ]; then\n      echo \"Service was aborted abnormally; clean the PID file and continue...\"\n      rm -f \"\\$PID\"\n    else\n      echo \"Service already started; skip...\"\n      exit 1\n    fi\n  fi\n}\n\nstart() {\n  service_checks \\$1\n  sudo -H -u \\$USER   \\$DAEMON \\$DAEMON_OPTS  \u003e \\$LOG 2\u003e\u00261  \u0026\n  RETVAL=\\$?\n  echo \\$! \u003e \\$PID\n}\n\nstop() {\n  killproc -p \\$PID -b \\$DAEMON  \\$NAME\n  RETVAL=\\$?\n}\n\nreload() {\n  #-- sorry but node_exporter doesn't handle -HUP signal...\n  #killproc -p \\$PID -b \\$DAEMON  \\$NAME -HUP\n  #RETVAL=\\$?\n  stop\n  start\n}\n\ncase \"\\$1\" in\n  start)\n    echo -n \\$\"Starting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    start\n    ;;\n\n  stop)\n    echo -n \\$\"Stopping \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    ;;\n\n  reload)\n    echo -n \\$\"Reloading \\$DESC configuration -\" \"\\$NAME\" \\$'\\n'\n    reload\n    ;;\n\n  restart|force-reload)\n    echo -n \\$\"Restarting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    start\n    ;;\n\n  syntax)\n    \\$DAEMON --help\n    ;;\n\n  status)\n    status -p \\$PID \\$DAEMON\n    ;;\n\n  *)\n    echo -n \\$\"Usage: /etc/init.d/\\$NAME {start|stop|reload|restart|force-reload|syntax|status}\" \\$'\\n'\n    ;;\nesac\n\nexit \\$RETVAL\nEND\n\nchmod +x /etc/init.d/node_exporter\nservice node_exporter start\nchkconfig node_exporter on\n\n",
                            "template": "    #!/bin/bash\n    echo ECS_CLUSTER=quorum-network-cocroaches-attack \u003e\u003e /etc/ecs/ecs.config\n\n    # node_exporter part\n    set -e\n    cd /tmp\n    curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz\n    tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz\n    mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/\n    useradd -rs /bin/false node_exporter\n\n\n    tee -a /etc/init.d/node_exporter \u003c\u003c END\n#!/bin/bash\n\n### BEGIN INIT INFO\n# processname:       node_exporter\n# Short-Description: Exporter for machine metrics.\n# Description:       Prometheus exporter for machine metrics,\n#                    written in Go with pluggable metric collectors.\n#\n# chkconfig: 2345 80 80\n# pidfile: /var/run/node_exporter/node_exporter.pid\n#\n#\n### END INIT INFO\n\n# Source function library.\n. /etc/init.d/functions\n\nNAME=node_exporter\nDESC=\"Exporter for machine metrics\"\nDAEMON=/usr/local/bin/node_exporter\nUSER=node_exporter\nCONFIG=\nPID=/var/run/node_exporter/\\$NAME.pid\nLOG=/var/log/node_exporter/\\$NAME.log\n\nDAEMON_OPTS=\nRETVAL=0\n\n# Check if DAEMON binary exist\n[ -f \\$DAEMON ] || exit 0\n\n[ -f /etc/default/node_exporter ]  \u0026\u0026  . /etc/default/node_exporter\n\nservice_checks() {\n  # Prepare directories\n  mkdir -p /var/run/node_exporter /var/log/node_exporter\n  chown -R \\$USER /var/run/node_exporter /var/log/node_exporter\n\n  # Check if PID exists\n  if [ -f \"\\$PID\" ]; then\n    PID_NUMBER=\\$(cat \\$PID)\n    if [ -z \"\\$(ps axf | grep \\$PID_NUMBER | grep -v grep)\" ]; then\n      echo \"Service was aborted abnormally; clean the PID file and continue...\"\n      rm -f \"\\$PID\"\n    else\n      echo \"Service already started; skip...\"\n      exit 1\n    fi\n  fi\n}\n\nstart() {\n  service_checks \\$1\n  sudo -H -u \\$USER   \\$DAEMON \\$DAEMON_OPTS  \u003e \\$LOG 2\u003e\u00261  \u0026\n  RETVAL=\\$?\n  echo \\$! \u003e \\$PID\n}\n\nstop() {\n  killproc -p \\$PID -b \\$DAEMON  \\$NAME\n  RETVAL=\\$?\n}\n\nreload() {\n  #-- sorry but node_exporter doesn't handle -HUP signal...\n  #killproc -p \\$PID -b \\$DAEMON  \\$NAME -HUP\n  #RETVAL=\\$?\n  stop\n  start\n}\n\ncase \"\\$1\" in\n  start)\n    echo -n \\$\"Starting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    start\n    ;;\n\n  stop)\n    echo -n \\$\"Stopping \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    ;;\n\n  reload)\n    echo -n \\$\"Reloading \\$DESC configuration -\" \"\\$NAME\" \\$'\\n'\n    reload\n    ;;\n\n  restart|force-reload)\n    echo -n \\$\"Restarting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    start\n    ;;\n\n  syntax)\n    \\$DAEMON --help\n    ;;\n\n  status)\n    status -p \\$PID \\$DAEMON\n    ;;\n\n  *)\n    echo -n \\$\"Usage: /etc/init.d/\\$NAME {start|stop|reload|restart|force-reload|syntax|status}\" \\$'\\n'\n    ;;\nesac\n\nexit \\$RETVAL\nEND\n\nchmod +x /etc/init.d/node_exporter\nservice node_exporter start\nchkconfig node_exporter on\n\n",
                            "vars.%": "1",
                            "vars.ecs_cluster_name": "quorum-network-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.template"
                },
                "local_file.bootstrap": {
                    "type": "local_file",
                    "depends_on": [
                        "aws_ecs_service.quorum.*",
                        "aws_ecs_task_definition.quorum",
                        "local.bastion_bucket",
                        "local.ecs_cluster_name",
                        "local.hosts_folder",
                        "local.normalized_host_ip",
                        "local.privacy_addresses_folder",
                        "local.quorum_docker_image",
                        "local.quorum_rpc_port",
                        "local.quorum_run_container_name",
                        "local.s3_revision_folder",
                        "local.shared_volume_container_path",
                        "local.tessera_thirdparty_port",
                        "random_string.random"
                    ],
                    "primary": {
                        "id": "09bccb7dc18c01e3a1c96853e9a3c885cb472081",
                        "attributes": {
                            "content": "#!/bin/bash\n\nset -e\n\nexport AWS_DEFAULT_REGION=us-east-2\nexport TASK_REVISION=2\nsudo rm -rf /qdata\nsudo mkdir -p /qdata/mappings\nsudo mkdir -p /qdata/privacyaddresses\n\n# Faketime array ( ClockSkew )\nold_IFS=$IFS\nIFS=',' faketime=(1 -3 2)\nIFS=${old_IFS}\ncounter=\"${#faketime[@]}\"\n\nwhile [ $counter -gt 0 ]\ndo\n    echo -n \"${faketime[-1]}\" \u003e ./$counter\n    faketime=(${faketime[@]::$counter})\n    sudo aws s3 cp ./$counter s3://us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/clockSkew/\n    counter=$((counter - 1))\ndone\n\ncount=0\nwhile [ $count -lt 0 ]\ndo\n  count=$(ls /qdata/privacyaddresses | grep ^ip | wc -l)\n  sudo aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/ /qdata/ \u003e /dev/null 2\u003e\u00261 \\\n    | echo Wait for nodes in Quorum network being up ... $count/0\n  sleep 1\ndone\n\nif which jq \u003e/dev/null; then\n  echo \"Found jq\"\nelse\n  echo \"jq not found. Instaling ...\"\n  sudo yum -y install jq\nfi\n\nfor t in $(aws ecs list-tasks --cluster quorum-network-cocroaches-attack | jq -r .taskArns[])\ndo\n  task_metadata=$(aws ecs describe-tasks --cluster quorum-network-cocroaches-attack --tasks $t)\n  HOST_IP=$(echo $task_metadata | jq -r '.tasks[0] | .containers[] | select(.name == \"quorum-run\") | .networkInterfaces[] | .privateIpv4Address')\n  if [ \"EC2\" == \"EC2\" ]\n  then\n    CONTAINER_INSTANCE_ARN=$(aws ecs describe-tasks --tasks $t --cluster quorum-network-cocroaches-attack | jq -r '.tasks[] | .containerInstanceArn')\n    EC2_INSTANCE_ID=$(aws ecs  describe-container-instances --container-instances $CONTAINER_INSTANCE_ARN --cluster quorum-network-cocroaches-attack |jq -r '.containerInstances[] | .ec2InstanceId')\n    HOST_IP=$(aws ec2 describe-instances --instance-ids $EC2_INSTANCE_ID | jq -r '.Reservations[0] | .Instances[] | .PublicIpAddress')\n  fi\n  group=$(echo $task_metadata | jq -r '.tasks[0] | .group')\n  taskArn=$(echo $task_metadata | jq -r '.tasks[0] | .taskDefinitionArn')\n  # only care about new task\n  if [[ \"$taskArn\" == *:$TASK_REVISION ]]; then\n     echo $group | sudo tee /qdata/mappings/ip_$(echo $HOST_IP | sed -e 's/\\./_/g')\n  fi\ndone\n\ncat \u003c\u003cSS | sudo tee /qdata/quorum_metadata\nquorum:\n  nodes:\nSS\nnodes=()\ncd /qdata/mappings\nfor idx in \"${!nodes[@]}\"\ndo\n  f=$(grep -l ${nodes[$idx]} *)\n  ip=$(cat /qdata/hosts/$f)\n  nodeIdx=$((idx+1))\n  script=\"/usr/local/bin/Node$nodeIdx\"\n  cat \u003c\u003cSS | sudo tee $script\n#!/bin/bash\n\nsudo docker run --rm -it quorumengineering/quorum:latest attach http://$ip:22000 $@\nSS\n  sudo chmod +x $script\n  cat \u003c\u003cSS | sudo tee -a /qdata/quorum_metadata\n    Node$nodeIdx:\n      privacy-address: $(cat /qdata/privacyaddresses/$f)\n      url: http://$ip:22000\n      third-party-url: http://$ip:9080\nSS\ndone\n\ncat \u003c\u003cSS | sudo tee /opt/prometheus/prometheus.yml\nglobal:\n  scrape_interval:     15s # By default, scrape targets every 15 seconds.\n\n  # Attach these labels to any time series or alerts when communicating with\n  # external systems (federation, remote storage, Alertmanager).\n  external_labels:\n    monitor: 'monitor'\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it's Prometheus itself.\nscrape_configs:\n- job_name: geth\n  metrics_path: /debug/metrics/prometheus\n  scheme: http\n  static_configs:\n  - targets:\n    - geth:6060\n- job_name: 'node'\n  static_configs:\n  - targets: [ node-exporter:9100 ]\n  file_sd_configs:\n  - files:\n    - 'targets.json'\nSS\n\ncat \u003c\u003cSS | sudo tee /opt/prometheus/docker-compose.yml\n# docker-compose.yml\nversion: '2'\nservices:\n    prometheus:\n        image: prom/prometheus:latest\n        volumes:\n            - /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml\n            - /opt/prometheus/targets.json:/etc/prometheus/targets.json\n        command:\n            - '--config.file=/etc/prometheus/prometheus.yml'\n        ports:\n            - '9090:9090'\n    node-exporter:\n        image: prom/node-exporter:latest\n        ports:\n            - '9100:9100'\n    grafana:\n        image: grafana/grafana:latest\n        volumes:\n            - /opt/grafana/dashboards:/var/lib/grafana/dashboards\n            - /opt/grafana/provisioning/dashboards/all.yml:/etc/grafana/provisioning/dashboards/all.yml\n            - /opt/grafana/provisioning/datasources/all.yml:/etc/grafana/provisioning/datasources/all.yml\n        environment:\n            - GF_SECURITY_ADMIN_PASSWORD=-\u003cND{!FA)tLFDoGB\n        depends_on:\n            - prometheus\n        ports:\n            - '3001:3000'\n    geth:\n        image: ethereum/client-go:latest\n        ports:\n            - '6060:6060'\n        command: --goerli --metrics --metrics.expensive --pprof --pprofaddr=0.0.0.0\n\nSS\n\ncount=$(ls /qdata/privacyaddresses | grep ^ip | wc -l)\ntarget_file=/tmp/targets.json\ni=0\necho '[' \u003e $target_file\nfor idx in \"${!nodes[@]}\"\ndo\n  f=$(grep -l ${nodes[$idx]} *)\n  ip=$(cat /qdata/hosts/$f)\n  i=$(($i+1))\n  if [ $i -lt \"$count\" ]; then\n    echo '{ \"targets\": [\"'$ip':9100\"] },' \u003e\u003e $target_file\n  else\n    echo '{ \"targets\": [\"'$ip':9100\"] }'  \u003e\u003e $target_file\n  fi\ndone\necho ']' \u003e\u003e $target_file\nsudo mv $target_file /opt/prometheus/\n\ncat \u003c\u003cSS | sudo tee /opt/grafana/provisioning/datasources/all.yml\ndatasources:\n- name: 'prometheus'\n  type: 'prometheus'\n  access: 'proxy'\n  org_id: 1\n  url: 'http://prometheus:9090'\n  is_default: true\n  version: 1\n  editable: true\nSS\n\ncat \u003c\u003cSS | sudo tee /opt/grafana/provisioning/dashboards/all.yml\n- name: 'default'\n  org_id: 1\n  folder: ''\n  type: 'file'\n  options:\n    folder: '/var/lib/grafana/dashboards'\nSS\n\nsudo sed -i s'/datasource\":.*/datasource\" :\"prometheus\",/' /opt/grafana/dashboards/dashboard-geth.json\nsudo sed -i s'/datasource\":.*/datasource\" :\"prometheus\",/' /opt/grafana/dashboards/dashboard-node-exporter.json\nsudo /usr/local/bin/docker-compose -f /opt/prometheus/docker-compose.yml up -d --force-recreate\n",
                            "filename": "/tmp/.terranova273240257/generated-bootstrap.sh",
                            "id": "09bccb7dc18c01e3a1c96853e9a3c885cb472081"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.local"
                },
                "local_file.private_key": {
                    "type": "local_file",
                    "depends_on": [
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "0619bff6918e2e978bb10f0b57b549297a9eaad0",
                        "attributes": {
                            "content": "---DUMMY PRIVATE KEY---",
                            "filename": "/tmp/.terranova273240257/quorum-cocroaches-attack.pem",
                            "id": "0619bff6918e2e978bb10f0b57b549297a9eaad0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.local"
                },
                "null_resource.bastion_remote_exec": {
                    "type": "null_resource",
                    "depends_on": [
                        "aws_ecs_task_definition.quorum",
                        "aws_instance.bastion",
                        "local_file.bootstrap",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "354840277283586537",
                        "attributes": {
                            "id": "354840277283586537",
                            "triggers.%": "3",
                            "triggers.bastion": "",
                            "triggers.ecs_task_definition": "2",
                            "triggers.script": "3972838542e1726b0d045ff3d42cc764"
                        },
                        "meta": {},
                        "tainted": true
                    },
                    "deposed": [],
                    "provider": "provider.null"
                },
                "random_id.bucket_postfix": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "6Ma2xnZmArE",
                        "attributes": {
                            "b64": "6Ma2xnZmArE",
                            "b64_std": "6Ma2xnZmArE=",
                            "b64_url": "6Ma2xnZmArE",
                            "byte_length": "8",
                            "dec": "16773294825694167729",
                            "hex": "e8c6b6c6766602b1",
                            "id": "6Ma2xnZmArE"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_id.code": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "3M--hw",
                        "attributes": {
                            "b64": "3M--hw",
                            "b64_std": "3M++hw==",
                            "b64_url": "3M--hw",
                            "byte_length": "4",
                            "dec": "3704602247",
                            "hex": "dccfbe87",
                            "id": "3M--hw"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_id.ethstat_secret": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "tEAgROn3tc3oGqT3CgVSqA",
                        "attributes": {
                            "b64": "tEAgROn3tc3oGqT3CgVSqA",
                            "b64_std": "tEAgROn3tc3oGqT3CgVSqA==",
                            "b64_url": "tEAgROn3tc3oGqT3CgVSqA",
                            "byte_length": "16",
                            "dec": "239594000737262924426798755770955813544",
                            "hex": "b4402044e9f7b5cde81aa4f70a0552a8",
                            "id": "tEAgROn3tc3oGqT3CgVSqA"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_integer.network_id": {
                    "type": "random_integer",
                    "depends_on": [],
                    "primary": {
                        "id": "4856",
                        "attributes": {
                            "id": "4856",
                            "keepers.%": "1",
                            "keepers.changes_when": "cocroaches-attack",
                            "max": "9999",
                            "min": "2018",
                            "result": "4856"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_string.random": {
                    "type": "random_string",
                    "depends_on": [],
                    "primary": {
                        "id": "none",
                        "attributes": {
                            "id": "none",
                            "length": "16",
                            "lower": "true",
                            "min_lower": "0",
                            "min_numeric": "0",
                            "min_special": "0",
                            "min_upper": "0",
                            "number": "true",
                            "result": "-\u003cND{!FA)tLFDoGB",
                            "special": "true",
                            "upper": "true"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "tls_private_key.ssh": {
                    "type": "tls_private_key",
                    "depends_on": [],
                    "primary": {
                        "id": "e28f0d026fcd4faea3dd1ae386c7f918484f1273",
                        "attributes": {
                            "algorithm": "RSA",
                            "ecdsa_curve": "P224",
                            "id": "e28f0d026fcd4faea3dd1ae386c7f918484f1273",
                            "private_key_pem": "---DUMMY PRIVATE KEY---",
                            "public_key_fingerprint_md5": "44:43:cb:61:29:71:33:19:21:23:d3:69:3a:b2:3e:00",
                            "public_key_openssh": "dummyKeySsh",
                            "public_key_pem": "---DUMMY PUBLIC KEY---",
                            "rsa_bits": "2048"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.tls"
                }
            },
            "depends_on": []
        }
    ]
}`
	MultipleValuesOutputFixture = `{
    "version": 3,
    "terraform_version": "0.11.13",
    "serial": 3,
    "lineage": "56a6b1ae-aa91-ccb9-bb23-cfd3ab19176f",
    "modules": [
        {
            "path": [
                "root"
            ],
            "outputs": {
                "_status": {
                    "sensitive": false,
                    "type": "string",
                    "value": "Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n"
                },
                "bastion_host_dns": {
                    "sensitive": false,
                    "type": "array",
                    "value": ["", ""]
                },
                "bastion_host_ip": {
                    "sensitive": false,
                    "type": "map",
                    "value": {"ip": "invalid.ip.666"}
                },
                "bucket_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1"
                },
                "chain_id": {
                    "sensitive": false,
                    "type": "string",
                    "value": "4856"
                },
                "ecs_cluster_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "quorum-network-cocroaches-attack"
                },
                "grafana_host_url": {
                    "sensitive": false,
                    "type": "string",
                    "value": "http://invalid.ip.666:3001"
                },
                "grafana_password": {
                    "sensitive": false,
                    "type": "string",
                    "value": "-\u003cND{!FA)tLFDoGB"
                },
                "network_name": {
                    "sensitive": false,
                    "type": "string",
                    "value": "cocroaches-attack"
                },
                "private_key_file": {
                    "sensitive": false,
                    "type": "string",
                    "value": "/tmp/.terranova273240257/quorum-cocroaches-attack.pem"
                }
            },
            "resources": {
                "aws_autoscaling_group.asg": {
                    "type": "aws_autoscaling_group",
                    "depends_on": [
                        "aws_launch_configuration.lc",
                        "aws_subnet.public.*",
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "quorum-network-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:autoscaling:us-east-2:051582052996:autoScalingGroup:30919fe0-83af-4d9f-b570-8399e4c432eb:autoScalingGroupName/quorum-network-cocroaches-attack",
                            "availability_zones.#": "1",
                            "availability_zones.4293815384": "us-east-2a",
                            "default_cooldown": "30",
                            "desired_capacity": "0",
                            "enabled_metrics.#": "0",
                            "force_delete": "false",
                            "health_check_grace_period": "300",
                            "health_check_type": "EC2",
                            "id": "quorum-network-cocroaches-attack",
                            "launch_configuration": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "launch_template.#": "0",
                            "load_balancers.#": "0",
                            "max_size": "0",
                            "metrics_granularity": "1Minute",
                            "min_size": "0",
                            "name": "quorum-network-cocroaches-attack",
                            "placement_group": "",
                            "protect_from_scale_in": "false",
                            "service_linked_role_arn": "arn:aws:iam::051582052996:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling",
                            "suspended_processes.#": "0",
                            "tags.#": "2",
                            "tags.0.%": "3",
                            "tags.0.key": "ecs_cluster",
                            "tags.0.propagate_at_launch": "1",
                            "tags.0.value": "quorum-network-cocroaches-attack",
                            "tags.1.%": "3",
                            "tags.1.key": "created_by",
                            "tags.1.propagate_at_launch": "1",
                            "tags.1.value": "terraform",
                            "target_group_arns.#": "0",
                            "termination_policies.#": "1",
                            "termination_policies.0": "Default",
                            "vpc_zone_identifier.#": "1",
                            "vpc_zone_identifier.2389729966": "subnet-035251cd0096068f1",
                            "wait_for_capacity_timeout": "10m"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "delete": 600000000000
                            }
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_cloudwatch_log_group.quorum": {
                    "type": "aws_cloudwatch_log_group",
                    "depends_on": [
                        "local.common_tags"
                    ],
                    "primary": {
                        "id": "/ecs/quorum/cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:logs:us-east-2:051582052996:log-group:/ecs/quorum/cocroaches-attack:*",
                            "id": "/ecs/quorum/cocroaches-attack",
                            "kms_key_id": "",
                            "name": "/ecs/quorum/cocroaches-attack",
                            "retention_in_days": "7",
                            "tags.%": "4",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_ecs_cluster.quorum": {
                    "type": "aws_ecs_cluster",
                    "depends_on": [
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                            "id": "arn:aws:ecs:us-east-2:051582052996:cluster/quorum-network-cocroaches-attack",
                            "name": "quorum-network-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_ecs_task_definition.quorum": {
                    "type": "aws_ecs_task_definition",
                    "depends_on": [
                        "aws_iam_role.ecs_task",
                        "local.container_definitions",
                        "local.shared_volume_name"
                    ],
                    "primary": {
                        "id": "quorum-istanbul-tessera-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:ecs:us-east-2:051582052996:task-definition/quorum-istanbul-tessera-cocroaches-attack:2",
                            "container_definitions": "[{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"mkdir -p /qdata/dd/geth\\necho \\\"\\\" \u003e /qdata/passwords.txt\\nbootnode -genkey /qdata/dd/geth/nodekey\\nexport NODE_ID=$(bootnode -nodekey /qdata/dd/geth/nodekey -writeaddress)\\necho Creating an account for this node\\ngeth --datadir /qdata/dd account new --password /qdata/passwords.txt\\nexport KEYSTORE_FILE=$(ls /qdata/dd/keystore/ | head -n1)\\nexport ACCOUNT_ADDRESS=$(cat /qdata/dd/keystore/$KEYSTORE_FILE | sed 's/^.*\\\"address\\\":\\\"\\\\([^\\\"]*\\\\)\\\".*$/\\\\1/g')\\necho Writing account address $ACCOUNT_ADDRESS to /qdata/first_account_address\\necho $ACCOUNT_ADDRESS \u003e /qdata/first_account_address\\necho Writing Node Id [$NODE_ID] to /qdata/node_id\\necho $NODE_ID \u003e /qdata/node_id\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -f /qdata/node_id ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/quorum:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"node-key-bootstrap\",\"portMappings\":[],\"volumesFrom\":[]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until Node Key is ready ...\\nwhile [ ! -f \\\"/qdata/node_id\\\" ]; do sleep 1; done\\napk update\\napk add curl jq\\nexport TASK_REVISION=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq '.Revision' -r)\\necho \\\"Task Revision: $TASK_REVISION\\\"\\necho $TASK_REVISION \u003e /qdata/task_revision\\nexport HOST_IP=$(/usr/bin/curl http://169.254.169.254/latest/meta-data/public-ipv4)\\necho \\\"Host IP: $HOST_IP\\\"\\necho $HOST_IP \u003e /qdata/host_ip\\nexport TASK_ARN=$(curl -s $ECS_CONTAINER_METADATA_URI/task | jq -r '.TaskARN')\\nexport REGION=$(echo $TASK_ARN | awk -F: '{ print $4}')\\naws ecs describe-tasks --region $REGION --cluster quorum-network-cocroaches-attack --tasks $TASK_ARN | jq -r '.tasks[0] | .group' \u003e /qdata/service\\nmkdir -p /qdata/hosts\\nmkdir -p /qdata/nodeids\\nmkdir -p /qdata/accounts\\nmkdir -p /qdata/lib\\naws s3 cp s3://us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/libs/libfaketime.so /qdata/lib/libfaketime.so\\naws s3 cp /qdata/node_id s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/nodeids/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\naws s3 cp /qdata/host_ip s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/hosts/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\naws s3 cp /qdata/first_account_address s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/accounts/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/hosts | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/hosts /qdata/hosts \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other containers to report their IPs ... $count/0\\\"; sleep 1; done\\necho \\\"All containers have reported their IPs\\\"\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/accounts | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/accounts /qdata/accounts \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other nodes to report their accounts ... $count/0\\\"; sleep 1; done\\necho \\\"All nodes have registered accounts\\\"\\ncount=0; while [ $count -lt 0 ]; do count=$(ls /qdata/nodeids | grep ^ip | wc -l); aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/nodeids /qdata/nodeids \u003e /dev/null 2\u003e\u00261 | echo \\\"Wait for other nodes to report their IDs ... $count/0\\\"; sleep 1; done\\necho \\\"All nodes have registered their IDs\\\"\\nalloc=\\\"\\\"; for f in $(ls /qdata/accounts); do address=$(cat /qdata/accounts/$f); alloc=\\\"$alloc,\\\\\\\"$address\\\\\\\": { \\\"balance\\\": \\\"\\\\\\\"1000000000000000000000000000\\\\\\\"\\\"}\\\"; done\\nalloc=\\\"{${alloc:1}}\\\"\\nextraData=\\\"\\\\\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\\\\\"\\\"\\napk add --repository http://dl-cdn.alpinelinux.org/alpine/v3.7/community go=1.9.4-r0\\napk add git gcc musl-dev linux-headers\\ngit clone https://github.com/getamis/istanbul-tools /istanbul-tools/src/github.com/getamis/istanbul-tools\\nexport GOPATH=/istanbul-tools\\nexport GOROOT=/usr/lib/go\\necho 'package main\\n\\nimport (\\n\\t\\\"encoding/hex\\\"\\n\\t\\\"fmt\\\"\\n\\t\\\"os\\\"\\n\\n\\t\\\"github.com/ethereum/go-ethereum/crypto\\\"\\n\\t\\\"github.com/ethereum/go-ethereum/p2p/discover\\\"\\n)\\n\\nfunc main() {\\n\\tif len(os.Args) \u003c 2 {\\n\\t\\tfmt.Println(\\\"missing enode value\\\")\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tenode := os.Args[1]\\n\\tnodeId, err := discover.HexID(enode)\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tpub, err := nodeId.Pubkey()\\n\\tif err != nil {\\n\\t\\tfmt.Println(err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tfmt.Printf(\\\"0x%s\\\\n\\\", hex.EncodeToString(crypto.PubkeyToAddress(*pub).Bytes()))\\n}\\n' \u003e /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go\\nall=\\\"\\\"; for f in $(ls /qdata/nodeids); do address=$(cat /qdata/nodeids/$f); all=\\\"$all,$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/extra.go $address)\\\"; done\\nall=\\\"${all:1}\\\"\\necho Validator Addresses: $all\\nextraData=\\\"\\\\\\\"$(go run /istanbul-tools/src/github.com/getamis/istanbul-tools/cmd/istanbul/main.go extra encode --validators $all | awk -F: '{print $2}' | tr -d ' ')\\\\\\\"\\\"\\nmixHash=\\\"\\\\\\\"0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365\\\\\\\"\\\"\\ndifficulty=\\\"\\\\\\\"0x01\\\\\\\"\\\"\\necho '{\\\"alloc\\\":{},\\\"coinbase\\\":\\\"0x0000000000000000000000000000000000000000\\\",\\\"config\\\":{\\\"byzantiumBlock\\\":1,\\\"chainId\\\":4856,\\\"eip150Block\\\":1,\\\"eip150Hash\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"eip155Block\\\":0,\\\"eip158Block\\\":1,\\\"homesteadBlock\\\":0,\\\"isQuorum\\\":true},\\\"difficulty\\\":\\\"0x0\\\",\\\"extraData\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"gasLimit\\\":\\\"0xE0000000\\\",\\\"mixHash\\\":\\\"0x00000000000000000000000000000000000000647572616c65787365646c6578\\\",\\\"nonce\\\":\\\"0x0\\\",\\\"parentHash\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"timestamp\\\":\\\"0x00\\\"}' | jq \\\". + { alloc : $alloc, extraData: $extraData, mixHash: $mixHash, difficulty: $difficulty} | .config=.config + {istanbul: {epoch: 30000, policy: 0} }\\\" \u003e /qdata/genesis.json\\ncat /qdata/genesis.json\\necho \\\"Done!\\\" \u003e /qdata/metadata_bootstrap_container_status\\necho Wait until privacy engine initialized ...\\nwhile [ ! -f \\\"/qdata/.pub\\\" ]; do sleep 1; done\\naws s3 cp /qdata/.pub s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/privacyaddresses/ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g') --sse aws:kms --sse-kms-key-id arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -f /qdata/metadata_bootstrap_container_status ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"senseyeio/alpine-aws-cli:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"metamain-bootstrap\",\"portMappings\":[],\"volumesFrom\":[{\"sourceContainer\":\"node-key-bootstrap\"}]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until metadata bootstrap completed ...\\nwhile [ ! -f \\\"/qdata/metadata_bootstrap_container_status\\\" ]; do sleep 1; done\\necho Wait until tessera is ready ...\\nwhile [ ! -S \\\"/qdata/tm.ipc\\\" ]; do sleep 1; done\\nmkdir -p /qdata/dd/geth\\necho \\\"\\\" \u003e /qdata/passwords.txt\\necho \\\"Creating /qdata/dd/static-nodes.json and /qdata/dd/permissioned-nodes.json\\\"\\nall=\\\"\\\"; for f in $(ls /qdata/nodeids); do nodeid=$(cat /qdata/nodeids/$f); ip=$(cat /qdata/hosts/$f); all=\\\"$all,\\\\\\\"enode://$nodeid@$ip:21000?discport=0\u0026\\\\\\\"\\\"; done; all=${all:1}\\necho \\\"[$all]\\\" \u003e /qdata/dd/static-nodes.json\\necho \\\"[$all]\\\" \u003e /qdata/dd/permissioned-nodes.json\\necho Permissioned Nodes: $(cat /qdata/dd/permissioned-nodes.json)\\ngeth --datadir /qdata/dd init /qdata/genesis.json\\nexport IDENTITY=$(cat /qdata/service | awk -F: '{print $2}')\\necho 'Running geth with: --datadir /qdata/dd --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --rpcport 22000 --port 21000 --unlock 0 --password /qdata/passwords.txt --nodiscover --networkid 4856 --verbosity 5 --debug --identity $IDENTITY --ethstats \\\"$IDENTITY:b4402044e9f7b5cde81aa4f70a0552a8@10.0.0.239:3000\\\" --istanbul.blockperiod 1 --emitcheckpoints --syncmode full --mine --minerthreads 1'\\ngeth --datadir /qdata/dd --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --rpcport 22000 --port 21000 --unlock 0 --password /qdata/passwords.txt --nodiscover --networkid 4856 --verbosity 5 --debug --identity $IDENTITY --ethstats \\\"$IDENTITY:b4402044e9f7b5cde81aa4f70a0552a8@10.0.0.239:3000\\\" --istanbul.blockperiod 1 --emitcheckpoints --syncmode full --mine --minerthreads 1\"],\"environment\":[{\"name\":\"LD_PRELOAD\",\"value\":\"/qdata/lib/libfaketime.so\"},{\"name\":\"PRIVATE_CONFIG\",\"value\":\"/qdata/tm.ipc\"}],\"essential\":true,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -S /qdata/dd/geth.ipc ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/quorum:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"quorum-run\",\"portMappings\":[{\"containerPort\":22000,\"hostPort\":22000,\"protocol\":\"tcp\"},{\"containerPort\":21000,\"hostPort\":21000,\"protocol\":\"tcp\"},{\"containerPort\":50400,\"hostPort\":50400,\"protocol\":\"tcp\"}],\"volumesFrom\":[{\"sourceContainer\":\"metamain-bootstrap\"},{\"sourceContainer\":\"tessera-run\"}]},{\"cpu\":0,\"dockerLabels\":{\"DockerImage.PrivacyEngine\":\"quorumengineering/tessera:latest\",\"DockerImage.Quorum\":\"quorumengineering/quorum:latest\",\"ECSClusterName\":\"quorum-network-cocroaches-attack\",\"NetworkName\":\"cocroaches-attack\"},\"entryPoint\":[\"/bin/sh\",\"-c\",\"set -e\\necho Wait until metadata bootstrap completed ...\\nwhile [ ! -f \\\"/qdata/metadata_bootstrap_container_status\\\" ]; do sleep 1; done\\napk update\\napk add jq\\ncd /qdata; echo \\\"\\n\\\" | java -jar /tessera/tessera-app.jar -keygen /qdata/\\nexport HOST_IP=$(cat /qdata/host_ip)\\nexport TM_PUB=$(cat /qdata/.pub)\\nexport TM_KEY=$(cat /qdata/.key)\\necho \\\"\\nHost IP: $HOST_IP\\\"\\necho \\\"Public Key: $TM_PUB\\\"\\nall=\\\"\\\"; for f in $(ls /qdata/hosts | grep -v ip_$(echo $HOST_IP | sed -e 's/\\\\./_/g')); do ip=$(cat /qdata/hosts/$f); all=\\\"$all,{ \\\\\\\"url\\\\\\\": \\\\\\\"http://$ip:9000/\\\\\\\" }\\\"; done\\nall=\\\"[{ \\\\\\\"url\\\\\\\": \\\\\\\"http://$HOST_IP:9000/\\\\\\\" }$all]\\\"\\nexport TESSERA_VERSION=latest\\nexport V=$(echo -e \\\"0.8\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\necho \\\"Creating /qdata/tessera.cfg\\\"\\nDDIR=/qdata/dd\\nunzip -p /tessera/tessera-app.jar META-INF/MANIFEST.MF | grep Tessera-Version | cut -d: -f2 | xargs\\necho \\\"Tessera Version: $TESSERA_VERSION\\\"\\nV08=$(echo -e \\\"0.8\\\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\nV09=$(echo -e \\\"0.9\\\\n$TESSERA_VERSION\\\" | sort -n -r -t '.' -k 1,1 -k 2,2 | head -n1)\\ncase \\\"$TESSERA_VERSION\\\" in\\n    \\\"$V09\\\"|latest)\\n    # use new config\\n    cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n  \\\"useWhiteList\\\": false,\\n  \\\"jdbc\\\": {\\n    \\\"username\\\": \\\"sa\\\",\\n    \\\"password\\\": \\\"\\\",\\n    \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n    \\\"autoCreateTables\\\": true\\n  },\\n  \\\"serverConfigs\\\":[\\n  {\\n    \\\"app\\\":\\\"ThirdParty\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"http://$HOST_IP:9080\\\",\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"Q2T\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"unix:/qdata/tm.ipc\\\",\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"P2P\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverAddress\\\": \\\"http://$HOST_IP:9000\\\",\\n    \\\"sslConfig\\\": {\\n      \\\"tls\\\": \\\"OFF\\\",\\n      \\\"generateKeyStoreIfNotExisted\\\": true,\\n      \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n      \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n      \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n      \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n      \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n      \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  }\\n  ],\\n  \\\"peer\\\": $all,\\n  \\\"keys\\\": {\\n    \\\"passwords\\\": [],\\n    \\\"keyData\\\": [\\n      {\\n        \\\"config\\\": $TM_KEY,\\n        \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n      }\\n    ]\\n  },\\n  \\\"alwaysSendTo\\\": []\\n}\\nEOF    \\n      ;;\\n    \\\"$V08\\\")\\n      # use enhanced config\\n      cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n  \\\"useWhiteList\\\": false,\\n  \\\"jdbc\\\": {\\n    \\\"username\\\": \\\"sa\\\",\\n    \\\"password\\\": \\\"\\\",\\n    \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n    \\\"autoCreateTables\\\": true\\n  },\\n  \\\"serverConfigs\\\":[\\n  {\\n    \\\"app\\\":\\\"ThirdParty\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"INET\\\",\\n      \\\"port\\\": 9080,\\n      \\\"hostName\\\": \\\"http://$HOST_IP\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"Q2T\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"UNIX\\\",\\n      \\\"path\\\":\\\"/qdata/tm.ipc\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"UNIX_SOCKET\\\"\\n  },\\n  {\\n    \\\"app\\\":\\\"P2P\\\",\\n    \\\"enabled\\\": true,\\n    \\\"serverSocket\\\":{\\n      \\\"type\\\":\\\"INET\\\",\\n      \\\"port\\\": 9000,\\n      \\\"hostName\\\": \\\"http://$HOST_IP\\\"\\n    },\\n    \\\"sslConfig\\\": {\\n      \\\"tls\\\": \\\"OFF\\\",\\n      \\\"generateKeyStoreIfNotExisted\\\": true,\\n      \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n      \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n      \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n      \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n      \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n      \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n      \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n      \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n    },\\n    \\\"communicationType\\\" : \\\"REST\\\"\\n  }\\n  ],\\n  \\\"peer\\\": $all,\\n  \\\"keys\\\": {\\n    \\\"passwords\\\": [],\\n    \\\"keyData\\\": [\\n      {\\n        \\\"config\\\": $TM_KEY,\\n        \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n      }\\n    ]\\n  },\\n  \\\"alwaysSendTo\\\": []\\n}\\nEOF\\n      ;;\\n    *)\\n    # use old config\\n    cat \u003c\u003cEOF \u003e /qdata/tessera.cfg\\n{\\n    \\\"useWhiteList\\\": false,\\n    \\\"jdbc\\\": {\\n        \\\"username\\\": \\\"sa\\\",\\n        \\\"password\\\": \\\"\\\",\\n        \\\"url\\\": \\\"jdbc:h2:./${DDIR}/db;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0\\\",\\n        \\\"autoCreateTables\\\": true\\n    },\\n    \\\"server\\\": {\\n        \\\"port\\\": 9000,\\n        \\\"hostName\\\": \\\"http://$HOST_IP\\\",\\n        \\\"sslConfig\\\": {\\n            \\\"tls\\\": \\\"OFF\\\",\\n            \\\"generateKeyStoreIfNotExisted\\\": true,\\n            \\\"serverKeyStore\\\": \\\"${DDIR}/server-keystore\\\",\\n            \\\"serverKeyStorePassword\\\": \\\"quorum\\\",\\n            \\\"serverTrustStore\\\": \\\"${DDIR}/server-truststore\\\",\\n            \\\"serverTrustStorePassword\\\": \\\"quorum\\\",\\n            \\\"serverTrustMode\\\": \\\"TOFU\\\",\\n            \\\"knownClientsFile\\\": \\\"${DDIR}/knownClients\\\",\\n            \\\"clientKeyStore\\\": \\\"${DDIR}/client-keystore\\\",\\n            \\\"clientKeyStorePassword\\\": \\\"quorum\\\",\\n            \\\"clientTrustStore\\\": \\\"${DDIR}/client-truststore\\\",\\n            \\\"clientTrustStorePassword\\\": \\\"quorum\\\",\\n            \\\"clientTrustMode\\\": \\\"TOFU\\\",\\n            \\\"knownServersFile\\\": \\\"${DDIR}/knownServers\\\"\\n        }\\n    },\\n    \\\"peer\\\": $all,\\n    \\\"keys\\\": {\\n        \\\"passwords\\\": [],\\n        \\\"keyData\\\": [\\n            {\\n                \\\"config\\\": $TM_KEY,\\n                \\\"publicKey\\\": \\\"$TM_PUB\\\"\\n            }\\n        ]\\n    },\\n    \\\"alwaysSendTo\\\": [],\\n    \\\"unixSocketFile\\\": \\\"/qdata/tm.ipc\\\"\\n}\\nEOF\\n      ;;\\nesac\\ncat /qdata/tessera.cfg\\n\\njava -jar /tessera/tessera-app.jar -configfile /qdata/tessera.cfg\"],\"environment\":[],\"essential\":false,\"healthCheck\":{\"command\":[\"CMD-SHELL\",\"[ -S /qdata/tm.ipc ];\"],\"interval\":30,\"retries\":10,\"startPeriod\":300,\"timeout\":60},\"image\":\"quorumengineering/tessera:latest\",\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"/ecs/quorum/cocroaches-attack\",\"awslogs-region\":\"us-east-2\",\"awslogs-stream-prefix\":\"logs\"}},\"mountPoints\":[{\"containerPath\":\"/qdata\",\"sourceVolume\":\"quorum_shared_volume\"}],\"name\":\"tessera-run\",\"portMappings\":[{\"containerPort\":9000,\"hostPort\":9000,\"protocol\":\"tcp\"},{\"containerPort\":9080,\"hostPort\":9080,\"protocol\":\"tcp\"}],\"volumesFrom\":[{\"sourceContainer\":\"metamain-bootstrap\"}]}]",
                            "cpu": "4096",
                            "execution_role_arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "family": "quorum-istanbul-tessera-cocroaches-attack",
                            "id": "quorum-istanbul-tessera-cocroaches-attack",
                            "memory": "8192",
                            "network_mode": "bridge",
                            "placement_constraints.#": "0",
                            "requires_compatibilities.#": "1",
                            "requires_compatibilities.2737437151": "EC2",
                            "revision": "2",
                            "task_role_arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "volume.#": "1",
                            "volume.4175909684.host_path": "",
                            "volume.4175909684.name": "quorum_shared_volume"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_instance_profile.bastion": {
                    "type": "aws_iam_instance_profile",
                    "depends_on": [
                        "aws_iam_role.bastion",
                        "local.default_bastion_resource_name"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:instance-profile/quorum-bastion-cocroaches-attack",
                            "create_date": "2019-09-05T10:13:33Z",
                            "id": "quorum-bastion-cocroaches-attack",
                            "name": "quorum-bastion-cocroaches-attack",
                            "path": "/",
                            "role": "quorum-bastion-cocroaches-attack",
                            "roles.#": "0",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_instance_profile.ecsInstanceProfile": {
                    "type": "aws_iam_instance_profile",
                    "depends_on": [
                        "aws_iam_role.ecsInstanceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceProfile-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:instance-profile/ecsInstanceProfile-dccfbe87",
                            "create_date": "2019-09-05T10:13:33Z",
                            "id": "ecsInstanceProfile-dccfbe87",
                            "name": "ecsInstanceProfile-dccfbe87",
                            "path": "/",
                            "role": "ecsInstanceRole-dccfbe87",
                            "roles.#": "0",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_policy.bastion": {
                    "type": "aws_iam_policy",
                    "depends_on": [
                        "data.aws_iam_policy_document.bastion"
                    ],
                    "primary": {
                        "id": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "description": "This policy allows task to access S3 bucket and ECS",
                            "id": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "name": "quorum-bastion-policy-cocroaches-attack",
                            "path": "/",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:*\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowEC2\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ec2:*\",\n      \"Resource\": \"*\"\n    }\n  ]\n}"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_policy.ecs_task": {
                    "type": "aws_iam_policy",
                    "depends_on": [
                        "data.aws_iam_policy_document.ecs_task"
                    ],
                    "primary": {
                        "id": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "description": "This policy allows task to access S3 bucket",
                            "id": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "name": "quorum-ecs-task-policy-cocroaches-attack",
                            "path": "/",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3Access\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:DescribeTasks\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    }\n  ]\n}"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.bastion": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "local.default_bastion_resource_name"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/quorum-bastion-cocroaches-attack",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "quorum-bastion-cocroaches-attack",
                            "max_session_duration": "3600",
                            "name": "quorum-bastion-cocroaches-attack",
                            "path": "/",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecsInstanceRole": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceRole-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecsInstanceRole-dccfbe87",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"\",\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "ecsInstanceRole-dccfbe87",
                            "max_session_duration": "3600",
                            "name": "ecsInstanceRole-dccfbe87",
                            "path": "/",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": "***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecsServiceRole": {
                    "type": "aws_iam_role",
                    "depends_on": [
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsServiceRole-dccfbe87",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecsServiceRole-dccfbe87",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Sid\":\"\",\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ecs.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "ecsServiceRole-dccfbe87",
                            "max_session_duration": "3600",
                            "name": "ecsServiceRole-dccfbe87",
                            "path": "/",
                            "unique_id": ""
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role.ecs_task": {
                    "type": "aws_iam_role",
                    "depends_on": [],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack",
                        "attributes": {
                            "arn": "arn:aws:iam::051582052996:role/ecs/quorum-ecs-task-cocroaches-attack",
                            "assume_role_policy": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":[\"ecs-tasks.amazonaws.com\",\"ecs.amazonaws.com\"]},\"Action\":\"sts:AssumeRole\"}]}",
                            "create_date": "2019-09-05T10:13:31Z",
                            "force_detach_policies": "false",
                            "id": "quorum-ecs-task-cocroaches-attack",
                            "max_session_duration": "3600",
                            "name": "quorum-ecs-task-cocroaches-attack",
                            "path": "/ecs/",
<<<<<<< HEAD
                            "unique_id": ""
=======
                            "unique_id": ***REMOVED***"
>>>>>>> c0d0a771565e9813484c4e6d7719dec5b12dd682
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy.ecsInstanceRolePolicy": {
                    "type": "aws_iam_role_policy",
                    "depends_on": [
                        "aws_iam_role.ecsInstanceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsInstanceRole-dccfbe87:ecsInstanceRolePolicy-dccfbe87",
                        "attributes": {
                            "id": "ecsInstanceRole-dccfbe87:ecsInstanceRolePolicy-dccfbe87",
                            "name": "ecsInstanceRolePolicy-dccfbe87",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ecs:CreateCluster\",\n        \"ecs:DeregisterContainerInstance\",\n        \"ecs:DiscoverPollEndpoint\",\n        \"ecs:Poll\",\n        \"ecs:RegisterContainerInstance\",\n        \"ecs:StartTelemetrySession\",\n        \"ecs:Submit*\",\n        \"ecr:GetAuthorizationToken\",\n        \"ecr:BatchCheckLayerAvailability\",\n        \"ecr:GetDownloadUrlForLayer\",\n        \"ecr:BatchGetImage\",\n        \"logs:CreateLogStream\",\n        \"logs:PutLogEvents\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\n",
                            "role": "ecsInstanceRole-dccfbe87"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy.ecsServiceRolePolicy": {
                    "type": "aws_iam_role_policy",
                    "depends_on": [
                        "aws_iam_role.ecsServiceRole",
                        "random_id.code"
                    ],
                    "primary": {
                        "id": "ecsServiceRole-dccfbe87:ecsServiceRolePolicy-dccfbe87",
                        "attributes": {
                            "id": "ecsServiceRole-dccfbe87:ecsServiceRolePolicy-dccfbe87",
                            "name": "ecsServiceRolePolicy-dccfbe87",
                            "policy": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ec2:AuthorizeSecurityGroupIngress\",\n        \"ec2:Describe*\",\n        \"elasticloadbalancing:DeregisterInstancesFromLoadBalancer\",\n        \"elasticloadbalancing:DeregisterTargets\",\n        \"elasticloadbalancing:Describe*\",\n        \"elasticloadbalancing:RegisterInstancesWithLoadBalancer\",\n        \"elasticloadbalancing:RegisterTargets\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\n",
                            "role": "ecsServiceRole-dccfbe87"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.bastion": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_policy.bastion",
                        "aws_iam_role.bastion"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack-20190905101338541400000004",
                        "attributes": {
                            "id": "quorum-bastion-cocroaches-attack-20190905101338541400000004",
                            "policy_arn": "arn:aws:iam::051582052996:policy/quorum-bastion-policy-cocroaches-attack",
                            "role": "quorum-bastion-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_cloudwatch": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101333138700000002",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101333138700000002",
                            "policy_arn": "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_execution": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101333138100000001",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101333138100000001",
                            "policy_arn": "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_iam_role_policy_attachment.ecs_task_s3": {
                    "type": "aws_iam_role_policy_attachment",
                    "depends_on": [
                        "aws_iam_policy.ecs_task",
                        "aws_iam_role.ecs_task"
                    ],
                    "primary": {
                        "id": "quorum-ecs-task-cocroaches-attack-20190905101338522200000003",
                        "attributes": {
                            "id": "quorum-ecs-task-cocroaches-attack-20190905101338522200000003",
                            "policy_arn": "arn:aws:iam::051582052996:policy/quorum-ecs-task-policy-cocroaches-attack",
                            "role": "quorum-ecs-task-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_instance.bastion": {
                    "type": "aws_instance",
                    "depends_on": [
                        "aws_iam_instance_profile.bastion",
                        "aws_key_pair.ssh",
                        "aws_security_group.bastion-ethstats",
                        "aws_security_group.bastion-ssh",
                        "aws_security_group.quorum",
                        "aws_subnet.public",
                        "data.aws_ami.this",
                        "local.bastion_bucket",
                        "local.common_tags",
                        "local.default_bastion_resource_name",
                        "local.ethstats_docker_image",
                        "local.ethstats_port",
                        "local.quorum_docker_image",
                        "random_id.ethstat_secret",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "i-092a0879925d45d82",
                        "attributes": {
                            "ami": "ami-00c03f7f7f2ec15c3",
                            "associate_public_ip_address": "true",
                            "availability_zone": "us-east-2a",
                            "cpu_core_count": "2",
                            "cpu_threads_per_core": "1",
                            "credit_specification.#": "1",
                            "credit_specification.0.cpu_credits": "standard",
                            "disable_api_termination": "false",
                            "ebs_block_device.#": "0",
                            "ebs_optimized": "false",
                            "ephemeral_block_device.#": "0",
                            "get_password_data": "false",
                            "iam_instance_profile": "quorum-bastion-cocroaches-attack",
                            "id": "i-092a0879925d45d82",
                            "instance_state": "running",
                            "instance_type": "t2.large",
                            "ipv6_addresses.#": "0",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "monitoring": "false",
                            "network_interface.#": "0",
                            "network_interface_id": "eni-0e00d78db53cf56c9",
                            "password_data": "",
                            "placement_group": "",
                            "primary_network_interface_id": "eni-0e00d78db53cf56c9",
                            "private_dns": "ip-10-0-0-239.us-east-2.compute.internal",
                            "private_ip": "10.0.0.239",
                            "public_dns": "",
                            "public_ip": "invalid.ip.666",
                            "root_block_device.#": "1",
                            "root_block_device.0.delete_on_termination": "true",
                            "root_block_device.0.iops": "100",
                            "root_block_device.0.volume_id": "vol-0efe095602110db6f",
                            "root_block_device.0.volume_size": "8",
                            "root_block_device.0.volume_type": "gp2",
                            "security_groups.#": "0",
                            "source_dest_check": "true",
                            "subnet_id": "subnet-035251cd0096068f1",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "tenancy": "default",
                            "user_data": "c388a9205eb078013748e9431a2c713189dc038c",
                            "volume_tags.%": "0",
                            "vpc_security_group_ids.#": "3",
                            "vpc_security_group_ids.1114044214": "sg-02d4a1108481c7db4",
                            "vpc_security_group_ids.1424535282": "sg-047de8aa5ccdc10f5",
                            "vpc_security_group_ids.1485421960": "sg-04e6635eb212ff555"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 1200000000000,
                                "update": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_internet_gateway.this": {
                    "type": "aws_internet_gateway",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "igw-0d383894b8f6b5e94",
                        "attributes": {
                            "id": "igw-0d383894b8f6b5e94",
                            "tags.%": "1",
                            "tags.Name": "",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_key_pair.ssh": {
                    "type": "aws_key_pair",
                    "depends_on": [
                        "local.default_bastion_resource_name",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "quorum-bastion-cocroaches-attack",
                        "attributes": {
                            "id": "quorum-bastion-cocroaches-attack",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC775eGcJ2GZBfUfpiQnXqW/WfhQD8C7mnI1qz1COnR3yft7fpwVNP1wJ9lC8DB6HUdsKl0WISydrMrkwk5aIh4Hr5WJh2hcZq73FFUvQhEwnKvjce6rGTjl5pIRFwLgA9KB4hXfwCii6J+ul5CZs9zIteqaKovQbXcWh/6y0u0yvppP16Y2NuoDWqJjNNDo1QeiS27Ft88jbxVqX/B35cIzkjBiWvoLeoG6wD+K6r/kX0OMXomwyu4Ofc088WON4cXIiGHudDDMJPZD/h0hKAseDyflwpsY4qCG4bJpij3IcruvrMB1G+z0cpcERf3MtxFs+ob+LPZCD6v4QI2Eszv"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_kms_key.bucket": {
                    "type": "aws_kms_key",
                    "depends_on": [
                        "data.aws_iam_policy_document.kms_policy",
                        "local.common_tags"
                    ],
                    "primary": {
                        "id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                        "attributes": {
                            "arn": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "deletion_window_in_days": "7",
                            "description": "Used to encrypt/decrypt objects stored inside bucket created for this deployment",
                            "enable_key_rotation": "false",
                            "id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "is_enabled": "true",
                            "key_id": "bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "key_usage": "ENCRYPT_DECRYPT",
                            "policy": "{\"Statement\":[{\"Action\":\"kms:*\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::051582052996:root\"},\"Resource\":\"*\",\"Sid\":\"AllowAccess\"}],\"Version\":\"2012-10-17\"}",
                            "tags.%": "4",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_launch_configuration.lc": {
                    "type": "aws_launch_configuration",
                    "depends_on": [
                        "aws_iam_instance_profile.ecsInstanceProfile",
                        "aws_key_pair.ssh",
                        "aws_security_group.quorum",
                        "data.aws_ami.ecs_ami",
                        "data.template_file.user_data",
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "quorum-network-cocroaches-attack20190905101403318200000005",
                        "attributes": {
                            "associate_public_ip_address": "false",
                            "ebs_block_device.#": "0",
                            "ebs_optimized": "false",
                            "enable_monitoring": "true",
                            "ephemeral_block_device.#": "0",
                            "iam_instance_profile": "ecsInstanceProfile-dccfbe87",
                            "id": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "image_id": "ami-035a1bdaf0e4bf265",
                            "instance_type": "t2.xlarge",
                            "key_name": "quorum-bastion-cocroaches-attack",
                            "name": "quorum-network-cocroaches-attack20190905101403318200000005",
                            "name_prefix": "quorum-network-cocroaches-attack",
                            "root_block_device.#": "1",
                            "root_block_device.0.delete_on_termination": "true",
                            "root_block_device.0.iops": "0",
                            "root_block_device.0.volume_size": "16",
                            "root_block_device.0.volume_type": "",
                            "security_groups.#": "1",
                            "security_groups.1114044214": "sg-02d4a1108481c7db4",
                            "spot_price": "",
                            "user_data": "08fe58a67d425d4ad60356e369bf6e3c353f403b",
                            "vpc_classic_link_id": "",
                            "vpc_classic_link_security_groups.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route.public_internet_gateway": {
                    "type": "aws_route",
                    "depends_on": [
                        "aws_internet_gateway.this",
                        "aws_route_table.public"
                    ],
                    "primary": {
                        "id": "r-rtb-04cfca31968dae4771080289494",
                        "attributes": {
                            "destination_cidr_block": "0.0.0.0/0",
                            "destination_prefix_list_id": "",
                            "egress_only_gateway_id": "",
                            "gateway_id": "igw-0d383894b8f6b5e94",
                            "id": "r-rtb-04cfca31968dae4771080289494",
                            "instance_id": "",
                            "instance_owner_id": "",
                            "nat_gateway_id": "",
                            "network_interface_id": "",
                            "origin": "CreateRoute",
                            "route_table_id": "rtb-04cfca31968dae477",
                            "state": "active",
                            "vpc_peering_connection_id": ""
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 300000000000,
                                "delete": 300000000000
                            }
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.0": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-0539e90b9bba7657f",
                        "attributes": {
                            "id": "rtb-0539e90b9bba7657f",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.1": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-03b7807afd2862da6",
                        "attributes": {
                            "id": "rtb-03b7807afd2862da6",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2b",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.private.2": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.max_subnet_length",
                        "local.nat_gateway_count",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-0d69c07019210736f",
                        "attributes": {
                            "id": "rtb-0d69c07019210736f",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2c",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table.public": {
                    "type": "aws_route_table",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "rtb-04cfca31968dae477",
                        "attributes": {
                            "id": "rtb-04cfca31968dae477",
                            "propagating_vgws.#": "0",
                            "route.#": "0",
                            "tags.%": "1",
                            "tags.Name": "-public",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.0": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-0b1ffd059fd0a8a08",
                        "attributes": {
                            "id": "rtbassoc-0b1ffd059fd0a8a08",
                            "route_table_id": "rtb-0539e90b9bba7657f",
                            "subnet_id": "subnet-08c762a35be7b0024"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.1": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-00f88936a60c4cce2",
                        "attributes": {
                            "id": "rtbassoc-00f88936a60c4cce2",
                            "route_table_id": "rtb-03b7807afd2862da6",
                            "subnet_id": "subnet-0360554e5daf7b5d4"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.private.2": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.private.*",
                        "aws_subnet.private.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-05da0e76079fcfc35",
                        "attributes": {
                            "id": "rtbassoc-05da0e76079fcfc35",
                            "route_table_id": "rtb-0d69c07019210736f",
                            "subnet_id": "subnet-0dfb4a947cdf1ae22"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_route_table_association.public": {
                    "type": "aws_route_table_association",
                    "depends_on": [
                        "aws_route_table.public",
                        "aws_subnet.public.*"
                    ],
                    "primary": {
                        "id": "rtbassoc-010c308d38c127217",
                        "attributes": {
                            "id": "rtbassoc-010c308d38c127217",
                            "route_table_id": "rtb-04cfca31968dae477",
                            "subnet_id": "subnet-035251cd0096068f1"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_s3_bucket.bastion": {
                    "type": "aws_s3_bucket",
                    "depends_on": [
                        "local.bastion_bucket"
                    ],
                    "primary": {
                        "id": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                        "attributes": {
                            "acceleration_status": "",
                            "acl": "private",
                            "arn": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket_domain_name": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1.s3.amazonaws.com",
                            "bucket_regional_domain_name": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1.s3.us-east-2.amazonaws.com",
                            "cors_rule.#": "0",
                            "force_destroy": "true",
                            "hosted_zone_id": "Z2O1EMRO9K5GLX",
                            "id": "us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "logging.#": "0",
                            "region": "us-east-2",
                            "replication_configuration.#": "0",
                            "request_payer": "BucketOwner",
                            "server_side_encryption_configuration.#": "0",
                            "tags.%": "0",
                            "versioning.#": "1",
                            "versioning.0.enabled": "true",
                            "versioning.0.mfa_delete": "false",
                            "website.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_s3_bucket.quorum": {
                    "type": "aws_s3_bucket",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "data.aws_iam_policy_document.bucket_policy",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                        "attributes": {
                            "acceleration_status": "",
                            "acl": "private",
                            "arn": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "bucket_domain_name": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1.s3.amazonaws.com",
                            "bucket_regional_domain_name": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1.s3.us-east-2.amazonaws.com",
                            "cors_rule.#": "0",
                            "force_destroy": "true",
                            "hosted_zone_id": "Z2O1EMRO9K5GLX",
                            "id": "us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "logging.#": "0",
                            "policy": "{\"Statement\":[{\"Action\":\"s3:*\",\"Effect\":\"Allow\",\"Principal\":{\"AWS\":\"arn:aws:iam::051582052996:root\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"AllowAccess\"},{\"Action\":\"s3:PutObject\",\"Condition\":{\"Null\":{\"s3:x-amz-server-side-encryption\":\"true\"}},\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"DenyAccess1\"},{\"Action\":\"s3:PutObject\",\"Condition\":{\"StringNotEquals\":{\"s3:x-amz-server-side-encryption\":\"aws:kms\"}},\"Effect\":\"Deny\",\"Principal\":{\"AWS\":\"*\"},\"Resource\":[\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"],\"Sid\":\"DenyAccess2\"}],\"Version\":\"2012-10-17\"}",
                            "region": "us-east-2",
                            "replication_configuration.#": "0",
                            "request_payer": "BucketOwner",
                            "server_side_encryption_configuration.#": "1",
                            "server_side_encryption_configuration.0.rule.#": "1",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.#": "1",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.0.kms_master_key_id": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "server_side_encryption_configuration.0.rule.0.apply_server_side_encryption_by_default.0.sse_algorithm": "aws:kms",
                            "tags.%": "0",
                            "versioning.#": "1",
                            "versioning.0.enabled": "true",
                            "versioning.0.mfa_delete": "false",
                            "website.#": "0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-ethstats": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "aws_subnet.public",
                        "local.common_tags",
                        "local.quorum_rpc_port",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-047de8aa5ccdc10f5",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-047de8aa5ccdc10f5",
                            "description": "Security group used by external to access ethstats for Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-047de8aa5ccdc10f5",
                            "ingress.#": "1",
                            "ingress.2872103538.cidr_blocks.#": "1",
                            "ingress.2872103538.cidr_blocks.0": "10.0.0.0/24",
                            "ingress.2872103538.description": "Allow geth console",
                            "ingress.2872103538.from_port": "22000",
                            "ingress.2872103538.ipv6_cidr_blocks.#": "0",
                            "ingress.2872103538.protocol": "tcp",
                            "ingress.2872103538.security_groups.#": "0",
                            "ingress.2872103538.self": "false",
                            "ingress.2872103538.to_port": "22000",
                            "name": "quorum-bastion-ethstats-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "client-bastion-geth-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-geth": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-05f12b4adeda993e5",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-05f12b4adeda993e5",
                            "description": "Security group used by external to access geth for quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-05f12b4adeda993e5",
                            "ingress.#": "1",
                            "ingress.338493142.cidr_blocks.#": "1",
                            "ingress.338493142.cidr_blocks.0": "0.0.0.0/0",
                            "ingress.338493142.description": "Allow ethstats",
                            "ingress.338493142.from_port": "3000",
                            "ingress.338493142.ipv6_cidr_blocks.#": "0",
                            "ingress.338493142.protocol": "tcp",
                            "ingress.338493142.security_groups.#": "0",
                            "ingress.338493142.self": "false",
                            "ingress.338493142.to_port": "3000",
                            "name": "quorum-bastion-get-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-ethstats-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.bastion-ssh": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-04e6635eb212ff555",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-04e6635eb212ff555",
                            "description": "Security group used by Bastion node to access Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-04e6635eb212ff555",
                            "ingress.#": "1",
                            "ingress.1248900448.cidr_blocks.#": "1",
                            "ingress.1248900448.cidr_blocks.0": "0.0.0.0/0",
                            "ingress.1248900448.description": "Allow SSH",
                            "ingress.1248900448.from_port": "22",
                            "ingress.1248900448.ipv6_cidr_blocks.#": "0",
                            "ingress.1248900448.protocol": "tcp",
                            "ingress.1248900448.security_groups.#": "0",
                            "ingress.1248900448.self": "false",
                            "ingress.1248900448.to_port": "22",
                            "name": "quorum-bastion-ssh-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-bastion-ssh-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group.quorum": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "local.common_tags",
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "sg-02d4a1108481c7db4",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-02d4a1108481c7db4",
                            "description": "Security group used in Quorum network cocroaches-attack",
                            "egress.#": "1",
                            "egress.2405122738.cidr_blocks.#": "1",
                            "egress.2405122738.cidr_blocks.0": "0.0.0.0/0",
                            "egress.2405122738.description": "Allow all",
                            "egress.2405122738.from_port": "0",
                            "egress.2405122738.ipv6_cidr_blocks.#": "0",
                            "egress.2405122738.prefix_list_ids.#": "0",
                            "egress.2405122738.protocol": "-1",
                            "egress.2405122738.security_groups.#": "0",
                            "egress.2405122738.self": "false",
                            "egress.2405122738.to_port": "0",
                            "id": "sg-02d4a1108481c7db4",
                            "ingress.#": "0",
                            "name": "quorum-sg-cocroaches-attack",
                            "owner_id": "051582052996",
                            "revoke_rules_on_delete": "false",
                            "tags.%": "5",
                            "tags.DockerImage.PrivacyEngine": "quorumengineering/tessera:latest",
                            "tags.DockerImage.Quorum": "quorumengineering/quorum:latest",
                            "tags.ECSClusterName": "quorum-network-cocroaches-attack",
                            "tags.Name": "quorum-sg-cocroaches-attack",
                            "tags.NetworkName": "cocroaches-attack",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "e2bfb730-ecaa-11e6-8f88-34363bc7c4c0": {
                                "create": 600000000000,
                                "delete": 600000000000
                            },
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.ethstats": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.ethstats_port"
                    ],
                    "primary": {
                        "id": "sgrule-2136424293",
                        "attributes": {
                            "description": "ethstats traffic",
                            "from_port": "3000",
                            "id": "sgrule-2136424293",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "3000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.ethstats-bastion": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.bastion-ethstats",
                        "local.ethstats_port"
                    ],
                    "primary": {
                        "id": "sgrule-435918442",
                        "attributes": {
                            "description": "ethstats traffic",
                            "from_port": "3000",
                            "id": "sgrule-435918442",
                            "protocol": "tcp",
                            "security_group_id": "sg-047de8aa5ccdc10f5",
                            "self": "true",
                            "to_port": "3000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.geth_admin_rpc": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.quorum_rpc_port"
                    ],
                    "primary": {
                        "id": "sgrule-3810161799",
                        "attributes": {
                            "description": "Geth Admin RPC traffic",
                            "from_port": "22000",
                            "id": "sgrule-3810161799",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "22000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.geth_p2p": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.quorum_p2p_port"
                    ],
                    "primary": {
                        "id": "sgrule-256513769",
                        "attributes": {
                            "description": "Geth P2P traffic",
                            "from_port": "21000",
                            "id": "sgrule-256513769",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "21000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.open-all-ingress-research": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum"
                    ],
                    "primary": {
                        "id": "sgrule-2623329713",
                        "attributes": {
                            "cidr_blocks.#": "1",
                            "cidr_blocks.0": "0.0.0.0/0",
                            "description": "Open all ports",
                            "from_port": "0",
                            "id": "sgrule-2623329713",
                            "protocol": "-1",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "false",
                            "to_port": "0",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.tessera": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.tessera_port"
                    ],
                    "primary": {
                        "id": "sgrule-637454889",
                        "attributes": {
                            "description": "Tessera API traffic",
                            "from_port": "9000",
                            "id": "sgrule-637454889",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "9000",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_security_group_rule.tessera_thirdparty": {
                    "type": "aws_security_group_rule",
                    "depends_on": [
                        "aws_security_group.quorum",
                        "local.tessera_thirdparty_port"
                    ],
                    "primary": {
                        "id": "sgrule-4039332832",
                        "attributes": {
                            "description": "Tessera Thirdparty API traffic",
                            "from_port": "9080",
                            "id": "sgrule-4039332832",
                            "protocol": "tcp",
                            "security_group_id": "sg-02d4a1108481c7db4",
                            "self": "true",
                            "to_port": "9080",
                            "type": "ingress"
                        },
                        "meta": {
                            "schema_version": "2"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.0": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-08c762a35be7b0024",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2a",
                            "cidr_block": "10.0.1.0/24",
                            "id": "subnet-08c762a35be7b0024",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.1": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-0360554e5daf7b5d4",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2b",
                            "cidr_block": "10.0.2.0/24",
                            "id": "subnet-0360554e5daf7b5d4",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2b",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.private.2": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-0dfb4a947cdf1ae22",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2c",
                            "cidr_block": "10.0.3.0/24",
                            "id": "subnet-0dfb4a947cdf1ae22",
                            "map_public_ip_on_launch": "false",
                            "tags.%": "1",
                            "tags.Name": "-private-us-east-2c",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_subnet.public": {
                    "type": "aws_subnet",
                    "depends_on": [
                        "local.vpc_id"
                    ],
                    "primary": {
                        "id": "subnet-035251cd0096068f1",
                        "attributes": {
                            "assign_ipv6_address_on_creation": "false",
                            "availability_zone": "us-east-2a",
                            "cidr_block": "10.0.0.0/24",
                            "id": "subnet-035251cd0096068f1",
                            "map_public_ip_on_launch": "true",
                            "tags.%": "1",
                            "tags.Name": "-public-us-east-2a",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "aws_vpc.this": {
                    "type": "aws_vpc",
                    "depends_on": [],
                    "primary": {
                        "id": "vpc-010e95f77f8c7f7ee",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:vpc/vpc-010e95f77f8c7f7ee",
                            "assign_generated_ipv6_cidr_block": "false",
                            "cidr_block": "10.0.0.0/16",
                            "default_network_acl_id": "acl-0f9342c12ab343237",
                            "default_route_table_id": "rtb-08cc5b75ea27736b1",
                            "default_security_group_id": "sg-0825371596c4eebbb",
                            "dhcp_options_id": "dopt-d710e5bc",
                            "enable_dns_hostnames": "false",
                            "enable_dns_support": "true",
                            "id": "vpc-010e95f77f8c7f7ee",
                            "instance_tenancy": "default",
                            "main_route_table_id": "rtb-08cc5b75ea27736b1",
                            "tags.%": "1",
                            "tags.Name": ""
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_ami.ecs_ami": {
                    "type": "aws_ami",
                    "depends_on": [],
                    "primary": {
                        "id": "ami-035a1bdaf0e4bf265",
                        "attributes": {
                            "architecture": "x86_64",
                            "block_device_mappings.#": "2",
                            "block_device_mappings.2538004115.device_name": "/dev/xvdcz",
                            "block_device_mappings.2538004115.ebs.%": "5",
                            "block_device_mappings.2538004115.ebs.delete_on_termination": "true",
                            "block_device_mappings.2538004115.ebs.encrypted": "false",
                            "block_device_mappings.2538004115.ebs.iops": "0",
                            "block_device_mappings.2538004115.ebs.volume_size": "22",
                            "block_device_mappings.2538004115.ebs.volume_type": "gp2",
                            "block_device_mappings.2538004115.no_device": "",
                            "block_device_mappings.2538004115.virtual_name": "",
                            "block_device_mappings.340275815.device_name": "/dev/xvda",
                            "block_device_mappings.340275815.ebs.%": "6",
                            "block_device_mappings.340275815.ebs.delete_on_termination": "true",
                            "block_device_mappings.340275815.ebs.encrypted": "false",
                            "block_device_mappings.340275815.ebs.iops": "0",
                            "block_device_mappings.340275815.ebs.snapshot_id": "snap-09a5a92a34c4cf8c3",
                            "block_device_mappings.340275815.ebs.volume_size": "8",
                            "block_device_mappings.340275815.ebs.volume_type": "gp2",
                            "block_device_mappings.340275815.no_device": "",
                            "block_device_mappings.340275815.virtual_name": "",
                            "creation_date": "2019-08-16T22:34:38.000Z",
                            "description": "Amazon Linux AMI 2018.03.w x86_64 ECS HVM GP2",
                            "filter.#": "1",
                            "filter.3350713981.name": "name",
                            "filter.3350713981.values.#": "1",
                            "filter.3350713981.values.0": "amzn-ami-*-amazon-ecs-optimized",
                            "hypervisor": "xen",
                            "id": "ami-035a1bdaf0e4bf265",
                            "image_id": "ami-035a1bdaf0e4bf265",
                            "image_location": "amazon/amzn-ami-2018.03.w-amazon-ecs-optimized",
                            "image_owner_alias": "amazon",
                            "image_type": "machine",
                            "most_recent": "true",
                            "name": "amzn-ami-2018.03.w-amazon-ecs-optimized",
                            "owner_id": "591542846629",
                            "owners.#": "1",
                            "owners.0": "amazon",
                            "product_codes.#": "0",
                            "public": "true",
                            "root_device_name": "/dev/xvda",
                            "root_device_type": "ebs",
                            "root_snapshot_id": "snap-09a5a92a34c4cf8c3",
                            "sriov_net_support": "simple",
                            "state": "available",
                            "state_reason.%": "2",
                            "state_reason.code": "UNSET",
                            "state_reason.message": "UNSET",
                            "tags.%": "0",
                            "virtualization_type": "hvm"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_ami.this": {
                    "type": "aws_ami",
                    "depends_on": [],
                    "primary": {
                        "id": "ami-00c03f7f7f2ec15c3",
                        "attributes": {
                            "architecture": "x86_64",
                            "block_device_mappings.#": "1",
                            "block_device_mappings.340275815.device_name": "/dev/xvda",
                            "block_device_mappings.340275815.ebs.%": "6",
                            "block_device_mappings.340275815.ebs.delete_on_termination": "true",
                            "block_device_mappings.340275815.ebs.encrypted": "false",
                            "block_device_mappings.340275815.ebs.iops": "0",
                            "block_device_mappings.340275815.ebs.snapshot_id": "snap-0e43f2c4403c9a35b",
                            "block_device_mappings.340275815.ebs.volume_size": "8",
                            "block_device_mappings.340275815.ebs.volume_type": "gp2",
                            "block_device_mappings.340275815.no_device": "",
                            "block_device_mappings.340275815.virtual_name": "",
                            "creation_date": "2019-08-30T07:06:00.000Z",
                            "description": "Amazon Linux 2 AMI 2.0.20190823.1 x86_64 HVM gp2",
                            "filter.#": "3",
                            "filter.2026626658.name": "name",
                            "filter.2026626658.values.#": "1",
                            "filter.2026626658.values.0": "amzn2-ami-hvm-*",
                            "filter.3386043752.name": "architecture",
                            "filter.3386043752.values.#": "1",
                            "filter.3386043752.values.0": "x86_64",
                            "filter.490168357.name": "virtualization-type",
                            "filter.490168357.values.#": "1",
                            "filter.490168357.values.0": "hvm",
                            "hypervisor": "xen",
                            "id": "ami-00c03f7f7f2ec15c3",
                            "image_id": "ami-00c03f7f7f2ec15c3",
                            "image_location": "amazon/amzn2-ami-hvm-2.0.20190823.1-x86_64-gp2",
                            "image_owner_alias": "amazon",
                            "image_type": "machine",
                            "most_recent": "true",
                            "name": "amzn2-ami-hvm-2.0.20190823.1-x86_64-gp2",
                            "owner_id": "137112412989",
                            "owners.#": "1",
                            "owners.0": "137112412989",
                            "product_codes.#": "0",
                            "public": "true",
                            "root_device_name": "/dev/xvda",
                            "root_device_type": "ebs",
                            "root_snapshot_id": "snap-0e43f2c4403c9a35b",
                            "sriov_net_support": "simple",
                            "state": "available",
                            "state_reason.%": "2",
                            "state_reason.code": "UNSET",
                            "state_reason.message": "UNSET",
                            "tags.%": "0",
                            "virtualization_type": "hvm"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_caller_identity.this": {
                    "type": "aws_caller_identity",
                    "depends_on": [],
                    "primary": {
                        "id": "2019-09-05 10:13:23.699585902 +0000 UTC",
                        "attributes": {
                           
                            "arn": "arn:aws:iam::051582052996:user/bkrzakala",
                            "id": "2019-09-05 10:13:23.699585902 +0000 UTC",
                            "user_id": "***REMOVED***"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.bastion": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "local.bastion_bucket",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "3317039311",
                        "attributes": {
                            "id": "3317039311",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:*\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowEC2\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ec2:*\",\n      \"Resource\": \"*\"\n    }\n  ]\n}",
                            "statement.#": "5",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "0",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowS3",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.1834123015": "s3:*",
                            "statement.1.condition.#": "0",
                            "statement.1.effect": "Allow",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "0",
                            "statement.1.resources.#": "2",
                            "statement.1.resources.2997558646": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.1.resources.94925287": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.1.sid": "AllowS3Bastion",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.1011197950": "kms:*",
                            "statement.2.condition.#": "0",
                            "statement.2.effect": "Allow",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "0",
                            "statement.2.resources.#": "1",
                            "statement.2.resources.3285651814": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "statement.2.sid": "AllowKMSAccess",
                            "statement.3.actions.#": "1",
                            "statement.3.actions.3112138991": "ecs:*",
                            "statement.3.condition.#": "0",
                            "statement.3.effect": "Allow",
                            "statement.3.not_actions.#": "0",
                            "statement.3.not_principals.#": "0",
                            "statement.3.not_resources.#": "0",
                            "statement.3.principals.#": "0",
                            "statement.3.resources.#": "1",
                            "statement.3.resources.2679715827": "*",
                            "statement.3.sid": "AllowECS",
                            "statement.4.actions.#": "1",
                            "statement.4.actions.2597799863": "ec2:*",
                            "statement.4.condition.#": "0",
                            "statement.4.effect": "Allow",
                            "statement.4.not_actions.#": "0",
                            "statement.4.not_principals.#": "0",
                            "statement.4.not_resources.#": "0",
                            "statement.4.principals.#": "0",
                            "statement.4.resources.#": "1",
                            "statement.4.resources.2679715827": "*",
                            "statement.4.sid": "AllowEC2"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.bucket_policy": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "data.aws_caller_identity.this",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "2088926263",
                        "attributes": {
                            "id": "2088926263",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"arn:aws:iam::051582052996:root\"\n      }\n    },\n    {\n      \"Sid\": \"DenyAccess1\",\n      \"Effect\": \"Deny\",\n      \"Action\": \"s3:PutObject\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"*\"\n      },\n      \"Condition\": {\n        \"Null\": {\n          \"s3:x-amz-server-side-encryption\": \"true\"\n        }\n      }\n    },\n    {\n      \"Sid\": \"DenyAccess2\",\n      \"Effect\": \"Deny\",\n      \"Action\": \"s3:PutObject\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ],\n      \"Principal\": {\n        \"AWS\": \"*\"\n      },\n      \"Condition\": {\n        \"StringNotEquals\": {\n          \"s3:x-amz-server-side-encryption\": \"aws:kms\"\n        }\n      }\n    }\n  ]\n}",
                            "statement.#": "3",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "1",
                            "statement.0.principals.3691871349.identifiers.#": "1",
                            "statement.0.principals.3691871349.identifiers.2401438501": "arn:aws:iam::051582052996:root",
                            "statement.0.principals.3691871349.type": "AWS",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowAccess",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.315547055": "s3:PutObject",
                            "statement.1.condition.#": "1",
                            "statement.1.condition.3673734150.test": "Null",
                            "statement.1.condition.3673734150.values.#": "1",
                            "statement.1.condition.3673734150.values.4043113848": "true",
                            "statement.1.condition.3673734150.variable": "s3:x-amz-server-side-encryption",
                            "statement.1.effect": "Deny",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "1",
                            "statement.1.principals.636693895.identifiers.#": "1",
                            "statement.1.principals.636693895.identifiers.2679715827": "*",
                            "statement.1.principals.636693895.type": "AWS",
                            "statement.1.resources.#": "2",
                            "statement.1.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.1.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.1.sid": "DenyAccess1",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.315547055": "s3:PutObject",
                            "statement.2.condition.#": "1",
                            "statement.2.condition.3059326814.test": "StringNotEquals",
                            "statement.2.condition.3059326814.values.#": "1",
                            "statement.2.condition.3059326814.values.800761281": "aws:kms",
                            "statement.2.condition.3059326814.variable": "s3:x-amz-server-side-encryption",
                            "statement.2.effect": "Deny",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "1",
                            "statement.2.principals.636693895.identifiers.#": "1",
                            "statement.2.principals.636693895.identifiers.2679715827": "*",
                            "statement.2.principals.636693895.type": "AWS",
                            "statement.2.resources.#": "2",
                            "statement.2.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.2.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.2.sid": "DenyAccess2"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.ecs_task": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "aws_kms_key.bucket",
                        "local.bastion_bucket",
                        "local.quorum_bucket"
                    ],
                    "primary": {
                        "id": "2507052418",
                        "attributes": {
                            "id": "2507052418",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowS3Access\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    },\n    {\n      \"Sid\": \"AllowKMSAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda\"\n    },\n    {\n      \"Sid\": \"AllowECS\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"ecs:DescribeTasks\",\n      \"Resource\": \"*\"\n    },\n    {\n      \"Sid\": \"AllowS3Bastion\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"s3:*\",\n      \"Resource\": [\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*\",\n        \"arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1\"\n      ]\n    }\n  ]\n}",
                            "statement.#": "4",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1834123015": "s3:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "0",
                            "statement.0.resources.#": "2",
                            "statement.0.resources.1362820624": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.0.resources.2488377391": "arn:aws:s3:::us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.0.sid": "AllowS3Access",
                            "statement.1.actions.#": "1",
                            "statement.1.actions.1011197950": "kms:*",
                            "statement.1.condition.#": "0",
                            "statement.1.effect": "Allow",
                            "statement.1.not_actions.#": "0",
                            "statement.1.not_principals.#": "0",
                            "statement.1.not_resources.#": "0",
                            "statement.1.principals.#": "0",
                            "statement.1.resources.#": "1",
                            "statement.1.resources.3285651814": "arn:aws:kms:us-east-2:051582052996:key/bfaef9bd-cecb-43d5-9508-eeaa9b0e5fda",
                            "statement.1.sid": "AllowKMSAccess",
                            "statement.2.actions.#": "1",
                            "statement.2.actions.974674342": "ecs:DescribeTasks",
                            "statement.2.condition.#": "0",
                            "statement.2.effect": "Allow",
                            "statement.2.not_actions.#": "0",
                            "statement.2.not_principals.#": "0",
                            "statement.2.not_resources.#": "0",
                            "statement.2.principals.#": "0",
                            "statement.2.resources.#": "1",
                            "statement.2.resources.2679715827": "*",
                            "statement.2.sid": "AllowECS",
                            "statement.3.actions.#": "1",
                            "statement.3.actions.1834123015": "s3:*",
                            "statement.3.condition.#": "0",
                            "statement.3.effect": "Allow",
                            "statement.3.not_actions.#": "0",
                            "statement.3.not_principals.#": "0",
                            "statement.3.not_resources.#": "0",
                            "statement.3.principals.#": "0",
                            "statement.3.resources.#": "2",
                            "statement.3.resources.2997558646": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/*",
                            "statement.3.resources.94925287": "arn:aws:s3:::us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1",
                            "statement.3.sid": "AllowS3Bastion"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_iam_policy_document.kms_policy": {
                    "type": "aws_iam_policy_document",
                    "depends_on": [
                        "data.aws_caller_identity.this"
                    ],
                    "primary": {
                        "id": "1405806754",
                        "attributes": {
                            "id": "1405806754",
                            "json": "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"AllowAccess\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"kms:*\",\n      \"Resource\": \"*\",\n      \"Principal\": {\n        \"AWS\": \"arn:aws:iam::051582052996:root\"\n      }\n    }\n  ]\n}",
                            "statement.#": "1",
                            "statement.0.actions.#": "1",
                            "statement.0.actions.1011197950": "kms:*",
                            "statement.0.condition.#": "0",
                            "statement.0.effect": "Allow",
                            "statement.0.not_actions.#": "0",
                            "statement.0.not_principals.#": "0",
                            "statement.0.not_resources.#": "0",
                            "statement.0.principals.#": "1",
                            "statement.0.principals.3691871349.identifiers.#": "1",
                            "statement.0.principals.3691871349.identifiers.2401438501": "arn:aws:iam::051582052996:root",
                            "statement.0.principals.3691871349.type": "AWS",
                            "statement.0.resources.#": "1",
                            "statement.0.resources.2679715827": "*",
                            "statement.0.sid": "AllowAccess"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.aws_security_group.default": {
                    "type": "aws_security_group",
                    "depends_on": [
                        "aws_vpc.this"
                    ],
                    "primary": {
                        "id": "sg-0825371596c4eebbb",
                        "attributes": {
                            "arn": "arn:aws:ec2:us-east-2:051582052996:security-group/sg-0825371596c4eebbb",
                            "description": "default VPC security group",
                            "id": "sg-0825371596c4eebbb",
                            "name": "default",
                            "tags.%": "0",
                            "vpc_id": "vpc-010e95f77f8c7f7ee"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.aws"
                },
                "data.template_file.user_data": {
                    "type": "template_file",
                    "depends_on": [
                        "local.ecs_cluster_name"
                    ],
                    "primary": {
                        "id": "c08a5a7b7fca738108790bf3d0b997c3ccffebc143541469ec3785f8122811c8",
                        "attributes": {
                            "id": "c08a5a7b7fca738108790bf3d0b997c3ccffebc143541469ec3785f8122811c8",
                            "rendered": "    #!/bin/bash\n    echo ECS_CLUSTER=quorum-network-cocroaches-attack \u003e\u003e /etc/ecs/ecs.config\n\n    # node_exporter part\n    set -e\n    cd /tmp\n    curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz\n    tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz\n    mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/\n    useradd -rs /bin/false node_exporter\n\n\n    tee -a /etc/init.d/node_exporter \u003c\u003c END\n#!/bin/bash\n\n### BEGIN INIT INFO\n# processname:       node_exporter\n# Short-Description: Exporter for machine metrics.\n# Description:       Prometheus exporter for machine metrics,\n#                    written in Go with pluggable metric collectors.\n#\n# chkconfig: 2345 80 80\n# pidfile: /var/run/node_exporter/node_exporter.pid\n#\n#\n### END INIT INFO\n\n# Source function library.\n. /etc/init.d/functions\n\nNAME=node_exporter\nDESC=\"Exporter for machine metrics\"\nDAEMON=/usr/local/bin/node_exporter\nUSER=node_exporter\nCONFIG=\nPID=/var/run/node_exporter/\\$NAME.pid\nLOG=/var/log/node_exporter/\\$NAME.log\n\nDAEMON_OPTS=\nRETVAL=0\n\n# Check if DAEMON binary exist\n[ -f \\$DAEMON ] || exit 0\n\n[ -f /etc/default/node_exporter ]  \u0026\u0026  . /etc/default/node_exporter\n\nservice_checks() {\n  # Prepare directories\n  mkdir -p /var/run/node_exporter /var/log/node_exporter\n  chown -R \\$USER /var/run/node_exporter /var/log/node_exporter\n\n  # Check if PID exists\n  if [ -f \"\\$PID\" ]; then\n    PID_NUMBER=\\$(cat \\$PID)\n    if [ -z \"\\$(ps axf | grep \\$PID_NUMBER | grep -v grep)\" ]; then\n      echo \"Service was aborted abnormally; clean the PID file and continue...\"\n      rm -f \"\\$PID\"\n    else\n      echo \"Service already started; skip...\"\n      exit 1\n    fi\n  fi\n}\n\nstart() {\n  service_checks \\$1\n  sudo -H -u \\$USER   \\$DAEMON \\$DAEMON_OPTS  \u003e \\$LOG 2\u003e\u00261  \u0026\n  RETVAL=\\$?\n  echo \\$! \u003e \\$PID\n}\n\nstop() {\n  killproc -p \\$PID -b \\$DAEMON  \\$NAME\n  RETVAL=\\$?\n}\n\nreload() {\n  #-- sorry but node_exporter doesn't handle -HUP signal...\n  #killproc -p \\$PID -b \\$DAEMON  \\$NAME -HUP\n  #RETVAL=\\$?\n  stop\n  start\n}\n\ncase \"\\$1\" in\n  start)\n    echo -n \\$\"Starting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    start\n    ;;\n\n  stop)\n    echo -n \\$\"Stopping \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    ;;\n\n  reload)\n    echo -n \\$\"Reloading \\$DESC configuration -\" \"\\$NAME\" \\$'\\n'\n    reload\n    ;;\n\n  restart|force-reload)\n    echo -n \\$\"Restarting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    start\n    ;;\n\n  syntax)\n    \\$DAEMON --help\n    ;;\n\n  status)\n    status -p \\$PID \\$DAEMON\n    ;;\n\n  *)\n    echo -n \\$\"Usage: /etc/init.d/\\$NAME {start|stop|reload|restart|force-reload|syntax|status}\" \\$'\\n'\n    ;;\nesac\n\nexit \\$RETVAL\nEND\n\nchmod +x /etc/init.d/node_exporter\nservice node_exporter start\nchkconfig node_exporter on\n\n",
                            "template": "    #!/bin/bash\n    echo ECS_CLUSTER=quorum-network-cocroaches-attack \u003e\u003e /etc/ecs/ecs.config\n\n    # node_exporter part\n    set -e\n    cd /tmp\n    curl -L -O https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz\n    tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz\n    mv node_exporter-0.18.1.linux-amd64/node_exporter /usr/local/bin/\n    useradd -rs /bin/false node_exporter\n\n\n    tee -a /etc/init.d/node_exporter \u003c\u003c END\n#!/bin/bash\n\n### BEGIN INIT INFO\n# processname:       node_exporter\n# Short-Description: Exporter for machine metrics.\n# Description:       Prometheus exporter for machine metrics,\n#                    written in Go with pluggable metric collectors.\n#\n# chkconfig: 2345 80 80\n# pidfile: /var/run/node_exporter/node_exporter.pid\n#\n#\n### END INIT INFO\n\n# Source function library.\n. /etc/init.d/functions\n\nNAME=node_exporter\nDESC=\"Exporter for machine metrics\"\nDAEMON=/usr/local/bin/node_exporter\nUSER=node_exporter\nCONFIG=\nPID=/var/run/node_exporter/\\$NAME.pid\nLOG=/var/log/node_exporter/\\$NAME.log\n\nDAEMON_OPTS=\nRETVAL=0\n\n# Check if DAEMON binary exist\n[ -f \\$DAEMON ] || exit 0\n\n[ -f /etc/default/node_exporter ]  \u0026\u0026  . /etc/default/node_exporter\n\nservice_checks() {\n  # Prepare directories\n  mkdir -p /var/run/node_exporter /var/log/node_exporter\n  chown -R \\$USER /var/run/node_exporter /var/log/node_exporter\n\n  # Check if PID exists\n  if [ -f \"\\$PID\" ]; then\n    PID_NUMBER=\\$(cat \\$PID)\n    if [ -z \"\\$(ps axf | grep \\$PID_NUMBER | grep -v grep)\" ]; then\n      echo \"Service was aborted abnormally; clean the PID file and continue...\"\n      rm -f \"\\$PID\"\n    else\n      echo \"Service already started; skip...\"\n      exit 1\n    fi\n  fi\n}\n\nstart() {\n  service_checks \\$1\n  sudo -H -u \\$USER   \\$DAEMON \\$DAEMON_OPTS  \u003e \\$LOG 2\u003e\u00261  \u0026\n  RETVAL=\\$?\n  echo \\$! \u003e \\$PID\n}\n\nstop() {\n  killproc -p \\$PID -b \\$DAEMON  \\$NAME\n  RETVAL=\\$?\n}\n\nreload() {\n  #-- sorry but node_exporter doesn't handle -HUP signal...\n  #killproc -p \\$PID -b \\$DAEMON  \\$NAME -HUP\n  #RETVAL=\\$?\n  stop\n  start\n}\n\ncase \"\\$1\" in\n  start)\n    echo -n \\$\"Starting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    start\n    ;;\n\n  stop)\n    echo -n \\$\"Stopping \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    ;;\n\n  reload)\n    echo -n \\$\"Reloading \\$DESC configuration -\" \"\\$NAME\" \\$'\\n'\n    reload\n    ;;\n\n  restart|force-reload)\n    echo -n \\$\"Restarting \\$DESC -\" \"\\$NAME\" \\$'\\n'\n    stop\n    start\n    ;;\n\n  syntax)\n    \\$DAEMON --help\n    ;;\n\n  status)\n    status -p \\$PID \\$DAEMON\n    ;;\n\n  *)\n    echo -n \\$\"Usage: /etc/init.d/\\$NAME {start|stop|reload|restart|force-reload|syntax|status}\" \\$'\\n'\n    ;;\nesac\n\nexit \\$RETVAL\nEND\n\nchmod +x /etc/init.d/node_exporter\nservice node_exporter start\nchkconfig node_exporter on\n\n",
                            "vars.%": "1",
                            "vars.ecs_cluster_name": "quorum-network-cocroaches-attack"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.template"
                },
                "local_file.bootstrap": {
                    "type": "local_file",
                    "depends_on": [
                        "aws_ecs_service.quorum.*",
                        "aws_ecs_task_definition.quorum",
                        "local.bastion_bucket",
                        "local.ecs_cluster_name",
                        "local.hosts_folder",
                        "local.normalized_host_ip",
                        "local.privacy_addresses_folder",
                        "local.quorum_docker_image",
                        "local.quorum_rpc_port",
                        "local.quorum_run_container_name",
                        "local.s3_revision_folder",
                        "local.shared_volume_container_path",
                        "local.tessera_thirdparty_port",
                        "random_string.random"
                    ],
                    "primary": {
                        "id": "09bccb7dc18c01e3a1c96853e9a3c885cb472081",
                        "attributes": {
                            "content": "#!/bin/bash\n\nset -e\n\nexport AWS_DEFAULT_REGION=us-east-2\nexport TASK_REVISION=2\nsudo rm -rf /qdata\nsudo mkdir -p /qdata/mappings\nsudo mkdir -p /qdata/privacyaddresses\n\n# Faketime array ( ClockSkew )\nold_IFS=$IFS\nIFS=',' faketime=(1 -3 2)\nIFS=${old_IFS}\ncounter=\"${#faketime[@]}\"\n\nwhile [ $counter -gt 0 ]\ndo\n    echo -n \"${faketime[-1]}\" \u003e ./$counter\n    faketime=(${faketime[@]::$counter})\n    sudo aws s3 cp ./$counter s3://us-east-2-bastion-cocroaches-attack-e8c6b6c6766602b1/clockSkew/\n    counter=$((counter - 1))\ndone\n\ncount=0\nwhile [ $count -lt 0 ]\ndo\n  count=$(ls /qdata/privacyaddresses | grep ^ip | wc -l)\n  sudo aws s3 cp --recursive s3://us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1/rev_$TASK_REVISION/ /qdata/ \u003e /dev/null 2\u003e\u00261 \\\n    | echo Wait for nodes in Quorum network being up ... $count/0\n  sleep 1\ndone\n\nif which jq \u003e/dev/null; then\n  echo \"Found jq\"\nelse\n  echo \"jq not found. Instaling ...\"\n  sudo yum -y install jq\nfi\n\nfor t in $(aws ecs list-tasks --cluster quorum-network-cocroaches-attack | jq -r .taskArns[])\ndo\n  task_metadata=$(aws ecs describe-tasks --cluster quorum-network-cocroaches-attack --tasks $t)\n  HOST_IP=$(echo $task_metadata | jq -r '.tasks[0] | .containers[] | select(.name == \"quorum-run\") | .networkInterfaces[] | .privateIpv4Address')\n  if [ \"EC2\" == \"EC2\" ]\n  then\n    CONTAINER_INSTANCE_ARN=$(aws ecs describe-tasks --tasks $t --cluster quorum-network-cocroaches-attack | jq -r '.tasks[] | .containerInstanceArn')\n    EC2_INSTANCE_ID=$(aws ecs  describe-container-instances --container-instances $CONTAINER_INSTANCE_ARN --cluster quorum-network-cocroaches-attack |jq -r '.containerInstances[] | .ec2InstanceId')\n    HOST_IP=$(aws ec2 describe-instances --instance-ids $EC2_INSTANCE_ID | jq -r '.Reservations[0] | .Instances[] | .PublicIpAddress')\n  fi\n  group=$(echo $task_metadata | jq -r '.tasks[0] | .group')\n  taskArn=$(echo $task_metadata | jq -r '.tasks[0] | .taskDefinitionArn')\n  # only care about new task\n  if [[ \"$taskArn\" == *:$TASK_REVISION ]]; then\n     echo $group | sudo tee /qdata/mappings/ip_$(echo $HOST_IP | sed -e 's/\\./_/g')\n  fi\ndone\n\ncat \u003c\u003cSS | sudo tee /qdata/quorum_metadata\nquorum:\n  nodes:\nSS\nnodes=()\ncd /qdata/mappings\nfor idx in \"${!nodes[@]}\"\ndo\n  f=$(grep -l ${nodes[$idx]} *)\n  ip=$(cat /qdata/hosts/$f)\n  nodeIdx=$((idx+1))\n  script=\"/usr/local/bin/Node$nodeIdx\"\n  cat \u003c\u003cSS | sudo tee $script\n#!/bin/bash\n\nsudo docker run --rm -it quorumengineering/quorum:latest attach http://$ip:22000 $@\nSS\n  sudo chmod +x $script\n  cat \u003c\u003cSS | sudo tee -a /qdata/quorum_metadata\n    Node$nodeIdx:\n      privacy-address: $(cat /qdata/privacyaddresses/$f)\n      url: http://$ip:22000\n      third-party-url: http://$ip:9080\nSS\ndone\n\ncat \u003c\u003cSS | sudo tee /opt/prometheus/prometheus.yml\nglobal:\n  scrape_interval:     15s # By default, scrape targets every 15 seconds.\n\n  # Attach these labels to any time series or alerts when communicating with\n  # external systems (federation, remote storage, Alertmanager).\n  external_labels:\n    monitor: 'monitor'\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it's Prometheus itself.\nscrape_configs:\n- job_name: geth\n  metrics_path: /debug/metrics/prometheus\n  scheme: http\n  static_configs:\n  - targets:\n    - geth:6060\n- job_name: 'node'\n  static_configs:\n  - targets: [ node-exporter:9100 ]\n  file_sd_configs:\n  - files:\n    - 'targets.json'\nSS\n\ncat \u003c\u003cSS | sudo tee /opt/prometheus/docker-compose.yml\n# docker-compose.yml\nversion: '2'\nservices:\n    prometheus:\n        image: prom/prometheus:latest\n        volumes:\n            - /opt/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml\n            - /opt/prometheus/targets.json:/etc/prometheus/targets.json\n        command:\n            - '--config.file=/etc/prometheus/prometheus.yml'\n        ports:\n            - '9090:9090'\n    node-exporter:\n        image: prom/node-exporter:latest\n        ports:\n            - '9100:9100'\n    grafana:\n        image: grafana/grafana:latest\n        volumes:\n            - /opt/grafana/dashboards:/var/lib/grafana/dashboards\n            - /opt/grafana/provisioning/dashboards/all.yml:/etc/grafana/provisioning/dashboards/all.yml\n            - /opt/grafana/provisioning/datasources/all.yml:/etc/grafana/provisioning/datasources/all.yml\n        environment:\n            - GF_SECURITY_ADMIN_PASSWORD=-\u003cND{!FA)tLFDoGB\n        depends_on:\n            - prometheus\n        ports:\n            - '3001:3000'\n    geth:\n        image: ethereum/client-go:latest\n        ports:\n            - '6060:6060'\n        command: --goerli --metrics --metrics.expensive --pprof --pprofaddr=0.0.0.0\n\nSS\n\ncount=$(ls /qdata/privacyaddresses | grep ^ip | wc -l)\ntarget_file=/tmp/targets.json\ni=0\necho '[' \u003e $target_file\nfor idx in \"${!nodes[@]}\"\ndo\n  f=$(grep -l ${nodes[$idx]} *)\n  ip=$(cat /qdata/hosts/$f)\n  i=$(($i+1))\n  if [ $i -lt \"$count\" ]; then\n    echo '{ \"targets\": [\"'$ip':9100\"] },' \u003e\u003e $target_file\n  else\n    echo '{ \"targets\": [\"'$ip':9100\"] }'  \u003e\u003e $target_file\n  fi\ndone\necho ']' \u003e\u003e $target_file\nsudo mv $target_file /opt/prometheus/\n\ncat \u003c\u003cSS | sudo tee /opt/grafana/provisioning/datasources/all.yml\ndatasources:\n- name: 'prometheus'\n  type: 'prometheus'\n  access: 'proxy'\n  org_id: 1\n  url: 'http://prometheus:9090'\n  is_default: true\n  version: 1\n  editable: true\nSS\n\ncat \u003c\u003cSS | sudo tee /opt/grafana/provisioning/dashboards/all.yml\n- name: 'default'\n  org_id: 1\n  folder: ''\n  type: 'file'\n  options:\n    folder: '/var/lib/grafana/dashboards'\nSS\n\nsudo sed -i s'/datasource\":.*/datasource\" :\"prometheus\",/' /opt/grafana/dashboards/dashboard-geth.json\nsudo sed -i s'/datasource\":.*/datasource\" :\"prometheus\",/' /opt/grafana/dashboards/dashboard-node-exporter.json\nsudo /usr/local/bin/docker-compose -f /opt/prometheus/docker-compose.yml up -d --force-recreate\n",
                            "filename": "/tmp/.terranova273240257/generated-bootstrap.sh",
                            "id": "09bccb7dc18c01e3a1c96853e9a3c885cb472081"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.local"
                },
                "local_file.private_key": {
                    "type": "local_file",
                    "depends_on": [
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "0619bff6918e2e978bb10f0b57b549297a9eaad0",
                        "attributes": {
                            "content": "---DUMMY PRIVATE KEY---",
                            "filename": "/tmp/.terranova273240257/quorum-cocroaches-attack.pem",
                            "id": "0619bff6918e2e978bb10f0b57b549297a9eaad0"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.local"
                },
                "null_resource.bastion_remote_exec": {
                    "type": "null_resource",
                    "depends_on": [
                        "aws_ecs_task_definition.quorum",
                        "aws_instance.bastion",
                        "local_file.bootstrap",
                        "tls_private_key.ssh"
                    ],
                    "primary": {
                        "id": "354840277283586537",
                        "attributes": {
                            "id": "354840277283586537",
                            "triggers.%": "3",
                            "triggers.bastion": "",
                            "triggers.ecs_task_definition": "2",
                            "triggers.script": "3972838542e1726b0d045ff3d42cc764"
                        },
                        "meta": {},
                        "tainted": true
                    },
                    "deposed": [],
                    "provider": "provider.null"
                },
                "random_id.bucket_postfix": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "6Ma2xnZmArE",
                        "attributes": {
                            "b64": "6Ma2xnZmArE",
                            "b64_std": "6Ma2xnZmArE=",
                            "b64_url": "6Ma2xnZmArE",
                            "byte_length": "8",
                            "dec": "16773294825694167729",
                            "hex": "e8c6b6c6766602b1",
                            "id": "6Ma2xnZmArE"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_id.code": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "3M--hw",
                        "attributes": {
                            "b64": "3M--hw",
                            "b64_std": "3M++hw==",
                            "b64_url": "3M--hw",
                            "byte_length": "4",
                            "dec": "3704602247",
                            "hex": "dccfbe87",
                            "id": "3M--hw"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_id.ethstat_secret": {
                    "type": "random_id",
                    "depends_on": [],
                    "primary": {
                        "id": "tEAgROn3tc3oGqT3CgVSqA",
                        "attributes": {
                            "b64": "tEAgROn3tc3oGqT3CgVSqA",
                            "b64_std": "tEAgROn3tc3oGqT3CgVSqA==",
                            "b64_url": "tEAgROn3tc3oGqT3CgVSqA",
                            "byte_length": "16",
                            "dec": "239594000737262924426798755770955813544",
                            "hex": "b4402044e9f7b5cde81aa4f70a0552a8",
                            "id": "tEAgROn3tc3oGqT3CgVSqA"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_integer.network_id": {
                    "type": "random_integer",
                    "depends_on": [],
                    "primary": {
                        "id": "4856",
                        "attributes": {
                            "id": "4856",
                            "keepers.%": "1",
                            "keepers.changes_when": "cocroaches-attack",
                            "max": "9999",
                            "min": "2018",
                            "result": "4856"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "random_string.random": {
                    "type": "random_string",
                    "depends_on": [],
                    "primary": {
                        "id": "none",
                        "attributes": {
                            "id": "none",
                            "length": "16",
                            "lower": "true",
                            "min_lower": "0",
                            "min_numeric": "0",
                            "min_special": "0",
                            "min_upper": "0",
                            "number": "true",
                            "result": "-\u003cND{!FA)tLFDoGB",
                            "special": "true",
                            "upper": "true"
                        },
                        "meta": {
                            "schema_version": "1"
                        },
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.random"
                },
                "tls_private_key.ssh": {
                    "type": "tls_private_key",
                    "depends_on": [],
                    "primary": {
                        "id": "e28f0d026fcd4faea3dd1ae386c7f918484f1273",
                        "attributes": {
                            "algorithm": "RSA",
                            "ecdsa_curve": "P224",
                            "id": "e28f0d026fcd4faea3dd1ae386c7f918484f1273",
                            "private_key_pem": "---DUMMY PRIVATE KEY---",
                            "public_key_fingerprint_md5": "44:43:cb:61:29:71:33:19:21:23:d3:69:3a:b2:3e:00",
                            "public_key_openssh": "dummyKeySsh",
                            "public_key_pem": "---DUMMY PUBLIC KEY---",
                            "rsa_bits": "2048"
                        },
                        "meta": {},
                        "tainted": false
                    },
                    "deposed": [],
                    "provider": "provider.tls"
                }
            },
            "depends_on": []
        }
    ]
}`
	YamlV01Fixture = `version: 0.1
resourceType: variables
variables: 
  simpleKey: variable
`
	YamlV02Fixture = `version: 0.2
resourceType: variables
variables: 
  network_name: variable`

	YamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  region: "us-east-1"
  network_name: "dum-net"
  number_of_nodes: '5'
  quorum_docker_image_tag: '2.2.5'
  faketime: ["+2d", "-30h", "+120", "-120", "120", "44", "-44", "+44", "+0", "-0", "+0s", "-0h"]`

	IncorrectSignYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  faketime: ["@2s"]`

	IncorrectNodesNumberYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  number_of_nodes: 'Abc'`

	IncorrectQuorumDockerImageTagYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  quorum_docker_image_tag: '2.X.5'`

	IncorrectNetworkNameStringYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  network_name: "d$#sa!"`

	IncorrectNetworkNameLengthYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  network_name: 'dummy-network-dummy-network'`

	IncorrectStringVariablesYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  region: 'dummy region 2'`

	IncorrectAwsRegionType = `version: 0.3
resourceType: variables
variables: 
  region: 'kaz-uk-2'`

	IncorrectAwsInstanceType = `version: 0.3
resourceType: variables
variables: 
  asg_instance_type: 'someInvalid-instance-type.large'`

	IncorrectConsensusMechanismType = `version: 0.3
resourceType: variables
variables: 
  consensus_mechanism: 'brambory'`

	IncorrectUnitYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  faketime: ["+2x"]`

	IncorrectValueYamlV03Fixture = `version: 0.3
resourceType: variables
variables: 
  faketime: ["+1A24s"]`

	YamlFixtureConfigurable = `version: %v
resourceType: %s
variables:
  simpleKey: variable
`

	YamlFixtureNoVariables = `version: 0.1
resourceType: variables
`
	YamlFixtureGasLimitGreaterThanMinGasLimit = `version: 0.3
resourceType: variables
variables: 
  genesis_min_gas_limit: 20
  genesis_gas_limit: 21
`
	YamlFixtureGasLimitWithoutMinGas = `version: 0.3
resourceType: variables
variables: 
  genesis_gas_limit: '0x58f7'
`
	YamlFixtureGasLimitLowetHanMinGasLimit = `version: 0.3
resourceType: variables
variables: 
  genesis_min_gas_limit: 21
  genesis_gas_limit: 20
`

	YamlFixtureWithHexUtils = `version: 0.1
resourceType: variables
variables:
  simpleKey: variable
  region:                'us-east-1'     ## You can set region for deployment here
  default_region:        'us-west-1'     ## If key region is not present it is default region setter
  profile:               'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  aws_access_key_id:     'dummyValue'    ## It overrides access key id env variable. If omitted system env is used
  aws_secret_access_key: 'dummyValue'    ## It overrides secret access key env variable. If omitted system env is used
  genesis_gas_limit:      25		     ## Used to set genesis gas limit
  genesis_timestamp:      38	         ## Used to set genesis timestamp
  genesis_difficulty:     12             ## Used to set genesis difficulty
  genesis_nonce:          0              ## Used to set genesis nonce
  consensus_mechanism:    "instanbul"    ## Used to set consensus mechanism supported values are raft/istanbul
`

	YamlFixtureWithInvalidHexUtils = `version: 0.3
resourceType: variables
variables:
  simpleKey: variable
  region:                'us-east-1'     ## You can set region for deployment here
  default_region:        'us-west-1'     ## If key region is not present it is default region setter
  profile:               'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  aws_access_key_id:     'dummyValue'    ## It overrides access key id env variable. If omitted system env is used
  aws_secret_access_key: 'dummyValue'    ## It overrides secret access key env variable. If omitted system env is used
  genesis_gas_limit:     'invalidVal'    ## Used to set genesis gas limit
  genesis_timestamp:     '3n8'	         ## Used to set genesis timestamp
  genesis_difficulty:     ['12']         ## Used to set genesis difficulty
  genesis_nonce:          {}	         ## Used to set genesis nonce
  consensus_mechanism:    "instanbul"    ## Used to set consensus mechanism supported values are raft/istanbul
`

	NoSuchFileOrDirectoryMsg = "\n[ERR] Yaml Validation error: open %s: no such file or directory"
	NotValidExtMsg           = "\n[ERR] Yaml Validation error: %s is not in supported file types. Valid are: [.yml .yaml]"
	DummyRecipeBodyFail      = `variable "count"    { default = 2 }
  variable "key_name" {}
  variable "region" {}
  provider "aws" {
    region        =  "${var.region}"
  }
  resource "aws_instance" "server" {
    instance_type = "t2.micro"
    ami           = "ami-6e1a0117"
    count         = "${var.count}"
    key_name      = "${var.key_name}"
  }`
	OutputAsAStringWithoutHeaderFixture   = "_status = Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n\nbastion_host_dns = \nbastion_host_ip = invalid.ip.666\nbucket_name = us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\nchain_id = 4856\necs_cluster_name = quorum-network-cocroaches-attack\ngrafana_host_url = http://invalid.ip.666:3001\ngrafana_password = -<ND{!FA)tLFDoGB\nnetwork_name = cocroaches-attack\nprivate_key_file = /tmp/.terranova273240257/quorum-cocroaches-attack.pem"
	OutputAsAStringFromMultipleValueTypes = "_status = Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n\nbastion_host_dns = [\"\", \"\"]\nbastion_host_ip = {\"ip\": \"invalid.ip.666\"}\nbucket_name = us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\nchain_id = 4856\necs_cluster_name = quorum-network-cocroaches-attack\ngrafana_host_url = http://invalid.ip.666:3001\ngrafana_password = -<ND{!FA)tLFDoGB\nnetwork_name = cocroaches-attack\nprivate_key_file = /tmp/.terranova273240257/quorum-cocroaches-attack.pem"
	OutputAsAStringWithInvalidValues      = "_status = Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n\nbastion_host_dns = [\"\", \"\"]\nbastion_host_ip =\nbucket_name = us-east-2-ecs-cocroaches-attack-e8c6b6c6766602b1\nchain_id = 4856\necs_cluster_name = quorum-network-cocroaches-attack\ngrafana_host_url = http://invalid.ip.666:3001\ngrafana_password = -<ND{!FA)tLFDoGB\nnetwork_name = \nprivate_key_file = /tmp/.terranova273240257/quorum-cocroaches-attack.pem"
	PrivateKeyPairBody                    = "---DUMMY PRIVATE KEY---"
	PublicKeyPairBody                     = "---DUMMY PUBLIC KEY---"
	OpenSshKeyBody                        = "dummyKeySsh"
	DefaulStateFileName                   = "terraform.tfstate"
	DefaultStateFileBody                  = `{
    "version": 3,
    "terraform_version": "0.11.13",
    "serial": 1,
    "outputs": {},
    "resources": []
}`
)
