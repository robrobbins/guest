// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multi

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
	_ = abi.ConvertType
)

// MultiMetaData contains all meta data concerning the Multi contract.
var MultiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InputLengthMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"MultiTransfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506105238061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806375992d3c14610030575b600080fd5b61004a6004803603810190610045919061030e565b610060565b60405161005791906103be565b60405180910390f35b6000808686905090508085859050146100a5576040517faaad13f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600083905060005b82811015610181578173ffffffffffffffffffffffffffffffffffffffff166323b872dd338b8b858181106100e5576100e46103d9565b5b90506020020160208101906100fa9190610408565b8a8a8681811061010d5761010c6103d9565b5b905060200201356040518463ffffffff1660e01b81526004016101329392919061045d565b6020604051808303816000875af1158015610151573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017591906104c0565b508060010190506100ad565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f3fb7fba08f5ceba290b54d7b90d0ca84ed80e78724e0dd7c522f2587c71906fb60405160405180910390a360019250505095945050505050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261021a576102196101f5565b5b8235905067ffffffffffffffff811115610237576102366101fa565b5b602083019150836020820283011115610253576102526101ff565b5b9250929050565b60008083601f8401126102705761026f6101f5565b5b8235905067ffffffffffffffff81111561028d5761028c6101fa565b5b6020830191508360208202830111156102a9576102a86101ff565b5b9250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102db826102b0565b9050919050565b6102eb816102d0565b81146102f657600080fd5b50565b600081359050610308816102e2565b92915050565b60008060008060006060868803121561032a576103296101eb565b5b600086013567ffffffffffffffff811115610348576103476101f0565b5b61035488828901610204565b9550955050602086013567ffffffffffffffff811115610377576103766101f0565b5b6103838882890161025a565b93509350506040610396888289016102f9565b9150509295509295909350565b60008115159050919050565b6103b8816103a3565b82525050565b60006020820190506103d360008301846103af565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561041e5761041d6101eb565b5b600061042c848285016102f9565b91505092915050565b61043e816102d0565b82525050565b6000819050919050565b61045781610444565b82525050565b60006060820190506104726000830186610435565b61047f6020830185610435565b61048c604083018461044e565b949350505050565b61049d816103a3565b81146104a857600080fd5b50565b6000815190506104ba81610494565b92915050565b6000602082840312156104d6576104d56101eb565b5b60006104e4848285016104ab565b9150509291505056fea2646970667358221220acca6f6e7d59efb23b5bc8e602b8fcc4a42abf2ac65895fc74bfb0ee3c00631664736f6c63430008190033",
}

// MultiABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiMetaData.ABI instead.
var MultiABI = MultiMetaData.ABI

// MultiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiMetaData.Bin instead.
var MultiBin = MultiMetaData.Bin

// DeployMulti deploys a new Ethereum contract, binding an instance of Multi to it.
func DeployMulti(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multi, error) {
	parsed, err := MultiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multi{MultiCaller: MultiCaller{contract: contract}, MultiTransactor: MultiTransactor{contract: contract}, MultiFilterer: MultiFilterer{contract: contract}}, nil
}

// Multi is an auto generated Go binding around an Ethereum contract.
type Multi struct {
	MultiCaller     // Read-only binding to the contract
	MultiTransactor // Write-only binding to the contract
	MultiFilterer   // Log filterer for contract events
}

// MultiCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSession struct {
	Contract     *Multi            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallerSession struct {
	Contract *MultiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MultiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiTransactorSession struct {
	Contract     *MultiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiRaw struct {
	Contract *Multi // Generic contract binding to access the raw methods on
}

// MultiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallerRaw struct {
	Contract *MultiCaller // Generic read-only contract binding to access the raw methods on
}

// MultiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiTransactorRaw struct {
	Contract *MultiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulti creates a new instance of Multi, bound to a specific deployed contract.
func NewMulti(address common.Address, backend bind.ContractBackend) (*Multi, error) {
	contract, err := bindMulti(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multi{MultiCaller: MultiCaller{contract: contract}, MultiTransactor: MultiTransactor{contract: contract}, MultiFilterer: MultiFilterer{contract: contract}}, nil
}

// NewMultiCaller creates a new read-only instance of Multi, bound to a specific deployed contract.
func NewMultiCaller(address common.Address, caller bind.ContractCaller) (*MultiCaller, error) {
	contract, err := bindMulti(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCaller{contract: contract}, nil
}

// NewMultiTransactor creates a new write-only instance of Multi, bound to a specific deployed contract.
func NewMultiTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiTransactor, error) {
	contract, err := bindMulti(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiTransactor{contract: contract}, nil
}

// NewMultiFilterer creates a new log filterer instance of Multi, bound to a specific deployed contract.
func NewMultiFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiFilterer, error) {
	contract, err := bindMulti(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiFilterer{contract: contract}, nil
}

// bindMulti binds a generic wrapper to an already deployed contract.
func bindMulti(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multi *MultiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multi.Contract.MultiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multi *MultiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multi.Contract.MultiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multi *MultiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multi.Contract.MultiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multi *MultiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multi *MultiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multi *MultiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multi.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x75992d3c.
//
// Solidity: function transfer(address[] recipients, uint256[] amounts, address token) returns(bool)
func (_Multi *MultiTransactor) Transfer(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Multi.contract.Transact(opts, "transfer", recipients, amounts, token)
}

// Transfer is a paid mutator transaction binding the contract method 0x75992d3c.
//
// Solidity: function transfer(address[] recipients, uint256[] amounts, address token) returns(bool)
func (_Multi *MultiSession) Transfer(recipients []common.Address, amounts []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Multi.Contract.Transfer(&_Multi.TransactOpts, recipients, amounts, token)
}

// Transfer is a paid mutator transaction binding the contract method 0x75992d3c.
//
// Solidity: function transfer(address[] recipients, uint256[] amounts, address token) returns(bool)
func (_Multi *MultiTransactorSession) Transfer(recipients []common.Address, amounts []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Multi.Contract.Transfer(&_Multi.TransactOpts, recipients, amounts, token)
}

// MultiMultiTransferIterator is returned from FilterMultiTransfer and is used to iterate over the raw logs and unpacked data for MultiTransfer events raised by the Multi contract.
type MultiMultiTransferIterator struct {
	Event *MultiMultiTransfer // Event containing the contract specifics and raw log

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
func (it *MultiMultiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiMultiTransfer)
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
		it.Event = new(MultiMultiTransfer)
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
func (it *MultiMultiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiMultiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiMultiTransfer represents a MultiTransfer event raised by the Multi contract.
type MultiMultiTransfer struct {
	Caller common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMultiTransfer is a free log retrieval operation binding the contract event 0x3fb7fba08f5ceba290b54d7b90d0ca84ed80e78724e0dd7c522f2587c71906fb.
//
// Solidity: event MultiTransfer(address indexed caller, address indexed token)
func (_Multi *MultiFilterer) FilterMultiTransfer(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*MultiMultiTransferIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Multi.contract.FilterLogs(opts, "MultiTransfer", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiMultiTransferIterator{contract: _Multi.contract, event: "MultiTransfer", logs: logs, sub: sub}, nil
}

// WatchMultiTransfer is a free log subscription operation binding the contract event 0x3fb7fba08f5ceba290b54d7b90d0ca84ed80e78724e0dd7c522f2587c71906fb.
//
// Solidity: event MultiTransfer(address indexed caller, address indexed token)
func (_Multi *MultiFilterer) WatchMultiTransfer(opts *bind.WatchOpts, sink chan<- *MultiMultiTransfer, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Multi.contract.WatchLogs(opts, "MultiTransfer", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiMultiTransfer)
				if err := _Multi.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
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

// ParseMultiTransfer is a log parse operation binding the contract event 0x3fb7fba08f5ceba290b54d7b90d0ca84ed80e78724e0dd7c522f2587c71906fb.
//
// Solidity: event MultiTransfer(address indexed caller, address indexed token)
func (_Multi *MultiFilterer) ParseMultiTransfer(log types.Log) (*MultiMultiTransfer, error) {
	event := new(MultiMultiTransfer)
	if err := _Multi.contract.UnpackLog(event, "MultiTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
