package account

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HexFormatAcctAddress(acctAddress string) string {
	if len(acctAddress) >= 2 && acctAddress[0] == '0' && (acctAddress[1] == 'x' || acctAddress[1] == 'X') {
		acctAddress = acctAddress[2:]
	}
	if len(acctAddress) != 64 {
		// Pad with leading zeros if necessary
		acctAddress = fmt.Sprintf("%0*s", 64, acctAddress)
	}
	return acctAddress
}

func GetAccountAuth(client *ethclient.Client, acctAddress string) *bind.TransactOpts {
	fmt.Println("acctAddress=", acctAddress)
	privateKeyHex := os.Getenv("PRIVATE_KEY")[2:]
	privateKeyBytes, _ := hex.DecodeString(privateKeyHex)

	fmt.Println("privateKeyBytes=", privateKeyBytes)
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	// privateKey, err := crypto.HexToECDSA(acctAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("privateKey=", privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// fromAddress := common.HexToAddress(privateKey)
	fmt.Println("fromAddress=", fromAddress)
	// fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nonce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("chainID=", chainID)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("gasPrice=", gasPrice)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	fmt.Println("Sender account:", auth.From.Hex())
	fmt.Printf("Account balance: %v wei\n", balance)
	fmt.Println("auth=", auth)
	return auth
}
