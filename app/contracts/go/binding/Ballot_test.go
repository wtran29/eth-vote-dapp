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

	acc1PrivK, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	acc1PubK := acc1PrivK.Public()
	acc1PubKeyECDSA, ok := acc1PubK.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	accountOne := crypto.PubkeyToAddress(*acc1PubKeyECDSA)

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		fromAddress: {Balance: big.NewInt(1000000000000000000)},
		accountOne:  {Balance: big.NewInt(1000000000000000000)},
		// accountTwo:  {Balance: big.NewInt(1000000000000000000)},
	}, 3000000)

	chainID := sim.Blockchain().Config().ChainID
	fmt.Println(chainID)

	acct1Auth, err := bind.NewKeyedTransactorWithChainID(acc1PrivK, chainID)
	if err != nil {
		t.Fatalf("unable to create account one auth: %s", err)
	}
	fmt.Println("acct1Auth=", acct1Auth)

	// acct2Auth, err := bind.NewKeyedTransactorWithChainID(acct2Pk, chainID)
	// if err != nil {
	// 	t.Fatalf("unable to create account two auth: %s", err)
	// }
	// fmt.Println("acct2Auth=", acct2Auth)

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
	t.Run("allows a voter to cast a vote", func(t *testing.T) {
		fcandidate, err := contract.Candidates(nil, big.NewInt(1))
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(0).String(), fcandidate.VoteCount.String(), "vote count should be zero")

		// Increment the nonce to match the prepopulated task
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		tx, err := contract.Vote(auth, fcandidate.Id)

		sim.Commit()
		fmt.Println(tx)
		assert.NoError(t, err)

		// Wait for the transaction to be mined
		receipt, err := bind.WaitMined(context.Background(), sim, tx)
		assert.NoError(t, err)
		assert.NotNil(t, receipt)

		fcandidate, err = contract.Candidates(nil, big.NewInt(1))
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(1), fcandidate.VoteCount, "vote count should be one")
		sim.Commit()

		voted, err := contract.Voters(nil, fromAddress)
		assert.NoError(t, err)
		assert.True(t, voted, "voter should be marked as voted")

		scandidate, err := contract.Candidates(nil, big.NewInt(2))
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(0).String(), scandidate.VoteCount.String(), "candidate 2 did not receive any votes")
		sim.Commit()
	})
	t.Run("throws an exception for invalid candidates", func(t *testing.T) {

		// Increment the nonce to match the prepopulated task
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		invalidCandidate, err := contract.Candidates(nil, big.NewInt(99))
		if err != nil {
			t.Fatalf("Invalid candidate 99: %v", err)
		}

		assert.Equal(t, big.NewInt(0).String(), invalidCandidate.Id.String(), "candidate 99 does not exist")

		candidate1, err := contract.Candidates(nil, big.NewInt(1))
		if err != nil {
			t.Fatalf("Error getting candidate 1: %v", err)
		}
		assert.Equal(t, big.NewInt(1), candidate1.VoteCount)

		candidate2, err := contract.Candidates(nil, big.NewInt(2))
		if err != nil {
			t.Fatalf("Error getting candidate 2: %v", err)
		}
		assert.Equal(t, big.NewInt(0).String(), candidate2.VoteCount.String())

	})
	t.Run("throws an exception for double voting", func(t *testing.T) {
		candidate1, err := contract.Candidates(nil, big.NewInt(1))
		if err != nil {
			t.Fatalf("Error getting candidate 1: %v", err)
		}

		candidate2, err := contract.Candidates(nil, big.NewInt(2))
		if err != nil {
			t.Fatalf("Error getting candidate 2: %v", err)
		}

		// Check if the sender has already voted
		hasAlreadyVoted, err := contract.Voters(nil, acct1Auth.From)
		fmt.Println("has voted? =", hasAlreadyVoted)
		if err != nil {
			t.Fatalf("Error getting voter's vote status: %v", err)
		}
		assert.False(t, hasAlreadyVoted, "Sender has already voted")
		// Cast a vote for candidate 1
		// candidateId1 := big.NewInt(2)
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		tx1, err := contract.Vote(acct1Auth, candidate1.Id)
		sim.Commit()
		fmt.Println("first vote=", tx1)
		assert.NoError(t, err)
		assert.NotNil(t, tx1)

		// Check that the sender has been marked as having voted
		hasAlreadyVoted, err = contract.Voters(nil, acct1Auth.From)
		fmt.Println("has voted?=", hasAlreadyVoted)
		if err != nil {
			t.Fatalf("Error getting voter's vote status: %v", err)
		}
		assert.True(t, hasAlreadyVoted, "Sender has been marked as having voted")

		// candidate1, err := contract.Candidates(nil, candidateId1)
		fmt.Println("candidate1=", candidate1)
		voteCount := candidate1.VoteCount
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(1), voteCount)

		// Check that the sender has been marked as having voted
		hasAlreadyVoted, err = contract.Voters(nil, acct1Auth.From)
		fmt.Println("has voted?=", hasAlreadyVoted)
		if err != nil {
			t.Fatalf("Error getting voter's vote status: %v", err)
		}
		assert.True(t, hasAlreadyVoted, "Sender has been marked as having voted")

		// Attempt to cast another vote for candidate 2 using the same account
		// candidateId2 := big.NewInt(2)
		if !hasAlreadyVoted {
			auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
			tx2, err := contract.Vote(acct1Auth, candidate2.Id)
			sim.Commit()
			fmt.Println("second vote=", tx2)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "revert")
			assert.Nil(t, tx2)
		}

		// candidate2, err := contract.Candidates(nil, candidateId2)
		fmt.Println("candidate 2 vote=", candidate2.VoteCount)
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(0).String(), candidate2.VoteCount.String())

		// Ensure that candidate 1 still has only 1 vote
		// candidate1, err = contract.Candidates(nil, candidateId1)
		fmt.Println("candidate 1 vote=", candidate1.VoteCount)
		assert.NoError(t, err)
		assert.Equal(t, big.NewInt(1).String(), candidate1.VoteCount.String())
	})

}
