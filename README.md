# Protochannel

Protochannel is a prototype project to demonstrate the usage of Ethereum state channel
to do off-chain transactions settleabe at any point in time forward manner.

It's a work in progress TicTacToe implementation.

## How it works

Player interacts through the Golang API services by pushing signed
message (personalSign) with MetaMask.
They exchange signed message with state composed
of memory, nonce, counter, state function transition.

State:

| Offset    | Byte length  | Description |
|:----------:|:----------:|-------------|
| `0` | 1 | Opcode for the state transition function: *NOP*(0), *PUT(x,y)*(1) |
| `1` | 1 | Operation counter: nonce |
| `2` | 1 | Parameter 1|
| `3` | 1 | Parameter 2|
| `4` | 9 | Tic Tac Toe Board game memory |

State transition is enforced by games rules in Solidity Smart Contract

TODO:

 * Specifications
 * Messaging
 * Settlement
 * Smart contract
 * Documentation
 * docs/protochannel.yml
 * Tests
 * Use libp2p

## How to run the prototype ?

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

The API server log shows when event are fired by the smart contract
The browser console log contains various debug informations

## Environment variables
*WS_URI* is the url of a capable WebSocket RPC node, if private blockchain, the node networkid must match genesis chainid for Metamask to work correctly
  - WS_URI=ws://1.1.3.7:8546

*DB_DSN* is the DSN of a PostgreSQL database (embedded with docker-compose)
  - DB_DSN=host=postgres user=f3b1be969686afb4520ce dbname=protochannel sslmode=disable password=839169c0ea5e59146a

*PRIVATE_KEY* is a private key in hex format, used if needed to emit transaction to or deploy smart contract
  - PRIVATE_KEY=8e973e8e3d06a321d7285695560a45bd4946eb6f59410732b048c59a0174cc09 # 0xcd5bb3a1E6d7676DAD570566A455dC38b7e3EaDf

*CONTRACT_ADDRESS* is the contract address of the deployed contract that the service should interact with, empty will trigger a new deployment
  - CONTRACT_ADDRESS=

*RETRY* is the number of retry the service will perfom before giving up
  - RETRY=3

## FAQ

**How to generate API server ?**
```
swagger generate server -f docs/protochannel.yml
```
See: [Go Swagger][2]

**How to generate API Smart Contract ABI ?**
```

abigen -sol contracts/tictactoe.sol --pkg internal -out internal/tictactoe.go
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

## Resources

- Martin Koeppelmann (Oct. 2015, blog) - [How offchain trading will work](http://forum.groupgnosis.com/t/how-offchain-trading-will-work/63)
- Robert Mccone (Oct. 2015, blog) - [Ethereum Lightning Network and Beyond](http://www.arcturnus.com/ethereum-lightning-network-and-beyond/)
- Jeff Coleman (Nov. 2015, blog) - [State Channels](http://www.jeffcoleman.ca/state-channels/) (see also: [discussion on /r/ethereum](https://www.reddit.com/r/ethereum/comments/3tcu82/state_channels_an_explanation/))
- Heiko Hees (Dec. 2015, talk) - [Raiden: Scaling Out With Offchain State
Networks](https://www.youtube.com/watch?v=h791zjvf3uQ)
- Jeff Coleman (Dec. 2015, interview) - [Epicenter Bitcoin: State Networks](https://www.youtube.com/watch?v=v0ZJDsRYnbA)
- Jehan Tremback (Dec. 2015, blog) - [Universal Payment Channels](http://altheamesh.com/blog/universal-payment-channels/)
- Martin Koeppelmann (Jan. 2016, slides) - [Scalability via State Channels](http://de.slideshare.net/MartinKppelmann/state-channels-and-scalibility)
- Vitalik Buterin (Jun. 2016, paper) - [Ethereum: Platform Review (page 30)](http://static1.squarespace.com/static/55f73743e4b051cfcc0b02cf/t/57506f387da24ff6bdecb3c1/1464889147417/Ethereum_Paper.pdf)
- Dennis Peterson (July 2016, code) - [Sparky: A Lightning Network in Two Pages of Solidity](http://www.blunderingcode.com/a-lightning-network-in-two-pages-of-solidity/)
- Ameen Soleimani (Sept. 2016, talk) - [An Introduction to State Channels in
Depth](https://www.youtube.com/watch?v=MEL50CVOcH4)
- Jeff Coleman (ongoing, wiki) - [State Channels Wiki](https://github.com/ledgerlabs/state-channels/wiki)
- Dennis Peterson (Oct. 2016, code) - [Gamble Channels: Fast Verifiable Off-Chain Gambling](http://www.blunderingcode.com/gamble-channels-fast-verifiable-off-chain-gambling/)
- Jeff Coleman (ongoing, code) - [Toy State Channels](https://github.com/ledgerlabs/toy-state-channels/tree/master/contracts)
- Heiko Hees (ongoing, code) - [Raiden Network](https://github.com/raiden-network/raiden)
- Sergey Ukustov (ongoing, code) - [Machinomy](https://github.com/machinomy/machinomy)

Source: https://www.reddit.com/r/ethereum/comments/6fde8t/ethereum_payment_channels_in_50_lines_of_solidity/dihe3tj/
