#!/bin/bash 

set -x

function run_folder {
   local  dir=$1 f=
   cd $dir/run
   # sequential execution
   for f in */*.sh ; do
       bash $f -H
   done
}

# parallel execution 
for j in * ; do
   run_folder $j &
done
wait