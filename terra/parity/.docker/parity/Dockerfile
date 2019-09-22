FROM parity/parity:latest
WORKDIR /var/app
ADD config/genesis.json /var/config/genesis.json
ADD config/node.pwds /var/config/node.pwds
ADD config/node1-config.toml /var/config/node1-config.toml
ADD config/node2-config.toml /var/config/node2-config.toml
ADD config/authorities /var/app/keys/ApolloPOA
ENTRYPOINT ["parity"]