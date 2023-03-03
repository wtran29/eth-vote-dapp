package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/wtran29/eth-vote-dapp/app/block/account"
	Ballot "github.com/wtran29/eth-vote-dapp/app/contracts/go/binding"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	acctAddress := os.Getenv("ACCT_ADDRESS")
	clientURL := os.Getenv("GANACHE_RPC")
	// Connect to Ethereum client
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("client=", client)
	// defer client.Close()
	// deploy smart contract
	fmt.Println("Running migrations...")
	// migrations.DeployMigrations(client)
	// migrations.DeployTodoList(client)
	auth := account.GetAccountAuth(client, acctAddress)
	address, tx, instance, err := Ballot.DeployBallot(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy TodoList contract: %v", err)
	}
	txHex := tx.Hash().Hex()
	fmt.Printf("tx= %v\n", txHex)
	fmt.Printf("migrations= %v\n", instance)
	fmt.Printf("Contract deployed: %s\n", address.Hex())

	fmt.Println("Migrations complete!")
	account.CheckTrans(txHex)
}
