// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package homework

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

// HomeworkMetaData contains all meta data concerning the Homework contract.
var HomeworkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"addNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101d68061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80634e70b1dc1461004357806367e0badb14610061578063efcfe81a1461007f575b5f5ffd5b61004b610089565b604051610058919061012b565b60405180910390f35b6100696100a1565b604051610076919061012b565b60405180910390f35b6100876100bc565b005b5f5f9054906101000a900467ffffffffffffffff1681565b5f5f5f9054906101000a900467ffffffffffffffff16905090565b5f5f81819054906101000a900467ffffffffffffffff16809291906100e090610171565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b5f67ffffffffffffffff82169050919050565b61012581610109565b82525050565b5f60208201905061013e5f83018461011c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61017b82610109565b915067ffffffffffffffff820361019557610194610144565b5b60018201905091905056fea26469706673582212204eaff5a72c34e662ca4a64a794b6c6a9f77675309d6d85eb2dd02d97807dfecb64736f6c63430008210033",
}

// HomeworkABI is the input ABI used to generate the binding from.
// Deprecated: Use HomeworkMetaData.ABI instead.
var HomeworkABI = HomeworkMetaData.ABI

// HomeworkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HomeworkMetaData.Bin instead.
var HomeworkBin = HomeworkMetaData.Bin

// DeployHomework deploys a new Ethereum contract, binding an instance of Homework to it.
func DeployHomework(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Homework, error) {
	parsed, err := HomeworkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HomeworkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Homework{HomeworkCaller: HomeworkCaller{contract: contract}, HomeworkTransactor: HomeworkTransactor{contract: contract}, HomeworkFilterer: HomeworkFilterer{contract: contract}}, nil
}

// Homework is an auto generated Go binding around an Ethereum contract.
type Homework struct {
	HomeworkCaller     // Read-only binding to the contract
	HomeworkTransactor // Write-only binding to the contract
	HomeworkFilterer   // Log filterer for contract events
}

// HomeworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type HomeworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HomeworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HomeworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HomeworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HomeworkSession struct {
	Contract     *Homework         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HomeworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HomeworkCallerSession struct {
	Contract *HomeworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HomeworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HomeworkTransactorSession struct {
	Contract     *HomeworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HomeworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type HomeworkRaw struct {
	Contract *Homework // Generic contract binding to access the raw methods on
}

// HomeworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HomeworkCallerRaw struct {
	Contract *HomeworkCaller // Generic read-only contract binding to access the raw methods on
}

// HomeworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HomeworkTransactorRaw struct {
	Contract *HomeworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHomework creates a new instance of Homework, bound to a specific deployed contract.
func NewHomework(address common.Address, backend bind.ContractBackend) (*Homework, error) {
	contract, err := bindHomework(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Homework{HomeworkCaller: HomeworkCaller{contract: contract}, HomeworkTransactor: HomeworkTransactor{contract: contract}, HomeworkFilterer: HomeworkFilterer{contract: contract}}, nil
}

// NewHomeworkCaller creates a new read-only instance of Homework, bound to a specific deployed contract.
func NewHomeworkCaller(address common.Address, caller bind.ContractCaller) (*HomeworkCaller, error) {
	contract, err := bindHomework(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HomeworkCaller{contract: contract}, nil
}

// NewHomeworkTransactor creates a new write-only instance of Homework, bound to a specific deployed contract.
func NewHomeworkTransactor(address common.Address, transactor bind.ContractTransactor) (*HomeworkTransactor, error) {
	contract, err := bindHomework(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HomeworkTransactor{contract: contract}, nil
}

// NewHomeworkFilterer creates a new log filterer instance of Homework, bound to a specific deployed contract.
func NewHomeworkFilterer(address common.Address, filterer bind.ContractFilterer) (*HomeworkFilterer, error) {
	contract, err := bindHomework(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HomeworkFilterer{contract: contract}, nil
}

// bindHomework binds a generic wrapper to an already deployed contract.
func bindHomework(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HomeworkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Homework *HomeworkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Homework.Contract.HomeworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Homework *HomeworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Homework.Contract.HomeworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Homework *HomeworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Homework.Contract.HomeworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Homework *HomeworkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Homework.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Homework *HomeworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Homework.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Homework *HomeworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Homework.Contract.contract.Transact(opts, method, params...)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint64)
func (_Homework *HomeworkCaller) GetNum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Homework.contract.Call(opts, &out, "getNum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint64)
func (_Homework *HomeworkSession) GetNum() (uint64, error) {
	return _Homework.Contract.GetNum(&_Homework.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint64)
func (_Homework *HomeworkCallerSession) GetNum() (uint64, error) {
	return _Homework.Contract.GetNum(&_Homework.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint64)
func (_Homework *HomeworkCaller) Num(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Homework.contract.Call(opts, &out, "num")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint64)
func (_Homework *HomeworkSession) Num() (uint64, error) {
	return _Homework.Contract.Num(&_Homework.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint64)
func (_Homework *HomeworkCallerSession) Num() (uint64, error) {
	return _Homework.Contract.Num(&_Homework.CallOpts)
}

// AddNum is a paid mutator transaction binding the contract method 0xefcfe81a.
//
// Solidity: function addNum() returns()
func (_Homework *HomeworkTransactor) AddNum(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Homework.contract.Transact(opts, "addNum")
}

// AddNum is a paid mutator transaction binding the contract method 0xefcfe81a.
//
// Solidity: function addNum() returns()
func (_Homework *HomeworkSession) AddNum() (*types.Transaction, error) {
	return _Homework.Contract.AddNum(&_Homework.TransactOpts)
}

// AddNum is a paid mutator transaction binding the contract method 0xefcfe81a.
//
// Solidity: function addNum() returns()
func (_Homework *HomeworkTransactorSession) AddNum() (*types.Transaction, error) {
	return _Homework.Contract.AddNum(&_Homework.TransactOpts)
}
