 
jmeter/apache-jmeter-5.2.1/bin/./jmeter -n -t deploy.jmx \
    -Jurl=54.144.244.216    -Jport=22000 -Jaccount=0xe822d7a7001e847d0d78761719be30953b7aca71 \
    -Jthread=10 -Jloop=10000 -Jjmeter.save.saveservice.output_format=xml \
    -Jjmeter.save.saveservice.response_data=true -Jjmeter.save.saveservice.samplerData=true  \
    -Jjmeter.save.saveservice.requestHeaders=true -Jjmeter.save.saveservice.url=true     \
    -Jjmeter.save.saveservice.responseHeaders=true



cat chainhammer/hammer/config.py | grep RPCaddress= | cut -f2 -d"=http://" | sed -e 's/^[ \t]*//'

Install java 

sudo apt-get -y update
sudo apt-get -y install default-jre

web3.fromWei(eth.getBalance(eth.accounts[0]), "ether")