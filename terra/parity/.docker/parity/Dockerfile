FROM parity/parity:latest
WORKDIR /var/parity
ADD config/genesis.json /var/config/genesis.json
ADD config/node1-config.toml /var/config/node1-config.toml
ADD config/node2-config.toml /var/config/node2-config.toml
ADD config/authorities $HOME/.local/share/io.parity.ethereum/keys/ethereum
ENTRYPOINT ["parity"]