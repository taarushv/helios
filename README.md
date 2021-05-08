# Helios

**Note**: working on v0.2, check back soon!

Helios is a tool built on-top of [go-ethereum](https://github.com/ethereum/go-ethereum) and the [ELK stack](https://www.elastic.co/what-is/elk-stack) to query and monitor the mempool. 

The goal of this project is to create a unified JSON API to classify various trades, transfers, loans, liquidations, etc. and seamlessly interact with the pool (before the transactions are mined). Hoping to extend this to include local EVM execution + atomic on-chain arbitrage + dynamic gas adjustment to capture said value. 


Currently supports the following txs: 
* ETH direct transfers
* ERC20 approves and transfers
* Contract deploys
* Uniswap trades
* Chainlink oracle updates


## Why?
TL;DR [Miner extractable Value](https://arxiv.org/abs/1904.05234)

Different kinds of MEV: 
* Front-running
	* Ex: If token A is trading for a different price on two different markets (cross-DEX or just against a centralized exchange), arbitrageurs can profit by bringing the prices back to equilibrium. The first person to have their tx mined makes money by frontrunning the rest. 
* Back-running
	* Ex: If a position (either borrowing against collateral or perpetual futures) is close to it's liquidation price, one can make money by executing their tx right behind an oracle update that renders the position underwater. 
* Re-org
	* The chaotic kind of MEV, miners could re-org the entire chain up to 'n' number of blocks if the net MEV > block rewards. Ex: If ETH goes from 100$ to 200$, a miner can maliciously pick a block where a uniswap pool (say ETHUSDC) has liquidity priced below current market rate, buy from the pool, and exclude any other txs that'd cause a state change to the liquidity pool (non-trivial to pull off). 

However, there's not enough flexible tools to work with the mempool. 
## How does it work?
We subscribe to new txs and new blocks received by the client. Each tx is classified based on it's function signature or the "to" address. We check each tx against our pre-built filters and mark them as mined accordingly. Similarly, after each block is mined, we look for reordering taint (deviation from the standard go-ethereum client rules wrt gas prices and priority in block) to also see if the miner is potentially engaging in extracting MEV. We also note for transactions that get mined without appearing in our mempool (either because our peer simply didn't receive the tx or the miner doesn't broadcast the tx before mining it). 

Written in Go to directly integrate with Geth (via IPC socket) and it's TxPool. It leverages elastic search as a document store for fast queries and the underlying search engine. Elastic search exposes a simple JSON API clients can query to formulate strategies. 

Example: Query to list all chainlink oracle update txs in the mempool (along with relevant parsed data such as pair, current price, and oracle submission) 

![enter image description here](https://i.imgur.com/KToewcG.png)
## Quick start

![](helios.gif)

Helios also supports infura so the fastest way to get up and running would be generating a infura API key, creating a .env file (with INFURA_WS_URL="wss://mainnet.infura.io/ws/v3/....") and doing the following in the root folder: 
	
`go build && ./helios -client=infura -mode=quick`

"quick" mode does not require/initiate a elastic search client (no inserts, just the console preview)

[More instructions on using infura](https://github.com/taarushv/helios/issues/2)

## Complete Guide
Note: Publishing a docker image to make this seamless and easy soon!

Follow [this](https://www.digitalocean.com/community/tutorials/how-to-install-elasticsearch-logstash-and-kibana-elastic-stack-on-ubuntu-20-04) guide to setup elastic search on the same instance as your local node. 

If you intend to use this with your local node, add `GETH_IPC_PATH="/home/userName/.ethereum/geth.ipc"` to your .env file. 

Also highly recommend using the following command to take advantage of optimized config. Increased the number of default peers and max transactions held in the mempool before being dropped: 

 `geth --config optimized_geth_config.toml`
 
 ./helios:
 * -mode=quick vs -mode=full
	* use full mode after you setup elastic and make sure the inserts work. quick is light weight and fast but the JSON API exposed by -mode=full is very useful. 
 * -client=local vs -client=infura
	* see above
 * -flush=indexName
	* `./helios -flush=transactions` and `./helios -flush=blocks` would erase the respective indexes and the documents within. This overrides other flags, be careful!
	
## TODO:
 * Source tx.Time() directly from the client by modifying the API (more accurate than doing it on the fly)
 * Maker and dydx use the same oracles (dydx has high incentives for perp futures liquidators too) so adding support to them should be next. Along with all other relevant 
 * Integrate everything into a single client build (alongside geth)
 * Multi-trade arbitrage opportunities and taking advantage of ETH atomicity + flash loan liquidity. 

Apart from adding more filters, Function signatures and to/from addresses don't go far enough because most sophisticated arb bots use complex mechanisms (create2=>execute=>self destruct) to obscure their strategies in order to avoid being front run. Currently experimenting on a frontrunner prototype (`panther/src/index.js`, WIP) that can execute transactions locally on a ganache-fork (updated every time a new block is mined) to access events emitted before a tx is even mined. Building a agnostic frontrunner + backrunner (that simply takes in txHash it needs to outrun/backrun, while having a dynamic mechanism to keep bidding going until the opportunity is no longer worthwhile). 
