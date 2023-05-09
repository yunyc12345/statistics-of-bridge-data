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

// MailerMetaData contains all meta data concerning the Mailer contract.
var MailerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zkBridgeEntrypointAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"MessageSend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkBridgeEntrypoint\",\"outputs\":[{\"internalType\":\"contractIZKBridgeEntrypoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MailerABI is the input ABI used to generate the binding from.
// Deprecated: Use MailerMetaData.ABI instead.
var MailerABI = MailerMetaData.ABI

// Mailer is an auto generated Go binding around an Ethereum contract.
type Mailer struct {
	MailerCaller     // Read-only binding to the contract
	MailerTransactor // Write-only binding to the contract
	MailerFilterer   // Log filterer for contract events
}

// MailerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailerSession struct {
	Contract     *Mailer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailerCallerSession struct {
	Contract *MailerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MailerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailerTransactorSession struct {
	Contract     *MailerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailerRaw struct {
	Contract *Mailer // Generic contract binding to access the raw methods on
}

// MailerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailerCallerRaw struct {
	Contract *MailerCaller // Generic read-only contract binding to access the raw methods on
}

// MailerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailerTransactorRaw struct {
	Contract *MailerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailer creates a new instance of Mailer, bound to a specific deployed contract.
func NewMailer(address common.Address, backend bind.ContractBackend) (*Mailer, error) {
	contract, err := bindMailer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailer{MailerCaller: MailerCaller{contract: contract}, MailerTransactor: MailerTransactor{contract: contract}, MailerFilterer: MailerFilterer{contract: contract}}, nil
}

// NewMailerCaller creates a new read-only instance of Mailer, bound to a specific deployed contract.
func NewMailerCaller(address common.Address, caller bind.ContractCaller) (*MailerCaller, error) {
	contract, err := bindMailer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailerCaller{contract: contract}, nil
}

// NewMailerTransactor creates a new write-only instance of Mailer, bound to a specific deployed contract.
func NewMailerTransactor(address common.Address, transactor bind.ContractTransactor) (*MailerTransactor, error) {
	contract, err := bindMailer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailerTransactor{contract: contract}, nil
}

// NewMailerFilterer creates a new log filterer instance of Mailer, bound to a specific deployed contract.
func NewMailerFilterer(address common.Address, filterer bind.ContractFilterer) (*MailerFilterer, error) {
	contract, err := bindMailer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailerFilterer{contract: contract}, nil
}

// bindMailer binds a generic wrapper to an already deployed contract.
func bindMailer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailer *MailerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailer.Contract.MailerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailer *MailerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailer.Contract.MailerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailer *MailerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailer.Contract.MailerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailer *MailerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailer *MailerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailer *MailerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailer.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mailer *MailerCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailer.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mailer *MailerSession) Fee() (*big.Int, error) {
	return _Mailer.Contract.Fee(&_Mailer.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Mailer *MailerCallerSession) Fee() (*big.Int, error) {
	return _Mailer.Contract.Fee(&_Mailer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailer *MailerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailer *MailerSession) Owner() (common.Address, error) {
	return _Mailer.Contract.Owner(&_Mailer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailer *MailerCallerSession) Owner() (common.Address, error) {
	return _Mailer.Contract.Owner(&_Mailer.CallOpts)
}

// ZkBridgeEntrypoint is a free data retrieval call binding the contract method 0x904c8116.
//
// Solidity: function zkBridgeEntrypoint() view returns(address)
func (_Mailer *MailerCaller) ZkBridgeEntrypoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailer.contract.Call(opts, &out, "zkBridgeEntrypoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkBridgeEntrypoint is a free data retrieval call binding the contract method 0x904c8116.
//
// Solidity: function zkBridgeEntrypoint() view returns(address)
func (_Mailer *MailerSession) ZkBridgeEntrypoint() (common.Address, error) {
	return _Mailer.Contract.ZkBridgeEntrypoint(&_Mailer.CallOpts)
}

// ZkBridgeEntrypoint is a free data retrieval call binding the contract method 0x904c8116.
//
// Solidity: function zkBridgeEntrypoint() view returns(address)
func (_Mailer *MailerCallerSession) ZkBridgeEntrypoint() (common.Address, error) {
	return _Mailer.Contract.ZkBridgeEntrypoint(&_Mailer.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Mailer *MailerTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailer.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Mailer *MailerSession) ClaimFees() (*types.Transaction, error) {
	return _Mailer.Contract.ClaimFees(&_Mailer.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Mailer *MailerTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Mailer.Contract.ClaimFees(&_Mailer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailer *MailerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailer *MailerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailer.Contract.RenounceOwnership(&_Mailer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailer *MailerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailer.Contract.RenounceOwnership(&_Mailer.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x40925e95.
//
// Solidity: function sendMessage(uint16 dstChainId, address dstAddress, address recipient, string message) payable returns()
func (_Mailer *MailerTransactor) SendMessage(opts *bind.TransactOpts, dstChainId uint16, dstAddress common.Address, recipient common.Address, message string) (*types.Transaction, error) {
	return _Mailer.contract.Transact(opts, "sendMessage", dstChainId, dstAddress, recipient, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x40925e95.
//
// Solidity: function sendMessage(uint16 dstChainId, address dstAddress, address recipient, string message) payable returns()
func (_Mailer *MailerSession) SendMessage(dstChainId uint16, dstAddress common.Address, recipient common.Address, message string) (*types.Transaction, error) {
	return _Mailer.Contract.SendMessage(&_Mailer.TransactOpts, dstChainId, dstAddress, recipient, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x40925e95.
//
// Solidity: function sendMessage(uint16 dstChainId, address dstAddress, address recipient, string message) payable returns()
func (_Mailer *MailerTransactorSession) SendMessage(dstChainId uint16, dstAddress common.Address, recipient common.Address, message string) (*types.Transaction, error) {
	return _Mailer.Contract.SendMessage(&_Mailer.TransactOpts, dstChainId, dstAddress, recipient, message)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Mailer *MailerTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Mailer.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Mailer *MailerSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Mailer.Contract.SetFee(&_Mailer.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Mailer *MailerTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Mailer.Contract.SetFee(&_Mailer.TransactOpts, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailer *MailerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mailer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailer *MailerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailer.Contract.TransferOwnership(&_Mailer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailer *MailerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailer.Contract.TransferOwnership(&_Mailer.TransactOpts, newOwner)
}

// MailerMessageSendIterator is returned from FilterMessageSend and is used to iterate over the raw logs and unpacked data for MessageSend events raised by the Mailer contract.
type MailerMessageSendIterator struct {
	Event *MailerMessageSend // Event containing the contract specifics and raw log

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
func (it *MailerMessageSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailerMessageSend)
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
		it.Event = new(MailerMessageSend)
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
func (it *MailerMessageSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailerMessageSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailerMessageSend represents a MessageSend event raised by the Mailer contract.
type MailerMessageSend struct {
	Sequence   uint64
	DstChainId uint32
	DstAddress common.Address
	Sender     common.Address
	Recipient  common.Address
	Message    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageSend is a free log retrieval operation binding the contract event 0xbdb4daeca4a1c274318631b2e6569f5609c49b073f20d5925b06f069b71c84d4.
//
// Solidity: event MessageSend(uint64 indexed sequence, uint32 indexed dstChainId, address indexed dstAddress, address sender, address recipient, string message)
func (_Mailer *MailerFilterer) FilterMessageSend(opts *bind.FilterOpts, sequence []uint64, dstChainId []uint32, dstAddress []common.Address) (*MailerMessageSendIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _Mailer.contract.FilterLogs(opts, "MessageSend", sequenceRule, dstChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &MailerMessageSendIterator{contract: _Mailer.contract, event: "MessageSend", logs: logs, sub: sub}, nil
}

// WatchMessageSend is a free log subscription operation binding the contract event 0xbdb4daeca4a1c274318631b2e6569f5609c49b073f20d5925b06f069b71c84d4.
//
// Solidity: event MessageSend(uint64 indexed sequence, uint32 indexed dstChainId, address indexed dstAddress, address sender, address recipient, string message)
func (_Mailer *MailerFilterer) WatchMessageSend(opts *bind.WatchOpts, sink chan<- *MailerMessageSend, sequence []uint64, dstChainId []uint32, dstAddress []common.Address) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _Mailer.contract.WatchLogs(opts, "MessageSend", sequenceRule, dstChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailerMessageSend)
				if err := _Mailer.contract.UnpackLog(event, "MessageSend", log); err != nil {
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

// ParseMessageSend is a log parse operation binding the contract event 0xbdb4daeca4a1c274318631b2e6569f5609c49b073f20d5925b06f069b71c84d4.
//
// Solidity: event MessageSend(uint64 indexed sequence, uint32 indexed dstChainId, address indexed dstAddress, address sender, address recipient, string message)
func (_Mailer *MailerFilterer) ParseMessageSend(log types.Log) (*MailerMessageSend, error) {
	event := new(MailerMessageSend)
	if err := _Mailer.contract.UnpackLog(event, "MessageSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mailer contract.
type MailerOwnershipTransferredIterator struct {
	Event *MailerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MailerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailerOwnershipTransferred)
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
		it.Event = new(MailerOwnershipTransferred)
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
func (it *MailerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailerOwnershipTransferred represents a OwnershipTransferred event raised by the Mailer contract.
type MailerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailer *MailerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MailerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MailerOwnershipTransferredIterator{contract: _Mailer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailer *MailerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MailerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailerOwnershipTransferred)
				if err := _Mailer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailer *MailerFilterer) ParseOwnershipTransferred(log types.Log) (*MailerOwnershipTransferred, error) {
	event := new(MailerOwnershipTransferred)
	if err := _Mailer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
