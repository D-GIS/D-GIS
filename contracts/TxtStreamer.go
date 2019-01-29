// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TxtStreamerABI is the input ABI used to generate the binding from.
const TxtStreamerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"resp\",\"type\":\"uint256\"}],\"name\":\"Tell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"mResponse\",\"type\":\"uint256\"}],\"name\":\"TxtResponse\",\"type\":\"event\"}]"

// TxtStreamerBin is the compiled bytecode used for deploying new contracts.
const TxtStreamerBin = `{
	"linkReferences": {},
	"object": "608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610135806100606000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063740179f714610046575b600080fd5b34801561005257600080fd5b5061007160048036038101908080359060200190929190505050610073565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515156100cf57600080fd5b7f622c5f888423fcffc5ee7e195d932d9a46c6cc0eed08e12d48e9eb78dd8e425a816040518082815260200191505060405180910390a1505600a165627a7a7230582094204d2646c2c7e40601e6e12f9714c7bd7c9a8c7dd115ae14574766f37ebb300029",
	"opcodes": "PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP CALLER PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP PUSH2 0x135 DUP1 PUSH2 0x60 PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x41 JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x740179F7 EQ PUSH2 0x46 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST CALLVALUE DUP1 ISZERO PUSH2 0x52 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH2 0x71 PUSH1 0x4 DUP1 CALLDATASIZE SUB DUP2 ADD SWAP1 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP3 SWAP2 SWAP1 POP POP POP PUSH2 0x73 JUMP JUMPDEST STOP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO ISZERO PUSH2 0xCF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH32 0x622C5F888423FCFFC5EE7E195D932D9A46C6CC0EED08E12D48E9EB78DD8E425A DUP2 PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG1 POP JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 SWAP5 KECCAK256 0x4d 0x26 0x46 0xc2 0xc7 0xe4 MOD ADD 0xe6 0xe1 0x2f SWAP8 EQ 0xc7 0xbd PUSH29 0x9A8C7DD115AE14574766F37EBB30002900000000000000000000000000 ",
	"sourceMap": "27:303:0:-;;;152:56;8:9:-1;5:2;;;30:1;27;20:12;5:2;152:56:0;191:10;183:5;;:18;;;;;;;;;;;;;;;;;;27:303;;;;;;"
}`

// DeployTxtStreamer deploys a new Ethereum contract, binding an instance of TxtStreamer to it.
func DeployTxtStreamer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TxtStreamer, error) {
	parsed, err := abi.JSON(strings.NewReader(TxtStreamerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TxtStreamerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TxtStreamer{TxtStreamerCaller: TxtStreamerCaller{contract: contract}, TxtStreamerTransactor: TxtStreamerTransactor{contract: contract}, TxtStreamerFilterer: TxtStreamerFilterer{contract: contract}}, nil
}

// TxtStreamer is an auto generated Go binding around an Ethereum contract.
type TxtStreamer struct {
	TxtStreamerCaller     // Read-only binding to the contract
	TxtStreamerTransactor // Write-only binding to the contract
	TxtStreamerFilterer   // Log filterer for contract events
}

// TxtStreamerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxtStreamerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStreamerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxtStreamerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStreamerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxtStreamerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxtStreamerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxtStreamerSession struct {
	Contract     *TxtStreamer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxtStreamerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxtStreamerCallerSession struct {
	Contract *TxtStreamerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TxtStreamerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxtStreamerTransactorSession struct {
	Contract     *TxtStreamerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TxtStreamerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxtStreamerRaw struct {
	Contract *TxtStreamer // Generic contract binding to access the raw methods on
}

// TxtStreamerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxtStreamerCallerRaw struct {
	Contract *TxtStreamerCaller // Generic read-only contract binding to access the raw methods on
}

// TxtStreamerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxtStreamerTransactorRaw struct {
	Contract *TxtStreamerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxtStreamer creates a new instance of TxtStreamer, bound to a specific deployed contract.
func NewTxtStreamer(address common.Address, backend bind.ContractBackend) (*TxtStreamer, error) {
	contract, err := bindTxtStreamer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxtStreamer{TxtStreamerCaller: TxtStreamerCaller{contract: contract}, TxtStreamerTransactor: TxtStreamerTransactor{contract: contract}, TxtStreamerFilterer: TxtStreamerFilterer{contract: contract}}, nil
}

// NewTxtStreamerCaller creates a new read-only instance of TxtStreamer, bound to a specific deployed contract.
func NewTxtStreamerCaller(address common.Address, caller bind.ContractCaller) (*TxtStreamerCaller, error) {
	contract, err := bindTxtStreamer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxtStreamerCaller{contract: contract}, nil
}

// NewTxtStreamerTransactor creates a new write-only instance of TxtStreamer, bound to a specific deployed contract.
func NewTxtStreamerTransactor(address common.Address, transactor bind.ContractTransactor) (*TxtStreamerTransactor, error) {
	contract, err := bindTxtStreamer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxtStreamerTransactor{contract: contract}, nil
}

// NewTxtStreamerFilterer creates a new log filterer instance of TxtStreamer, bound to a specific deployed contract.
func NewTxtStreamerFilterer(address common.Address, filterer bind.ContractFilterer) (*TxtStreamerFilterer, error) {
	contract, err := bindTxtStreamer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxtStreamerFilterer{contract: contract}, nil
}

// bindTxtStreamer binds a generic wrapper to an already deployed contract.
func bindTxtStreamer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxtStreamerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxtStreamer *TxtStreamerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxtStreamer.Contract.TxtStreamerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxtStreamer *TxtStreamerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStreamer.Contract.TxtStreamerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxtStreamer *TxtStreamerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxtStreamer.Contract.TxtStreamerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxtStreamer *TxtStreamerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxtStreamer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxtStreamer *TxtStreamerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxtStreamer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxtStreamer *TxtStreamerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxtStreamer.Contract.contract.Transact(opts, method, params...)
}

// Tell is a paid mutator transaction binding the contract method 0x740179f7.
//
// Solidity: function Tell(resp uint256) returns()
func (_TxtStreamer *TxtStreamerTransactor) Tell(opts *bind.TransactOpts, resp *big.Int) (*types.Transaction, error) {
	return _TxtStreamer.contract.Transact(opts, "Tell", resp)
}

// Tell is a paid mutator transaction binding the contract method 0x740179f7.
//
// Solidity: function Tell(resp uint256) returns()
func (_TxtStreamer *TxtStreamerSession) Tell(resp *big.Int) (*types.Transaction, error) {
	return _TxtStreamer.Contract.Tell(&_TxtStreamer.TransactOpts, resp)
}

// Tell is a paid mutator transaction binding the contract method 0x740179f7.
//
// Solidity: function Tell(resp uint256) returns()
func (_TxtStreamer *TxtStreamerTransactorSession) Tell(resp *big.Int) (*types.Transaction, error) {
	return _TxtStreamer.Contract.Tell(&_TxtStreamer.TransactOpts, resp)
}

// TxtStreamerTxtResponseIterator is returned from FilterTxtResponse and is used to iterate over the raw logs and unpacked data for TxtResponse events raised by the TxtStreamer contract.
type TxtStreamerTxtResponseIterator struct {
	Event *TxtStreamerTxtResponse // Event containing the contract specifics and raw log

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
func (it *TxtStreamerTxtResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TxtStreamerTxtResponse)
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
		it.Event = new(TxtStreamerTxtResponse)
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
func (it *TxtStreamerTxtResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TxtStreamerTxtResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TxtStreamerTxtResponse represents a TxtResponse event raised by the TxtStreamer contract.
type TxtStreamerTxtResponse struct {
	MResponse *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTxtResponse is a free log retrieval operation binding the contract event 0x622c5f888423fcffc5ee7e195d932d9a46c6cc0eed08e12d48e9eb78dd8e425a.
//
// Solidity: e TxtResponse(mResponse uint256)
func (_TxtStreamer *TxtStreamerFilterer) FilterTxtResponse(opts *bind.FilterOpts) (*TxtStreamerTxtResponseIterator, error) {

	logs, sub, err := _TxtStreamer.contract.FilterLogs(opts, "TxtResponse")
	if err != nil {
		return nil, err
	}
	return &TxtStreamerTxtResponseIterator{contract: _TxtStreamer.contract, event: "TxtResponse", logs: logs, sub: sub}, nil
}

// WatchTxtResponse is a free log subscription operation binding the contract event 0x622c5f888423fcffc5ee7e195d932d9a46c6cc0eed08e12d48e9eb78dd8e425a.
//
// Solidity: e TxtResponse(mResponse uint256)
func (_TxtStreamer *TxtStreamerFilterer) WatchTxtResponse(opts *bind.WatchOpts, sink chan<- *TxtStreamerTxtResponse) (event.Subscription, error) {

	logs, sub, err := _TxtStreamer.contract.WatchLogs(opts, "TxtResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TxtStreamerTxtResponse)
				if err := _TxtStreamer.contract.UnpackLog(event, "TxtResponse", log); err != nil {
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
