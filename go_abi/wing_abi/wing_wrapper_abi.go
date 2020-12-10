// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wing_wrapper_abi

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"name\":\"_toContract\",\"type\":\"bytes\"},{\"name\":\"_method\",\"type\":\"bytes\"},{\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IEthCrossChainManagerProxyABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IEthCrossChainManagerProxyFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerProxyFuncSigs = map[string]string{
	"87939a7f": "getEthCrossChainManager()",
}

// IEthCrossChainManagerProxy is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManagerProxy struct {
	IEthCrossChainManagerProxyCaller     // Read-only binding to the contract
	IEthCrossChainManagerProxyTransactor // Write-only binding to the contract
	IEthCrossChainManagerProxyFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerProxySession struct {
	Contract     *IEthCrossChainManagerProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerProxyCallerSession struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IEthCrossChainManagerProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerProxyTransactorSession struct {
	Contract     *IEthCrossChainManagerProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyRaw struct {
	Contract *IEthCrossChainManagerProxy // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyCallerRaw struct {
	Contract *IEthCrossChainManagerProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerProxyTransactorRaw struct {
	Contract *IEthCrossChainManagerProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManagerProxy creates a new instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxy(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManagerProxy, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxy{IEthCrossChainManagerProxyCaller: IEthCrossChainManagerProxyCaller{contract: contract}, IEthCrossChainManagerProxyTransactor: IEthCrossChainManagerProxyTransactor{contract: contract}, IEthCrossChainManagerProxyFilterer: IEthCrossChainManagerProxyFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerProxyCaller creates a new read-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerProxyCaller, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyTransactor creates a new write-only instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerProxyTransactor, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerProxyFilterer creates a new log filterer instance of IEthCrossChainManagerProxy, bound to a specific deployed contract.
func NewIEthCrossChainManagerProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerProxyFilterer, error) {
	contract, err := bindIEthCrossChainManagerProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerProxyFilterer{contract: contract}, nil
}

// bindIEthCrossChainManagerProxy binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManagerProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.IEthCrossChainManagerProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManagerProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManagerProxy.Contract.contract.Transact(opts, method, params...)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCaller) GetEthCrossChainManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IEthCrossChainManagerProxy.contract.Call(opts, out, "getEthCrossChainManager")
	return *ret0, err
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxySession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_IEthCrossChainManagerProxy *IEthCrossChainManagerProxyCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _IEthCrossChainManagerProxy.Contract.GetEthCrossChainManager(&_IEthCrossChainManagerProxy.CallOpts)
}

// IEthWingWrapperABI is the input ABI used to generate the binding from.
const IEthWingWrapperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"ontAddr\",\"type\":\"bytes\"}],\"name\":\"supply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"bytes\"}],\"name\":\"withdrawWing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_assetContract\",\"type\":\"bytes\"},{\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthWingWrapperFuncSigs maps the 4-byte function signature to its string representation.
var IEthWingWrapperFuncSigs = map[string]string{
	"3fec5c61": "supply(address,uint256,bytes)",
	"9813954e": "withdraw(bytes,uint64)",
	"57632392": "withdrawWing(bytes)",
}

// IEthWingWrapper is an auto generated Go binding around an Ethereum contract.
type IEthWingWrapper struct {
	IEthWingWrapperCaller     // Read-only binding to the contract
	IEthWingWrapperTransactor // Write-only binding to the contract
	IEthWingWrapperFilterer   // Log filterer for contract events
}

// IEthWingWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthWingWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthWingWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthWingWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthWingWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthWingWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthWingWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthWingWrapperSession struct {
	Contract     *IEthWingWrapper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthWingWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthWingWrapperCallerSession struct {
	Contract *IEthWingWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IEthWingWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthWingWrapperTransactorSession struct {
	Contract     *IEthWingWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IEthWingWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthWingWrapperRaw struct {
	Contract *IEthWingWrapper // Generic contract binding to access the raw methods on
}

// IEthWingWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthWingWrapperCallerRaw struct {
	Contract *IEthWingWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// IEthWingWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthWingWrapperTransactorRaw struct {
	Contract *IEthWingWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthWingWrapper creates a new instance of IEthWingWrapper, bound to a specific deployed contract.
func NewIEthWingWrapper(address common.Address, backend bind.ContractBackend) (*IEthWingWrapper, error) {
	contract, err := bindIEthWingWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthWingWrapper{IEthWingWrapperCaller: IEthWingWrapperCaller{contract: contract}, IEthWingWrapperTransactor: IEthWingWrapperTransactor{contract: contract}, IEthWingWrapperFilterer: IEthWingWrapperFilterer{contract: contract}}, nil
}

// NewIEthWingWrapperCaller creates a new read-only instance of IEthWingWrapper, bound to a specific deployed contract.
func NewIEthWingWrapperCaller(address common.Address, caller bind.ContractCaller) (*IEthWingWrapperCaller, error) {
	contract, err := bindIEthWingWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthWingWrapperCaller{contract: contract}, nil
}

// NewIEthWingWrapperTransactor creates a new write-only instance of IEthWingWrapper, bound to a specific deployed contract.
func NewIEthWingWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthWingWrapperTransactor, error) {
	contract, err := bindIEthWingWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthWingWrapperTransactor{contract: contract}, nil
}

// NewIEthWingWrapperFilterer creates a new log filterer instance of IEthWingWrapper, bound to a specific deployed contract.
func NewIEthWingWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthWingWrapperFilterer, error) {
	contract, err := bindIEthWingWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthWingWrapperFilterer{contract: contract}, nil
}

// bindIEthWingWrapper binds a generic wrapper to an already deployed contract.
func bindIEthWingWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthWingWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthWingWrapper *IEthWingWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthWingWrapper.Contract.IEthWingWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthWingWrapper *IEthWingWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.IEthWingWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthWingWrapper *IEthWingWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.IEthWingWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthWingWrapper *IEthWingWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthWingWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthWingWrapper *IEthWingWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthWingWrapper *IEthWingWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.contract.Transact(opts, method, params...)
}

// Supply is a paid mutator transaction binding the contract method 0x3fec5c61.
//
// Solidity: function supply(address token, uint256 amount, bytes ontAddr) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactor) Supply(opts *bind.TransactOpts, token common.Address, amount *big.Int, ontAddr []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.contract.Transact(opts, "supply", token, amount, ontAddr)
}

// Supply is a paid mutator transaction binding the contract method 0x3fec5c61.
//
// Solidity: function supply(address token, uint256 amount, bytes ontAddr) returns()
func (_IEthWingWrapper *IEthWingWrapperSession) Supply(token common.Address, amount *big.Int, ontAddr []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.Supply(&_IEthWingWrapper.TransactOpts, token, amount, ontAddr)
}

// Supply is a paid mutator transaction binding the contract method 0x3fec5c61.
//
// Solidity: function supply(address token, uint256 amount, bytes ontAddr) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactorSession) Supply(token common.Address, amount *big.Int, ontAddr []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.Supply(&_IEthWingWrapper.TransactOpts, token, amount, ontAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9813954e.
//
// Solidity: function withdraw(bytes _assetContract, uint64 _amount) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactor) Withdraw(opts *bind.TransactOpts, _assetContract []byte, _amount uint64) (*types.Transaction, error) {
	return _IEthWingWrapper.contract.Transact(opts, "withdraw", _assetContract, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9813954e.
//
// Solidity: function withdraw(bytes _assetContract, uint64 _amount) returns()
func (_IEthWingWrapper *IEthWingWrapperSession) Withdraw(_assetContract []byte, _amount uint64) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.Withdraw(&_IEthWingWrapper.TransactOpts, _assetContract, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9813954e.
//
// Solidity: function withdraw(bytes _assetContract, uint64 _amount) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactorSession) Withdraw(_assetContract []byte, _amount uint64) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.Withdraw(&_IEthWingWrapper.TransactOpts, _assetContract, _amount)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0x57632392.
//
// Solidity: function withdrawWing(bytes targetAddress) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactor) WithdrawWing(opts *bind.TransactOpts, targetAddress []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.contract.Transact(opts, "withdrawWing", targetAddress)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0x57632392.
//
// Solidity: function withdrawWing(bytes targetAddress) returns()
func (_IEthWingWrapper *IEthWingWrapperSession) WithdrawWing(targetAddress []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.WithdrawWing(&_IEthWingWrapper.TransactOpts, targetAddress)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0x57632392.
//
// Solidity: function withdrawWing(bytes targetAddress) returns()
func (_IEthWingWrapper *IEthWingWrapperTransactorSession) WithdrawWing(targetAddress []byte) (*types.Transaction, error) {
	return _IEthWingWrapper.Contract.WithdrawWing(&_IEthWingWrapper.TransactOpts, targetAddress)
}

// ILockProxyABI is the input ABI used to generate the binding from.
const ILockProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"toAssetHash\",\"type\":\"bytes\"}],\"name\":\"bindAssetHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"targetProxyHash\",\"type\":\"bytes\"}],\"name\":\"bindProxyHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"assetHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getBalanceFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromAssetHash\",\"type\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint64\"},{\"name\":\"toAddress\",\"type\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proxyHashMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILockProxyFuncSigs maps the 4-byte function signature to its string representation.
var ILockProxyFuncSigs = map[string]string{
	"4f7d9808": "assetHashMap(address,uint64)",
	"3348f63b": "bindAssetHash(address,uint64,bytes)",
	"379b98f6": "bindProxyHash(uint64,bytes)",
	"59c589a1": "getBalanceFor(address)",
	"84a6d055": "lock(address,uint64,bytes,uint256)",
	"d798f881": "managerProxyContract()",
	"9e5767aa": "proxyHashMap(uint64)",
}

// ILockProxy is an auto generated Go binding around an Ethereum contract.
type ILockProxy struct {
	ILockProxyCaller     // Read-only binding to the contract
	ILockProxyTransactor // Write-only binding to the contract
	ILockProxyFilterer   // Log filterer for contract events
}

// ILockProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILockProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILockProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILockProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILockProxySession struct {
	Contract     *ILockProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILockProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILockProxyCallerSession struct {
	Contract *ILockProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ILockProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILockProxyTransactorSession struct {
	Contract     *ILockProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ILockProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILockProxyRaw struct {
	Contract *ILockProxy // Generic contract binding to access the raw methods on
}

// ILockProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILockProxyCallerRaw struct {
	Contract *ILockProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ILockProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILockProxyTransactorRaw struct {
	Contract *ILockProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILockProxy creates a new instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxy(address common.Address, backend bind.ContractBackend) (*ILockProxy, error) {
	contract, err := bindILockProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILockProxy{ILockProxyCaller: ILockProxyCaller{contract: contract}, ILockProxyTransactor: ILockProxyTransactor{contract: contract}, ILockProxyFilterer: ILockProxyFilterer{contract: contract}}, nil
}

// NewILockProxyCaller creates a new read-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyCaller(address common.Address, caller bind.ContractCaller) (*ILockProxyCaller, error) {
	contract, err := bindILockProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyCaller{contract: contract}, nil
}

// NewILockProxyTransactor creates a new write-only instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ILockProxyTransactor, error) {
	contract, err := bindILockProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILockProxyTransactor{contract: contract}, nil
}

// NewILockProxyFilterer creates a new log filterer instance of ILockProxy, bound to a specific deployed contract.
func NewILockProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ILockProxyFilterer, error) {
	contract, err := bindILockProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILockProxyFilterer{contract: contract}, nil
}

// bindILockProxy binds a generic wrapper to an already deployed contract.
func bindILockProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILockProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.ILockProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.ILockProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockProxy *ILockProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ILockProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockProxy *ILockProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockProxy *ILockProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockProxy.Contract.contract.Transact(opts, method, params...)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCaller) AssetHashMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ILockProxy.contract.Call(opts, out, "assetHashMap", arg0, arg1)
	return *ret0, err
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxySession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.CallOpts, arg0, arg1)
}

// AssetHashMap is a free data retrieval call binding the contract method 0x4f7d9808.
//
// Solidity: function assetHashMap(address , uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCallerSession) AssetHashMap(arg0 common.Address, arg1 uint64) ([]byte, error) {
	return _ILockProxy.Contract.AssetHashMap(&_ILockProxy.CallOpts, arg0, arg1)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxyCaller) GetBalanceFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ILockProxy.contract.Call(opts, out, "getBalanceFor", arg0)
	return *ret0, err
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxySession) GetBalanceFor(arg0 common.Address) (*big.Int, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.CallOpts, arg0)
}

// GetBalanceFor is a free data retrieval call binding the contract method 0x59c589a1.
//
// Solidity: function getBalanceFor(address ) view returns(uint256)
func (_ILockProxy *ILockProxyCallerSession) GetBalanceFor(arg0 common.Address) (*big.Int, error) {
	return _ILockProxy.Contract.GetBalanceFor(&_ILockProxy.CallOpts, arg0)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxyCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ILockProxy.contract.Call(opts, out, "managerProxyContract")
	return *ret0, err
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxySession) ManagerProxyContract() (common.Address, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_ILockProxy *ILockProxyCallerSession) ManagerProxyContract() (common.Address, error) {
	return _ILockProxy.Contract.ManagerProxyContract(&_ILockProxy.CallOpts)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCaller) ProxyHashMap(opts *bind.CallOpts, arg0 uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ILockProxy.contract.Call(opts, out, "proxyHashMap", arg0)
	return *ret0, err
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxySession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.CallOpts, arg0)
}

// ProxyHashMap is a free data retrieval call binding the contract method 0x9e5767aa.
//
// Solidity: function proxyHashMap(uint64 ) view returns(bytes)
func (_ILockProxy *ILockProxyCallerSession) ProxyHashMap(arg0 uint64) ([]byte, error) {
	return _ILockProxy.Contract.ProxyHashMap(&_ILockProxy.CallOpts, arg0)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindAssetHash(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindAssetHash", fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindAssetHash is a paid mutator transaction binding the contract method 0x3348f63b.
//
// Solidity: function bindAssetHash(address fromAssetHash, uint64 toChainId, bytes toAssetHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindAssetHash(fromAssetHash common.Address, toChainId uint64, toAssetHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindAssetHash(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAssetHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactor) BindProxyHash(opts *bind.TransactOpts, toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "bindProxyHash", toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxySession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// BindProxyHash is a paid mutator transaction binding the contract method 0x379b98f6.
//
// Solidity: function bindProxyHash(uint64 toChainId, bytes targetProxyHash) returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) BindProxyHash(toChainId uint64, targetProxyHash []byte) (*types.Transaction, error) {
	return _ILockProxy.Contract.BindProxyHash(&_ILockProxy.TransactOpts, toChainId, targetProxyHash)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxyTransactor) Lock(opts *bind.TransactOpts, fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.contract.Transact(opts, "lock", fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxySession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// Lock is a paid mutator transaction binding the contract method 0x84a6d055.
//
// Solidity: function lock(address fromAssetHash, uint64 toChainId, bytes toAddress, uint256 amount) payable returns(bool)
func (_ILockProxy *ILockProxyTransactorSession) Lock(fromAssetHash common.Address, toChainId uint64, toAddress []byte, amount *big.Int) (*types.Transaction, error) {
	return _ILockProxy.Contract.Lock(&_ILockProxy.TransactOpts, fromAssetHash, toChainId, toAddress, amount)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820625353f4c9e9f2bcd4304018d21791a5cf770202850ef8ea55b8acf7b394104f0029"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582098b831a563987e87571c0f2ce9f15cd36a69d48a882b78124bcc9edf3d7e193c0029"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b8f3e82eaa218b45b03801be7700fb2fc46659a698b57d7d4534629f8fc20ba10029"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// WingwrapperABI is the input ABI used to generate the binding from.
const WingwrapperABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"polyLockProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateRequestStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"toChainId\",\"type\":\"uint64\"}],\"name\":\"supply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ontProxy\",\"type\":\"bytes\"}],\"name\":\"updateOntLockProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userAgentContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawFeeThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyFeeThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requestIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"toChainId\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ontLockProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"managerProxyContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint64\"}],\"name\":\"withdrawWing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"updatePolyProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_polyLockProxy\",\"type\":\"address\"},{\"name\":\"_userAgentContract\",\"type\":\"bytes\"},{\"name\":\"_managerProxyContract\",\"type\":\"address\"},{\"name\":\"_ontLockproxy\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldProxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"PolyLockProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldProxy\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"newProxy\",\"type\":\"bytes\"}],\"name\":\"OntLockProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"SupplyFeeThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"WithdrawFeeThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"target\",\"type\":\"bytes\"}],\"name\":\"WithdrawWing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"oeprationType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"RequestUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFee\",\"type\":\"event\"}]"

// WingwrapperFuncSigs maps the 4-byte function signature to its string representation.
var WingwrapperFuncSigs = map[string]string{
	"f851a440": "admin()",
	"d798f881": "managerProxyContract()",
	"a1057f86": "ontLockProxy()",
	"21526180": "polyLockProxy()",
	"8163ade3": "requestIndex(address)",
	"301bf698": "supply(address,uint256,uint64)",
	"6115c522": "supplyFeeThreshold()",
	"40d34aa0": "updateOntLockProxy(bytes)",
	"f572f9b0": "updatePolyProxy(address)",
	"2b00754f": "updateRequestStatus(address,uint256,uint8)",
	"445adb1d": "userAgentContract()",
	"8693a290": "withdraw(address,uint256,uint64)",
	"55f1f1ed": "withdrawFeeThreshold()",
	"da15ae4e": "withdrawWing(bytes,uint64)",
}

// WingwrapperBin is the compiled bytecode used for deploying new contracts.
var WingwrapperBin = "0x60806040523480156200001157600080fd5b506040516200332238038062003322833981018060405260a08110156200003757600080fd5b81516020830151604084018051929491938201926401000000008111156200005e57600080fd5b820160208101848111156200007257600080fd5b81516401000000008111828201871017156200008d57600080fd5b50506020820151604090920180519194929391640100000000811115620000b357600080fd5b82016020810184811115620000c757600080fd5b8151640100000000811182820187101715620000e257600080fd5b505060068054600160a060020a03808b16600160a060020a03199283161790925560028054928a169290911691909117905585519093506200012e92506003915060208601906200016b565b5060058054600160a060020a031916600160a060020a03841617905580516200015f9060049060208401906200016b565b50505050505062000210565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001ae57805160ff1916838001178555620001de565b82800160010185558215620001de579182015b82811115620001de578251825591602001919060010190620001c1565b50620001ec929150620001f0565b5090565b6200020d91905b80821115620001ec5760008155600101620001f7565b90565b61310280620002206000396000f3fe6080604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632152618081146100d45780632b00754f14610105578063301bf6981461014957806340d34aa014610185578063445adb1d1461023857806355f1f1ed146102c25780636115c522146102e95780638163ade3146102fe5780638693a29014610331578063a1057f861461037a578063d798f8811461038f578063da15ae4e146103a4578063f572f9b014610463578063f851a44014610496575b600080fd5b3480156100e057600080fd5b506100e96104ab565b60408051600160a060020a039092168252519081900360200190f35b34801561011157600080fd5b506101476004803603606081101561012857600080fd5b508035600160a060020a0316906020810135906040013560ff166104ba565b005b6101476004803603606081101561015f57600080fd5b508035600160a060020a0316906020810135906040013567ffffffffffffffff166106bb565b34801561019157600080fd5b50610147600480360360208110156101a857600080fd5b8101906020810181356401000000008111156101c357600080fd5b8201836020820111156101d557600080fd5b803590602001918460018302840111640100000000831117156101f757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611169945050505050565b34801561024457600080fd5b5061024d61138d565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561028757818101518382015260200161026f565b50505050905090810190601f1680156102b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102ce57600080fd5b506102d761141b565b60408051918252519081900360200190f35b3480156102f557600080fd5b506102d7611421565b34801561030a57600080fd5b506102d76004803603602081101561032157600080fd5b5035600160a060020a0316611427565b34801561033d57600080fd5b506101476004803603606081101561035457600080fd5b508035600160a060020a0316906020810135906040013567ffffffffffffffff16611439565b34801561038657600080fd5b5061024d611998565b34801561039b57600080fd5b506100e96119f3565b3480156103b057600080fd5b50610147600480360360408110156103c757600080fd5b8101906020810181356401000000008111156103e257600080fd5b8201836020820111156103f457600080fd5b8035906020019184600183028401116401000000008311171561041657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903567ffffffffffffffff169150611a029050565b34801561046f57600080fd5b506101476004803603602081101561048657600080fd5b5035600160a060020a0316611e66565b3480156104a257600080fd5b506100e9611f38565b600254600160a060020a031681565b600654600160a060020a0316331461051c576040805160e560020a62461bcd02815260206004820152600960248201527f6f6e6c7920706f6c790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0383166000908152600160209081526040808320858452909152812090815460ff16600281111561055057fe5b14156105a6576040805160e560020a62461bcd02815260206004820152601660248201527f72657175657374206973206e6f74206578697374656400000000000000000000604482015290519081900360640190fd5b8160038111156105b257fe5b815460ff1660028111156105c257fe5b10610617576040805160e560020a62461bcd02815260206004820152601960248201527f73746174757320636f756c64206f6e6c7920666f727761726400000000000000604482015290519081900360640190fd5b80548290829061ff00191661010083600381111561063157fe5b0217905550805460408051600160a060020a0387168152602081018690527fcfe2e12491510e84790328447af892f6fcfc8d3ab0ca63b02b8e98f4b637b8de928792879260ff909216918791810183600281111561068b57fe5b60ff16815260200182600381111561069f57fe5b60ff16815260200194505050505060405180910390a150505050565b600254600160a060020a0316151561071d576040805160e560020a62461bcd02815260206004820152600a60248201527f706f6c7920656d70747900000000000000000000000000000000000000000000604482015290519081900360640190fd5b600254600160a060020a0316821515610780576040805160e560020a62461bcd02815260206004820152601f60248201527f616d6f756e742073686f756c642062652067726561746572207468616e203000604482015290519081900360640190fd5b600160a060020a0384161515610a125734151561080d576040805160e560020a62461bcd02815260206004820152602160248201527f7472616e736665727265642065746865722063616e6e6f74206265207a65726f60448201527f2100000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b34831461088a576040805160e560020a62461bcd02815260206004820152602960248201527f7472616e73666572726564206574686572206973206e6f7420657175616c207460448201527f6f20616d6f756e74210000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6040517f84a6d055000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830190815267ffffffffffffffff8516602484015260648301869052608060448401908152600380546002600019610100600184161502019091160460848601819052938616946384a6d0559489948b948a9493879360a490910190859080156109685780601f1061093d57610100808354040283529160200191610968565b820191906000526020600020905b81548152906001019060200180831161094b57829003601f168201915b5050955050505050506020604051808303818588803b15801561098a57600080fd5b505af115801561099e573d6000803e3d6000fd5b50505050506040513d60208110156109b557600080fd5b50511515610a0d576040805160e560020a62461bcd02815260206004820152600e60248201527f6c70206c6f636b206661696c6564000000000000000000000000000000000000604482015290519081900360640190fd5b610bd4565b8330610a2f600160a060020a03831633838863ffffffff611f4716565b600254610a4f90600160a060020a0384811691168763ffffffff611fd516565b6040517f84a6d055000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830190815267ffffffffffffffff8716602484015260648301889052608060448401908152600380546002600019610100600184161502019091160460848601819052938816946384a6d055948c948b948d939192909160a4019085908015610b2c5780601f10610b0157610100808354040283529160200191610b2c565b820191906000526020600020905b815481529060010190602001808311610b0f57829003601f168201915b505095505050505050602060405180830381600087803b158015610b4f57600080fd5b505af1158015610b63573d6000803e3d6000fd5b505050506040513d6020811015610b7957600080fd5b50511515610bd1576040805160e560020a62461bcd02815260206004820152600e60248201527f6c70206c6f636b206661696c6564000000000000000000000000000000000000604482015290519081900360640190fd5b50505b600554604080517f87939a7f0000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169160009183916387939a7f91600480820192602092909190829003018186803b158015610c3757600080fd5b505afa158015610c4b573d6000803e3d6000fd5b505050506040513d6020811015610c6157600080fd5b5051604080517f4f7d9808000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015267ffffffffffffffff8816602483015291519293508392606092871691634f7d9808916044808301926000929190829003018186803b158015610cdc57600080fd5b505afa158015610cf0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610d1957600080fd5b810190808051640100000000811115610d3157600080fd5b82016020810184811115610d4457600080fd5b8151640100000000811182820187101715610d5e57600080fd5b505080519094506000109250610dc1915050576040805160e560020a62461bcd02815260206004820152601460248201527f6e6f20746f41737365744861736820666f756e64000000000000000000000000604482015290519081900360640190fd5b6060610dce33838a612175565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152939450606093610e6c9390929091830182828015610e615780601f10610e3657610100808354040283529160200191610e61565b820191906000526020600020905b815481529060010190602001808311610e4457829003601f168201915b505050505083612555565b600454909150600260001961010060018416150201909116041515610edb576040805160e560020a62461bcd02815260206004820152601960248201527f656d70747920696c6c6567616c20746f50726f78794861736800000000000000604482015290519081900360640190fd5b6040517fbd5cf62500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8916600482810191825260806024840190815281546002600019610100600184161502019091160460848501819052600160a060020a0389169463bd5cf625948e94938893919290916044820191606481019160a49091019087908015610fb25780601f10610f8757610100808354040283529160200191610fb2565b820191906000526020600020905b815481529060010190602001808311610f9557829003601f168201915b5050848103835260128152602001807f696e766f6b655761736d436f6e74726163740000000000000000000000000000815250602001848103825285818151815260200191508051906020019080838360005b8381101561101d578181015183820152602001611005565b50505050905090810190601f16801561104a5780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561106d57600080fd5b505af1158015611081573d6000803e3d6000fd5b505050506040513d602081101561109757600080fd5b50511515611115576040805160e560020a62461bcd02815260206004820152602f60248201527f45746843726f7373436861696e4d616e616765722063726f7373436861696e2060448201527f6578656375746564206572726f72210000000000000000000000000000000000606482015290519081900360840190fd5b60408051338152600160a060020a038c1660208201528082018b905290517fd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e9181900360600190a150505050505050505050565b600654600160a060020a031633146111cb576040805160e560020a62461bcd02815260206004820152600a60248201527f6f6e6c792061646d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156112575780601f1061122c57610100808354040283529160200191611257565b820191906000526020600020905b81548152906001019060200180831161123a57829003601f168201915b505085519394506112739360049350602087019250905061303b565b507f91e64ed86813347dc94dfcd96757c0d1fb02d91a9fedc91ea9bbcdedbc08a576816004604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156112d95781810151838201526020016112c1565b50505050905090810190601f1680156113065780820380516001836020036101000a031916815260200191505b508381038252845460026000196101006001841615020190911604808252602090910190859080156113795780601f1061134e57610100808354040283529160200191611379565b820191906000526020600020905b81548152906001019060200180831161135c57829003601f168201915b505094505050505060405180910390a15050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156114135780601f106113e857610100808354040283529160200191611413565b820191906000526020600020905b8154815290600101906020018083116113f657829003601f168201915b505050505081565b60085481565b60075481565b60006020819052908152604090205481565b600254600160a060020a0316151561149b576040805160e560020a62461bcd02815260206004820152600a60248201527f706f6c7920656d70747900000000000000000000000000000000000000000000604482015290519081900360640190fd5b600554604080517f87939a7f0000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169160009183916387939a7f91600480820192602092909190829003018186803b1580156114fe57600080fd5b505afa158015611512573d6000803e3d6000fd5b505050506040513d602081101561152857600080fd5b5051600254604080517f4f7d9808000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015267ffffffffffffffff881660248301529151939450849391909216916060918391634f7d9808916044808301926000929190829003018186803b1580156115aa57600080fd5b505afa1580156115be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156115e757600080fd5b8101908080516401000000008111156115ff57600080fd5b8201602081018481111561161257600080fd5b815164010000000081118282018710171561162c57600080fd5b50508051909450600010925061168f915050576040805160e560020a62461bcd02815260206004820152601460248201527f6e6f20746f41737365744861736820666f756e64000000000000000000000000604482015290519081900360640190fd5b606061169c33838a612625565b60045490915060026000196101006001841615020190911604151561170b576040805160e560020a62461bcd02815260206004820152601960248201527f656d70747920696c6c6567616c20746f50726f78794861736800000000000000604482015290519081900360640190fd5b6040517fbd5cf62500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8816600482810191825260806024840190815281546002600019610100600184161502019091160460848501819052600160a060020a0389169463bd5cf625948d94938893919290916044820191606481019160a490910190879080156117e25780601f106117b7576101008083540402835291602001916117e2565b820191906000526020600020905b8154815290600101906020018083116117c557829003601f168201915b5050848103835260128152602001807f696e766f6b655761736d436f6e74726163740000000000000000000000000000815250602001848103825285818151815260200191508051906020019080838360005b8381101561184d578181015183820152602001611835565b50505050905090810190601f16801561187a5780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561189d57600080fd5b505af11580156118b1573d6000803e3d6000fd5b505050506040513d60208110156118c757600080fd5b50511515611945576040805160e560020a62461bcd02815260206004820152602f60248201527f45746843726f7373436861696e4d616e616765722063726f7373436861696e2060448201527f6578656375746564206572726f72210000000000000000000000000000000000606482015290519081900360840190fd5b60408051338152600160a060020a038b1660208201528082018a905290517f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9181900360600190a1505050505050505050565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156114135780601f106113e857610100808354040283529160200191611413565b600554600160a060020a031681565b600254600160a060020a03161515611a64576040805160e560020a62461bcd02815260206004820152600a60248201527f706f6c7920656d70747900000000000000000000000000000000000000000000604482015290519081900360640190fd5b600554604080517f87939a7f0000000000000000000000000000000000000000000000000000000081529051600160a060020a039092169160009183916387939a7f91600480820192602092909190829003018186803b158015611ac757600080fd5b505afa158015611adb573d6000803e3d6000fd5b505050506040513d6020811015611af157600080fd5b50519050806060611b0233876126a6565b600454909150600260001961010060018416150201909116041515611b71576040805160e560020a62461bcd02815260206004820152601960248201527f656d70747920696c6c6567616c20746f50726f78794861736800000000000000604482015290519081900360640190fd5b6040517fbd5cf62500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482810191825260806024840190815281546002600019610100600184161502019091160460848501819052600160a060020a0387169463bd5cf625948b94938893919290916044820191606481019160a49091019087908015611c485780601f10611c1d57610100808354040283529160200191611c48565b820191906000526020600020905b815481529060010190602001808311611c2b57829003601f168201915b5050848103835260128152602001807f696e766f6b655761736d436f6e74726163740000000000000000000000000000815250602001848103825285818151815260200191508051906020019080838360005b83811015611cb3578181015183820152602001611c9b565b50505050905090810190601f168015611ce05780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b158015611d0357600080fd5b505af1158015611d17573d6000803e3d6000fd5b505050506040513d6020811015611d2d57600080fd5b50511515611dab576040805160e560020a62461bcd02815260206004820152602f60248201527f45746843726f7373436861696e4d616e616765722063726f7373436861696e2060448201527f6578656375746564206572726f72210000000000000000000000000000000000606482015290519081900360840190fd5b7f91ac22144848c281e728c9336aa88e73dfde7c9829514adc019418766ec9dfa333876040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611e23578181015183820152602001611e0b565b50505050905090810190601f168015611e505780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1505050505050565b600654600160a060020a03163314611ec8576040805160e560020a62461bcd02815260206004820152600a60248201527f6f6e6c792061646d696e00000000000000000000000000000000000000000000604482015290519081900360640190fd5b60028054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040805191909216808252602082019390935281517f40bf952ebb271dccb7dd1b9244d2db50205559185e04bcc10bad51308cc62a4d929181900390910190a15050565b600654600160a060020a031681565b60408051600160a060020a0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611fcf908590612a04565b50505050565b8015806120745750604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152306004820152600160a060020a03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b15801561204657600080fd5b505afa15801561205a573d6000803e3d6000fd5b505050506040513d602081101561207057600080fd5b5051155b15156120f0576040805160e560020a62461bcd02815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000606482015290519081900360840190fd5b60408051600160a060020a038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052612170908490612a04565b505050565b604080518082018252600681527f737570706c79000000000000000000000000000000000000000000000000000060208083019190915282516c01000000000000000000000000600160a060020a03881602918101919091528251808203601401815260349091019092526060918290600060fc60020a6121f66004612c01565b600160f860020a026122088751612c01565b61221188612c46565b600160f860020a026122238851612c01565b88600260f860020a028f604051602001808c600160f860020a031916600160f860020a03191681526001018b600160f860020a031916600160f860020a03191681526001018a805190602001908083835b602083106122935780518252601f199092019160209182019101612274565b51815160209384036101000a6000190180199092169116179052600160f860020a03198d16919093019081528a516001909101928b0191508083835b602083106122ee5780518252601f1990920191602091820191016122cf565b51815160209384036101000a60001901801990921691161790528a5191909301928a0191508083835b602083106123365780518252601f199092019160209182019101612317565b51815160209384036101000a6000190180199092169116179052600160f860020a03198a16919093019081528751600190910192880191508083835b602083106123915780518252601f199092019160209182019101612372565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106123d95780518252601f1990920191602091820191016123ba565b51815160209384036101000a6000190180199092169116179052600160f860020a03198716919093019081528451600190910192850191508083835b602083106124345780518252601f199092019160209182019101612415565b6001836020036101000a0380198251168184511680821785525050505050509050019b505050505050505050505050604051602081830303815290604052915081600460f860020a0261248687612d0c565b6040516020018084805190602001908083835b602083106124b85780518252601f199092019160209182019101612499565b51815160209384036101000a6000190180199092169116179052600160f860020a03198716919093019081528451600190910192850191508083835b602083106125135780518252601f1990920191602091820191016124f4565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405293505050509392505050565b60608061256184612c46565b61256a84612c46565b6040516020018083805190602001908083835b6020831061259c5780518252601f19909201916020918201910161257d565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106125e45780518252601f1990920191602091820191016125c5565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290508091505092915050565b604080518082018252600881527f776974686472617700000000000000000000000000000000000000000000000060208083019190915282516c01000000000000000000000000600160a060020a03881602918101919091528251808203601401815260349091019092526060918290600060fc60020a6121f66004612c01565b604080518082018252600c81527f776974686472617757696e67000000000000000000000000000000000000000060208083019190915282516c01000000000000000000000000600160a060020a03871602918101919091528251808203601401815260349091019092526060918290600060fc60020a6127276003612c01565b600160f860020a026127398751612c01565b61274288612c46565b600160f860020a026127548851612c01565b88600160f860020a026127678f51612c01565b8f604051602001808d600160f860020a031916600160f860020a03191681526001018c600160f860020a031916600160f860020a03191681526001018b805190602001908083835b602083106127ce5780518252601f1990920191602091820191016127af565b51815160209384036101000a6000190180199092169116179052600160f860020a03198e16919093019081528b516001909101928c0191508083835b602083106128295780518252601f19909201916020918201910161280a565b51815160209384036101000a60001901801990921691161790528b5191909301928b0191508083835b602083106128715780518252601f199092019160209182019101612852565b51815160209384036101000a6000190180199092169116179052600160f860020a03198b16919093019081528851600190910192890191508083835b602083106128cc5780518252601f1990920191602091820191016128ad565b51815160209384036101000a600019018019909216911617905288519190930192880191508083835b602083106129145780518252601f1990920191602091820191016128f5565b51815160209384036101000a6000190180199092169116179052600160f860020a03198816919093019081528551600190910192860191508083835b6020831061296f5780518252601f199092019160209182019101612950565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106129b75780518252601f199092019160209182019101612998565b6001836020036101000a0380198251168184511680821785525050505050509050019c50505050505050505050505050604051602081830303815290604052915081935050505092915050565b612a0d82612dc5565b1515612a63576040805160e560020a62461bcd02815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b6000606083600160a060020a0316836040518082805190602001908083835b60208310612aa15780518252601f199092019160209182019101612a82565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612b03576040519150601f19603f3d011682016040523d82523d6000602084013e612b08565b606091505b5091509150811515612b64576040805160e560020a62461bcd02815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b600081511115611fcf57808060200190516020811015612b8357600080fd5b50511515611fcf576040805160e560020a62461bcd02815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015290519081900360840190fd5b6040516002808252606091906000601f5b82821015612c345785811a826020860101536001919091019060001901612c12565b5050506022810160405290505b919050565b8051606090612c5481612e01565b836040516020018083805190602001908083835b60208310612c875780518252601f199092019160209182019101612c68565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310612ccf5780518252601f199092019160209182019101612cb0565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b60607f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115612d86576040805160e560020a62461bcd02815260206004820152601b60248201527f56616c756520657863656564732075696e743235352072616e67650000000000604482015290519081900360640190fd5b60405160208082526000601f5b82821015612db55785811a826020860101536001919091019060001901612d93565b5050506040818101905292915050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708115801590612df95750808214155b949350505050565b606060fd8267ffffffffffffffff161015612e2657612e1f82612f82565b9050612c41565b61ffff67ffffffffffffffff831611612f2157612e627ffd00000000000000000000000000000000000000000000000000000000000000612f9e565b612e6b83612c01565b6040516020018083805190602001908083835b60208310612e9d5780518252601f199092019160209182019101612e7e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310612ee55780518252601f199092019160209182019101612ec6565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050612c41565b63ffffffff67ffffffffffffffff831611612f6857612f5f7ffe00000000000000000000000000000000000000000000000000000000000000612f9e565b612e6b83612fb5565b612f79600160f860020a0319612f9e565b612e6b83612ff8565b604080516001815260f89290921b602083015260218201905290565b6060612faf60f860020a8304612f82565b92915050565b6040516004808252606091906000601f5b82821015612fe85785811a826020860101536001919091019060001901612fc6565b5050506024810160405292915050565b6040516008808252606091906000601f5b8282101561302b5785811a826020860101536001919091019060001901613009565b5050506028810160405292915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061307c57805160ff19168380011785556130a9565b828001600101855582156130a9579182015b828111156130a957825182559160200191906001019061308e565b506130b59291506130b9565b5090565b6130d391905b808211156130b557600081556001016130bf565b9056fea165627a7a723058205d87950c0d8b1091bd6e0d0c65b0043b440aed1482ebcaeb0eb8736359e5b2b20029"

// DeployWingwrapper deploys a new Ethereum contract, binding an instance of Wingwrapper to it.
func DeployWingwrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _polyLockProxy common.Address, _userAgentContract []byte, _managerProxyContract common.Address, _ontLockproxy []byte) (common.Address, *types.Transaction, *Wingwrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(WingwrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WingwrapperBin), backend, _admin, _polyLockProxy, _userAgentContract, _managerProxyContract, _ontLockproxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wingwrapper{WingwrapperCaller: WingwrapperCaller{contract: contract}, WingwrapperTransactor: WingwrapperTransactor{contract: contract}, WingwrapperFilterer: WingwrapperFilterer{contract: contract}}, nil
}

// Wingwrapper is an auto generated Go binding around an Ethereum contract.
type Wingwrapper struct {
	WingwrapperCaller     // Read-only binding to the contract
	WingwrapperTransactor // Write-only binding to the contract
	WingwrapperFilterer   // Log filterer for contract events
}

// WingwrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type WingwrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WingwrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WingwrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WingwrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WingwrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WingwrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WingwrapperSession struct {
	Contract     *Wingwrapper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WingwrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WingwrapperCallerSession struct {
	Contract *WingwrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WingwrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WingwrapperTransactorSession struct {
	Contract     *WingwrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WingwrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type WingwrapperRaw struct {
	Contract *Wingwrapper // Generic contract binding to access the raw methods on
}

// WingwrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WingwrapperCallerRaw struct {
	Contract *WingwrapperCaller // Generic read-only contract binding to access the raw methods on
}

// WingwrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WingwrapperTransactorRaw struct {
	Contract *WingwrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWingwrapper creates a new instance of Wingwrapper, bound to a specific deployed contract.
func NewWingwrapper(address common.Address, backend bind.ContractBackend) (*Wingwrapper, error) {
	contract, err := bindWingwrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wingwrapper{WingwrapperCaller: WingwrapperCaller{contract: contract}, WingwrapperTransactor: WingwrapperTransactor{contract: contract}, WingwrapperFilterer: WingwrapperFilterer{contract: contract}}, nil
}

// NewWingwrapperCaller creates a new read-only instance of Wingwrapper, bound to a specific deployed contract.
func NewWingwrapperCaller(address common.Address, caller bind.ContractCaller) (*WingwrapperCaller, error) {
	contract, err := bindWingwrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WingwrapperCaller{contract: contract}, nil
}

// NewWingwrapperTransactor creates a new write-only instance of Wingwrapper, bound to a specific deployed contract.
func NewWingwrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*WingwrapperTransactor, error) {
	contract, err := bindWingwrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WingwrapperTransactor{contract: contract}, nil
}

// NewWingwrapperFilterer creates a new log filterer instance of Wingwrapper, bound to a specific deployed contract.
func NewWingwrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*WingwrapperFilterer, error) {
	contract, err := bindWingwrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WingwrapperFilterer{contract: contract}, nil
}

// bindWingwrapper binds a generic wrapper to an already deployed contract.
func bindWingwrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WingwrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wingwrapper *WingwrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wingwrapper.Contract.WingwrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wingwrapper *WingwrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wingwrapper.Contract.WingwrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wingwrapper *WingwrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wingwrapper.Contract.WingwrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wingwrapper *WingwrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wingwrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wingwrapper *WingwrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wingwrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wingwrapper *WingwrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wingwrapper.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wingwrapper *WingwrapperCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wingwrapper *WingwrapperSession) Admin() (common.Address, error) {
	return _Wingwrapper.Contract.Admin(&_Wingwrapper.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Wingwrapper *WingwrapperCallerSession) Admin() (common.Address, error) {
	return _Wingwrapper.Contract.Admin(&_Wingwrapper.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Wingwrapper *WingwrapperCaller) ManagerProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "managerProxyContract")
	return *ret0, err
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Wingwrapper *WingwrapperSession) ManagerProxyContract() (common.Address, error) {
	return _Wingwrapper.Contract.ManagerProxyContract(&_Wingwrapper.CallOpts)
}

// ManagerProxyContract is a free data retrieval call binding the contract method 0xd798f881.
//
// Solidity: function managerProxyContract() view returns(address)
func (_Wingwrapper *WingwrapperCallerSession) ManagerProxyContract() (common.Address, error) {
	return _Wingwrapper.Contract.ManagerProxyContract(&_Wingwrapper.CallOpts)
}

// OntLockProxy is a free data retrieval call binding the contract method 0xa1057f86.
//
// Solidity: function ontLockProxy() view returns(bytes)
func (_Wingwrapper *WingwrapperCaller) OntLockProxy(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "ontLockProxy")
	return *ret0, err
}

// OntLockProxy is a free data retrieval call binding the contract method 0xa1057f86.
//
// Solidity: function ontLockProxy() view returns(bytes)
func (_Wingwrapper *WingwrapperSession) OntLockProxy() ([]byte, error) {
	return _Wingwrapper.Contract.OntLockProxy(&_Wingwrapper.CallOpts)
}

// OntLockProxy is a free data retrieval call binding the contract method 0xa1057f86.
//
// Solidity: function ontLockProxy() view returns(bytes)
func (_Wingwrapper *WingwrapperCallerSession) OntLockProxy() ([]byte, error) {
	return _Wingwrapper.Contract.OntLockProxy(&_Wingwrapper.CallOpts)
}

// PolyLockProxy is a free data retrieval call binding the contract method 0x21526180.
//
// Solidity: function polyLockProxy() view returns(address)
func (_Wingwrapper *WingwrapperCaller) PolyLockProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "polyLockProxy")
	return *ret0, err
}

// PolyLockProxy is a free data retrieval call binding the contract method 0x21526180.
//
// Solidity: function polyLockProxy() view returns(address)
func (_Wingwrapper *WingwrapperSession) PolyLockProxy() (common.Address, error) {
	return _Wingwrapper.Contract.PolyLockProxy(&_Wingwrapper.CallOpts)
}

// PolyLockProxy is a free data retrieval call binding the contract method 0x21526180.
//
// Solidity: function polyLockProxy() view returns(address)
func (_Wingwrapper *WingwrapperCallerSession) PolyLockProxy() (common.Address, error) {
	return _Wingwrapper.Contract.PolyLockProxy(&_Wingwrapper.CallOpts)
}

// RequestIndex is a free data retrieval call binding the contract method 0x8163ade3.
//
// Solidity: function requestIndex(address ) view returns(uint256)
func (_Wingwrapper *WingwrapperCaller) RequestIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "requestIndex", arg0)
	return *ret0, err
}

// RequestIndex is a free data retrieval call binding the contract method 0x8163ade3.
//
// Solidity: function requestIndex(address ) view returns(uint256)
func (_Wingwrapper *WingwrapperSession) RequestIndex(arg0 common.Address) (*big.Int, error) {
	return _Wingwrapper.Contract.RequestIndex(&_Wingwrapper.CallOpts, arg0)
}

// RequestIndex is a free data retrieval call binding the contract method 0x8163ade3.
//
// Solidity: function requestIndex(address ) view returns(uint256)
func (_Wingwrapper *WingwrapperCallerSession) RequestIndex(arg0 common.Address) (*big.Int, error) {
	return _Wingwrapper.Contract.RequestIndex(&_Wingwrapper.CallOpts, arg0)
}

// SupplyFeeThreshold is a free data retrieval call binding the contract method 0x6115c522.
//
// Solidity: function supplyFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperCaller) SupplyFeeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "supplyFeeThreshold")
	return *ret0, err
}

// SupplyFeeThreshold is a free data retrieval call binding the contract method 0x6115c522.
//
// Solidity: function supplyFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperSession) SupplyFeeThreshold() (*big.Int, error) {
	return _Wingwrapper.Contract.SupplyFeeThreshold(&_Wingwrapper.CallOpts)
}

// SupplyFeeThreshold is a free data retrieval call binding the contract method 0x6115c522.
//
// Solidity: function supplyFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperCallerSession) SupplyFeeThreshold() (*big.Int, error) {
	return _Wingwrapper.Contract.SupplyFeeThreshold(&_Wingwrapper.CallOpts)
}

// UserAgentContract is a free data retrieval call binding the contract method 0x445adb1d.
//
// Solidity: function userAgentContract() view returns(bytes)
func (_Wingwrapper *WingwrapperCaller) UserAgentContract(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "userAgentContract")
	return *ret0, err
}

// UserAgentContract is a free data retrieval call binding the contract method 0x445adb1d.
//
// Solidity: function userAgentContract() view returns(bytes)
func (_Wingwrapper *WingwrapperSession) UserAgentContract() ([]byte, error) {
	return _Wingwrapper.Contract.UserAgentContract(&_Wingwrapper.CallOpts)
}

// UserAgentContract is a free data retrieval call binding the contract method 0x445adb1d.
//
// Solidity: function userAgentContract() view returns(bytes)
func (_Wingwrapper *WingwrapperCallerSession) UserAgentContract() ([]byte, error) {
	return _Wingwrapper.Contract.UserAgentContract(&_Wingwrapper.CallOpts)
}

// WithdrawFeeThreshold is a free data retrieval call binding the contract method 0x55f1f1ed.
//
// Solidity: function withdrawFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperCaller) WithdrawFeeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wingwrapper.contract.Call(opts, out, "withdrawFeeThreshold")
	return *ret0, err
}

// WithdrawFeeThreshold is a free data retrieval call binding the contract method 0x55f1f1ed.
//
// Solidity: function withdrawFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperSession) WithdrawFeeThreshold() (*big.Int, error) {
	return _Wingwrapper.Contract.WithdrawFeeThreshold(&_Wingwrapper.CallOpts)
}

// WithdrawFeeThreshold is a free data retrieval call binding the contract method 0x55f1f1ed.
//
// Solidity: function withdrawFeeThreshold() view returns(uint256)
func (_Wingwrapper *WingwrapperCallerSession) WithdrawFeeThreshold() (*big.Int, error) {
	return _Wingwrapper.Contract.WithdrawFeeThreshold(&_Wingwrapper.CallOpts)
}

// Supply is a paid mutator transaction binding the contract method 0x301bf698.
//
// Solidity: function supply(address token, uint256 amount, uint64 toChainId) payable returns()
func (_Wingwrapper *WingwrapperTransactor) Supply(opts *bind.TransactOpts, token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "supply", token, amount, toChainId)
}

// Supply is a paid mutator transaction binding the contract method 0x301bf698.
//
// Solidity: function supply(address token, uint256 amount, uint64 toChainId) payable returns()
func (_Wingwrapper *WingwrapperSession) Supply(token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.Supply(&_Wingwrapper.TransactOpts, token, amount, toChainId)
}

// Supply is a paid mutator transaction binding the contract method 0x301bf698.
//
// Solidity: function supply(address token, uint256 amount, uint64 toChainId) payable returns()
func (_Wingwrapper *WingwrapperTransactorSession) Supply(token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.Supply(&_Wingwrapper.TransactOpts, token, amount, toChainId)
}

// UpdateOntLockProxy is a paid mutator transaction binding the contract method 0x40d34aa0.
//
// Solidity: function updateOntLockProxy(bytes _ontProxy) returns()
func (_Wingwrapper *WingwrapperTransactor) UpdateOntLockProxy(opts *bind.TransactOpts, _ontProxy []byte) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "updateOntLockProxy", _ontProxy)
}

// UpdateOntLockProxy is a paid mutator transaction binding the contract method 0x40d34aa0.
//
// Solidity: function updateOntLockProxy(bytes _ontProxy) returns()
func (_Wingwrapper *WingwrapperSession) UpdateOntLockProxy(_ontProxy []byte) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdateOntLockProxy(&_Wingwrapper.TransactOpts, _ontProxy)
}

// UpdateOntLockProxy is a paid mutator transaction binding the contract method 0x40d34aa0.
//
// Solidity: function updateOntLockProxy(bytes _ontProxy) returns()
func (_Wingwrapper *WingwrapperTransactorSession) UpdateOntLockProxy(_ontProxy []byte) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdateOntLockProxy(&_Wingwrapper.TransactOpts, _ontProxy)
}

// UpdatePolyProxy is a paid mutator transaction binding the contract method 0xf572f9b0.
//
// Solidity: function updatePolyProxy(address newProxy) returns()
func (_Wingwrapper *WingwrapperTransactor) UpdatePolyProxy(opts *bind.TransactOpts, newProxy common.Address) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "updatePolyProxy", newProxy)
}

// UpdatePolyProxy is a paid mutator transaction binding the contract method 0xf572f9b0.
//
// Solidity: function updatePolyProxy(address newProxy) returns()
func (_Wingwrapper *WingwrapperSession) UpdatePolyProxy(newProxy common.Address) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdatePolyProxy(&_Wingwrapper.TransactOpts, newProxy)
}

// UpdatePolyProxy is a paid mutator transaction binding the contract method 0xf572f9b0.
//
// Solidity: function updatePolyProxy(address newProxy) returns()
func (_Wingwrapper *WingwrapperTransactorSession) UpdatePolyProxy(newProxy common.Address) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdatePolyProxy(&_Wingwrapper.TransactOpts, newProxy)
}

// UpdateRequestStatus is a paid mutator transaction binding the contract method 0x2b00754f.
//
// Solidity: function updateRequestStatus(address user, uint256 index, uint8 status) returns()
func (_Wingwrapper *WingwrapperTransactor) UpdateRequestStatus(opts *bind.TransactOpts, user common.Address, index *big.Int, status uint8) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "updateRequestStatus", user, index, status)
}

// UpdateRequestStatus is a paid mutator transaction binding the contract method 0x2b00754f.
//
// Solidity: function updateRequestStatus(address user, uint256 index, uint8 status) returns()
func (_Wingwrapper *WingwrapperSession) UpdateRequestStatus(user common.Address, index *big.Int, status uint8) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdateRequestStatus(&_Wingwrapper.TransactOpts, user, index, status)
}

// UpdateRequestStatus is a paid mutator transaction binding the contract method 0x2b00754f.
//
// Solidity: function updateRequestStatus(address user, uint256 index, uint8 status) returns()
func (_Wingwrapper *WingwrapperTransactorSession) UpdateRequestStatus(user common.Address, index *big.Int, status uint8) (*types.Transaction, error) {
	return _Wingwrapper.Contract.UpdateRequestStatus(&_Wingwrapper.TransactOpts, user, index, status)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8693a290.
//
// Solidity: function withdraw(address token, uint256 amount, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "withdraw", token, amount, toChainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8693a290.
//
// Solidity: function withdraw(address token, uint256 amount, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperSession) Withdraw(token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.Withdraw(&_Wingwrapper.TransactOpts, token, amount, toChainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8693a290.
//
// Solidity: function withdraw(address token, uint256 amount, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperTransactorSession) Withdraw(token common.Address, amount *big.Int, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.Withdraw(&_Wingwrapper.TransactOpts, token, amount, toChainId)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0xda15ae4e.
//
// Solidity: function withdrawWing(bytes targetAddress, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperTransactor) WithdrawWing(opts *bind.TransactOpts, targetAddress []byte, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.contract.Transact(opts, "withdrawWing", targetAddress, toChainId)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0xda15ae4e.
//
// Solidity: function withdrawWing(bytes targetAddress, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperSession) WithdrawWing(targetAddress []byte, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.WithdrawWing(&_Wingwrapper.TransactOpts, targetAddress, toChainId)
}

// WithdrawWing is a paid mutator transaction binding the contract method 0xda15ae4e.
//
// Solidity: function withdrawWing(bytes targetAddress, uint64 toChainId) returns()
func (_Wingwrapper *WingwrapperTransactorSession) WithdrawWing(targetAddress []byte, toChainId uint64) (*types.Transaction, error) {
	return _Wingwrapper.Contract.WithdrawWing(&_Wingwrapper.TransactOpts, targetAddress, toChainId)
}

// WingwrapperOntLockProxyUpdatedIterator is returned from FilterOntLockProxyUpdated and is used to iterate over the raw logs and unpacked data for OntLockProxyUpdated events raised by the Wingwrapper contract.
type WingwrapperOntLockProxyUpdatedIterator struct {
	Event *WingwrapperOntLockProxyUpdated // Event containing the contract specifics and raw log

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
func (it *WingwrapperOntLockProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperOntLockProxyUpdated)
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
		it.Event = new(WingwrapperOntLockProxyUpdated)
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
func (it *WingwrapperOntLockProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperOntLockProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperOntLockProxyUpdated represents a OntLockProxyUpdated event raised by the Wingwrapper contract.
type WingwrapperOntLockProxyUpdated struct {
	OldProxy []byte
	NewProxy []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOntLockProxyUpdated is a free log retrieval operation binding the contract event 0x91e64ed86813347dc94dfcd96757c0d1fb02d91a9fedc91ea9bbcdedbc08a576.
//
// Solidity: event OntLockProxyUpdated(bytes oldProxy, bytes newProxy)
func (_Wingwrapper *WingwrapperFilterer) FilterOntLockProxyUpdated(opts *bind.FilterOpts) (*WingwrapperOntLockProxyUpdatedIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "OntLockProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &WingwrapperOntLockProxyUpdatedIterator{contract: _Wingwrapper.contract, event: "OntLockProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchOntLockProxyUpdated is a free log subscription operation binding the contract event 0x91e64ed86813347dc94dfcd96757c0d1fb02d91a9fedc91ea9bbcdedbc08a576.
//
// Solidity: event OntLockProxyUpdated(bytes oldProxy, bytes newProxy)
func (_Wingwrapper *WingwrapperFilterer) WatchOntLockProxyUpdated(opts *bind.WatchOpts, sink chan<- *WingwrapperOntLockProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "OntLockProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperOntLockProxyUpdated)
				if err := _Wingwrapper.contract.UnpackLog(event, "OntLockProxyUpdated", log); err != nil {
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

// ParseOntLockProxyUpdated is a log parse operation binding the contract event 0x91e64ed86813347dc94dfcd96757c0d1fb02d91a9fedc91ea9bbcdedbc08a576.
//
// Solidity: event OntLockProxyUpdated(bytes oldProxy, bytes newProxy)
func (_Wingwrapper *WingwrapperFilterer) ParseOntLockProxyUpdated(log types.Log) (*WingwrapperOntLockProxyUpdated, error) {
	event := new(WingwrapperOntLockProxyUpdated)
	if err := _Wingwrapper.contract.UnpackLog(event, "OntLockProxyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperPolyLockProxyUpdatedIterator is returned from FilterPolyLockProxyUpdated and is used to iterate over the raw logs and unpacked data for PolyLockProxyUpdated events raised by the Wingwrapper contract.
type WingwrapperPolyLockProxyUpdatedIterator struct {
	Event *WingwrapperPolyLockProxyUpdated // Event containing the contract specifics and raw log

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
func (it *WingwrapperPolyLockProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperPolyLockProxyUpdated)
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
		it.Event = new(WingwrapperPolyLockProxyUpdated)
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
func (it *WingwrapperPolyLockProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperPolyLockProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperPolyLockProxyUpdated represents a PolyLockProxyUpdated event raised by the Wingwrapper contract.
type WingwrapperPolyLockProxyUpdated struct {
	OldProxy common.Address
	NewProxy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPolyLockProxyUpdated is a free log retrieval operation binding the contract event 0x40bf952ebb271dccb7dd1b9244d2db50205559185e04bcc10bad51308cc62a4d.
//
// Solidity: event PolyLockProxyUpdated(address oldProxy, address newProxy)
func (_Wingwrapper *WingwrapperFilterer) FilterPolyLockProxyUpdated(opts *bind.FilterOpts) (*WingwrapperPolyLockProxyUpdatedIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "PolyLockProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &WingwrapperPolyLockProxyUpdatedIterator{contract: _Wingwrapper.contract, event: "PolyLockProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchPolyLockProxyUpdated is a free log subscription operation binding the contract event 0x40bf952ebb271dccb7dd1b9244d2db50205559185e04bcc10bad51308cc62a4d.
//
// Solidity: event PolyLockProxyUpdated(address oldProxy, address newProxy)
func (_Wingwrapper *WingwrapperFilterer) WatchPolyLockProxyUpdated(opts *bind.WatchOpts, sink chan<- *WingwrapperPolyLockProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "PolyLockProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperPolyLockProxyUpdated)
				if err := _Wingwrapper.contract.UnpackLog(event, "PolyLockProxyUpdated", log); err != nil {
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

// ParsePolyLockProxyUpdated is a log parse operation binding the contract event 0x40bf952ebb271dccb7dd1b9244d2db50205559185e04bcc10bad51308cc62a4d.
//
// Solidity: event PolyLockProxyUpdated(address oldProxy, address newProxy)
func (_Wingwrapper *WingwrapperFilterer) ParsePolyLockProxyUpdated(log types.Log) (*WingwrapperPolyLockProxyUpdated, error) {
	event := new(WingwrapperPolyLockProxyUpdated)
	if err := _Wingwrapper.contract.UnpackLog(event, "PolyLockProxyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperRequestUpdatedIterator is returned from FilterRequestUpdated and is used to iterate over the raw logs and unpacked data for RequestUpdated events raised by the Wingwrapper contract.
type WingwrapperRequestUpdatedIterator struct {
	Event *WingwrapperRequestUpdated // Event containing the contract specifics and raw log

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
func (it *WingwrapperRequestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperRequestUpdated)
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
		it.Event = new(WingwrapperRequestUpdated)
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
func (it *WingwrapperRequestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperRequestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperRequestUpdated represents a RequestUpdated event raised by the Wingwrapper contract.
type WingwrapperRequestUpdated struct {
	User          common.Address
	RequestIndex  *big.Int
	OeprationType uint8
	Status        uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestUpdated is a free log retrieval operation binding the contract event 0xcfe2e12491510e84790328447af892f6fcfc8d3ab0ca63b02b8e98f4b637b8de.
//
// Solidity: event RequestUpdated(address user, uint256 requestIndex, uint8 oeprationType, uint8 status)
func (_Wingwrapper *WingwrapperFilterer) FilterRequestUpdated(opts *bind.FilterOpts) (*WingwrapperRequestUpdatedIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return &WingwrapperRequestUpdatedIterator{contract: _Wingwrapper.contract, event: "RequestUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestUpdated is a free log subscription operation binding the contract event 0xcfe2e12491510e84790328447af892f6fcfc8d3ab0ca63b02b8e98f4b637b8de.
//
// Solidity: event RequestUpdated(address user, uint256 requestIndex, uint8 oeprationType, uint8 status)
func (_Wingwrapper *WingwrapperFilterer) WatchRequestUpdated(opts *bind.WatchOpts, sink chan<- *WingwrapperRequestUpdated) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperRequestUpdated)
				if err := _Wingwrapper.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
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

// ParseRequestUpdated is a log parse operation binding the contract event 0xcfe2e12491510e84790328447af892f6fcfc8d3ab0ca63b02b8e98f4b637b8de.
//
// Solidity: event RequestUpdated(address user, uint256 requestIndex, uint8 oeprationType, uint8 status)
func (_Wingwrapper *WingwrapperFilterer) ParseRequestUpdated(log types.Log) (*WingwrapperRequestUpdated, error) {
	event := new(WingwrapperRequestUpdated)
	if err := _Wingwrapper.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Wingwrapper contract.
type WingwrapperSupplyIterator struct {
	Event *WingwrapperSupply // Event containing the contract specifics and raw log

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
func (it *WingwrapperSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperSupply)
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
		it.Event = new(WingwrapperSupply)
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
func (it *WingwrapperSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperSupply represents a Supply event raised by the Wingwrapper contract.
type WingwrapperSupply struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) FilterSupply(opts *bind.FilterOpts) (*WingwrapperSupplyIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &WingwrapperSupplyIterator{contract: _Wingwrapper.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *WingwrapperSupply) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperSupply)
				if err := _Wingwrapper.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) ParseSupply(log types.Log) (*WingwrapperSupply, error) {
	event := new(WingwrapperSupply)
	if err := _Wingwrapper.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperSupplyFeeThresholdUpdatedIterator is returned from FilterSupplyFeeThresholdUpdated and is used to iterate over the raw logs and unpacked data for SupplyFeeThresholdUpdated events raised by the Wingwrapper contract.
type WingwrapperSupplyFeeThresholdUpdatedIterator struct {
	Event *WingwrapperSupplyFeeThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *WingwrapperSupplyFeeThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperSupplyFeeThresholdUpdated)
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
		it.Event = new(WingwrapperSupplyFeeThresholdUpdated)
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
func (it *WingwrapperSupplyFeeThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperSupplyFeeThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperSupplyFeeThresholdUpdated represents a SupplyFeeThresholdUpdated event raised by the Wingwrapper contract.
type WingwrapperSupplyFeeThresholdUpdated struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupplyFeeThresholdUpdated is a free log retrieval operation binding the contract event 0x3532e02de9efe8317bc516ef05db8cfc5bcc7c27974ab56035589dd82746f305.
//
// Solidity: event SupplyFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) FilterSupplyFeeThresholdUpdated(opts *bind.FilterOpts) (*WingwrapperSupplyFeeThresholdUpdatedIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "SupplyFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &WingwrapperSupplyFeeThresholdUpdatedIterator{contract: _Wingwrapper.contract, event: "SupplyFeeThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSupplyFeeThresholdUpdated is a free log subscription operation binding the contract event 0x3532e02de9efe8317bc516ef05db8cfc5bcc7c27974ab56035589dd82746f305.
//
// Solidity: event SupplyFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) WatchSupplyFeeThresholdUpdated(opts *bind.WatchOpts, sink chan<- *WingwrapperSupplyFeeThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "SupplyFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperSupplyFeeThresholdUpdated)
				if err := _Wingwrapper.contract.UnpackLog(event, "SupplyFeeThresholdUpdated", log); err != nil {
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

// ParseSupplyFeeThresholdUpdated is a log parse operation binding the contract event 0x3532e02de9efe8317bc516ef05db8cfc5bcc7c27974ab56035589dd82746f305.
//
// Solidity: event SupplyFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) ParseSupplyFeeThresholdUpdated(log types.Log) (*WingwrapperSupplyFeeThresholdUpdated, error) {
	event := new(WingwrapperSupplyFeeThresholdUpdated)
	if err := _Wingwrapper.contract.UnpackLog(event, "SupplyFeeThresholdUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Wingwrapper contract.
type WingwrapperWithdrawIterator struct {
	Event *WingwrapperWithdraw // Event containing the contract specifics and raw log

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
func (it *WingwrapperWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperWithdraw)
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
		it.Event = new(WingwrapperWithdraw)
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
func (it *WingwrapperWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperWithdraw represents a Withdraw event raised by the Wingwrapper contract.
type WingwrapperWithdraw struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) FilterWithdraw(opts *bind.FilterOpts) (*WingwrapperWithdrawIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &WingwrapperWithdrawIterator{contract: _Wingwrapper.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WingwrapperWithdraw) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperWithdraw)
				if err := _Wingwrapper.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address user, address token, uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) ParseWithdraw(log types.Log) (*WingwrapperWithdraw, error) {
	event := new(WingwrapperWithdraw)
	if err := _Wingwrapper.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperWithdrawFeeIterator is returned from FilterWithdrawFee and is used to iterate over the raw logs and unpacked data for WithdrawFee events raised by the Wingwrapper contract.
type WingwrapperWithdrawFeeIterator struct {
	Event *WingwrapperWithdrawFee // Event containing the contract specifics and raw log

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
func (it *WingwrapperWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperWithdrawFee)
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
		it.Event = new(WingwrapperWithdrawFee)
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
func (it *WingwrapperWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperWithdrawFee represents a WithdrawFee event raised by the Wingwrapper contract.
type WingwrapperWithdrawFee struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFee is a free log retrieval operation binding the contract event 0xa01cb43de193eb3a80b373fb949c09d0eedb01f39f3b6063ace0cb6b067cc123.
//
// Solidity: event WithdrawFee(uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) FilterWithdrawFee(opts *bind.FilterOpts) (*WingwrapperWithdrawFeeIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "WithdrawFee")
	if err != nil {
		return nil, err
	}
	return &WingwrapperWithdrawFeeIterator{contract: _Wingwrapper.contract, event: "WithdrawFee", logs: logs, sub: sub}, nil
}

// WatchWithdrawFee is a free log subscription operation binding the contract event 0xa01cb43de193eb3a80b373fb949c09d0eedb01f39f3b6063ace0cb6b067cc123.
//
// Solidity: event WithdrawFee(uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) WatchWithdrawFee(opts *bind.WatchOpts, sink chan<- *WingwrapperWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "WithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperWithdrawFee)
				if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
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

// ParseWithdrawFee is a log parse operation binding the contract event 0xa01cb43de193eb3a80b373fb949c09d0eedb01f39f3b6063ace0cb6b067cc123.
//
// Solidity: event WithdrawFee(uint256 amount)
func (_Wingwrapper *WingwrapperFilterer) ParseWithdrawFee(log types.Log) (*WingwrapperWithdrawFee, error) {
	event := new(WingwrapperWithdrawFee)
	if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperWithdrawFeeThresholdUpdatedIterator is returned from FilterWithdrawFeeThresholdUpdated and is used to iterate over the raw logs and unpacked data for WithdrawFeeThresholdUpdated events raised by the Wingwrapper contract.
type WingwrapperWithdrawFeeThresholdUpdatedIterator struct {
	Event *WingwrapperWithdrawFeeThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *WingwrapperWithdrawFeeThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperWithdrawFeeThresholdUpdated)
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
		it.Event = new(WingwrapperWithdrawFeeThresholdUpdated)
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
func (it *WingwrapperWithdrawFeeThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperWithdrawFeeThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperWithdrawFeeThresholdUpdated represents a WithdrawFeeThresholdUpdated event raised by the Wingwrapper contract.
type WingwrapperWithdrawFeeThresholdUpdated struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFeeThresholdUpdated is a free log retrieval operation binding the contract event 0x248dfe484b859ea95b0ea8847a68c2865b86feaf8a5c3ad456afea9a647093fe.
//
// Solidity: event WithdrawFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) FilterWithdrawFeeThresholdUpdated(opts *bind.FilterOpts) (*WingwrapperWithdrawFeeThresholdUpdatedIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "WithdrawFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &WingwrapperWithdrawFeeThresholdUpdatedIterator{contract: _Wingwrapper.contract, event: "WithdrawFeeThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawFeeThresholdUpdated is a free log subscription operation binding the contract event 0x248dfe484b859ea95b0ea8847a68c2865b86feaf8a5c3ad456afea9a647093fe.
//
// Solidity: event WithdrawFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) WatchWithdrawFeeThresholdUpdated(opts *bind.WatchOpts, sink chan<- *WingwrapperWithdrawFeeThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "WithdrawFeeThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperWithdrawFeeThresholdUpdated)
				if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawFeeThresholdUpdated", log); err != nil {
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

// ParseWithdrawFeeThresholdUpdated is a log parse operation binding the contract event 0x248dfe484b859ea95b0ea8847a68c2865b86feaf8a5c3ad456afea9a647093fe.
//
// Solidity: event WithdrawFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (_Wingwrapper *WingwrapperFilterer) ParseWithdrawFeeThresholdUpdated(log types.Log) (*WingwrapperWithdrawFeeThresholdUpdated, error) {
	event := new(WingwrapperWithdrawFeeThresholdUpdated)
	if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawFeeThresholdUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WingwrapperWithdrawWingIterator is returned from FilterWithdrawWing and is used to iterate over the raw logs and unpacked data for WithdrawWing events raised by the Wingwrapper contract.
type WingwrapperWithdrawWingIterator struct {
	Event *WingwrapperWithdrawWing // Event containing the contract specifics and raw log

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
func (it *WingwrapperWithdrawWingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WingwrapperWithdrawWing)
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
		it.Event = new(WingwrapperWithdrawWing)
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
func (it *WingwrapperWithdrawWingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WingwrapperWithdrawWingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WingwrapperWithdrawWing represents a WithdrawWing event raised by the Wingwrapper contract.
type WingwrapperWithdrawWing struct {
	User   common.Address
	Target []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWing is a free log retrieval operation binding the contract event 0x91ac22144848c281e728c9336aa88e73dfde7c9829514adc019418766ec9dfa3.
//
// Solidity: event WithdrawWing(address user, bytes target)
func (_Wingwrapper *WingwrapperFilterer) FilterWithdrawWing(opts *bind.FilterOpts) (*WingwrapperWithdrawWingIterator, error) {

	logs, sub, err := _Wingwrapper.contract.FilterLogs(opts, "WithdrawWing")
	if err != nil {
		return nil, err
	}
	return &WingwrapperWithdrawWingIterator{contract: _Wingwrapper.contract, event: "WithdrawWing", logs: logs, sub: sub}, nil
}

// WatchWithdrawWing is a free log subscription operation binding the contract event 0x91ac22144848c281e728c9336aa88e73dfde7c9829514adc019418766ec9dfa3.
//
// Solidity: event WithdrawWing(address user, bytes target)
func (_Wingwrapper *WingwrapperFilterer) WatchWithdrawWing(opts *bind.WatchOpts, sink chan<- *WingwrapperWithdrawWing) (event.Subscription, error) {

	logs, sub, err := _Wingwrapper.contract.WatchLogs(opts, "WithdrawWing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WingwrapperWithdrawWing)
				if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawWing", log); err != nil {
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

// ParseWithdrawWing is a log parse operation binding the contract event 0x91ac22144848c281e728c9336aa88e73dfde7c9829514adc019418766ec9dfa3.
//
// Solidity: event WithdrawWing(address user, bytes target)
func (_Wingwrapper *WingwrapperFilterer) ParseWithdrawWing(log types.Log) (*WingwrapperWithdrawWing, error) {
	event := new(WingwrapperWithdrawWing)
	if err := _Wingwrapper.contract.UnpackLog(event, "WithdrawWing", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582032e12bb45cb5808e573af73f83959386bd9a09af5acc1f7a7e0a0fdadebd90c80029"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820e74a59cca5bfdf7acd52645f205feda980a3e02044f5d2687f3c81428497bc9b0029"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}

