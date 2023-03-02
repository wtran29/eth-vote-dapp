package account

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func WeiToEth(wei uint64, client *ethclient.Client) *big.Float {
	ethValue := new(big.Float)
	weiValue := new(big.Float).SetUint64(wei)
	ethValue = ethValue.Quo(weiValue, big.NewFloat(params.Ether))
	return ethValue

}
