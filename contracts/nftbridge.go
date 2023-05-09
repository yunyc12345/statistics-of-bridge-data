package contracts

import "errors"

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NftBridgeMetaData contains all meta data concerning the NftBridge contract.
var NftBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewPendingImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"sourceChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"sendChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ReceiveNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftBridge\",\"type\":\"address\"}],\"name\":\"RegisterChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"TransferNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_LOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId_\",\"type\":\"uint16\"}],\"name\":\"bridgeContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWrappedAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"}],\"name\":\"setLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"submitContractUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"}],\"name\":\"transferNFT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16\",\"name\":\"recipientChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"}],\"name\":\"transferNFTBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tokenChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"tokenAddress\",\"type\":\"bytes32\"}],\"name\":\"wrappedAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkBridge\",\"outputs\":[{\"internalType\":\"contractIZKBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"zkReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NftBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use NftBridgeMetaData.ABI instead.
var NftBridgeABI = NftBridgeMetaData.ABI

// NftBridge is an auto generated Go binding around an Ethereum contract.
type NftBridge struct {
	NftBridgeCaller     // Read-only binding to the contract
	NftBridgeTransactor // Write-only binding to the contract
	NftBridgeFilterer   // Log filterer for contract events
}

// NftBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftBridgeSession struct {
	Contract     *NftBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftBridgeCallerSession struct {
	Contract *NftBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NftBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftBridgeTransactorSession struct {
	Contract     *NftBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NftBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftBridgeRaw struct {
	Contract *NftBridge // Generic contract binding to access the raw methods on
}

// NftBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftBridgeCallerRaw struct {
	Contract *NftBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NftBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftBridgeTransactorRaw struct {
	Contract *NftBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNftBridge creates a new instance of NftBridge, bound to a specific deployed contract.
func NewNftBridge(address common.Address, backend bind.ContractBackend) (*NftBridge, error) {
	contract, err := bindNftBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NftBridge{NftBridgeCaller: NftBridgeCaller{contract: contract}, NftBridgeTransactor: NftBridgeTransactor{contract: contract}, NftBridgeFilterer: NftBridgeFilterer{contract: contract}}, nil
}

// NewNftBridgeCaller creates a new read-only instance of NftBridge, bound to a specific deployed contract.
func NewNftBridgeCaller(address common.Address, caller bind.ContractCaller) (*NftBridgeCaller, error) {
	contract, err := bindNftBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftBridgeCaller{contract: contract}, nil
}

// NewNftBridgeTransactor creates a new write-only instance of NftBridge, bound to a specific deployed contract.
func NewNftBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NftBridgeTransactor, error) {
	contract, err := bindNftBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftBridgeTransactor{contract: contract}, nil
}

// NewNftBridgeFilterer creates a new log filterer instance of NftBridge, bound to a specific deployed contract.
func NewNftBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NftBridgeFilterer, error) {
	contract, err := bindNftBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftBridgeFilterer{contract: contract}, nil
}

// bindNftBridge binds a generic wrapper to an already deployed contract.
func bindNftBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftBridge *NftBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftBridge.Contract.NftBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftBridge *NftBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftBridge.Contract.NftBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftBridge *NftBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftBridge.Contract.NftBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NftBridge *NftBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NftBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NftBridge *NftBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NftBridge *NftBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NftBridge.Contract.contract.Transact(opts, method, params...)
}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_NftBridge *NftBridgeCaller) MINLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "MIN_LOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_NftBridge *NftBridgeSession) MINLOCKTIME() (*big.Int, error) {
	return _NftBridge.Contract.MINLOCKTIME(&_NftBridge.CallOpts)
}

// MINLOCKTIME is a free data retrieval call binding the contract method 0x3ff03207.
//
// Solidity: function MIN_LOCK_TIME() view returns(uint256)
func (_NftBridge *NftBridgeCallerSession) MINLOCKTIME() (*big.Int, error) {
	return _NftBridge.Contract.MINLOCKTIME(&_NftBridge.CallOpts)
}

// BridgeContracts is a free data retrieval call binding the contract method 0xad66a5f1.
//
// Solidity: function bridgeContracts(uint16 chainId_) view returns(address)
func (_NftBridge *NftBridgeCaller) BridgeContracts(opts *bind.CallOpts, chainId_ uint16) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "bridgeContracts", chainId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeContracts is a free data retrieval call binding the contract method 0xad66a5f1.
//
// Solidity: function bridgeContracts(uint16 chainId_) view returns(address)
func (_NftBridge *NftBridgeSession) BridgeContracts(chainId_ uint16) (common.Address, error) {
	return _NftBridge.Contract.BridgeContracts(&_NftBridge.CallOpts, chainId_)
}

// BridgeContracts is a free data retrieval call binding the contract method 0xad66a5f1.
//
// Solidity: function bridgeContracts(uint16 chainId_) view returns(address)
func (_NftBridge *NftBridgeCallerSession) BridgeContracts(chainId_ uint16) (common.Address, error) {
	return _NftBridge.Contract.BridgeContracts(&_NftBridge.CallOpts, chainId_)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_NftBridge *NftBridgeCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_NftBridge *NftBridgeSession) ChainId() (uint16, error) {
	return _NftBridge.Contract.ChainId(&_NftBridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_NftBridge *NftBridgeCallerSession) ChainId() (uint16, error) {
	return _NftBridge.Contract.ChainId(&_NftBridge.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_NftBridge *NftBridgeCaller) IsInitialized(opts *bind.CallOpts, impl common.Address) (bool, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "isInitialized", impl)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_NftBridge *NftBridgeSession) IsInitialized(impl common.Address) (bool, error) {
	return _NftBridge.Contract.IsInitialized(&_NftBridge.CallOpts, impl)
}

// IsInitialized is a free data retrieval call binding the contract method 0xd60b347f.
//
// Solidity: function isInitialized(address impl) view returns(bool)
func (_NftBridge *NftBridgeCallerSession) IsInitialized(impl common.Address) (bool, error) {
	return _NftBridge.Contract.IsInitialized(&_NftBridge.CallOpts, impl)
}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address token) view returns(bool)
func (_NftBridge *NftBridgeCaller) IsWrappedAsset(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "isWrappedAsset", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address token) view returns(bool)
func (_NftBridge *NftBridgeSession) IsWrappedAsset(token common.Address) (bool, error) {
	return _NftBridge.Contract.IsWrappedAsset(&_NftBridge.CallOpts, token)
}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address token) view returns(bool)
func (_NftBridge *NftBridgeCallerSession) IsWrappedAsset(token common.Address) (bool, error) {
	return _NftBridge.Contract.IsWrappedAsset(&_NftBridge.CallOpts, token)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_NftBridge *NftBridgeCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_NftBridge *NftBridgeSession) LockTime() (*big.Int, error) {
	return _NftBridge.Contract.LockTime(&_NftBridge.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_NftBridge *NftBridgeCallerSession) LockTime() (*big.Int, error) {
	return _NftBridge.Contract.LockTime(&_NftBridge.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) view returns(bytes4)
func (_NftBridge *NftBridgeCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "onERC721Received", operator, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) view returns(bytes4)
func (_NftBridge *NftBridgeSession) OnERC721Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NftBridge.Contract.OnERC721Received(&_NftBridge.CallOpts, operator, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address , uint256 , bytes ) view returns(bytes4)
func (_NftBridge *NftBridgeCallerSession) OnERC721Received(operator common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _NftBridge.Contract.OnERC721Received(&_NftBridge.CallOpts, operator, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftBridge *NftBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftBridge *NftBridgeSession) Owner() (common.Address, error) {
	return _NftBridge.Contract.Owner(&_NftBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NftBridge *NftBridgeCallerSession) Owner() (common.Address, error) {
	return _NftBridge.Contract.Owner(&_NftBridge.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_NftBridge *NftBridgeCaller) PendingImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "pendingImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_NftBridge *NftBridgeSession) PendingImplementation() (common.Address, error) {
	return _NftBridge.Contract.PendingImplementation(&_NftBridge.CallOpts)
}

// PendingImplementation is a free data retrieval call binding the contract method 0x396f7b23.
//
// Solidity: function pendingImplementation() view returns(address)
func (_NftBridge *NftBridgeCallerSession) PendingImplementation() (common.Address, error) {
	return _NftBridge.Contract.PendingImplementation(&_NftBridge.CallOpts)
}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_NftBridge *NftBridgeCaller) ToUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "toUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_NftBridge *NftBridgeSession) ToUpdateTime() (*big.Int, error) {
	return _NftBridge.Contract.ToUpdateTime(&_NftBridge.CallOpts)
}

// ToUpdateTime is a free data retrieval call binding the contract method 0x9f0a22a6.
//
// Solidity: function toUpdateTime() view returns(uint256)
func (_NftBridge *NftBridgeCallerSession) ToUpdateTime() (*big.Int, error) {
	return _NftBridge.Contract.ToUpdateTime(&_NftBridge.CallOpts)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x2f3a3d5d.
//
// Solidity: function tokenImplementation() view returns(address)
func (_NftBridge *NftBridgeCaller) TokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "tokenImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenImplementation is a free data retrieval call binding the contract method 0x2f3a3d5d.
//
// Solidity: function tokenImplementation() view returns(address)
func (_NftBridge *NftBridgeSession) TokenImplementation() (common.Address, error) {
	return _NftBridge.Contract.TokenImplementation(&_NftBridge.CallOpts)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x2f3a3d5d.
//
// Solidity: function tokenImplementation() view returns(address)
func (_NftBridge *NftBridgeCallerSession) TokenImplementation() (common.Address, error) {
	return _NftBridge.Contract.TokenImplementation(&_NftBridge.CallOpts)
}

// WrappedAsset is a free data retrieval call binding the contract method 0x1ff1e286.
//
// Solidity: function wrappedAsset(uint16 tokenChainId, bytes32 tokenAddress) view returns(address)
func (_NftBridge *NftBridgeCaller) WrappedAsset(opts *bind.CallOpts, tokenChainId uint16, tokenAddress [32]byte) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "wrappedAsset", tokenChainId, tokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedAsset is a free data retrieval call binding the contract method 0x1ff1e286.
//
// Solidity: function wrappedAsset(uint16 tokenChainId, bytes32 tokenAddress) view returns(address)
func (_NftBridge *NftBridgeSession) WrappedAsset(tokenChainId uint16, tokenAddress [32]byte) (common.Address, error) {
	return _NftBridge.Contract.WrappedAsset(&_NftBridge.CallOpts, tokenChainId, tokenAddress)
}

// WrappedAsset is a free data retrieval call binding the contract method 0x1ff1e286.
//
// Solidity: function wrappedAsset(uint16 tokenChainId, bytes32 tokenAddress) view returns(address)
func (_NftBridge *NftBridgeCallerSession) WrappedAsset(tokenChainId uint16, tokenAddress [32]byte) (common.Address, error) {
	return _NftBridge.Contract.WrappedAsset(&_NftBridge.CallOpts, tokenChainId, tokenAddress)
}

// ZkBridge is a free data retrieval call binding the contract method 0xb4244659.
//
// Solidity: function zkBridge() view returns(address)
func (_NftBridge *NftBridgeCaller) ZkBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NftBridge.contract.Call(opts, &out, "zkBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkBridge is a free data retrieval call binding the contract method 0xb4244659.
//
// Solidity: function zkBridge() view returns(address)
func (_NftBridge *NftBridgeSession) ZkBridge() (common.Address, error) {
	return _NftBridge.Contract.ZkBridge(&_NftBridge.CallOpts)
}

// ZkBridge is a free data retrieval call binding the contract method 0xb4244659.
//
// Solidity: function zkBridge() view returns(address)
func (_NftBridge *NftBridgeCallerSession) ZkBridge() (common.Address, error) {
	return _NftBridge.Contract.ZkBridge(&_NftBridge.CallOpts)
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_NftBridge *NftBridgeTransactor) ConfirmContractUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "confirmContractUpgrade")
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_NftBridge *NftBridgeSession) ConfirmContractUpgrade() (*types.Transaction, error) {
	return _NftBridge.Contract.ConfirmContractUpgrade(&_NftBridge.TransactOpts)
}

// ConfirmContractUpgrade is a paid mutator transaction binding the contract method 0xd1dc83c2.
//
// Solidity: function confirmContractUpgrade() returns()
func (_NftBridge *NftBridgeTransactorSession) ConfirmContractUpgrade() (*types.Transaction, error) {
	return _NftBridge.Contract.ConfirmContractUpgrade(&_NftBridge.TransactOpts)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb6dc1236.
//
// Solidity: function registerChain(uint16 chainId, address contractAddress) returns()
func (_NftBridge *NftBridgeTransactor) RegisterChain(opts *bind.TransactOpts, chainId uint16, contractAddress common.Address) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "registerChain", chainId, contractAddress)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb6dc1236.
//
// Solidity: function registerChain(uint16 chainId, address contractAddress) returns()
func (_NftBridge *NftBridgeSession) RegisterChain(chainId uint16, contractAddress common.Address) (*types.Transaction, error) {
	return _NftBridge.Contract.RegisterChain(&_NftBridge.TransactOpts, chainId, contractAddress)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xb6dc1236.
//
// Solidity: function registerChain(uint16 chainId, address contractAddress) returns()
func (_NftBridge *NftBridgeTransactorSession) RegisterChain(chainId uint16, contractAddress common.Address) (*types.Transaction, error) {
	return _NftBridge.Contract.RegisterChain(&_NftBridge.TransactOpts, chainId, contractAddress)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_NftBridge *NftBridgeTransactor) SetLockTime(opts *bind.TransactOpts, lockTime *big.Int) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "setLockTime", lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_NftBridge *NftBridgeSession) SetLockTime(lockTime *big.Int) (*types.Transaction, error) {
	return _NftBridge.Contract.SetLockTime(&_NftBridge.TransactOpts, lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 lockTime) returns()
func (_NftBridge *NftBridgeTransactorSession) SetLockTime(lockTime *big.Int) (*types.Transaction, error) {
	return _NftBridge.Contract.SetLockTime(&_NftBridge.TransactOpts, lockTime)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_NftBridge *NftBridgeTransactor) SubmitContractUpgrade(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "submitContractUpgrade", newImplementation)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_NftBridge *NftBridgeSession) SubmitContractUpgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _NftBridge.Contract.SubmitContractUpgrade(&_NftBridge.TransactOpts, newImplementation)
}

// SubmitContractUpgrade is a paid mutator transaction binding the contract method 0x3a552757.
//
// Solidity: function submitContractUpgrade(address newImplementation) returns()
func (_NftBridge *NftBridgeTransactorSession) SubmitContractUpgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _NftBridge.Contract.SubmitContractUpgrade(&_NftBridge.TransactOpts, newImplementation)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xac7b22dc.
//
// Solidity: function transferNFT(address token, uint256 tokenID, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeTransactor) TransferNFT(opts *bind.TransactOpts, token common.Address, tokenID *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "transferNFT", token, tokenID, recipientChain, recipient)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xac7b22dc.
//
// Solidity: function transferNFT(address token, uint256 tokenID, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeSession) TransferNFT(token common.Address, tokenID *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.Contract.TransferNFT(&_NftBridge.TransactOpts, token, tokenID, recipientChain, recipient)
}

// TransferNFT is a paid mutator transaction binding the contract method 0xac7b22dc.
//
// Solidity: function transferNFT(address token, uint256 tokenID, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeTransactorSession) TransferNFT(token common.Address, tokenID *big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.Contract.TransferNFT(&_NftBridge.TransactOpts, token, tokenID, recipientChain, recipient)
}

// TransferNFTBatch is a paid mutator transaction binding the contract method 0xaba5ca04.
//
// Solidity: function transferNFTBatch(address[] tokenList, uint256[] tokenIDList, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeTransactor) TransferNFTBatch(opts *bind.TransactOpts, tokenList []common.Address, tokenIDList []*big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "transferNFTBatch", tokenList, tokenIDList, recipientChain, recipient)
}

// TransferNFTBatch is a paid mutator transaction binding the contract method 0xaba5ca04.
//
// Solidity: function transferNFTBatch(address[] tokenList, uint256[] tokenIDList, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeSession) TransferNFTBatch(tokenList []common.Address, tokenIDList []*big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.Contract.TransferNFTBatch(&_NftBridge.TransactOpts, tokenList, tokenIDList, recipientChain, recipient)
}

// TransferNFTBatch is a paid mutator transaction binding the contract method 0xaba5ca04.
//
// Solidity: function transferNFTBatch(address[] tokenList, uint256[] tokenIDList, uint16 recipientChain, bytes32 recipient) payable returns(uint64 sequence)
func (_NftBridge *NftBridgeTransactorSession) TransferNFTBatch(tokenList []common.Address, tokenIDList []*big.Int, recipientChain uint16, recipient [32]byte) (*types.Transaction, error) {
	return _NftBridge.Contract.TransferNFTBatch(&_NftBridge.TransactOpts, tokenList, tokenIDList, recipientChain, recipient)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_NftBridge *NftBridgeTransactor) ZkReceive(opts *bind.TransactOpts, srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _NftBridge.contract.Transact(opts, "zkReceive", srcChainId, srcAddress, sequence, payload)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_NftBridge *NftBridgeSession) ZkReceive(srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _NftBridge.Contract.ZkReceive(&_NftBridge.TransactOpts, srcChainId, srcAddress, sequence, payload)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_NftBridge *NftBridgeTransactorSession) ZkReceive(srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _NftBridge.Contract.ZkReceive(&_NftBridge.TransactOpts, srcChainId, srcAddress, sequence, payload)
}

// NftBridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NftBridge contract.
type NftBridgeAdminChangedIterator struct {
	Event *NftBridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *NftBridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeAdminChanged)
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
		it.Event = new(NftBridgeAdminChanged)
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
func (it *NftBridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeAdminChanged represents a AdminChanged event raised by the NftBridge contract.
type NftBridgeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NftBridge *NftBridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NftBridgeAdminChangedIterator, error) {

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NftBridgeAdminChangedIterator{contract: _NftBridge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NftBridge *NftBridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NftBridgeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeAdminChanged)
				if err := _NftBridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NftBridge *NftBridgeFilterer) ParseAdminChanged(log types.Log) (*NftBridgeAdminChanged, error) {
	event := new(NftBridgeAdminChanged)
	if err := _NftBridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NftBridge contract.
type NftBridgeBeaconUpgradedIterator struct {
	Event *NftBridgeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NftBridgeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeBeaconUpgraded)
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
		it.Event = new(NftBridgeBeaconUpgraded)
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
func (it *NftBridgeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeBeaconUpgraded represents a BeaconUpgraded event raised by the NftBridge contract.
type NftBridgeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NftBridge *NftBridgeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NftBridgeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeBeaconUpgradedIterator{contract: _NftBridge.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NftBridge *NftBridgeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NftBridgeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeBeaconUpgraded)
				if err := _NftBridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NftBridge *NftBridgeFilterer) ParseBeaconUpgraded(log types.Log) (*NftBridgeBeaconUpgraded, error) {
	event := new(NftBridgeBeaconUpgraded)
	if err := _NftBridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeContractUpgradedIterator is returned from FilterContractUpgraded and is used to iterate over the raw logs and unpacked data for ContractUpgraded events raised by the NftBridge contract.
type NftBridgeContractUpgradedIterator struct {
	Event *NftBridgeContractUpgraded // Event containing the contract specifics and raw log

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
func (it *NftBridgeContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeContractUpgraded)
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
		it.Event = new(NftBridgeContractUpgraded)
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
func (it *NftBridgeContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeContractUpgraded represents a ContractUpgraded event raised by the NftBridge contract.
type NftBridgeContractUpgraded struct {
	OldContract common.Address
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgraded is a free log retrieval operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_NftBridge *NftBridgeFilterer) FilterContractUpgraded(opts *bind.FilterOpts, oldContract []common.Address, newContract []common.Address) (*NftBridgeContractUpgradedIterator, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeContractUpgradedIterator{contract: _NftBridge.contract, event: "ContractUpgraded", logs: logs, sub: sub}, nil
}

// WatchContractUpgraded is a free log subscription operation binding the contract event 0x2e4cc16c100f0b55e2df82ab0b1a7e294aa9cbd01b48fbaf622683fbc0507a49.
//
// Solidity: event ContractUpgraded(address indexed oldContract, address indexed newContract)
func (_NftBridge *NftBridgeFilterer) WatchContractUpgraded(opts *bind.WatchOpts, sink chan<- *NftBridgeContractUpgraded, oldContract []common.Address, newContract []common.Address) (event.Subscription, error) {

	var oldContractRule []interface{}
	for _, oldContractItem := range oldContract {
		oldContractRule = append(oldContractRule, oldContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "ContractUpgraded", oldContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeContractUpgraded)
				if err := _NftBridge.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
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
func (_NftBridge *NftBridgeFilterer) ParseContractUpgraded(log types.Log) (*NftBridgeContractUpgraded, error) {
	event := new(NftBridgeContractUpgraded)
	if err := _NftBridge.contract.UnpackLog(event, "ContractUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeNewPendingImplementationIterator is returned from FilterNewPendingImplementation and is used to iterate over the raw logs and unpacked data for NewPendingImplementation events raised by the NftBridge contract.
type NftBridgeNewPendingImplementationIterator struct {
	Event *NftBridgeNewPendingImplementation // Event containing the contract specifics and raw log

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
func (it *NftBridgeNewPendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeNewPendingImplementation)
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
		it.Event = new(NftBridgeNewPendingImplementation)
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
func (it *NftBridgeNewPendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeNewPendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeNewPendingImplementation represents a NewPendingImplementation event raised by the NftBridge contract.
type NftBridgeNewPendingImplementation struct {
	PendingImplementation common.Address
	NewImplementation     common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewPendingImplementation is a free log retrieval operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address indexed pendingImplementation, address indexed newImplementation)
func (_NftBridge *NftBridgeFilterer) FilterNewPendingImplementation(opts *bind.FilterOpts, pendingImplementation []common.Address, newImplementation []common.Address) (*NftBridgeNewPendingImplementationIterator, error) {

	var pendingImplementationRule []interface{}
	for _, pendingImplementationItem := range pendingImplementation {
		pendingImplementationRule = append(pendingImplementationRule, pendingImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "NewPendingImplementation", pendingImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeNewPendingImplementationIterator{contract: _NftBridge.contract, event: "NewPendingImplementation", logs: logs, sub: sub}, nil
}

// WatchNewPendingImplementation is a free log subscription operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address indexed pendingImplementation, address indexed newImplementation)
func (_NftBridge *NftBridgeFilterer) WatchNewPendingImplementation(opts *bind.WatchOpts, sink chan<- *NftBridgeNewPendingImplementation, pendingImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var pendingImplementationRule []interface{}
	for _, pendingImplementationItem := range pendingImplementation {
		pendingImplementationRule = append(pendingImplementationRule, pendingImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "NewPendingImplementation", pendingImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeNewPendingImplementation)
				if err := _NftBridge.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
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
func (_NftBridge *NftBridgeFilterer) ParseNewPendingImplementation(log types.Log) (*NftBridgeNewPendingImplementation, error) {
	event := new(NftBridgeNewPendingImplementation)
	if err := _NftBridge.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeReceiveNFTIterator is returned from FilterReceiveNFT and is used to iterate over the raw logs and unpacked data for ReceiveNFT events raised by the NftBridge contract.
type NftBridgeReceiveNFTIterator struct {
	Event *NftBridgeReceiveNFT // Event containing the contract specifics and raw log

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
func (it *NftBridgeReceiveNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeReceiveNFT)
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
		it.Event = new(NftBridgeReceiveNFT)
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
func (it *NftBridgeReceiveNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeReceiveNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeReceiveNFT represents a ReceiveNFT event raised by the NftBridge contract.
type NftBridgeReceiveNFT struct {
	Sequence    uint64
	SourceToken common.Address
	Token       common.Address
	TokenID     *big.Int
	SourceChain uint16
	SendChain   uint16
	Recipient   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceiveNFT is a free log retrieval operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed sequence, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NftBridge *NftBridgeFilterer) FilterReceiveNFT(opts *bind.FilterOpts, sequence []uint64) (*NftBridgeReceiveNFTIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "ReceiveNFT", sequenceRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeReceiveNFTIterator{contract: _NftBridge.contract, event: "ReceiveNFT", logs: logs, sub: sub}, nil
}

// WatchReceiveNFT is a free log subscription operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed sequence, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NftBridge *NftBridgeFilterer) WatchReceiveNFT(opts *bind.WatchOpts, sink chan<- *NftBridgeReceiveNFT, sequence []uint64) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "ReceiveNFT", sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeReceiveNFT)
				if err := _NftBridge.contract.UnpackLog(event, "ReceiveNFT", log); err != nil {
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

// ParseReceiveNFT is a log parse operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed sequence, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NftBridge *NftBridgeFilterer) ParseReceiveNFT(log types.Log) (*NftBridgeReceiveNFT, error) {
	event := new(NftBridgeReceiveNFT)
	if err := _NftBridge.contract.UnpackLog(event, "ReceiveNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeRegisterChainIterator is returned from FilterRegisterChain and is used to iterate over the raw logs and unpacked data for RegisterChain events raised by the NftBridge contract.
type NftBridgeRegisterChainIterator struct {
	Event *NftBridgeRegisterChain // Event containing the contract specifics and raw log

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
func (it *NftBridgeRegisterChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeRegisterChain)
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
		it.Event = new(NftBridgeRegisterChain)
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
func (it *NftBridgeRegisterChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeRegisterChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeRegisterChain represents a RegisterChain event raised by the NftBridge contract.
type NftBridgeRegisterChain struct {
	ChainId   uint16
	NftBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegisterChain is a free log retrieval operation binding the contract event 0xaf4e2924b78f8cba0a0e626999cf99bb7a50dba42aab94e3030ed34e54e5c12c.
//
// Solidity: event RegisterChain(uint16 chainId, address nftBridge)
func (_NftBridge *NftBridgeFilterer) FilterRegisterChain(opts *bind.FilterOpts) (*NftBridgeRegisterChainIterator, error) {

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "RegisterChain")
	if err != nil {
		return nil, err
	}
	return &NftBridgeRegisterChainIterator{contract: _NftBridge.contract, event: "RegisterChain", logs: logs, sub: sub}, nil
}

// WatchRegisterChain is a free log subscription operation binding the contract event 0xaf4e2924b78f8cba0a0e626999cf99bb7a50dba42aab94e3030ed34e54e5c12c.
//
// Solidity: event RegisterChain(uint16 chainId, address nftBridge)
func (_NftBridge *NftBridgeFilterer) WatchRegisterChain(opts *bind.WatchOpts, sink chan<- *NftBridgeRegisterChain) (event.Subscription, error) {

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "RegisterChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeRegisterChain)
				if err := _NftBridge.contract.UnpackLog(event, "RegisterChain", log); err != nil {
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

// ParseRegisterChain is a log parse operation binding the contract event 0xaf4e2924b78f8cba0a0e626999cf99bb7a50dba42aab94e3030ed34e54e5c12c.
//
// Solidity: event RegisterChain(uint16 chainId, address nftBridge)
func (_NftBridge *NftBridgeFilterer) ParseRegisterChain(log types.Log) (*NftBridgeRegisterChain, error) {
	event := new(NftBridgeRegisterChain)
	if err := _NftBridge.contract.UnpackLog(event, "RegisterChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeTransferNFTIterator is returned from FilterTransferNFT and is used to iterate over the raw logs and unpacked data for TransferNFT events raised by the NftBridge contract.
type NftBridgeTransferNFTIterator struct {
	Event *NftBridgeTransferNFT // Event containing the contract specifics and raw log

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
func (it *NftBridgeTransferNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeTransferNFT)
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
		it.Event = new(NftBridgeTransferNFT)
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
func (it *NftBridgeTransferNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeTransferNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeTransferNFT represents a TransferNFT event raised by the NftBridge contract.
type NftBridgeTransferNFT struct {
	Sequence       uint64
	Token          common.Address
	TokenID        *big.Int
	RecipientChain uint16
	Sender         common.Address
	Recipient      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferNFT is a free log retrieval operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed sequence, address token, uint256 tokenID, uint16 recipientChain, address sender, address recipient)
func (_NftBridge *NftBridgeFilterer) FilterTransferNFT(opts *bind.FilterOpts, sequence []uint64) (*NftBridgeTransferNFTIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "TransferNFT", sequenceRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeTransferNFTIterator{contract: _NftBridge.contract, event: "TransferNFT", logs: logs, sub: sub}, nil
}

// WatchTransferNFT is a free log subscription operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed sequence, address token, uint256 tokenID, uint16 recipientChain, address sender, address recipient)
func (_NftBridge *NftBridgeFilterer) WatchTransferNFT(opts *bind.WatchOpts, sink chan<- *NftBridgeTransferNFT, sequence []uint64) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "TransferNFT", sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeTransferNFT)
				if err := _NftBridge.contract.UnpackLog(event, "TransferNFT", log); err != nil {
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

// ParseTransferNFT is a log parse operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed sequence, address token, uint256 tokenID, uint16 recipientChain, address sender, address recipient)
func (_NftBridge *NftBridgeFilterer) ParseTransferNFT(log types.Log) (*NftBridgeTransferNFT, error) {
	event := new(NftBridgeTransferNFT)
	if err := _NftBridge.contract.UnpackLog(event, "TransferNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NftBridge contract.
type NftBridgeUpgradedIterator struct {
	Event *NftBridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *NftBridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBridgeUpgraded)
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
		it.Event = new(NftBridgeUpgraded)
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
func (it *NftBridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBridgeUpgraded represents a Upgraded event raised by the NftBridge contract.
type NftBridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NftBridge *NftBridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NftBridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NftBridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NftBridgeUpgradedIterator{contract: _NftBridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NftBridge *NftBridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NftBridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NftBridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBridgeUpgraded)
				if err := _NftBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NftBridge *NftBridgeFilterer) ParseUpgraded(log types.Log) (*NftBridgeUpgraded, error) {
	event := new(NftBridgeUpgraded)
	if err := _NftBridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
