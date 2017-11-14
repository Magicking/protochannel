// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"listed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Del\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DelWhitelist\",\"type\":\"event\"}]"

// WhitelistBin is the compiled bytecode used for deploying new contracts.
const WhitelistBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556102968061003b6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166364138230811461006857806387dc5eec1461009b5780638da5cb5b146100bc578063e9d6f57f146100eb578063f794062e1461010a57600080fd5b341561007357600080fd5b610087600160a060020a0360043516610129565b604051901515815260200160405180910390f35b34156100a657600080fd5b6100ba600160a060020a036004351661013e565b005b34156100c757600080fd5b6100cf6101c0565b604051600160a060020a03909116815260200160405180910390f35b34156100f657600080fd5b6100ba600160a060020a03600435166101cf565b341561011557600080fd5b610087600160a060020a036004351661024c565b60016020526000908152604090205460ff1681565b60005433600160a060020a0390811691161461015957600080fd5b600160a060020a038116600090815260016020819052604091829020805460ff191690911790557fe463fa6bdecb16f96f58191d902152633214e760ea443684105a7eef1ad16b9d90829051600160a060020a03909116815260200160405180910390a150565b600054600160a060020a031681565b60005433600160a060020a039081169116146101ea57600080fd5b600160a060020a03811660009081526001602052604090819020805460ff191690557f0cdc737c1f273caa5beeea96ca9136ce2fe05b1388871b463a16688b2d5e2d0d90829051600160a060020a03909116815260200160405180910390a150565b600160a060020a031660009081526001602052604090205460ff16905600a165627a7a723058208252a841519b87fa02791ad0a82674ef1c635a4adc6b883815c49a594e78e2e70029`

// DeployWhitelist deploys a new Ethereum contract, binding an instance of Whitelist to it.
func DeployWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Whitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}}, nil
}

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// IsListed is a free data retrieval call binding the contract method 0xf794062e.
//
// Solidity: function isListed(addr address) constant returns(bool)
func (_Whitelist *WhitelistCaller) IsListed(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "isListed", addr)
	return *ret0, err
}

// IsListed is a free data retrieval call binding the contract method 0xf794062e.
//
// Solidity: function isListed(addr address) constant returns(bool)
func (_Whitelist *WhitelistSession) IsListed(addr common.Address) (bool, error) {
	return _Whitelist.Contract.IsListed(&_Whitelist.CallOpts, addr)
}

// IsListed is a free data retrieval call binding the contract method 0xf794062e.
//
// Solidity: function isListed(addr address) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) IsListed(addr common.Address) (bool, error) {
	return _Whitelist.Contract.IsListed(&_Whitelist.CallOpts, addr)
}

// Listed is a free data retrieval call binding the contract method 0x64138230.
//
// Solidity: function listed( address) constant returns(bool)
func (_Whitelist *WhitelistCaller) Listed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "listed", arg0)
	return *ret0, err
}

// Listed is a free data retrieval call binding the contract method 0x64138230.
//
// Solidity: function listed( address) constant returns(bool)
func (_Whitelist *WhitelistSession) Listed(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Listed(&_Whitelist.CallOpts, arg0)
}

// Listed is a free data retrieval call binding the contract method 0x64138230.
//
// Solidity: function listed( address) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) Listed(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Listed(&_Whitelist.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistCallerSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x87dc5eec.
//
// Solidity: function Add(addr address) returns()
func (_Whitelist *WhitelistTransactor) Add(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "Add", addr)
}

// Add is a paid mutator transaction binding the contract method 0x87dc5eec.
//
// Solidity: function Add(addr address) returns()
func (_Whitelist *WhitelistSession) Add(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Add(&_Whitelist.TransactOpts, addr)
}

// Add is a paid mutator transaction binding the contract method 0x87dc5eec.
//
// Solidity: function Add(addr address) returns()
func (_Whitelist *WhitelistTransactorSession) Add(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Add(&_Whitelist.TransactOpts, addr)
}

// Del is a paid mutator transaction binding the contract method 0xe9d6f57f.
//
// Solidity: function Del(addr address) returns()
func (_Whitelist *WhitelistTransactor) Del(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "Del", addr)
}

// Del is a paid mutator transaction binding the contract method 0xe9d6f57f.
//
// Solidity: function Del(addr address) returns()
func (_Whitelist *WhitelistSession) Del(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Del(&_Whitelist.TransactOpts, addr)
}

// Del is a paid mutator transaction binding the contract method 0xe9d6f57f.
//
// Solidity: function Del(addr address) returns()
func (_Whitelist *WhitelistTransactorSession) Del(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Del(&_Whitelist.TransactOpts, addr)
}
