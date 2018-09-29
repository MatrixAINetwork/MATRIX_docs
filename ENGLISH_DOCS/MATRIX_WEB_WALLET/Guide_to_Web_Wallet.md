# MATRIX Web Wallet User Guide
---

## About
A simple web wallet is available for users who do not want to install additional software on their computers.

There are a couple of things you need to know about the web client before you use it:

- You cannot stake mine with it.
- Your wallet is kept entirely in the local storage of your web browser. This means that if you delete your local storage, you will delete your wallet and must recreate it from the seed.
- The security of your wallet depends entirely on the security of your web browser.
- You can put a pin on your wallet to prevent sending funds but any other access is dependent entirely on the access controls on your computer, not on the server or any login details. 

Web URL: Will be announced soon


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

Take minerDeposits for example, you will need to deposit 10,000 MAN (the same process for validatorDeposits, with the deposit amount of 100,000 MAN)

![](https://i.imgur.com/Zt3Q0fF.png)


![](https://i.imgur.com/gTMWsav.png)

![](https://i.imgur.com/RGOpQH6.png)

## WITHDRWAL TRANSACTIONS & REFUND TRANSACTIONS (Under 'Contracts')

The only difference with deposit transactions is:

![](https://i.imgur.com/Appem22.png)


## Check Transaction Status


![](https://i.imgur.com/ifyaebv.png)

#### If it displays "Transaction Found", the transaction has already been recorded on the blockchain.

![](https://i.imgur.com/rGzXRSX.png)

#### If it displays "Pending Transaction Found", the transaction is in the transaction pool, which can be cancelled or replaced by clicking 'Cancel or Replace Transaction'


![](https://i.imgur.com/DgnfqMh.png)

