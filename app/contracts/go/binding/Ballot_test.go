package Ballot_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	Ballot "github.com/wtran29/eth-vote-dapp/app/contracts/go/binding"
)

func TestBallot(t *testing.T) {
	// Create a simulated backend
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		fromAddress: {Balance: big.NewInt(1000000000000000000)},
	}, 3000000)

	chainID := sim.Blockchain().Config().ChainID
	fmt.Println(chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		t.Fatalf("unable to create auth: %s", err)
	}

	// Set the simulated backend for the transaction signer
	nonce, err := sim.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		panic(err)
	}
	// gasLimit := sim.Blockchain().GasLimit()
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice, err = sim.SuggestGasPrice(context.Background())
	require.NoError(t, err)
	auth.Context = context.WithValue(context.Background(), "backend", sim)

	// Deploy the contract
	address, _, contract, err := Ballot.DeployBallot(auth, sim)
	require.NoError(t, err)
	assert.NotNil(t, address)
	// Commit all pending transactions
	sim.Commit()

	// Test the contract
	t.Run("initializes two candidates", func(t *testing.T) {
		candidatesCount, err := contract.CandidatesCount(nil)
		if err != nil {
			fmt.Printf("could not get candidates count: %s", err)
		}
		require.NoError(t, err)
		assert.Equal(t, "2", candidatesCount.String())
	})
	t.Run("initialize the candidates with correct values", func(t *testing.T) {
		fcandidate, err := contract.Candidates(nil, big.NewInt(1))
		if err != nil {
			fmt.Printf("did not find first candidate: %s", err)
		}
		require.NoError(t, err)
		assert.Equal(t, "1", fcandidate.Id.String(), "contains the correct id")
		assert.Equal(t, "Candidate 1", fcandidate.Name, "contains the correct name")
		assert.Equal(t, "0", fcandidate.VoteCount.String(), "contains the correct votes count")
		scandidate, err := contract.Candidates(nil, big.NewInt(2))
		if err != nil {
			fmt.Printf("did not find second candidate: %s", err)
		}
		require.NoError(t, err)
		assert.Equal(t, "2", scandidate.Id.String(), "contains the correct id")
		assert.Equal(t, "Candidate 2", scandidate.Name, "contains the correct name")
		assert.Equal(t, "0", scandidate.VoteCount.String(), "contains the correct votes count")
	})

}
