// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ZKBridageMetaData contains all meta data concerning the ZKBridage contract.
var ZKBridageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"ExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"MessagePublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewPendingImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"MESSAGE_TOPIC\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"}],\"name\":\"blockUpdater\",\"outputs\":[{\"internalType\":\"contractIBlockUpdater\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isTransferCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"}],\"name\":\"mptVerifier\",\"outputs\":[{\"internalType\":\"contractIMptVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"}],\"name\":\"nextSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeContract\",\"type\":\"bytes32\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"blockUpdater\",\"type\":\"address\"}],\"name\":\"setBlockUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"}],\"name\":\"setLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"mptVerifier\",\"type\":\"address\"}],\"name\":\"setMptVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"submitContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"srcBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"mptProof\",\"type\":\"bytes\"}],\"name\":\"validateTransactionProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"}],\"name\":\"zkBridgeContracts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ZKBridageABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKBridageMetaData.ABI instead.
var ZKBridageABI = ZKBridageMetaData.ABI

// ZKBridage is an auto generated Go binding around an Ethereum contract.
type ZKBridage struct {
	ZKBridageCaller     // Read-only binding to the contract
	ZKBridageTransactor // Write-only binding to the contract
	ZKBridageFilterer   // Log filterer for contract events
}

// ZKBridageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKBridageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKBridageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKBridageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKBridageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKBridageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKBridageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKBridageSession struct {
	Contract     *ZKBridage        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKBridageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKBridageCallerSession struct {
	Contract *ZKBridageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZKBridageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKBridageTransactorSession struct {
	Contract     *ZKBridageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZKBridageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKBridageRaw struct {
	Contract *ZKBridage // Generic contract binding to access the raw methods on
}

// ZKBridageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKBridageCallerRaw struct {
	Contract *ZKBridageCaller // Generic read-only contract binding to access the raw methods on
}

// ZKBridageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKBridageTransactorRaw struct {
	Contract *ZKBridageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKBridage creates a new instance of ZKBridage, bound to a specific deployed contract.
func NewZKBridage(address common.Address, backend bind.ContractBackend) (*ZKBridage, error) {
	contract, err := bindZKBridage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKBridage{ZKBridageCaller: ZKBridageCaller{contract: contract}, ZKBridageTransactor: ZKBridageTransactor{contract: contract}, ZKBridageFilterer: ZKBridageFilterer{contract: contract}}, nil
}

// NewZKBridageCaller creates a new read-only instance of ZKBridage, bound to a specific deployed contract.
func NewZKBridageCaller(address common.Address, caller bind.ContractCaller) (*ZKBridageCaller, error) {
	contract, err := bindZKBridage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKBridageCaller{contract: contract}, nil
}

// NewZKBridageTransactor creates a new write-only instance of ZKBridage, bound to a specific deployed contract.
func NewZKBridageTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKBridageTransactor, error) {
	contract, err := bindZKBridage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKBridageTransactor{contract: contract}, nil
}

// NewZKBridageFilterer creates a new log filterer instance of ZKBridage, bound to a specific deployed contract.
func NewZKBridageFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKBridageFilterer, error) {
	contract, err := bindZKBridage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKBridageFilterer{contract: contract}, nil
}

// bindZKBridage binds a generic wrapper to an already deployed contract.
func bindZKBridage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZKBridageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKBridage *ZKBridageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKBridage.Contract.ZKBridageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKBridage *ZKBridageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKBridage.Contract.ZKBridageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKBridage *ZKBridageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKBridage.Contract.ZKBridageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKBridage *ZKBridageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKBridage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKBridage *ZKBridageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKBridage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKBridage *ZKBridageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKBridage.Contract.contract.Transact(opts, method, params...)
}

// MESSAGETOPIC is a free data retrieval call binding the contract method 0x3fe3da36.
//
// Solidity: function MESSAGE_TOPIC() view returns(bytes32)
func (_ZKBridage *ZKBridageCaller) MESSAGETOPIC(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "MESSAGE_TOPIC")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGETOPIC is a free data retrieval call binding the contract method 0x3fe3da36.
//
// Solidity: function MESSAGE_TOPIC() view returns(bytes32)
func (_ZKBridage *ZKBridageSession) MESSAGETOPIC() ([32]byte, error) {
	return _ZKBridage.Contract.MESSAGETOPIC(&_ZKBridage.CallOpts)
}

// MESSAGETOPIC is a free data retrieval call binding the contract method 0x3fe3da36.
//
// Solidity: function MESSAGE_TOPIC() view returns(bytes32)
func (_ZKBridage *ZKBridageCallerSession) MESSAGETOPIC() ([32]byte, error) {
	return _ZKBridage.Contract.MESSAGETOPIC(&_ZKBridage.CallOpts)
}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_ZKBridage *ZKBridageCaller) MINLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "MIN_LOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_ZKBridage *ZKBridageSession) MINLOCKTIME() (*big.Int, error) {
	return _ZKBridage.Contract.MINLOCKTIME(&_ZKBridage.CallOpts)
}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_ZKBridage *ZKBridageCallerSession) MINLOCKTIME() (*big.Int, error) {
	return _ZKBridage.Contract.MINLOCKTIME(&_ZKBridage.CallOpts)
}

// BlockUpdater is a free data retrieval call binding the contract method 0x66798c2c.
//
// Solidity: function blockUpdater(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageCaller) BlockUpdater(opts *bind.CallOpts, chainId uint16) (common.Address, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "blockUpdater", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockUpdater is a free data retrieval call binding the contract method 0x66798c2c.
//
// Solidity: function blockUpdater(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageSession) BlockUpdater(chainId uint16) (common.Address, error) {
	return _ZKBridage.Contract.BlockUpdater(&_ZKBridage.CallOpts, chainId)
}

// BlockUpdater is a free data retrieval call binding the contract method 0x66798c2c.
//
// Solidity: function blockUpdater(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageCallerSession) BlockUpdater(chainId uint16) (common.Address, error) {
	return _ZKBridage.Contract.BlockUpdater(&_ZKBridage.CallOpts, chainId)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_ZKBridage *ZKBridageCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_ZKBridage *ZKBridageSession) ChainId() (uint16, error) {
	return _ZKBridage.Contract.ChainId(&_ZKBridage.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_ZKBridage *ZKBridageCallerSession) ChainId() (uint16, error) {
	return _ZKBridage.Contract.ChainId(&_ZKBridage.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_ZKBridage *ZKBridageCaller) IsInitialized(opts *bind.CallOpts, impl common.Address) (bool, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "isInitialized", impl)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_ZKBridage *ZKBridageSession) IsInitialized(impl common.Address) (bool, error) {
	return _ZKBridage.Contract.IsInitialized(&_ZKBridage.CallOpts, impl)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_ZKBridage *ZKBridageCallerSession) IsInitialized(impl common.Address) (bool, error) {
	return _ZKBridage.Contract.IsInitialized(&_ZKBridage.CallOpts, impl)
}

// IsTransferCompleted is a free data retrieval call binding the contract method 0xaa4efa5b.
//
// Solidity: function isTransferCompleted(bytes32 hash) view returns(bool)
func (_ZKBridage *ZKBridageCaller) IsTransferCompleted(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "isTransferCompleted", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferCompleted is a free data retrieval call binding the contract method 0xaa4efa5b.
//
// Solidity: function isTransferCompleted(bytes32 hash) view returns(bool)
func (_ZKBridage *ZKBridageSession) IsTransferCompleted(hash [32]byte) (bool, error) {
	return _ZKBridage.Contract.IsTransferCompleted(&_ZKBridage.CallOpts, hash)
}

// IsTransferCompleted is a free data retrieval call binding the contract method 0xaa4efa5b.
//
// Solidity: function isTransferCompleted(bytes32 hash) view returns(bool)
func (_ZKBridage *ZKBridageCallerSession) IsTransferCompleted(hash [32]byte) (bool, error) {
	return _ZKBridage.Contract.IsTransferCompleted(&_ZKBridage.CallOpts, hash)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_ZKBridage *ZKBridageCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_ZKBridage *ZKBridageSession) LockTime() (*big.Int, error) {
	return _ZKBridage.Contract.LockTime(&_ZKBridage.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_ZKBridage *ZKBridageCallerSession) LockTime() (*big.Int, error) {
	return _ZKBridage.Contract.LockTime(&_ZKBridage.CallOpts)
}

// MptVerifier is a free data retrieval call binding the contract method 0x6b83252a.
//
// Solidity: function mptVerifier(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageCaller) MptVerifier(opts *bind.CallOpts, chainId uint16) (common.Address, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "mptVerifier", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MptVerifier is a free data retrieval call binding the contract method 0x6b83252a.
//
// Solidity: function mptVerifier(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageSession) MptVerifier(chainId uint16) (common.Address, error) {
	return _ZKBridage.Contract.MptVerifier(&_ZKBridage.CallOpts, chainId)
}

// MptVerifier is a free data retrieval call binding the contract method 0x6b83252a.
//
// Solidity: function mptVerifier(uint16 chainId) view returns(address)
func (_ZKBridage *ZKBridageCallerSession) MptVerifier(chainId uint16) (common.Address, error) {
	return _ZKBridage.Contract.MptVerifier(&_ZKBridage.CallOpts, chainId)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_ZKBridage *ZKBridageCaller) NextSequence(opts *bind.CallOpts, emitter common.Address) (uint64, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "nextSequence", emitter)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_ZKBridage *ZKBridageSession) NextSequence(emitter common.Address) (uint64, error) {
	return _ZKBridage.Contract.NextSequence(&_ZKBridage.CallOpts, emitter)
}

// NextSequence is a free data retrieval call binding the contract method 0x4cf842b5.
//
// Solidity: function nextSequence(address emitter) view returns(uint64)
func (_ZKBridage *ZKBridageCallerSession) NextSequence(emitter common.Address) (uint64, error) {
	return _ZKBridage.Contract.NextSequence(&_ZKBridage.CallOpts, emitter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKBridage *ZKBridageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKBridage *ZKBridageSession) Owner() (common.Address, error) {
	return _ZKBridage.Contract.Owner(&_ZKBridage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZKBridage *ZKBridageCallerSession) Owner() (common.Address, error) {
	return _ZKBridage.Contract.Owner(&_ZKBridage.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_ZKBridage *ZKBridageCaller) PendingImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "pendingImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_ZKBridage *ZKBridageSession) PendingImplementation() (common.Address, error) {
	return _ZKBridage.Contract.PendingImplementation(&_ZKBridage.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_ZKBridage *ZKBridageCallerSession) PendingImplementation() (common.Address, error) {
	return _ZKBridage.Contract.PendingImplementation(&_ZKBridage.CallOpts)
}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_ZKBridage *ZKBridageCaller) ToUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "toUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_ZKBridage *ZKBridageSession) ToUpdateTime() (*big.Int, error) {
	return _ZKBridage.Contract.ToUpdateTime(&_ZKBridage.CallOpts)
}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_ZKBridage *ZKBridageCallerSession) ToUpdateTime() (*big.Int, error) {
	return _ZKBridage.Contract.ToUpdateTime(&_ZKBridage.CallOpts)
}

// ZkBridgeContracts is a free data retrieval call binding the contract method 0xfe2db7d0.
//
// Solidity: function zkBridgeContracts(uint16 chainId) view returns(bytes32)
func (_ZKBridage *ZKBridageCaller) ZkBridgeContracts(opts *bind.CallOpts, chainId uint16) ([32]byte, error) {
	var out []interface{}
	err := _ZKBridage.contract.Call(opts, &out, "zkBridgeContracts", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZkBridgeContracts is a free data retrieval call binding the contract method 0xfe2db7d0.
//
// Solidity: function zkBridgeContracts(uint16 chainId) view returns(bytes32)
func (_ZKBridage *ZKBridageSession) ZkBridgeContracts(chainId uint16) ([32]byte, error) {
	return _ZKBridage.Contract.ZkBridgeContracts(&_ZKBridage.CallOpts, chainId)
}

// ZkBridgeContracts is a free data retrieval call binding the contract method 0xfe2db7d0.
//
// Solidity: function zkBridgeContracts(uint16 chainId) view returns(bytes32)
func (_ZKBridage *ZKBridageCallerSession) ZkBridgeContracts(chainId uint16) ([32]byte, error) {
	return _ZKBridage.Contract.ZkBridgeContracts(&_ZKBridage.CallOpts, chainId)
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_ZKBridage *ZKBridageTransactor) ConfirmContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "confirmContractUpgrade")
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_ZKBridage *ZKBridageSession) ConfirmContractUpgrade() (*types.Transaction, error) {
	return _ZKBridage.Contract.ConfirmContractUpgrade(&_ZKBridage.TransactOpts)
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_ZKBridage *ZKBridageTransactorSession) ConfirmContractUpgrade() (*types.Transaction, error) {
	return _ZKBridage.Contract.ConfirmContractUpgrade(&_ZKBridage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKBridage *ZKBridageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKBridage *ZKBridageSession) Initialize() (*types.Transaction, error) {
	return _ZKBridage.Contract.Initialize(&_ZKBridage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ZKBridage *ZKBridageTransactorSession) Initialize() (*types.Transaction, error) {
	return _ZKBridage.Contract.Initialize(&_ZKBridage.TransactOpts)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x65bb3ea7.
//
// Solidity: function registerChain(uint16 chainId, bytes32 bridgeContract) returns()
func (_ZKBridage *ZKBridageTransactor) RegisterChain(opts *bind.TransactOpts, chainId uint16, bridgeContract [32]byte) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "registerChain", chainId, bridgeContract)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x65bb3ea7.
//
// Solidity: function registerChain(uint16 chainId, bytes32 bridgeContract) returns()
func (_ZKBridage *ZKBridageSession) RegisterChain(chainId uint16, bridgeContract [32]byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.RegisterChain(&_ZKBridage.TransactOpts, chainId, bridgeContract)
}

// RegisterChain is a paid mutator transaction binding the contract method 0x65bb3ea7.
//
// Solidity: function registerChain(uint16 chainId, bytes32 bridgeContract) returns()
func (_ZKBridage *ZKBridageTransactorSession) RegisterChain(chainId uint16, bridgeContract [32]byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.RegisterChain(&_ZKBridage.TransactOpts, chainId, bridgeContract)
}

// Send is a paid mutator transaction binding the contract method 0xb1d995dd.
//
// Solidity: function send(uint16 dstChainId, address dstAddress, bytes payload) payable returns(uint64 sequence)
func (_ZKBridage *ZKBridageTransactor) Send(opts *bind.TransactOpts, dstChainId uint16, dstAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "send", dstChainId, dstAddress, payload)
}

// Send is a paid mutator transaction binding the contract method 0xb1d995dd.
//
// Solidity: function send(uint16 dstChainId, address dstAddress, bytes payload) payable returns(uint64 sequence)
func (_ZKBridage *ZKBridageSession) Send(dstChainId uint16, dstAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.Send(&_ZKBridage.TransactOpts, dstChainId, dstAddress, payload)
}

// Send is a paid mutator transaction binding the contract method 0xb1d995dd.
//
// Solidity: function send(uint16 dstChainId, address dstAddress, bytes payload) payable returns(uint64 sequence)
func (_ZKBridage *ZKBridageTransactorSession) Send(dstChainId uint16, dstAddress common.Address, payload []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.Send(&_ZKBridage.TransactOpts, dstChainId, dstAddress, payload)
}

// SetBlockUpdater is a paid mutator transaction binding the contract method 0x813d31c9.
//
// Solidity: function setBlockUpdater(uint16 chainId, address blockUpdater) returns()
func (_ZKBridage *ZKBridageTransactor) SetBlockUpdater(opts *bind.TransactOpts, chainId uint16, blockUpdater common.Address) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "setBlockUpdater", chainId, blockUpdater)
}

// SetBlockUpdater is a paid mutator transaction binding the contract method 0x813d31c9.
//
// Solidity: function setBlockUpdater(uint16 chainId, address blockUpdater) returns()
func (_ZKBridage *ZKBridageSession) SetBlockUpdater(chainId uint16, blockUpdater common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetBlockUpdater(&_ZKBridage.TransactOpts, chainId, blockUpdater)
}

// SetBlockUpdater is a paid mutator transaction binding the contract method 0x813d31c9.
//
// Solidity: function setBlockUpdater(uint16 chainId, address blockUpdater) returns()
func (_ZKBridage *ZKBridageTransactorSession) SetBlockUpdater(chainId uint16, blockUpdater common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetBlockUpdater(&_ZKBridage.TransactOpts, chainId, blockUpdater)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_ZKBridage *ZKBridageTransactor) SetLockTime(opts *bind.TransactOpts, lockTime *big.Int) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "setLockTime", lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_ZKBridage *ZKBridageSession) SetLockTime(lockTime *big.Int) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetLockTime(&_ZKBridage.TransactOpts, lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_ZKBridage *ZKBridageTransactorSession) SetLockTime(lockTime *big.Int) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetLockTime(&_ZKBridage.TransactOpts, lockTime)
}

// SetMptVerifier is a paid mutator transaction binding the contract method 0x7cf5744f.
//
// Solidity: function setMptVerifier(uint16 chainId, address mptVerifier) returns()
func (_ZKBridage *ZKBridageTransactor) SetMptVerifier(opts *bind.TransactOpts, chainId uint16, mptVerifier common.Address) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "setMptVerifier", chainId, mptVerifier)
}

// SetMptVerifier is a paid mutator transaction binding the contract method 0x7cf5744f.
//
// Solidity: function setMptVerifier(uint16 chainId, address mptVerifier) returns()
func (_ZKBridage *ZKBridageSession) SetMptVerifier(chainId uint16, mptVerifier common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetMptVerifier(&_ZKBridage.TransactOpts, chainId, mptVerifier)
}

// SetMptVerifier is a paid mutator transaction binding the contract method 0x7cf5744f.
//
// Solidity: function setMptVerifier(uint16 chainId, address mptVerifier) returns()
func (_ZKBridage *ZKBridageTransactorSession) SetMptVerifier(chainId uint16, mptVerifier common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SetMptVerifier(&_ZKBridage.TransactOpts, chainId, mptVerifier)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_ZKBridage *ZKBridageTransactor) SubmitContractUpgrade(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "submitContractUpgrade", newImplementation)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_ZKBridage *ZKBridageSession) SubmitContractUpgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SubmitContractUpgrade(&_ZKBridage.TransactOpts, newImplementation)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_ZKBridage *ZKBridageTransactorSession) SubmitContractUpgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _ZKBridage.Contract.SubmitContractUpgrade(&_ZKBridage.TransactOpts, newImplementation)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x4f64ca19.
//
// Solidity: function validateTransactionProof(uint16 srcChainId, bytes32 srcBlockHash, uint256 logIndex, bytes mptProof) returns()
func (_ZKBridage *ZKBridageTransactor) ValidateTransactionProof(opts *bind.TransactOpts, srcChainId uint16, srcBlockHash [32]byte, logIndex *big.Int, mptProof []byte) (*types.Transaction, error) {
	return _ZKBridage.contract.Transact(opts, "validateTransactionProof", srcChainId, srcBlockHash, logIndex, mptProof)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x4f64ca19.
//
// Solidity: function validateTransactionProof(uint16 srcChainId, bytes32 srcBlockHash, uint256 logIndex, bytes mptProof) returns()
func (_ZKBridage *ZKBridageSession) ValidateTransactionProof(srcChainId uint16, srcBlockHash [32]byte, logIndex *big.Int, mptProof []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.ValidateTransactionProof(&_ZKBridage.TransactOpts, srcChainId, srcBlockHash, logIndex, mptProof)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x4f64ca19.
//
// Solidity: function validateTransactionProof(uint16 srcChainId, bytes32 srcBlockHash, uint256 logIndex, bytes mptProof) returns()
func (_ZKBridage *ZKBridageTransactorSession) ValidateTransactionProof(srcChainId uint16, srcBlockHash [32]byte, logIndex *big.Int, mptProof []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.ValidateTransactionProof(&_ZKBridage.TransactOpts, srcChainId, srcBlockHash, logIndex, mptProof)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZKBridage *ZKBridageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ZKBridage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZKBridage *ZKBridageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.Fallback(&_ZKBridage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ZKBridage *ZKBridageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ZKBridage.Contract.Fallback(&_ZKBridage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZKBridage *ZKBridageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKBridage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZKBridage *ZKBridageSession) Receive() (*types.Transaction, error) {
	return _ZKBridage.Contract.Receive(&_ZKBridage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZKBridage *ZKBridageTransactorSession) Receive() (*types.Transaction, error) {
	return _ZKBridage.Contract.Receive(&_ZKBridage.TransactOpts)
}

// ZKBridageAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ZKBridage contract.
type ZKBridageAdminChangedIterator struct {
	Event *ZKBridageAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZKBridageAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageAdminChanged)
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
		it.Event = new(ZKBridageAdminChanged)
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
func (it *ZKBridageAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageAdminChanged represents a AdminChanged event raised by the ZKBridage contract.
type ZKBridageAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ZKBridage *ZKBridageFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ZKBridageAdminChangedIterator, error) {

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ZKBridageAdminChangedIterator{contract: _ZKBridage.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ZKBridage *ZKBridageFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKBridageAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageAdminChanged)
				if err := _ZKBridage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ZKBridage *ZKBridageFilterer) ParseAdminChanged(log types.Log) (*ZKBridageAdminChanged, error) {
	event := new(ZKBridageAdminChanged)
	if err := _ZKBridage.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ZKBridage contract.
type ZKBridageBeaconUpgradedIterator struct {
	Event *ZKBridageBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ZKBridageBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageBeaconUpgraded)
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
		it.Event = new(ZKBridageBeaconUpgraded)
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
func (it *ZKBridageBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageBeaconUpgraded represents a BeaconUpgraded event raised by the ZKBridage contract.
type ZKBridageBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ZKBridage *ZKBridageFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ZKBridageBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageBeaconUpgradedIterator{contract: _ZKBridage.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ZKBridage *ZKBridageFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ZKBridageBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageBeaconUpgraded)
				if err := _ZKBridage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ZKBridage *ZKBridageFilterer) ParseBeaconUpgraded(log types.Log) (*ZKBridageBeaconUpgraded, error) {
	event := new(ZKBridageBeaconUpgraded)
	if err := _ZKBridage.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageContractUpgradedIterator is returned from FilterContractUpgraded and is used to iterate over the raw logs and unpacked data for ContractUpgraded events raised by the ZKBridage contract.
type ZKBridageContractUpgradedIterator struct {
	Event *ZKBridageContractUpgraded // Event containing the contract specifics and raw log

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
func (it *ZKBridageContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageContractUpgraded)
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
		it.Event = new(ZKBridageContractUpgraded)
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
func (it *ZKBridageContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageContractUpgraded represents a ContractUpgraded event raised by the ZKBridage contract.
type ZKBridageContractUpgraded struct {
	OldContract common.Address
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgraded is a free log retrieval operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_ZKBridage *ZKBridageFilterer) FilterContractUpgraded(opts *bind.FilterOpts, oldContract []common.Address, newContract []common.Address) (*ZKBridageContractUpgradedIterator, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageContractUpgradedIterator{contract: _ZKBridage.contract, event: "ContractUpgraded", logs: logs, sub: sub}, nil
}

// WatchContractUpgraded is a free log subscription operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_ZKBridage *ZKBridageFilterer) WatchContractUpgraded(opts *bind.WatchOpts, sink chan<- *ZKBridageContractUpgraded, oldContract []common.Address, newContract []common.Address) (event.Subscription, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageContractUpgraded)
				if err := _ZKBridage.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
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

// ParseContractUpgraded is a log parse operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_ZKBridage *ZKBridageFilterer) ParseContractUpgraded(log types.Log) (*ZKBridageContractUpgraded, error) {
	event := new(ZKBridageContractUpgraded)
	if err := _ZKBridage.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageExecutedMessageIterator is returned from FilterExecutedMessage and is used to iterate over the raw logs and unpacked data for ExecutedMessage events raised by the ZKBridage contract.
type ZKBridageExecutedMessageIterator struct {
	Event *ZKBridageExecutedMessage // Event containing the contract specifics and raw log

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
func (it *ZKBridageExecutedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageExecutedMessage)
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
		it.Event = new(ZKBridageExecutedMessage)
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
func (it *ZKBridageExecutedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageExecutedMessage represents a ExecutedMessage event raised by the ZKBridage contract.
type ZKBridageExecutedMessage struct {
	Sender     common.Address
	SrcChainId uint16
	Sequence   uint64
	DstAddress common.Address
	Payload    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecutedMessage is a free log retrieval operation binding the contract event 0x4a008ac830958ba6fe8a6e667e2ab53a530eb6cdf93e55b27fc42d7a54cf25b7.
//
// Solidity: event ExecutedMessage(address indexed sender, uint16 indexed srcChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) FilterExecutedMessage(opts *bind.FilterOpts, sender []common.Address, srcChainId []uint16, sequence []uint64) (*ZKBridageExecutedMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "ExecutedMessage", senderRule, srcChainIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageExecutedMessageIterator{contract: _ZKBridage.contract, event: "ExecutedMessage", logs: logs, sub: sub}, nil
}

// WatchExecutedMessage is a free log subscription operation binding the contract event 0x4a008ac830958ba6fe8a6e667e2ab53a530eb6cdf93e55b27fc42d7a54cf25b7.
//
// Solidity: event ExecutedMessage(address indexed sender, uint16 indexed srcChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) WatchExecutedMessage(opts *bind.WatchOpts, sink chan<- *ZKBridageExecutedMessage, sender []common.Address, srcChainId []uint16, sequence []uint64) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "ExecutedMessage", senderRule, srcChainIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageExecutedMessage)
				if err := _ZKBridage.contract.UnpackLog(event, "ExecutedMessage", log); err != nil {
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

// ParseExecutedMessage is a log parse operation binding the contract event 0x4a008ac830958ba6fe8a6e667e2ab53a530eb6cdf93e55b27fc42d7a54cf25b7.
//
// Solidity: event ExecutedMessage(address indexed sender, uint16 indexed srcChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) ParseExecutedMessage(log types.Log) (*ZKBridageExecutedMessage, error) {
	event := new(ZKBridageExecutedMessage)
	if err := _ZKBridage.contract.UnpackLog(event, "ExecutedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageMessagePublishedIterator is returned from FilterMessagePublished and is used to iterate over the raw logs and unpacked data for MessagePublished events raised by the ZKBridage contract.
type ZKBridageMessagePublishedIterator struct {
	Event *ZKBridageMessagePublished // Event containing the contract specifics and raw log

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
func (it *ZKBridageMessagePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageMessagePublished)
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
		it.Event = new(ZKBridageMessagePublished)
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
func (it *ZKBridageMessagePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageMessagePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageMessagePublished represents a MessagePublished event raised by the ZKBridage contract.
type ZKBridageMessagePublished struct {
	Sender     common.Address
	DstChainId uint16
	Sequence   uint64
	DstAddress common.Address
	Payload    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessagePublished is a free log retrieval operation binding the contract event 0xb8abfd5c33667c7440a4fc1153ae39a24833dbe44f7eb19cbe5cd5f2583e4940.
//
// Solidity: event MessagePublished(address indexed sender, uint16 indexed dstChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) FilterMessagePublished(opts *bind.FilterOpts, sender []common.Address, dstChainId []uint16, sequence []uint64) (*ZKBridageMessagePublishedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "MessagePublished", senderRule, dstChainIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageMessagePublishedIterator{contract: _ZKBridage.contract, event: "MessagePublished", logs: logs, sub: sub}, nil
}

// WatchMessagePublished is a free log subscription operation binding the contract event 0xb8abfd5c33667c7440a4fc1153ae39a24833dbe44f7eb19cbe5cd5f2583e4940.
//
// Solidity: event MessagePublished(address indexed sender, uint16 indexed dstChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) WatchMessagePublished(opts *bind.WatchOpts, sink chan<- *ZKBridageMessagePublished, sender []common.Address, dstChainId []uint16, sequence []uint64) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "MessagePublished", senderRule, dstChainIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageMessagePublished)
				if err := _ZKBridage.contract.UnpackLog(event, "MessagePublished", log); err != nil {
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

// ParseMessagePublished is a log parse operation binding the contract event 0xb8abfd5c33667c7440a4fc1153ae39a24833dbe44f7eb19cbe5cd5f2583e4940.
//
// Solidity: event MessagePublished(address indexed sender, uint16 indexed dstChainId, uint64 indexed sequence, address dstAddress, bytes payload)
func (_ZKBridage *ZKBridageFilterer) ParseMessagePublished(log types.Log) (*ZKBridageMessagePublished, error) {
	event := new(ZKBridageMessagePublished)
	if err := _ZKBridage.contract.UnpackLog(event, "MessagePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageNewPendingImplementationIterator is returned from FilterNewPendingImplementation and is used to iterate over the raw logs and unpacked data for NewPendingImplementation events raised by the ZKBridage contract.
type ZKBridageNewPendingImplementationIterator struct {
	Event *ZKBridageNewPendingImplementation // Event containing the contract specifics and raw log

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
func (it *ZKBridageNewPendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageNewPendingImplementation)
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
		it.Event = new(ZKBridageNewPendingImplementation)
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
func (it *ZKBridageNewPendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageNewPendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageNewPendingImplementation represents a NewPendingImplementation event raised by the ZKBridage contract.
type ZKBridageNewPendingImplementation struct {
	PendingImplementation common.Address
	NewImplementation     common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewPendingImplementation is a free log retrieval operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address indexed pendingImplementation, address indexed newImplementation)
func (_ZKBridage *ZKBridageFilterer) FilterNewPendingImplementation(opts *bind.FilterOpts, pendingImplementation []common.Address, newImplementation []common.Address) (*ZKBridageNewPendingImplementationIterator, error) {

	var pendingImplementationRule []interface{}
	for _, pendingImplementationItem := range pendingImplementation {
		pendingImplementationRule = append(pendingImplementationRule, pendingImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "NewPendingImplementation", pendingImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageNewPendingImplementationIterator{contract: _ZKBridage.contract, event: "NewPendingImplementation", logs: logs, sub: sub}, nil
}

// WatchNewPendingImplementation is a free log subscription operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address indexed pendingImplementation, address indexed newImplementation)
func (_ZKBridage *ZKBridageFilterer) WatchNewPendingImplementation(opts *bind.WatchOpts, sink chan<- *ZKBridageNewPendingImplementation, pendingImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var pendingImplementationRule []interface{}
	for _, pendingImplementationItem := range pendingImplementation {
		pendingImplementationRule = append(pendingImplementationRule, pendingImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "NewPendingImplementation", pendingImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageNewPendingImplementation)
				if err := _ZKBridage.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
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

// ParseNewPendingImplementation is a log parse operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address indexed pendingImplementation, address indexed newImplementation)
func (_ZKBridage *ZKBridageFilterer) ParseNewPendingImplementation(log types.Log) (*ZKBridageNewPendingImplementation, error) {
	event := new(ZKBridageNewPendingImplementation)
	if err := _ZKBridage.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKBridageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZKBridage contract.
type ZKBridageUpgradedIterator struct {
	Event *ZKBridageUpgraded // Event containing the contract specifics and raw log

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
func (it *ZKBridageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKBridageUpgraded)
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
		it.Event = new(ZKBridageUpgraded)
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
func (it *ZKBridageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKBridageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKBridageUpgraded represents a Upgraded event raised by the ZKBridage contract.
type ZKBridageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZKBridage *ZKBridageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZKBridageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZKBridage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZKBridageUpgradedIterator{contract: _ZKBridage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZKBridage *ZKBridageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZKBridageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZKBridage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKBridageUpgraded)
				if err := _ZKBridage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZKBridage *ZKBridageFilterer) ParseUpgraded(log types.Log) (*ZKBridageUpgraded, error) {
	event := new(ZKBridageUpgraded)
	if err := _ZKBridage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
