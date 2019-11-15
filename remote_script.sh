 #!/bin/bash -e
 
 set -x

 # Clearing previous run results
cd chainhammer
scripts/install-initialize.sh
rm -rf results/runs/*
rm -rf reader/img/*
# Fire the Hammer 
echo "Summoning Thunder!!!"
CH_TXS=25000 CH_THREADING="threaded2 300" ./run.sh "$1"
echo "Creating remote copy of TPS"
cat results/runs/*.md | grep  peakTpsAv > results-$1.txt  |  sed -e 's/^[ \t]*//'
cat results/runs/*.md | grep  finalTpsAv >> results-$1.txt  |  sed -e 's/^[ \t]*//'
echo "Finished Remote Execution"