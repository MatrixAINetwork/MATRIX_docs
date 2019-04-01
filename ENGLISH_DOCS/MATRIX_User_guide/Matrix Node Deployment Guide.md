### Starting up your member nodes (Linux & Mac) - for deposited users

Step 1: Check out what you need to prepare (most of them can be obtaind from go-matrix repository)

    /gman: exe file

    /MANGenesis.json: genesis file

    /chaindata: a folder which you should create

    man.json: common profile which shall be put under /chaindata

Step 2: Run Initiate command

    ./gman  --datadir  ./chaindata/   init    ./MANGenesis.json

Step 3: Visit our web wallet to create a new wallet address, and save your keystore file as well as password.

Please refer to [['Guide to Web Wallet']](https://github.com/MatrixAINetwork/MATRIX_docs/blob/master/ENGLISH_DOCS/MATRIX_Web_Wallet/MATRIX%20Online%20Wallet%20Manual.pdf)

Carry out your deposit actions if you want to run for a miner or validator node (you can find steps on the above guide)

Step 4: Copy your keystore file to folder keystore which is generated at Step 2 (/chaindata/keystore)

Step 5: Create a file named signAccount.json under root, and its content is like:

    [
      {
        "Address":" MAN.gQAAHUeTBxvgbzf8tFgUtavDceJP ",
        "Password":" pass123456"
      }

    ]
Then, run: 

    ./gman --datadir ./chaindata aes --aesin ./signAccount.json --aesout entrust.json

Upon the window prompt, you will be asked to set a password (which should contain upper-case letter[s], lower-case letter[s], number[s] and special character[s])

Step 6: Copy the generated entrust.json to root

Step 7: Start gman

    ./gman --datadir ./chaindata --networkid 1 --debug --verbosity 5  --manAddress [your man.address here] --entrust ./entrust.json --gcmode archive --outputinfo 1 --syncmode full 

    for example, 

./gman --datadir ./chaindata --networkid 1 --debug --verbosity 5 --manAddress MAN.gQAAHUeTBxvgbzf8tFgUtavDceJP --entrust ./entrust.json --gcmode archive --outputinfo 1 --syncmode full

In this step, you will need to input the password set in step 5.

Step 8: Run 'Attach': ./gman attach /chaindata/gman.ipc (gman.ipc is generated under /chaindata when starting gman)


### Starting up your member nodes (Linux & Mac) - for non-deposited users

Step 1: Check out what you need to prepare (most of them can be obtaind from go-matrix repository)

    /gman: exe file

    /MANGenesis.json: genesis file

    /chaindata: a folder which you should create

    man.json: common profile which shall be put under /chaindata

Step 2: Run Initiate command

    ./gman  --datadir  ./chaindata/   init    ./MANGenesis.json

Step 3: Start
    ./gman --datadir ./chaindata --networkid 1  --outputinfo 1 --syncmode 'full'

### Starting up your member nodes (Windows) - for deposited users
Step 1: Check out what you need to prepare (most of them can be obtaind from go-matrix repository)

    /gman: exe file

    /MANGenesis.json: genesis file

    /chaindata: a folder which you should create

    man.json: common profile which shall be put under /chaindata

Step 2: Run Initiate command

    gman.exe --datadir chaindata\ init MANGenesis.json

Step 3: Create a file named signAccount.json, whose contents are:

    [
      {
        "Address":"MAN.2skMrkoEkecKjJLPz6qTdi8B3NgjU ",
        "Password":"haolin0123"
      }

    ]

Step 4: Run:

    gman.exe --datadir chaindata aes --aesin signAccount.json --aesout entrust.json

Upon the window prompt, you will be asked to set a password (which should contain upper-case letter[s], lower-case letter[s], number[s] and special character[s])

Step 5: Start gman

    gman --datadir chaindata  --networkid 1 --debug --verbosity 5  --manAddress  MAN.2skMrkoEkecKjJLPz6qTdi8B3NgjU --entrust entrust.json --gcmode archive --outputinfo 1 --syncmode full

In this step, you will need to input the password set in step 5.

Step 8: Open another window

    gman attach ipc:\\.\pipe\gman.ipc 

gman.ipc is generated under /chaindata when starting gman)

### Starting up your member nodes (Windows) - for non-deposited users

Step 1: Check out what you need to prepare (most of them can be obtaind from go-matrix repository)

    /gman: exe file

    /MANGenesis.json: genesis file

    /chaindata: a folder which you should create

    man.json: common profile which shall be put under /chaindata

Step 2: Run Initiate command

    gman.exe --datadir chaindata\ init MANGenesis.json

Step 3: Start gman

    gman --datadir chaindata  --networkid 1 --outputinfo 1 -- syncmode full