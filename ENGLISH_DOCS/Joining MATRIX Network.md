## Getting Started

We have presented how to operate a private network before. However, in order to test whether a blockchain network can achieve the best perfomance, there should be many nodes running together. 

Here, we sincerely invite you to join our network and enjoy fun mining on it.


## Terminology

| Terminology 	| Description                                                         	|
|-------------	|---------------------------------------------------------------------	|
| gman        	| MATRIX CLI client, which is the entry point into the MATRIX network 	|
| nodeID      	| Unique identifier of a node in the network                          	|


## Create a new MATRIX account

First, you need to get a new account in the MATRIX network, which means we shall generate a new wallet address. Entry into MATRIX network via gman is supported, but not so convenient enough. 

We recommend using our Online Wallet: testnet.matrix.io


**Please follow the steps to create the new wallet：**

#### Step 1: Click New Wallet, and input a password in the box, then click 'Create New Wallet'

![](https://i.imgur.com/DiZEJoM.png)

Please keep your password safe and well!

#### Step2: Save your Keystore File(UTC/JSON)

![](https://i.imgur.com/yOyrJkS.png)

Below is a sample JSON file:

    {"version":3,"id":"69dfbc60-7221-4f05-853d-749046b685bb","address":"de4ce30a9abc138408509c7628681fd08d931586","Crypto":{"ciphertext":"099827b94d7c64e4aa54bccbad342db7a375d558987c79ae921c9d7223b44bdc","cipherparams":{"iv":"07adb85fcfb94f487f4381fdc47e79b4"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"1d57cfc8fc5cb9a03a30ddc2cbe4949e17e220949a582ad8be10b41c02be9021","n":8192,"r":8,"p":1},"mac":"c85427a2d62b10d24647e06d2cda99c5576828d655a3d4382c49fa478fafebbb"}}

After clicking `I understand. Continue`:

![](https://i.imgur.com/5KpxHyD.png)


#### Step3: Click 'Save Your Address.'

#### Step4: Unlock your wallet Via keystore

![](https://i.imgur.com/oGJ4L3s.png)

![](https://i.imgur.com/qrl2qHk.png)

![](https://i.imgur.com/q4Tf1d1.png)



## Apply for MANs


As you can see, the initial balance of the newly created account is 0. You need to apply for MANs to go on, for example, to make deposits to run for Miner or Validator.


Just send your new account address to MATRIX, and we will recharge all accounts after that


## How to get your nodeID

nodeID is the unique identifier of a node in the network. It is required if you want to make deposits to run for Miner or Validator through our online wallet.

Let's see how to get this ID:

- Step1： Create a folder named 'chaindata' under root directory, where gman is also placed

- Step2: Create a new password.txt file under 'chaindata', which stores the password of your new wallet (This step is required if you want to perfrom transactions). Also place your keystore file here

- Step3: Fetch the specified genesis file (MANGenesis.json) and common profile (man.json) from Github ([https://github.com/MatrixAINetwork/MATRIX-TESTNET](https://github.com/MatrixAINetwork/MATRIX-TESTNET))

You can place MANGenesis.json anywhere, for example, under root; but man.json should be placed under /chaindata (Otherwise, you can't start gman due to the inability to read the profile)

- Step4: Initialize

/geth --datadir /chaindata/ init /MANGenesis.json

- Step5: Start GMAN client

/gman --datadir /chaindata --rpc --rpcaddr 0.0.0.0 --rpcapi admin,man,net,web3,personal,debug,miner,web3 --rpccorsdomain * --networkid 8 --debug --verbosity 5 --password /chaindata/password.txt --outputinfo 1 --gcmode archive --syncmode full 

- Step6: Run 'Attach': /gman attach /chaindata/gman.ipc (gman.ipc is generated under /chaindata when starting gman)

- Step7: Run the command 'admin.nodeInfo' and the node info will be listed as follows:

![](https://i.imgur.com/H00ERng.png)

## Make Deposits (Under 'Deposit/Withdrawal' of Online Wallet)


In order to be a candidate of MATRIX miners or validators, you must first make the required deposits, which can be accomplished via our online wallet.


There are two types of deposit transactions: deposit for miners (deposit should be between 10,000 and 100,000 MANs) and deposit for validators (deposit should be equal to or greater than 100,000 MANs).


### Deposits for Miners

![](https://i.imgur.com/8SemCkc.png)

Step 1: As the above screenshot indicates, you shall first select Deposit Transaction, and choose Miner Deposit from the Deposit Type dropdown list

Step 2: Input the amount and node ID in the box

Step 3: Unlock your wallet via Keystore / JSON File

![](https://i.imgur.com/A3thcHP.png)

Step 4: Click Submit, and it pops up a window like:

![](https://i.imgur.com/SUOtyV2.png)

Step 5: Finish this transaction by clicking 'Yes, I am sure! Make transaction.'

A transaction hash will be generated which you can copy to our Explorer to search for that transaction.


### View Deposit Status


After completing the deposit transaction, you can search for its status by inputting account address and clicking 'Search'


![](https://i.imgur.com/KnYfwRk.png)


This status means you are already in the deposit list for Miners and have opportunities to be elected as a Miner during the next election cycle.



### Deposits for Validators

Step 1: You shall first select Deposit Transaction, and choose Validator Deposit from the Deposit Type dropdown list

![](https://i.imgur.com/GDeUumy.png)


The remaining steps are similar to those of 'Deposits for Miners' except that the Deposit Amount should be >= 100,000 MANs.



After completing it successfully, its status will be like:

![](https://i.imgur.com/Q3gZ7d6.png)


This status means you are already in the deposit list for Validators and have opportunities to be elected as a Validator during the next election cycle.


### Withdraw from Deposit List

Now you may want to quit from deposit list and no longer act as a miner or validator. 

![](https://i.imgur.com/yR3ZtoD.png)


Step 1: First select Withdrawl Transaction, and choose deposit withdrawl from the Withdrawl Transaction dropdown list

Step 2: Unlock your wallet via Keystore / JSON File

Step 3: Complete the transaction like you have done previously.

Step 4: The status will be like: 

![](https://i.imgur.com/iEbCvBq.png)

You can see at which block height the withdrawl happens.

**Note: After withdrawl, your account will still exist in the deposit list. And you can make refunds after 600 blocks**


### Make Refunds

After withdrawl from Deposit List, you may want to get your deposit back. Here is how to:

Step 1: First select Withdrawl Transaction, and choose deposit refund from the Withdrawl Transaction dropdown list

Step 2: Unlock your wallet via Keystore / JSON File

Step 3: Complete the transaction like you have done previously.

After refund, your deposit status will disappear. And you can check your balance change.

Well, you can certainly re-join the deposit list.