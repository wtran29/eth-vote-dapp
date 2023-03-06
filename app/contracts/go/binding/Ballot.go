// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Ballot

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BallotMetaData contains all meta data concerning the Ballot contract.
var BallotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"votedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000586040518060400160405280600b81526020017f43616e6469646174652031000000000000000000000000000000000000000000815250620000a460201b60201c565b6200009e6040518060400160405280600b81526020017f43616e6469646174652032000000000000000000000000000000000000000000815250620000a460201b60201c565b62000500565b60026000815480929190620000b9906200015c565b919050555060405180606001604052806002548152602001828152602001600081525060016000600254815260200190815260200160002060008201518160000155602082015181600101908162000112919062000419565b506040820151816002015590505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000620001698262000152565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036200019e576200019d62000123565b5b600182019050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200022b57607f821691505b602082108103620002415762000240620001e3565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200026c565b620002b786836200026c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620002fa620002f4620002ee8462000152565b620002cf565b62000152565b9050919050565b6000819050919050565b6200031683620002d9565b6200032e620003258262000301565b84845462000279565b825550505050565b600090565b6200034562000336565b620003528184846200030b565b505050565b5b818110156200037a576200036e6000826200033b565b60018101905062000358565b5050565b601f821115620003c957620003938162000247565b6200039e846200025c565b81016020851015620003ae578190505b620003c6620003bd856200025c565b83018262000357565b50505b505050565b600082821c905092915050565b6000620003ee60001984600802620003ce565b1980831691505092915050565b6000620004098383620003db565b9150826002028217905092915050565b6200042482620001a9565b67ffffffffffffffff81111562000440576200043f620001b4565b5b6200044c825462000212565b620004598282856200037e565b600060209050601f8311600181146200049157600084156200047c578287015190505b620004888582620003fb565b865550620004f8565b601f198416620004a18662000247565b60005b82811015620004cb57848901518255600182019150602085019450602081019050620004a4565b86831015620004eb5784890151620004e7601f891682620003db565b8355505b6001600288020188555050505b505050505050565b61061780620005106000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630121b93f146100515780632d35a8a21461006d5780633477ee2e1461008b578063a3ec138d146100bd575b600080fd5b61006b60048036038101906100669190610324565b6100ed565b005b610075610211565b6040516100829190610360565b60405180910390f35b6100a560048036038101906100a09190610324565b610217565b6040516100b49392919061040b565b60405180910390f35b6100d760048036038101906100d291906104a7565b6102c9565b6040516100e491906104ef565b60405180910390f35b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561014357600080fd5b60008111801561015557506002548111155b61015e57600080fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600082815260200190815260200160002060020160008154809291906101dc90610539565b9190505550807ffff3c900d938d21d0990d786e819f29b8d05c1ef587b462b939609625b684b1660405160405180910390a250565b60025481565b6001602052806000526040600020600091509050806000015490806001018054610240906105b0565b80601f016020809104026020016040519081016040528092919081815260200182805461026c906105b0565b80156102b95780601f1061028e576101008083540402835291602001916102b9565b820191906000526020600020905b81548152906001019060200180831161029c57829003601f168201915b5050505050908060020154905083565b60006020528060005260406000206000915054906101000a900460ff1681565b600080fd5b6000819050919050565b610301816102ee565b811461030c57600080fd5b50565b60008135905061031e816102f8565b92915050565b60006020828403121561033a576103396102e9565b5b60006103488482850161030f565b91505092915050565b61035a816102ee565b82525050565b60006020820190506103756000830184610351565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156103b557808201518184015260208101905061039a565b60008484015250505050565b6000601f19601f8301169050919050565b60006103dd8261037b565b6103e78185610386565b93506103f7818560208601610397565b610400816103c1565b840191505092915050565b60006060820190506104206000830186610351565b818103602083015261043281856103d2565b90506104416040830184610351565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061047482610449565b9050919050565b61048481610469565b811461048f57600080fd5b50565b6000813590506104a18161047b565b92915050565b6000602082840312156104bd576104bc6102e9565b5b60006104cb84828501610492565b91505092915050565b60008115159050919050565b6104e9816104d4565b82525050565b600060208201905061050460008301846104e0565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610544826102ee565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036105765761057561050a565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806105c857607f821691505b6020821081036105db576105da610581565b5b5091905056fea264697066735822122028802a79305550bf2b7100c3caa07690e291e07136506beee741ad184ae7e07064736f6c63430008110033",
}

// BallotABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotMetaData.ABI instead.
var BallotABI = BallotMetaData.ABI

// BallotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BallotMetaData.Bin instead.
var BallotBin = BallotMetaData.Bin

// DeployBallot deploys a new Ethereum contract, binding an instance of Ballot to it.
func DeployBallot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ballot, error) {
	parsed, err := BallotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BallotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// Ballot is an auto generated Go binding around an Ethereum contract.
type Ballot struct {
	BallotCaller     // Read-only binding to the contract
	BallotTransactor // Write-only binding to the contract
	BallotFilterer   // Log filterer for contract events
}

// BallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotSession struct {
	Contract     *Ballot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotCallerSession struct {
	Contract *BallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTransactorSession struct {
	Contract     *BallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotRaw struct {
	Contract *Ballot // Generic contract binding to access the raw methods on
}

// BallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotCallerRaw struct {
	Contract *BallotCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTransactorRaw struct {
	Contract *BallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallot creates a new instance of Ballot, bound to a specific deployed contract.
func NewBallot(address common.Address, backend bind.ContractBackend) (*Ballot, error) {
	contract, err := bindBallot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// NewBallotCaller creates a new read-only instance of Ballot, bound to a specific deployed contract.
func NewBallotCaller(address common.Address, caller bind.ContractCaller) (*BallotCaller, error) {
	contract, err := bindBallot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotCaller{contract: contract}, nil
}

// NewBallotTransactor creates a new write-only instance of Ballot, bound to a specific deployed contract.
func NewBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTransactor, error) {
	contract, err := bindBallot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTransactor{contract: contract}, nil
}

// NewBallotFilterer creates a new log filterer instance of Ballot, bound to a specific deployed contract.
func NewBallotFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotFilterer, error) {
	contract, err := bindBallot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotFilterer{contract: contract}, nil
}

// bindBallot binds a generic wrapper to an already deployed contract.
func bindBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BallotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.BallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transact(opts, method, params...)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Ballot *BallotCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Ballot *BallotSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Candidates(&_Ballot.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(uint256 id, string name, uint256 voteCount)
func (_Ballot *BallotCallerSession) Candidates(arg0 *big.Int) (struct {
	Id        *big.Int
	Name      string
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Candidates(&_Ballot.CallOpts, arg0)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Ballot *BallotCaller) CandidatesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "candidatesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Ballot *BallotSession) CandidatesCount() (*big.Int, error) {
	return _Ballot.Contract.CandidatesCount(&_Ballot.CallOpts)
}

// CandidatesCount is a free data retrieval call binding the contract method 0x2d35a8a2.
//
// Solidity: function candidatesCount() view returns(uint256)
func (_Ballot *BallotCallerSession) CandidatesCount() (*big.Int, error) {
	return _Ballot.Contract.CandidatesCount(&_Ballot.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Ballot *BallotCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Ballot *BallotSession) Voters(arg0 common.Address) (bool, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool)
func (_Ballot *BallotCallerSession) Voters(arg0 common.Address) (bool, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Ballot *BallotTransactor) Vote(opts *bind.TransactOpts, _candidateId *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "vote", _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Ballot *BallotSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, _candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateId) returns()
func (_Ballot *BallotTransactorSession) Vote(_candidateId *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, _candidateId)
}

// BallotVotedEventIterator is returned from FilterVotedEvent and is used to iterate over the raw logs and unpacked data for VotedEvent events raised by the Ballot contract.
type BallotVotedEventIterator struct {
	Event *BallotVotedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BallotVotedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BallotVotedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BallotVotedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BallotVotedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BallotVotedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BallotVotedEvent represents a VotedEvent event raised by the Ballot contract.
type BallotVotedEvent struct {
	CandidateId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVotedEvent is a free log retrieval operation binding the contract event 0xfff3c900d938d21d0990d786e819f29b8d05c1ef587b462b939609625b684b16.
//
// Solidity: event votedEvent(uint256 indexed _candidateId)
func (_Ballot *BallotFilterer) FilterVotedEvent(opts *bind.FilterOpts, _candidateId []*big.Int) (*BallotVotedEventIterator, error) {

	var _candidateIdRule []interface{}
	for _, _candidateIdItem := range _candidateId {
		_candidateIdRule = append(_candidateIdRule, _candidateIdItem)
	}

	logs, sub, err := _Ballot.contract.FilterLogs(opts, "votedEvent", _candidateIdRule)
	if err != nil {
		return nil, err
	}
	return &BallotVotedEventIterator{contract: _Ballot.contract, event: "votedEvent", logs: logs, sub: sub}, nil
}

// WatchVotedEvent is a free log subscription operation binding the contract event 0xfff3c900d938d21d0990d786e819f29b8d05c1ef587b462b939609625b684b16.
//
// Solidity: event votedEvent(uint256 indexed _candidateId)
func (_Ballot *BallotFilterer) WatchVotedEvent(opts *bind.WatchOpts, sink chan<- *BallotVotedEvent, _candidateId []*big.Int) (event.Subscription, error) {

	var _candidateIdRule []interface{}
	for _, _candidateIdItem := range _candidateId {
		_candidateIdRule = append(_candidateIdRule, _candidateIdItem)
	}

	logs, sub, err := _Ballot.contract.WatchLogs(opts, "votedEvent", _candidateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BallotVotedEvent)
				if err := _Ballot.contract.UnpackLog(event, "votedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotedEvent is a log parse operation binding the contract event 0xfff3c900d938d21d0990d786e819f29b8d05c1ef587b462b939609625b684b16.
//
// Solidity: event votedEvent(uint256 indexed _candidateId)
func (_Ballot *BallotFilterer) ParseVotedEvent(log types.Log) (*BallotVotedEvent, error) {
	event := new(BallotVotedEvent)
	if err := _Ballot.contract.UnpackLog(event, "votedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
