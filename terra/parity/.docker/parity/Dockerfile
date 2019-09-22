FROM parity/parity:latest
WORKDIR /var/parity
ADD config/genesis.json /var/config/genesis.json
ADD config/node-config.toml /var/config/node-config.toml
ENTRYPOINT ["parity"]