### ACCOUNT PASSWORDS:
- `apollo`

### AUTHORITIES_ACCOUNTS:
- `0x80c79f22b0a4d1dd5fabd2149b0c2d366a41fd84`
- `0x185f74d0c50a1b2a56c4dc305d723e2a50f816df`

### User Accounts
- `0xadd97b160487fce7032d722d8564e287065c155e`

### Gather public enode (WIP) 
5. connecting the nodes
https://wiki.parity.io/Demo-PoA-tutorial.html
TODO: PUSH IT TO bin/cli
curl --data '{"jsonrpc":"2.0","method":"parity_enode","params":[],"id":0}' -H "Content-Type: application/json" -X POST node1:8545
// result of line 13
curl --data '{"jsonrpc":"2.0","method":"parity_addReservedPeer","params":["enode://RESULT"],"id":0}' -H "Content-Type: application/json" -X POST localhost:8541

