// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// ERC20TestMetaData contains all meta data concerning the ERC20Test contract.
var ERC20TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assertAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161097f38038061097f8339818101604052810190610032919061011d565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ea826100bf565b9050919050565b6100fa816100df565b811461010557600080fd5b50565b600081519050610117816100f1565b92915050565b60008060408385031215610134576101336100ba565b5b600061014285828601610108565b925050602061015385828601610108565b9150509250929050565b6108138061016c6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c040622614610030575b600080fd5b61003861003a565b005b61004261004c565b61004a6101bc565b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346bdca9a600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b8152600401600060405180830381865afa1580156100f5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061011e91906104a1565b6040518263ffffffff1660e01b815260040161013a919061058b565b602060405180830381865afa158015610157573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061017b91906105f8565b6101ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b190610671565b60405180910390fd5b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166330191626600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610265573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028991906106c7565b620f42406040518363ffffffff1660e01b81526004016102aa929190610748565b602060405180830381865afa1580156102c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102eb91906105f8565b61032a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610321906107bd565b60405180910390fd5b565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103938261034a565b810181811067ffffffffffffffff821117156103b2576103b161035b565b5b80604052505050565b60006103c561032c565b90506103d1828261038a565b919050565b600067ffffffffffffffff8211156103f1576103f061035b565b5b6103fa8261034a565b9050602081019050919050565b60005b8381101561042557808201518184015260208101905061040a565b60008484015250505050565b600061044461043f846103d6565b6103bb565b9050828152602081018484840111156104605761045f610345565b5b61046b848285610407565b509392505050565b600082601f83011261048857610487610340565b5b8151610498848260208601610431565b91505092915050565b6000602082840312156104b7576104b6610336565b5b600082015167ffffffffffffffff8111156104d5576104d461033b565b5b6104e184828501610473565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000610511826104ea565b61051b81856104f5565b935061052b818560208601610407565b6105348161034a565b840191505092915050565b7f4578616d706c6520546f6b656e00000000000000000000000000000000000000600082015250565b6000610575600d836104f5565b91506105808261053f565b602082019050919050565b600060408201905081810360008301526105a58184610506565b905081810360208301526105b881610568565b905092915050565b60008115159050919050565b6105d5816105c0565b81146105e057600080fd5b50565b6000815190506105f2816105cc565b92915050565b60006020828403121561060e5761060d610336565b5b600061061c848285016105e3565b91505092915050565b7f6e616d65206d69736d6174636800000000000000000000000000000000000000600082015250565b600061065b600d836104f5565b915061066682610625565b602082019050919050565b6000602082019050818103600083015261068a8161064e565b9050919050565b6000819050919050565b6106a481610691565b81146106af57600080fd5b50565b6000815190506106c18161069b565b92915050565b6000602082840312156106dd576106dc610336565b5b60006106eb848285016106b2565b91505092915050565b6106fd81610691565b82525050565b6000819050919050565b6000819050919050565b600061073261072d61072884610703565b61070d565b610691565b9050919050565b61074281610717565b82525050565b600060408201905061075d60008301856106f4565b61076a6020830184610739565b9392505050565b7f746f74616c20737570706c79206d69736d617463680000000000000000000000600082015250565b60006107a76015836104f5565b91506107b282610771565b602082019050919050565b600060208201905081810360008301526107d68161079a565b905091905056fea26469706673582212205d294012691a9f8224f001c0e6091a2221ae19fa6b1f4d14eba1de552840dc1564736f6c63430008190033",
}

// ERC20TestABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TestMetaData.ABI instead.
var ERC20TestABI = ERC20TestMetaData.ABI

// ERC20TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TestMetaData.Bin instead.
var ERC20TestBin = ERC20TestMetaData.Bin

// DeployERC20Test deploys a new Ethereum contract, binding an instance of ERC20Test to it.
func DeployERC20Test(auth *bind.TransactOpts, backend bind.ContractBackend, assertAddr common.Address, tokenAddr common.Address) (common.Address, *types.Transaction, *ERC20Test, error) {
	parsed, err := ERC20TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TestBin), backend, assertAddr, tokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Test{ERC20TestCaller: ERC20TestCaller{contract: contract}, ERC20TestTransactor: ERC20TestTransactor{contract: contract}, ERC20TestFilterer: ERC20TestFilterer{contract: contract}}, nil
}

// ERC20Test is an auto generated Go binding around an Ethereum contract.
type ERC20Test struct {
	ERC20TestCaller     // Read-only binding to the contract
	ERC20TestTransactor // Write-only binding to the contract
	ERC20TestFilterer   // Log filterer for contract events
}

// ERC20TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TestSession struct {
	Contract     *ERC20Test        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TestCallerSession struct {
	Contract *ERC20TestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TestTransactorSession struct {
	Contract     *ERC20TestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TestRaw struct {
	Contract *ERC20Test // Generic contract binding to access the raw methods on
}

// ERC20TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TestCallerRaw struct {
	Contract *ERC20TestCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TestTransactorRaw struct {
	Contract *ERC20TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Test creates a new instance of ERC20Test, bound to a specific deployed contract.
func NewERC20Test(address common.Address, backend bind.ContractBackend) (*ERC20Test, error) {
	contract, err := bindERC20Test(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Test{ERC20TestCaller: ERC20TestCaller{contract: contract}, ERC20TestTransactor: ERC20TestTransactor{contract: contract}, ERC20TestFilterer: ERC20TestFilterer{contract: contract}}, nil
}

// NewERC20TestCaller creates a new read-only instance of ERC20Test, bound to a specific deployed contract.
func NewERC20TestCaller(address common.Address, caller bind.ContractCaller) (*ERC20TestCaller, error) {
	contract, err := bindERC20Test(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TestCaller{contract: contract}, nil
}

// NewERC20TestTransactor creates a new write-only instance of ERC20Test, bound to a specific deployed contract.
func NewERC20TestTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TestTransactor, error) {
	contract, err := bindERC20Test(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TestTransactor{contract: contract}, nil
}

// NewERC20TestFilterer creates a new log filterer instance of ERC20Test, bound to a specific deployed contract.
func NewERC20TestFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TestFilterer, error) {
	contract, err := bindERC20Test(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TestFilterer{contract: contract}, nil
}

// bindERC20Test binds a generic wrapper to an already deployed contract.
func bindERC20Test(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Test *ERC20TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Test.Contract.ERC20TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Test *ERC20TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Test.Contract.ERC20TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Test *ERC20TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Test.Contract.ERC20TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Test *ERC20TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Test *ERC20TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Test *ERC20TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Test.Contract.contract.Transact(opts, method, params...)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_ERC20Test *ERC20TestCaller) Run(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ERC20Test.contract.Call(opts, &out, "run")

	if err != nil {
		return err
	}

	return err

}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_ERC20Test *ERC20TestSession) Run() error {
	return _ERC20Test.Contract.Run(&_ERC20Test.CallOpts)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_ERC20Test *ERC20TestCallerSession) Run() error {
	return _ERC20Test.Contract.Run(&_ERC20Test.CallOpts)
}
