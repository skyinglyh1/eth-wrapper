// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polywrapper

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPolyWrapperABI is the input ABI used to generate the binding from.
const IPolyWrapperABI = "[{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"toAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"txHash\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"speedUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IPolyWrapperFuncSigs maps the 4-byte function signature to its string representation.
var IPolyWrapperFuncSigs = map[string]string{
	"c415b95c": "feeCollector()",
	"f30f4889": "lock(address,uint64,bytes,uint256,uint256)",
	"9d4dc021": "lockProxy()",
	"5c975abb": "paused()",
	"d3ed7c76": "speedUp(address,bytes,uint256)",
}

// IPolyWrapper is an auto generated Go binding around an Ethereum contract.
type IPolyWrapper struct {
	IPolyWrapperCaller     // Read-only binding to the contract
	IPolyWrapperTransactor // Write-only binding to the contract
	IPolyWrapperFilterer   // Log filterer for contract events
}

// IPolyWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPolyWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPolyWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPolyWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPolyWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPolyWrapperSession struct {
	Contract     *IPolyWrapper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPolyWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPolyWrapperCallerSession struct {
	Contract *IPolyWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPolyWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPolyWrapperTransactorSession struct {
	Contract     *IPolyWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPolyWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPolyWrapperRaw struct {
	Contract *IPolyWrapper // Generic contract binding to access the raw methods on
}

// IPolyWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPolyWrapperCallerRaw struct {
	Contract *IPolyWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// IPolyWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPolyWrapperTransactorRaw struct {
	Contract *IPolyWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPolyWrapper creates a new instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapper(address common.Address, backend bind.ContractBackend) (*IPolyWrapper, error) {
	contract, err := bindIPolyWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapper{IPolyWrapperCaller: IPolyWrapperCaller{contract: contract}, IPolyWrapperTransactor: IPolyWrapperTransactor{contract: contract}, IPolyWrapperFilterer: IPolyWrapperFilterer{contract: contract}}, nil
}

// NewIPolyWrapperCaller creates a new read-only instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperCaller(address common.Address, caller bind.ContractCaller) (*IPolyWrapperCaller, error) {
	contract, err := bindIPolyWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperCaller{contract: contract}, nil
}

// NewIPolyWrapperTransactor creates a new write-only instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IPolyWrapperTransactor, error) {
	contract, err := bindIPolyWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperTransactor{contract: contract}, nil
}

// NewIPolyWrapperFilterer creates a new log filterer instance of IPolyWrapper, bound to a specific deployed contract.
func NewIPolyWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IPolyWrapperFilterer, error) {
	contract, err := bindIPolyWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPolyWrapperFilterer{contract: contract}, nil
}

// bindIPolyWrapper binds a generic wrapper to an already deployed contract.
func bindIPolyWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPolyWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolyWrapper *IPolyWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPolyWrapper.Contract.IPolyWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolyWrapper *IPolyWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.IPolyWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolyWrapper *IPolyWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.IPolyWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPolyWrapper *IPolyWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPolyWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPolyWrapper *IPolyWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPolyWrapper *IPolyWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.contract.Transact(opts, method, params...)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_IPolyWrapper *IPolyWrapperCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IPolyWrapper.contract.Call(opts, out, "feeCollector")
	return *ret0, err
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_IPolyWrapper *IPolyWrapperSession) FeeCollector() (common.Address, error) {
	return _IPolyWrapper.Contract.FeeCollector(&_IPolyWrapper.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_IPolyWrapper *IPolyWrapperCallerSession) FeeCollector() (common.Address, error) {
	return _IPolyWrapper.Contract.FeeCollector(&_IPolyWrapper.CallOpts)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_IPolyWrapper *IPolyWrapperCaller) LockProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IPolyWrapper.contract.Call(opts, out, "lockProxy")
	return *ret0, err
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_IPolyWrapper *IPolyWrapperSession) LockProxy() (common.Address, error) {
	return _IPolyWrapper.Contract.LockProxy(&_IPolyWrapper.CallOpts)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_IPolyWrapper *IPolyWrapperCallerSession) LockProxy() (common.Address, error) {
	return _IPolyWrapper.Contract.LockProxy(&_IPolyWrapper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IPolyWrapper *IPolyWrapperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IPolyWrapper.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IPolyWrapper *IPolyWrapperSession) Paused() (bool, error) {
	return _IPolyWrapper.Contract.Paused(&_IPolyWrapper.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IPolyWrapper *IPolyWrapperCallerSession) Paused() (bool, error) {
	return _IPolyWrapper.Contract.Paused(&_IPolyWrapper.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf30f4889.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperTransactor) Lock(opts *bind.TransactOpts, fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "lock", fromAsset, toChainId, toAddress, amount, fee)
}

// Lock is a paid mutator transaction binding the contract method 0xf30f4889.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Lock(&_IPolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee)
}

// Lock is a paid mutator transaction binding the contract method 0xf30f4889.
//
// Solidity: function lock(address fromAsset, uint64 toChainId, bytes toAddress, uint256 amount, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperTransactorSession) Lock(fromAsset common.Address, toChainId uint64, toAddress []byte, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.Lock(&_IPolyWrapper.TransactOpts, fromAsset, toChainId, toAddress, amount, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperTransactor) SpeedUp(opts *bind.TransactOpts, fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.contract.Transact(opts, "speedUp", fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.SpeedUp(&_IPolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

// SpeedUp is a paid mutator transaction binding the contract method 0xd3ed7c76.
//
// Solidity: function speedUp(address fromAsset, bytes txHash, uint256 fee) payable returns()
func (_IPolyWrapper *IPolyWrapperTransactorSession) SpeedUp(fromAsset common.Address, txHash []byte, fee *big.Int) (*types.Transaction, error) {
	return _IPolyWrapper.Contract.SpeedUp(&_IPolyWrapper.TransactOpts, fromAsset, txHash, fee)
}

