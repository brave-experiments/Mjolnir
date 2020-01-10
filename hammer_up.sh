title is_up.py
echo Loops until the node is answering on the expected port.
./is_up.py
echo Great, node is available now.
echo 

title tps.py
echo start listener tps.py, show here but also log into file $TPSLOG
echo this ENDS after send.py below writes a new INFOFILE $INFOFILE
unbuffer ./tps.py | tee "../$TPSLOG" &
echo

