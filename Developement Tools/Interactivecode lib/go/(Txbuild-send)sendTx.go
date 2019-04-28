/*
Package vehicle_btc main.go

@Leno Lee <yongli@matrix.io>
@copyright All rights reserved
*/
package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/MatrixAINetwork/go-AIMan/AIMan"
	"github.com/MatrixAINetwork/go-AIMan/Accounts"
	"github.com/MatrixAINetwork/go-AIMan/dto"
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"github.com/MatrixAINetwork/go-AIMan/providers"
	"github.com/MatrixAINetwork/go-AIMan/transactions"
	"github.com/MatrixAINetwork/go-AIMan/waiting"
	"github.com/MatrixAINetwork/go-matrix/base58"
	"github.com/MatrixAINetwork/go-matrix/core/types"
	"github.com/MatrixAINetwork/go-matrix/crypto"
	"math/big"
	"time"

	"github.com/MatrixAINetwork/go-matrix/accounts/keystore"
	"path/filepath"
)

var (
	KeystorePath = "keystore"
	Tom_Manager  = &manager.Manager{
		AIMan.NewAIMan(providers.NewHTTPProvider("api85.matrix.io", 100, false)),
		Accounts.NewKeystoreManager(KeystorePath, 1),
	}

	Jerry_Manager = &manager.Manager{
		AIMan.NewAIMan(providers.NewHTTPProvider("testnet.matrix.io", 100, true)),
		Accounts.NewKeystoreManager(KeystorePath, 3),
	}

)

//发送交易
func SendTx(from string, to string, money int64, usegas int, gasprice int64) (connection *manager.Manager, txID string) {

	connection = Jerry_Manager
	cid := *connection.ChainID
	fmt.Println(cid)
	types.NewEIP155Signer(connection.ChainID)

	amount := big.NewInt(money)
	gas := uint64(usegas)
	price := big.NewInt(gasprice)
	err := connection.Unlock(from, "xxx")
	if err != nil {
		return
	}

	//获取nonce
	nonce, err := connection.Man.GetTransactionCount(from, "latest")
	if err != nil {
		fmt.Println("GetTransactionCount:",err)
		return
	}

	//构建交易对象
	trans := transactions.NewTransaction(nonce.Uint64(), to, amount, gas, price,
		[]byte{}, 0, 0, 0)

	//对构建好的交易对象签名
	raw, err := connection.SignTx(trans, from)
	if err != nil {
		fmt.Println("SignTx:",err)
		return
	}

	//发送签名后的交易对象
	txID, err = connection.Man.SendRawTransaction(raw)
	if err != nil {
		fmt.Println("SendRawTransaction:",err)
		return
	}
	fmt.Println(txID)
	//fmt.Println("use",big.NewInt(0).Mul(trans.GasPrice.ToInt(),big.NewInt(trans.Gas)))
	var receipt *dto.TransactionReceipt
	wait2 := waiting.NewMultiWaiting(
		waiting.NewWaitTime(time.Second*60),
		waiting.NewWaitTxReceipt(connection, txID),
		//waiting.NewWaitBlockHeight(connection,blockNumber.Uint64()+3),
	)
	if index := wait2.Waiting(); index != 1 {
		//t.Error("timeout")
		//t.FailNow()
		fmt.Println("error")
	}
	receipt, err = connection.Man.GetTransactionReceipt(txID)
	if receipt.Status == false {
		fmt.Println("recipt_status == false")
	}
	fmt.Println(receipt)

	return
}

//发送交易（使用私钥进行签名）
func SendTxByPrivateKey(from string, to string, money int64, usegas int, gasprice int64,PrivateKey *ecdsa.PrivateKey) (connection *manager.Manager, txID string) {

	connection = Jerry_Manager

	amount := big.NewInt(money)
	gas := uint64(usegas)
	price := big.NewInt(gasprice)
	err := connection.Unlock(from, "xxx")
	if err != nil {
		return
	}

	//获取nonce
	nonce, err := connection.Man.GetTransactionCount(from, "latest")
	if err != nil {
		fmt.Println("GetTransactionCount:",err)
		return
	}

	//构建交易对象
	trans := transactions.NewTransaction(nonce.Uint64(), to, amount, gas, price,
		[]byte{}, 0, 0, 0)

	trans,err=connection.Man.SignTxByPrivate(trans,from,PrivateKey,connection.ChainID)
	//发送签名后的交易对象
	txID, err = connection.Man.SendRawTransaction(trans)
	if err != nil {
		fmt.Println("SendRawTransaction:",err)
		return
	}
	fmt.Println(txID)
	//fmt.Println("use",big.NewInt(0).Mul(trans.GasPrice.ToInt(),big.NewInt(trans.Gas)))
	var receipt *dto.TransactionReceipt
	wait2 := waiting.NewMultiWaiting(
		waiting.NewWaitTime(time.Second*60),
		waiting.NewWaitTxReceipt(connection, txID),
		//waiting.NewWaitBlockHeight(connection,blockNumber.Uint64()+3),
	)
	if index := wait2.Waiting(); index != 1 {
		//t.Error("timeout")
		//t.FailNow()
		fmt.Println("error")
	}
	receipt, err = connection.Man.GetTransactionReceipt(txID)
	if receipt.Status == false {
		fmt.Println("recipt_status == false")
	}
	fmt.Println(receipt)

	return
}

//创建账户(在本地文件夹中创建私钥文件)
func CreatKeystore() {
	// Create an encrypted keystore with standard crypto parameters
	ks := keystore.NewKeyStore(filepath.Join("", "keystore"), keystore.StandardScryptN, keystore.StandardScryptP)

	// Create a new account with the specified encryption passphrase
	newAcc, err := ks.NewAccount("Creation password")
	if err != nil {

	}
	manAddress := newAcc.ManAddress()
	fmt.Println(manAddress)
}

//创建账户
func GenManAddress()  {
	privateKey, err := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	if err != nil {
		return
	}
	//
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	fromMan := base58.Base58EncodeToString("MAN", from)
	fmt.Println(fromMan)
}

//获取账户余额
func GetBalance(addr string) *big.Int {
	connection := Jerry_Manager
	balance,err:=connection.Man.GetBalance(addr, "latest")
	if err!=nil {
	}
	fmt.Println(addr,":",balance[0].Balance.ToInt())
	return balance[0].Balance.ToInt()
}

//获取gasprice
func GetGasPrice() *big.Int  {
	connection := Jerry_Manager
	gasprice,_:=connection.Man.GetGasPrice()
	fmt.Println(gasprice)
	return gasprice
}

//获取区块
func GetBlockByNumber()  {
	connection := Jerry_Manager

	block,err:=connection.Man.GetBlockByNumber(big.NewInt(211),true)
	if err!=nil {
		fmt.Println("err:",err)
	}

	for _,txs := range block.Transactions {
		for _,tx := range txs{
			fmt.Println(tx)
		}
	}
}

func main() {
	//app, port := vehicle()
	//app.Run(iris.Addr("localhost:"+port), iris.WithoutServerError(iris.ErrServerClosed))

	from := "MAN.CrsnQSJJfGxpb2taGhChLuyZwZJo"
	to := "MAN.3qQQqfzBdwBjpauj6ght4G8E6o1yQ"
	SendTx(from, to, 1, 21000, 18e9)
	GetBalance(from)
	//CreatKeystore()
	//GenManAddress()
	GetBlockByNumber()
	GetGasPrice()
}