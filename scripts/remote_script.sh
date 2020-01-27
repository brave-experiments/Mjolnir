#!/bin/bash 
# @dev This script is responsible for executing chainhammer from the host machine .

set -x

cd chainhammer
# ./scripts/install-initialize.sh 
rm results/runs/*
rm reader/img/*
CH_TXS=25000 CH_THREADING="threaded2 25" ./run.sh "$@1"
# cd ~/



