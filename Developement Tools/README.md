| 交互代码库（语言） |                    |                                        |                                                          |
|--------------------|--------------------|----------------------------------------|----------------------------------------------------------|
|                    | js                 | 库和示例                               | 主要功能全部实现                                         |
|                    | go                 | 库和示例                               | 主要功能全部实现                                         |
|                    | java               | 库和示例                               | 主要功能全部实现                                             |
| dapp开发者工具     | docker环境单节启动 | 一键启动，跨平台，钱包，浏览器均可     | 使用手册                                                 |
|                    | Truffle开发框架    | 可以编译调试，不能一键部署到主链       | 核心是一键部署到matrix主链，协调truffle，remix，metamask |
|                    | remix ide          | 可以编译调试，需要metamask插件进行部署 | 核心是一键部署到matrix主链，协调truffle，remix，metamask |
|                    | 编写合约           | solidity官方手册                       | https://solidity.readthedocs.io/en/v0.5.7/#              |
| 钱包对外接口       | 云钱包节点         | rpc接口：api85.matrix.io               |                                                          |
|                    | app软件调用        | 实例                                   | 使用手册                                                 |




## Developement Tools
### Developer Smart Contracts
#### Smart Contract Languages
+ [Solidity](https://solidity.readthedocs.io/en/latest/)- EVM smart contracting language


#### Frameworks
+ [Truffle](https://truffleframework.com/) - Most popular smart contract development, testing, and deployment framework. The Truffle suite includes Truffle, and Drizzle and Ganache
+ [Dapp](https://dapp.tools/dapp/) - Framework for DApp development, successor to DApple
#### IDEs
+ [Remix](https://remix.ethereum.org/#optimize=false) - Web IDE with built in static analysis, test blockchain VM.

+ [Pragma](https://www.withpragma.com/) - Very simple web IDE for solidity, and auto-generated interfaces for smart contracts. autocomplete-solidity, and language-solidity packages

+ [Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=JuanBlanco.solidity) - Visual Studio Code extension that adds support for Solidity


### Transaction APIs

+ [Address format](https://github.com/MatrixAINetwork/TxSend-Sign-Demos/blob/master/Address%20Format.md)
+ [js SDK for transactions](https://github.com/MatrixAINetwork/TxSend-Sign-Demos/tree/master/js)
+ [Golang SDK for transactions](https://github.com/MatrixAINetwork/TxSend-Sign-Demos/tree/master/go)
+ [Java SDK for transactions](https://github.com/MatrixAINetwork/TxSend-Sign-Demos/tree/master/java)

## Smart Contracts
### Smart Contracts Overview
Smart contracts are autonomous agents deployed on a blockchain which will execute code and actions depending on given conditions.

### Usecases for smart contracts
Function as ‘multi-signature’ accounts, so that funds are spent only when a required percentage of people agree
Manage agreements between users, say, if one buys insurance from the other

Provide utility to other contracts (similar to how a software library works)

Store information about an application, such as domain registration information or membership records.
Solidity language

Solidity is a high-level programming language, that is used to implement smart contracts on MATRIX MAINNET.

It has Python, C++ and JavaScript influences and is used for Ethereum Virtual Machine ( EVM ). As a result, it is fairly convenient and easy to grasp for those that are already familiar with the Python, C++ or JavaScript. 





