# Guide to Blockchain Explorer

You can use the block explorer at： http://tom.matrix.io/

## Blockchain Platform

![](https://i.imgur.com/rqnruo5.png)

### Block Search

Information on all blocks—from the genesis block to all current blocks—can be found on this page, including block height, its previous and next blocks and the corresponding miners. You can also use the search bar to search by other criteria, such as, address, BlockNum, TxHash, etc.

![](https://i.imgur.com/R10AK5K.png)

### Transaction Search

You can search for transaction records on this page. Information on the sender and the recipient’s address, the amount of MAN transferred, block height of transaction record, corresponding Hash and production time can all be found. You can also use the search bar to look for a specific transaction by Hash or other information.

![](https://i.imgur.com/9rbSyqg.png)


### Account Information

On the Transaction record page, you can click on an address to direct to an Overview page.
![](https://i.imgur.com/RlU145U.png)

![](https://i.imgur.com/DyEw9tu.png)

On the above Overview page, we can see the balance of this address, all transactions involving this address (both From and To), etc.

### Check Transaction Statistics

Under Stats of the Home page, we can see four sub-tabs: HashRate Growth Chart, Miner Distribution Chart, Difficulty Chart and Blocktime Chart.

![](https://i.imgur.com/0cEnHPy.png)

![](https://i.imgur.com/fmcv7EA.png)

![](https://i.imgur.com/zINAHuk.png)

![](https://i.imgur.com/FZVsr5q.png)





## AI Server

When a user attached a picture to a transaction, it is automatically analyzed and processed by the AI Server. The analysis results and transaction information are then displayed in html in the MATRIX Blockchain Browser.

The AI Server supports several processing services including AI Pose Detection and AI Object Detection. The particular service used by the AI Server is selected according to image type.


    AI Pose Detection
The AI Pose Detection recognizes human posture. This AI service detects and marks different parts of the human body in static, flat pictures.

#### Before the AI Pose Detection

![](https://i.imgur.com/m6JcTkV.jpg)

#### After the AI Pose Detection

![](https://i.imgur.com/d80pm3I.png)


The AI Pose Detection processing is implemented by another automatically triggered transaction. Note that the “from” address of this transaction is the AI Server address. The “to” address is the “from” address in the original transaction.


    AI Object Detection
The AI Object Detection is designed to detect objects in pictures and provide measures of statistical similarity between multiple objects.


#### Before the AI Object Detection

![](https://i.imgur.com/vj0MGZ8.png)


#### After the AI Object Detection

![](https://i.imgur.com/Hd690nm.png)

As with the AI Pose Detection, the AI Object Detection processing is conducted automatically via another transaction. The “from” address of this transaction is the AI Server address. The “to” address is the “from” address in the original transaction.




