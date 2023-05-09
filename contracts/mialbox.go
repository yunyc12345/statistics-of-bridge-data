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

// MailboxMetaData contains all meta data concerning the Mailbox contract.
var MailboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zkBridgeReceiver\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"sourceChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"messagesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"zkReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MailboxMetaData.ABI instead.
var MailboxABI = MailboxMetaData.ABI

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// Messages is a free data retrieval call binding the contract method 0xc45d9bea.
//
// Solidity: function messages(address , uint256 ) view returns(address sender, string message)
func (_Mailbox *MailboxCaller) Messages(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "messages", arg0, arg1)

	outstruct := new(struct {
		Sender  common.Address
		Message string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0xc45d9bea.
//
// Solidity: function messages(address , uint256 ) view returns(address sender, string message)
func (_Mailbox *MailboxSession) Messages(arg0 common.Address, arg1 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _Mailbox.Contract.Messages(&_Mailbox.CallOpts, arg0, arg1)
}

// Messages is a free data retrieval call binding the contract method 0xc45d9bea.
//
// Solidity: function messages(address , uint256 ) view returns(address sender, string message)
func (_Mailbox *MailboxCallerSession) Messages(arg0 common.Address, arg1 *big.Int) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _Mailbox.Contract.Messages(&_Mailbox.CallOpts, arg0, arg1)
}

// MessagesLength is a free data retrieval call binding the contract method 0x66bf84d3.
//
// Solidity: function messagesLength(address recipient) view returns(uint256)
func (_Mailbox *MailboxCaller) MessagesLength(opts *bind.CallOpts, recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "messagesLength", recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessagesLength is a free data retrieval call binding the contract method 0x66bf84d3.
//
// Solidity: function messagesLength(address recipient) view returns(uint256)
func (_Mailbox *MailboxSession) MessagesLength(recipient common.Address) (*big.Int, error) {
	return _Mailbox.Contract.MessagesLength(&_Mailbox.CallOpts, recipient)
}

// MessagesLength is a free data retrieval call binding the contract method 0x66bf84d3.
//
// Solidity: function messagesLength(address recipient) view returns(uint256)
func (_Mailbox *MailboxCallerSession) MessagesLength(recipient common.Address) (*big.Int, error) {
	return _Mailbox.Contract.MessagesLength(&_Mailbox.CallOpts, recipient)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_Mailbox *MailboxTransactor) ZkReceive(opts *bind.TransactOpts, srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "zkReceive", srcChainId, srcAddress, sequence, payload)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_Mailbox *MailboxSession) ZkReceive(srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.ZkReceive(&_Mailbox.TransactOpts, srcChainId, srcAddress, sequence, payload)
}

// ZkReceive is a paid mutator transaction binding the contract method 0x2de9952a.
//
// Solidity: function zkReceive(uint16 srcChainId, address srcAddress, uint64 sequence, bytes payload) returns()
func (_Mailbox *MailboxTransactorSession) ZkReceive(srcChainId uint16, srcAddress common.Address, sequence uint64, payload []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.ZkReceive(&_Mailbox.TransactOpts, srcChainId, srcAddress, sequence, payload)
}

// MailboxMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the Mailbox contract.
type MailboxMessageReceivedIterator struct {
	Event *MailboxMessageReceived // Event containing the contract specifics and raw log

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
func (it *MailboxMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxMessageReceived)
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
		it.Event = new(MailboxMessageReceived)
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
func (it *MailboxMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxMessageReceived represents a MessageReceived event raised by the Mailbox contract.
type MailboxMessageReceived struct {
	Sequence      uint64
	SourceChainId uint32
	SourceAddress common.Address
	Sender        common.Address
	Recipient     common.Address
	Message       string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x348bed7710e1262dbcdd514b7ba2c2497294dbbd85ecb47e1f7c92b4d7d07ef7.
//
// Solidity: event MessageReceived(uint64 indexed sequence, uint32 indexed sourceChainId, address indexed sourceAddress, address sender, address recipient, string message)
func (_Mailbox *MailboxFilterer) FilterMessageReceived(opts *bind.FilterOpts, sequence []uint64, sourceChainId []uint32, sourceAddress []common.Address) (*MailboxMessageReceivedIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "MessageReceived", sequenceRule, sourceChainIdRule, sourceAddressRule)
	if err != nil {
		return nil, err
	}
	return &MailboxMessageReceivedIterator{contract: _Mailbox.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x348bed7710e1262dbcdd514b7ba2c2497294dbbd85ecb47e1f7c92b4d7d07ef7.
//
// Solidity: event MessageReceived(uint64 indexed sequence, uint32 indexed sourceChainId, address indexed sourceAddress, address sender, address recipient, string message)
func (_Mailbox *MailboxFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MailboxMessageReceived, sequence []uint64, sourceChainId []uint32, sourceAddress []common.Address) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var sourceAddressRule []interface{}
	for _, sourceAddressItem := range sourceAddress {
		sourceAddressRule = append(sourceAddressRule, sourceAddressItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "MessageReceived", sequenceRule, sourceChainIdRule, sourceAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxMessageReceived)
				if err := _Mailbox.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x348bed7710e1262dbcdd514b7ba2c2497294dbbd85ecb47e1f7c92b4d7d07ef7.
//
// Solidity: event MessageReceived(uint64 indexed sequence, uint32 indexed sourceChainId, address indexed sourceAddress, address sender, address recipient, string message)
func (_Mailbox *MailboxFilterer) ParseMessageReceived(log types.Log) (*MailboxMessageReceived, error) {
	event := new(MailboxMessageReceived)
	if err := _Mailbox.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
