# MATRIX Web Wallet User Guide
---

## About
A simple web wallet is available for users who do not want to install additional software on their computers.

There are a couple of things you need to know about the web client before you use it:

- You cannot stake mine with it.
- Your wallet is kept entirely in the local storage of your web browser. This means that if you delete your local storage, you will delete your wallet and must recreate it from the seed.
- The security of your wallet depends entirely on the security of your web browser.
- You can put a pin on your wallet to prevent sending funds but any other access is dependent entirely on the access controls on your computer, not on the server or any login details. 

Web URL: testnet1.matrix.io/


## Create New Wallet

![](https://i.imgur.com/DiZEJoM.png)


## Save your Keystore File(UTC/JSON)

![](https://i.imgur.com/yOyrJkS.png)

Below is a sample JSON file:

    {"version":3,"id":"69dfbc60-7221-4f05-853d-749046b685bb","address":"de4ce30a9abc138408509c7628681fd08d931586","Crypto":{"ciphertext":"099827b94d7c64e4aa54bccbad342db7a375d558987c79ae921c9d7223b44bdc","cipherparams":{"iv":"07adb85fcfb94f487f4381fdc47e79b4"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"1d57cfc8fc5cb9a03a30ddc2cbe4949e17e220949a582ad8be10b41c02be9021","n":8192,"r":8,"p":1},"mac":"c85427a2d62b10d24647e06d2cda99c5576828d655a3d4382c49fa478fafebbb"}}

After clicking `I understand. Continue`:

![](https://i.imgur.com/5KpxHyD.png)


## Click 'Save Your Address.'

![](https://i.imgur.com/oGJ4L3s.png)

![](https://i.imgur.com/qrl2qHk.png)

![](https://i.imgur.com/q4Tf1d1.png)

### VERY IMPORTANT

**PRIVATE KEY CAN BE USED TO GENERATE KEYSTORE AND PASSWORD, AS WELL AS CHANGE THE PASSWORD**


## ONLINE TRANSACTIONS

### One2One Transfer Transaction


![](https://i.imgur.com/x7SaiOS.png)


![](https://i.imgur.com/OVM5XKr.png)


![](https://i.imgur.com/yjv8gcS.png)


### One2Multiple Transafer Transaction

![](https://i.imgur.com/MRliRYA.png)

For the remaining actions, please refer to 'One2One Transfer Transaction'




## OFFLINE TRANSACTIONS


Compared to 'Online Transactions', the most distinct feature of 'Offline Transactions' is: **You can leave a message here**


Texts, image upload, as well as video upload (480 bytes at the most) are all supported


![](https://i.imgur.com/MJwquJp.png)


![](https://i.imgur.com/QBaBaad.png)



![](https://i.imgur.com/jHOzIK2.png)



You can scan the QR code to get the signed transaction.

Then, Click 'Send Transaction.

![](https://i.imgur.com/Met02ru.png)


Finally, confirm the transfer information, and click 'Yes, I am sure! Make transaction.'


## DEPOSIT TRANSACTIONS (Under 'Contracts')


In order to be a candidate of MATRIX miners or validators, you must first make the required deposits, which is generally accomplished via our smart contract.


There are two types of deposit transactions: deposit for miners (deposit should be between 10,000 and 100,000 MANs) and deposit for validators (deposit should be equal to or greater than 100,000 MANs).


Two actions will be needed when you want to quit from deposit: Withdraw and Refund.

When the withdraw is completed, it will display a withdraw height in the status bar. And when the refund is completed, the deposits will be returned to the original account.

> **Note: gasprice and gaslimit in deposit transactions are not required to be filled in manually, and set as default**

#### View the Deposit Status: Input the account address

![](https://i.imgur.com/rbeoetH.png)

 	Deposit Amount: The unit is WEI. The real amount is the result dividing the Deposit Amount by (1+e18) - ten to the power of 18.

 	NodeID: The ID of nodes qualified for deposits and election (You need to run gman when performing the testing, and obtain this id via commandline)

 	Deposit Status: well deposited; already quit from deposit at xx height but not refunded yet


##### How to carry out deposit transactions

For miner deposit: 1w<deposit amount<10w

![](https://i.imgur.com/PKdwssH.png)

For validator deposit: deposit amount>=10w
![](https://i.imgur.com/IX3MGt3.png)

## WITHDRWAL from DEPOSIT TRANSACTIONS  (Under 'Contracts')

You can withdraw from deposit and view the status after submitting the withdraw request.

**NOTE:** After the withdrawl, the account will still appear in the deposit list but with no reward. There are two requirements for a real withdrawl: reach a specified height (= latest height - original height); succees in refund


![](https://i.imgur.com/AIFWzWM.png)

![](https://i.imgur.com/PomNx9P.png)

## REFUND TRANSACTIONS (Under 'Contracts')

Similar to Withdrawl transactions, refund transactions will also be completed at a specified height. The original deposit will be returned to the account, and all related deposit information will be cleared

![](https://i.imgur.com/8Pts1Xj.png)




## Check Transaction Status


![](https://i.imgur.com/ifyaebv.png)

#### If it displays "Transaction Found", the transaction has already been recorded on the blockchain.

![](https://i.imgur.com/rGzXRSX.png)

#### If it displays "Pending Transaction Found", the transaction is in the transaction pool, which can be cancelled or replaced by clicking 'Cancel or Replace Transaction'


![](https://i.imgur.com/DgnfqMh.png)

