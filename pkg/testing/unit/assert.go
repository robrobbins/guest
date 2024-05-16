// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unit

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

// AssertMetaData contains all meta data concerning the Assert contract.
var AssertMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"equal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506103ef8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063301916261461004657806346bdca9a14610076578063df6ea8cf146100a6575b600080fd5b610060600480360381019061005b919061018f565b6100d6565b60405161006d91906101ea565b60405180910390f35b610090600480360381019061008b919061026a565b6100e3565b60405161009d91906101ea565b60405180910390f35b6100c060048036038101906100bb9190610321565b610142565b6040516100cd91906101ea565b60405180910390f35b6000818314905092915050565b600082826040516020016100f89291906103a0565b6040516020818303038152906040528051906020012085856040516020016101219291906103a0565b60405160208183030381529060405280519060200120149050949350505050565b6000818314905092915050565b600080fd5b600080fd5b6000819050919050565b61016c81610159565b811461017757600080fd5b50565b60008135905061018981610163565b92915050565b600080604083850312156101a6576101a561014f565b5b60006101b48582860161017a565b92505060206101c58582860161017a565b9150509250929050565b60008115159050919050565b6101e4816101cf565b82525050565b60006020820190506101ff60008301846101db565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261022a57610229610205565b5b8235905067ffffffffffffffff8111156102475761024661020a565b5b6020830191508360018202830111156102635761026261020f565b5b9250929050565b600080600080604085870312156102845761028361014f565b5b600085013567ffffffffffffffff8111156102a2576102a1610154565b5b6102ae87828801610214565b9450945050602085013567ffffffffffffffff8111156102d1576102d0610154565b5b6102dd87828801610214565b925092505092959194509250565b6000819050919050565b6102fe816102eb565b811461030957600080fd5b50565b60008135905061031b816102f5565b92915050565b600080604083850312156103385761033761014f565b5b60006103468582860161030c565b92505060206103578582860161030c565b9150509250929050565b600081905092915050565b82818337600083830152505050565b60006103878385610361565b935061039483858461036c565b82840190509392505050565b60006103ad82848661037b565b9150819050939250505056fea264697066735822122096b164e81e742e19f331ee95b2f98ae49f1b8b4eb8ea4bf27382a1dcd445ee6e64736f6c63430008190033",
}

// AssertABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertMetaData.ABI instead.
var AssertABI = AssertMetaData.ABI

// AssertBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssertMetaData.Bin instead.
var AssertBin = AssertMetaData.Bin

// DeployAssert deploys a new Ethereum contract, binding an instance of Assert to it.
func DeployAssert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Assert, error) {
	parsed, err := AssertMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assert{AssertCaller: AssertCaller{contract: contract}, AssertTransactor: AssertTransactor{contract: contract}, AssertFilterer: AssertFilterer{contract: contract}}, nil
}

// Assert is an auto generated Go binding around an Ethereum contract.
type Assert struct {
	AssertCaller     // Read-only binding to the contract
	AssertTransactor // Write-only binding to the contract
	AssertFilterer   // Log filterer for contract events
}

// AssertCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertSession struct {
	Contract     *Assert           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertCallerSession struct {
	Contract *AssertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertTransactorSession struct {
	Contract     *AssertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertRaw struct {
	Contract *Assert // Generic contract binding to access the raw methods on
}

// AssertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertCallerRaw struct {
	Contract *AssertCaller // Generic read-only contract binding to access the raw methods on
}

// AssertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertTransactorRaw struct {
	Contract *AssertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssert creates a new instance of Assert, bound to a specific deployed contract.
func NewAssert(address common.Address, backend bind.ContractBackend) (*Assert, error) {
	contract, err := bindAssert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assert{AssertCaller: AssertCaller{contract: contract}, AssertTransactor: AssertTransactor{contract: contract}, AssertFilterer: AssertFilterer{contract: contract}}, nil
}

// NewAssertCaller creates a new read-only instance of Assert, bound to a specific deployed contract.
func NewAssertCaller(address common.Address, caller bind.ContractCaller) (*AssertCaller, error) {
	contract, err := bindAssert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertCaller{contract: contract}, nil
}

// NewAssertTransactor creates a new write-only instance of Assert, bound to a specific deployed contract.
func NewAssertTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertTransactor, error) {
	contract, err := bindAssert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertTransactor{contract: contract}, nil
}

// NewAssertFilterer creates a new log filterer instance of Assert, bound to a specific deployed contract.
func NewAssertFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertFilterer, error) {
	contract, err := bindAssert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertFilterer{contract: contract}, nil
}

// bindAssert binds a generic wrapper to an already deployed contract.
func bindAssert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssertMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assert *AssertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assert.Contract.AssertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assert *AssertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assert.Contract.AssertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assert *AssertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assert.Contract.AssertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assert *AssertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assert *AssertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assert *AssertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assert.Contract.contract.Transact(opts, method, params...)
}

// Equal is a free data retrieval call binding the contract method 0x30191626.
//
// Solidity: function equal(uint256 a, uint256 b) pure returns(bool)
func (_Assert *AssertCaller) Equal(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _Assert.contract.Call(opts, &out, "equal", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equal is a free data retrieval call binding the contract method 0x30191626.
//
// Solidity: function equal(uint256 a, uint256 b) pure returns(bool)
func (_Assert *AssertSession) Equal(a *big.Int, b *big.Int) (bool, error) {
	return _Assert.Contract.Equal(&_Assert.CallOpts, a, b)
}

// Equal is a free data retrieval call binding the contract method 0x30191626.
//
// Solidity: function equal(uint256 a, uint256 b) pure returns(bool)
func (_Assert *AssertCallerSession) Equal(a *big.Int, b *big.Int) (bool, error) {
	return _Assert.Contract.Equal(&_Assert.CallOpts, a, b)
}

// Equal0 is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) pure returns(bool)
func (_Assert *AssertCaller) Equal0(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Assert.contract.Call(opts, &out, "equal0", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equal0 is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) pure returns(bool)
func (_Assert *AssertSession) Equal0(a string, b string) (bool, error) {
	return _Assert.Contract.Equal0(&_Assert.CallOpts, a, b)
}

// Equal0 is a free data retrieval call binding the contract method 0x46bdca9a.
//
// Solidity: function equal(string a, string b) pure returns(bool)
func (_Assert *AssertCallerSession) Equal0(a string, b string) (bool, error) {
	return _Assert.Contract.Equal0(&_Assert.CallOpts, a, b)
}

// Equal1 is a free data retrieval call binding the contract method 0xdf6ea8cf.
//
// Solidity: function equal(bytes32 a, bytes32 b) pure returns(bool)
func (_Assert *AssertCaller) Equal1(opts *bind.CallOpts, a [32]byte, b [32]byte) (bool, error) {
	var out []interface{}
	err := _Assert.contract.Call(opts, &out, "equal1", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equal1 is a free data retrieval call binding the contract method 0xdf6ea8cf.
//
// Solidity: function equal(bytes32 a, bytes32 b) pure returns(bool)
func (_Assert *AssertSession) Equal1(a [32]byte, b [32]byte) (bool, error) {
	return _Assert.Contract.Equal1(&_Assert.CallOpts, a, b)
}

// Equal1 is a free data retrieval call binding the contract method 0xdf6ea8cf.
//
// Solidity: function equal(bytes32 a, bytes32 b) pure returns(bool)
func (_Assert *AssertCallerSession) Equal1(a [32]byte, b [32]byte) (bool, error) {
	return _Assert.Contract.Equal1(&_Assert.CallOpts, a, b)
}
