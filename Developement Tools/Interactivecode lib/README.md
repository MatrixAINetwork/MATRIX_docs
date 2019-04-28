
### Account Address Generation Rules of MATRIX chain

+ Create a random private key (32 bytes)
+ ECDSA-secp256k1 (Elliptic Curve Digital Signature Algorithm) is used to map the private key to public key (64 bytes)
+ The last 20 bytes of the public key will be taken as basic address
+ The basic address will be Base58 encoded and then a pre-fix of ‘MAN.’ will be added
+ CRC8 check will be done on combined data and the check bit will be put at the end of the address to generate the MATRIX account


The MATRIX address looks like this:

    MAN.2uTgkPiGX9ziKuAKyeeWBt8duiBRH 



### How to Create an account - Two ways

+ From gman: `“personal.newAccount”`api, where keystore will be saved under `<gman path>\chaindata\keystore`
+ From Offical Web wallet (https://wallet.matrix.io/) or App

### How to View account balance

+ `man.getBalance` api

### Recharge

You need to know the following:

+ Only one mainnet exists for MATRIX blockchain. There’s no side chain. Fork will not happen.
+ All transactions recorded on the MATRIX blockchain are tamper-resistant, that is, once confirmed, the recharge will be successful.
+ MATRIX wallet account not only contains MAN assets, but also many other types of full assets issued by users (such as equity, Token, etc.). When recording users' recharges, it is necessary to determine the asset type of recharge, so as to avoid mistaking the recharge of other assets as MAN currency or confusing the recharge of MAN currency and Token.
+ MATRIX full nodes must be online in order to synchronize blocks. The block height can be viewed via the "man.blockNumber" API of the GMAN client.
+ Transfer between internal users in the Exchange does not need to go through the block chain. You can directly modify the user balance in the database. Only recharge and withdraw need to be recorded on the chain.

### Recharge History
You need to prepare codes to monitor each transaction in each block and record all recharge and withdraw transactions in the database. When there is a new recharge transaction, the user balance in the database should be modified accordingly.

The man.getblock (Number) method in GMAN API provides the function of obtaining block information. ‘Number’ in this method refers to block height.

The block information contains transaction details, and you need to record all your transactions as proof of user recharge and withdraw. If your address (which belongs to Exchange) is found in the output of the transaction, the MAN or Token balance corresponding to this recharge address should be modified in the database.

There is also another method: if an address belonging to the Exchange is found in the output of the transaction, record the recharge data in the database first, and then modify the user balance after several confirmations. This is not recommended.


### Transaction 

MATRIX provides demo projects of JS, GO and Java versions for transaction sending. Please get it from GitHub ([demos](https://github.com/MatrixAINetwork/MATRIX_docs/tree/master/Developement%20Tools/Interactivecode%20lib))

Transaction sending involves the following steps:

+ Unlock with keystore and the password
+ Obtain serial nonce of transaction’s ‘From’ account
+ Construct the transaction and remember to distinguish between ordinary transactions and contract transactions
+ Sign the transaction
+ Send the signed transaction data
+ Get the transaction receipt, confirm the transaction status and view the account balance


### Appendix

| API | blockNumber |
|-----------------|-------------------------------------------------|
| Description | Get the current block height |
| Input Parameter | N/A |
| Output | err: error message  blocknumber: current height |
| Example | blockNumber,err := man.GetBlockNumber() |

| API | GetTransactionCount |
|-----------------|--------------------------------------------------------------------|
| Description | Get current nonce |
| Input Parameter | address: source address  string: "latest", "earliest" or "pending" |
| Output | *big.int: transaction number of a specified address |
| Example | nonce, err := man.GetTransactionCount(from, "latest") |

| API | GetBalance |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Description | Get the balance of specified account address at a specified block height |
| Input Parameter | String: the address whose balances is to be obtained  Number|String - （optional ）If this value is not set, use the block set by man.defaultBlock, otherwise use the specified block |
| Output | err: error message  []common.RPCBalanceType: BigNumber instance that contains the current balance of specified adder, in wei |
| Example | balances,err := man.getBalance("MAN.47kRUvvpaQPHz3faEvesFcLpdYSim") |


| API | SendRawTransaction |
|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Description | Send a signed transaction |
| Input Parameter | type SendTxArgs1 struct {        From     string          `json:"from"`         To       *string         `json:"to"`         Gas      *hexutil.Uint64 `json:"gas"`          GasPrice *hexutil.Big    `json:"gasPrice"`         Value    *hexutil.Big    `json:"value"`           Nonce    *hexutil.Uint64 `json:"nonce"`           Data        *hexutil.Bytes `json:"data"`           Input       *hexutil.Bytes `json:"input"`           V           *hexutil.Big   `json:"v"`           R           *hexutil.Big   `json:"r"`           S           *hexutil.Big   `json:"s"`           Currency    *string        `json:"currency"`           TxType      byte           `json:"txType"`     //           LockHeight  uint64         `json:"lockHeight"` //           IsEntrustTx byte           `json:"isEntrustTx"`           CommitTime  uint64         `json:"commitTime"`           ExtraTo     []*ExtraTo_Mx1 `json:"extra_to"` //     }     Object – the transaction object to be sent     from: String – the specified sender’s address. If not specified, use man.defaultAccount     [if !supportLists]·           [endif]to: String – target address of transaction. If created by contract, leave it blank     [if !supportLists]·           [endif]value: Number|String|BigNumber - The amount of currency carried in the transaction, in Wei. If the transaction is created via contract, it is the initial fund     [if !supportLists]·           [endif]gas: Number|String|BigNumber - The gas used by the transaction; unused gas will be returned     [if !supportLists]·           [endif]gasPrice: Number|String|BigNumber – gas price of transaction     [if !supportLists]·           [endif]data: String – (optional) byte string containing relevant data. If created by contract, it is the code used for initialization.     [if !supportLists]·               [endif]nonce: Number - integer, using this value allows you to override your own pending transactions with the same nonce.     [if !supportLists]·           [endif]V ：signature result      [if !supportLists]·           [endif]R ：signature result      [if !supportLists]·           [endif]S ：signature result      [if !supportLists]·           [endif]Currency: type of currency     [if !supportLists]·           [endif]TxType ：transaction types (normal transactions are set as 0)      [if !supportLists]·           [endif]LockHeight: a reserved field [if !supportLists]·           [endif]IsEntrustTx  ：0-self-paid gas ，1- gas paid by others      [if !supportLists]·           [endif]CommitTime: commit time, available to scheduled transactions and revocable transactions only     ExtraTo: extended transaction (should be filled for one to many transactions) |
| Output | String - 32-byte transaction hash string in hexadecimal format |
| Example | txhash, err := man.SendRawTransaction(raw) |
