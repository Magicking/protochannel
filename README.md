# Protochannel repository

**How to run the prototype ?**

1. Install Docker and Docker Compose (Window 10, macOS, Linux)
2. Install git
3. Place yourself in the repository containing this file
4. Clone this repository with vendors using git
```git clone --recurse-submodules https://github.com/Magicking/protochannel```
5. Build the containers using docker-compose
```docker-compose build```
6. Setup environment variables in docker-compose.yml, see below
7. Run the prototype
```docker-compose up```
8. Launch your browser and unlock your account in your favorite wallet software
Tested with:
  - Metamask
8. Access interface at:
  - http://127.0.0.1:8000

Clicking on the button will trigger in this order:
  - Sign the message "TOTO" with your ethereum key
  - Post the message with the signature to the API server
  - The API server recover the address
  - Check IsListed with the smart contract by calling a RPC to the connected node

The API server log shows when event are fired by the smart contract
The browser console log contains various debug informations

**Environment variables*
*WS_URI* is the url of a capable WebSocket RPC node
  - WS_URI=ws://1.1.3.7:8546

*DB_DSN* is the DSN of a PostgreSQL database (embedded with docker-compose)
  - DB_DSN=host=postgres user=f3b1be969686afb4520ce dbname=protochannel sslmode=disable password=839169c0ea5e59146a

*PRIVATE_KEY* is a private key in hex format, used if needed to emit transaction to smart contract
  - PRIVATE_KEY=8e973e8e3d06a321d7285695560a45bd4946eb6f59410732b048c59a0174cc09 # 0xcd5bb3a1E6d7676DAD570566A455dC38b7e3EaDf

*CONTRACT_ADDRESS* is the contract address of the deployed contract that the service should interract with
  - CONTRACT_ADDRESS=0x59b9310dff261526b9526b0987ced3c9fcbd40e1 #todo deploy

*RETRY* is the number of retry the service will perfom before giving up
  - RETRY=3

**How to generate API server ?**
```
swagger generate server -f docs/protochannel.yml
```
See: [Go Swagger][2]

**How to generate API Smart Contract ABI ?**
```
abigen --sol contract/whitelist.sol --pkg internal --out internal/whitelist.go
```
See: [Abigen][3]

**How to install Docker and Docker Compose?**
See: [Docker Installation][1]
See: [Docker Compose Installation][4]

**How to install Git ?**
See: [Installing Git][5]

[1]: https://docker.github.io/engine/installation/
[2]: https://goswagger.io/
[3]: https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts#generating-the-bindings
[4]: https://docs.docker.com/compose/install/
[5]: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
