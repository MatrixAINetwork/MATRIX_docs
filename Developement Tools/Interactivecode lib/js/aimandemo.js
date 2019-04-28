
let AIMan = require('aiman');
let Tx = require('matrixjs-tx');

var keythereum = require("keythereum");
var polycrc = require('polycrc')
const bs58 = require('bs58')


var privateKey = new Buffer('c848014293983236e84472625165a3a0e7b66dead73c645ad393451e75e10660', 'hex')
let from = "MAN.3SPbc3M7bK8zCT8VbvjMGW2eCaBgY";

let aiman = new AIMan(new AIMan.providers.HttpProvider("https://testnet.matrix.io"));

genManAddress = function (address) {
  let crc8 = polycrc.crc(8, 0x07, 0x00, 0x00, false)
  const bytes = Buffer.from(address, 'hex')
  const manAddress = bs58.encode(bytes)
  let arr = ['1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P',
    'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
    'j', 'k', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'
  ];
  return ('MAN.' + manAddress) + arr[crc8('MAN.' + manAddress) % 58]
}

//Create Account(private key or password+keystore)
function CreatKeystore(password){
  var dk;// = keythereum.create();
  while (true) {
      dk = keythereum.create();
      if (dk.privateKey[0] === 0) {
        break;
      }
  }
  //private key
  console.log(dk.privateKey);
  //password + keystore
  var keyObject = keythereum.dump(password, dk.privateKey, dk.salt, dk.iv);
  keyObject.address = genManAddress(keyObject.address);
  console.log(keyObject.address);
  keythereum.exportToFile(keyObject, "./",function (result) {
    console.log(result);
});
}


//sendTx
function sendRawTransaction(from, privateKey) {

var rawTx = {
  to: 'MAN.2Uoz8g8jauMa2mtnwxrschj2qPJrE',
  value: '0x989680',
  gasPrice: '0x430e23400',
  gas: 21000,
  data: "0x00",
  nonce: 4503599627370496,
  TxEnterType: '',
  IsEntrustTx: '',
  CommitTime: 1547539401231,
  extra_to: [[0, 0, []]],
  chainId: 3
}
  aiman.man.getTransactionCount(from, function (err, result) {
    // console.log(result)
    if (err) {
      console.log(err.message);
    }
    var nonce = result;
    rawTx.nonce = aiman.toHex(nonce);
    const tx = new Tx(rawTx);
    // console.log("sss"+tx);
	//signTx
    tx.sign(privateKey);
	//signed data construct
    const serializedTx = tx.serialize();
    var data = "0x" + serializedTx.toString('hex');
   
    const txData = new Tx(data, true);
    var newTxData = {};
    newTxData.v = '0x' + txData.v.toString('hex');
    newTxData.r = '0x' + txData.r.toString('hex');
    newTxData.s = '0x' + txData.s.toString('hex');
    newTxData.data = '0x' + txData.data.toString('hex');
    newTxData.gasPrice = '0x' + txData.gasPrice.toString('hex');
    newTxData.gas = '0x' + txData.gasLimit.toString('hex');
    newTxData.value = '0x' + txData.value.toString('hex');
    if (newTxData.value == '0x') {
      newTxData.value = '0x0'
    }
    newTxData.nonce = '0x' + txData.nonce.toString('hex');
    if (txData.to.toString() != '') {
      newTxData.to = txData.to.toString();
      newTxData.currency = txData.to.toString().split('.')[0];
    } else {
      newTxData.currency = 'MAN';
    }
    var txType = '0x' + txData.extra_to[0].toString('hex');
    var lockHeight = '0x' + txData.extra_to[1].toString('hex');
    var isEntrustTx = '0x' + txData.IsEntrustTx.toString('hex');
    var txEnterType = '0x' + txData.TxEnterType.toString('hex');
    var commitTime = '0x' + txData.CommitTime.toString('hex');
    newTxData.txType = txType == '0x' ? 0 : Number(txType);
    newTxData.lockHeight = lockHeight == '0x' ? 0 : Number(lockHeight);
    newTxData.isEntrustTx = isEntrustTx == '0x' ? 0 : Number(isEntrustTx);
    newTxData.txEnterType = txEnterType == '0x' ? 0 : Number(txEnterType);
    newTxData.commitTime = commitTime == '0x' ? 0 : Number(commitTime);
    newTxData.extra_to = [];
    for (var i = 0, length = txData.extra_to[2].length; i < length; i++) {
      newTxData.extra_to.push({
        to: txData.extra_to[2][i][0].toString(),
        value: '0x' + txData.extra_to[2][i][1].toString('hex')
      });
    }
    aiman.man.sendRawTransaction(newTxData, function (err, result) {
      if (!err) {
        console.log(result);
      } else {
        console.log(err.message);
      }
    });
  });
}

//get gas price
function GetGasPrice() {
  let gasPrice = aiman.man.gasPrice;
  console.log(gasPrice.toString());
}

//get block info
function GetBlockByNumber(blocknum) {
  let blockinfo = aiman.man.getBlock(blocknum,false);
  console.log(blockinfo);
}

//get balance
function GetBalance(addr) {
  let balance = aiman.man.getBalance(addr);
  console.log(balance[0].balance.toString());
}


// sendRawTransaction(from, privateKey);
// GetGasPrice();
// GetBalance("MAN.2Uoz8g8jauMa2mtnwxrschj2qPJrE");
// GetBlockByNumber(121);
// CreatKeystore("1111111");
