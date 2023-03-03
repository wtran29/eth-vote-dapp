package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	Ballot "github.com/wtran29/eth-vote-dapp/app/contracts/go/binding"
)

type ContractDetails struct {
	Address          string            `json:"address"`
	ABI              string            `json:"abi"`
	Bytecode         string            `json:"bytecode"`
	DeployedBytecode string            `json:"deployedBytecode"`
	Methods          map[string]string `json:"methods"`
	Events           map[string]string `json:"events"`
}

func CheckTrans(txHex string) {
	// Connect to an Ethereum client
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	// Get the transaction receipt
	txHash := common.HexToHash(txHex)
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// // Convert the receipt to JSON format
	// receiptJSON, err := json.MarshalIndent(txReceipt, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Print the input data
	// fmt.Printf("Receipt: %v\n", string(receiptJSON))

	fmt.Printf("Block Number: %d\n", receipt.BlockNumber)
	fmt.Printf("Network ID: %d\n", networkID)
	fmt.Printf("Block Hash: %s\n", receipt.BlockHash.Hex())
	fmt.Printf("Transaction Hash: %s\n", receipt.TxHash.Hex())
	fmt.Printf("Contract Address: %s\n", receipt.ContractAddress.Hex())
	fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
	fmt.Printf("Cumulative Gas Used: %d\n", receipt.CumulativeGasUsed)
	fmt.Printf("Status: %d\n", receipt.Status)
	fmt.Printf("Logs Bloom: %s\n", receipt.Bloom.Big())
	fmt.Printf("Logs: %v\n", receipt.Logs)

	UpdateNetworks(networkID, receipt.ContractAddress.Hex())
}

func GetTodoListContractDetails() {
	// Connect to Ethereum network
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		fmt.Println("Failed to connect to Ethereum network:", err)
		return
	}

	// Get the network ID
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("Failed to get network ID:", err)
		return
	}

	// Load the contract ABI
	abiFile, err := ioutil.ReadFile("app/contracts/abi/ballot/Ballot.abi")
	if err != nil {
		fmt.Println("Failed to read ABI file:", err)
		return
	}
	abiString := string(abiFile)

	// Set the contract address
	contractAddress := common.HexToAddress(os.Getenv("ACCT_ADDRESS"))

	// Instantiate the contract

	contractInstance, err := Ballot.NewBallot(contractAddress, client)
	fmt.Printf("contract instance: %v\n", contractInstance)
	if err != nil {
		fmt.Println("Failed to instantiate contract:", err)
		return
	}

	// Get contract details
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		fmt.Println("Failed to get bytecode:", err)
		return
	}

	deployedBytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		fmt.Println("Failed to get deployed bytecode:", err)
		return
	}

	methods := make(map[string]string)
	abiParser, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		fmt.Println("Failed to parse ABI:", err)
		return
	}
	for _, method := range abiParser.Methods {
		methods[method.Name] = method.String()
	}

	events := make(map[string]string)
	for _, event := range abiParser.Events {
		events[event.Name] = event.String()
	}

	// Return contract details
	contractDetails := ContractDetails{
		Address:          contractAddress.Hex(),
		ABI:              abiString,
		Bytecode:         hexutil.Encode(bytecode),
		DeployedBytecode: hexutil.Encode(deployedBytecode),
		Methods:          methods,
		Events:           events,
	}

	fmt.Printf("Contract Details: %+v\n", contractDetails)
	fmt.Println("Network ID:", networkID.String())
	fmt.Println("Address:", contractDetails.Address)
	fmt.Println("ABI:", contractDetails.ABI)
	fmt.Println("Bytecode:", contractDetails.Bytecode)
	fmt.Println("DeployedBytecode:", contractDetails.DeployedBytecode)
	fmt.Println("Methods:", contractDetails.Methods)
	fmt.Println("Events:", contractDetails.Events)
	// taskCount, err := contractInstance.TaskCount(nil)
	// if err != nil {
	// 	fmt.Println("Failed to retrieve task count:", err)
	// 	return
	// }
	// fmt.Println("TaskCount:", taskCount)

	// tasks, err := contractInstance.Tasks(nil, big.NewInt(1))
	// if err != nil {
	// 	fmt.Println("Failed to retrieve task:", err)
	// 	return
	// }

	// if tasks.Id.Cmp(big.NewInt(0)) == 0 {
	// 	fmt.Println("Task with ID 1 does not exist.")
	// } else {
	// 	fmt.Println("Task with ID 1 exists and has content:", tasks.Content)
	// }
	// fmt.Println("Task:", tasks)
	// fmt.Println("Task id:", tasks.Id)
	// fmt.Println("Task content:", tasks.Content)
	// fmt.Println("Task completed:", tasks.Completed)
}

func UpdateNetworks(networkId *big.Int, address string) error {
	jsonFile, err := ioutil.ReadFile("app/frontend/src/Ballot.json")
	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err = json.Unmarshal(jsonFile, &data); err != nil {
		return err
	}

	if networks, ok := data["networks"].(map[string]interface{}); ok {
		networks[fmt.Sprintf("%d", networkId)] = map[string]interface{}{
			"address": address,
		}
	} else {
		data["networks"] = map[string]interface{}{
			fmt.Sprintf("%d", networkId): map[string]interface{}{
				"address": address,
			},
		}
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile("app/frontend/src/Ballot.json", jsonData, 0644); err != nil {
		return err
	}

	return nil
}
