// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TicTacToeABI is the input ABI used to generate the binding from.
const TicTacToeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"CreateBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"can_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"CheckAndApply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Judge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"op_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"GetPlayerMark\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Apply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"other\",\"type\":\"address\"}],\"name\":\"SetPlayer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"player2\",\"type\":\"address\"}],\"name\":\"NewGame\",\"type\":\"event\"}]"

// TicTacToeBin is the compiled bytecode used for deploying new contracts.
const TicTacToeBin = `0x60606040526004600255341561001457600080fd5b60038054600160a060020a033316600160a060020a0319918216178255600182905560009190915560048054821673533a245f03a1a46cacb933a3beef752fd8ff45c3179055600580549091167348d24715fe9f7daa90286141f9b17e184b5a148b179055610d4a806100886000396000f3006060604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631625059b81146100c95780631de21b761461015357806330332aea14610182578063355019db1461019557806338596d97146101fa5780633aa805bc1461025b578063712d215c1461028057806385893cb8146102935780638739da95146102a6578063df6a9805146102fe578063f2db77791461033a578063f39e9e751461039b578063f6bdd7f1146103f7575b600080fd5b34156100d457600080fd5b6100dc61040d565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610118578082015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015e57600080fd5b610166610642565b604051600160a060020a03909116815260200160405180910390f35b341561018d57600080fd5b610166610651565b34156101a057600080fd5b6101e660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061066095505050505050565b604051901515815260200160405180910390f35b341561020557600080fd5b6100dc60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff85351694602081013594506040013592506106ed915050565b341561026657600080fd5b61026e61072a565b60405190815260200160405180910390f35b341561028b57600080fd5b61026e610730565b341561029e57600080fd5b610166610736565b34156102b157600080fd5b6100dc60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505084359460200135935061074592505050565b341561030957600080fd5b61031d600160a060020a036004351661082b565b604051600160f860020a0319909116815260200160405180910390f35b341561034557600080fd5b6101e660046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250610895915050565b34156103a657600080fd5b6100dc60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506109e7915050565b61040b600160a060020a0360043516610ab6565b005b610415610d0c565b61041d610d0c565b600d60405180591061042c5750595b818152601f19601f830116810160200160405290509050600081818151811061045157fe5b906020010190600160f860020a031916908160001a90535060008160018151811061047857fe5b906020010190600160f860020a031916908160001a90535060008160028151811061049f57fe5b906020010190600160f860020a031916908160001a9053506000816003815181106104c657fe5b906020010190600160f860020a031916908160001a9053506000816004815181106104ed57fe5b906020010190600160f860020a031916908160001a90535060008160058151811061051457fe5b906020010190600160f860020a031916908160001a90535060008160068151811061053b57fe5b906020010190600160f860020a031916908160001a90535060008160078151811061056257fe5b906020010190600160f860020a031916908160001a90535060008160088151811061058957fe5b906020010190600160f860020a031916908160001a9053506000816009815181106105b057fe5b906020010190600160f860020a031916908160001a905350600081600a815181106105d757fe5b906020010190600160f860020a031916908160001a905350600081600b815181106105fe57fe5b906020010190600160f860020a031916908160001a905350600081600c8151811061062557fe5b906020010190600160f860020a031916908160001a905350919050565b600454600160a060020a031681565b600554600160a060020a031681565b60008060008360028151811061067257fe5b016020015160f860020a900460f860020a0260f860020a900491508360038151811061069a57fe5b016020015160f860020a900460f860020a0260f860020a90049050838282600054026002540101815181106106cb57fe5b016020015160f860020a9081900402600160f860020a03191615949350505050565b6106f5610d0c565b61070185858585610895565b151561070c57600080fd5b6107218561071c87878787610ba3565b6109e7565b95945050505050565b60015481565b60005481565b600354600160a060020a031681565b61074d610d0c565b60f860020a8460008151811061075f57fe5b906020010190600160f860020a031916908160001a9053508360018151811061078457fe5b016020015160f860020a900460f860020a0260f860020a900460010160f860020a02846001815181106107b357fe5b906020010190600160f860020a031916908160001a9053508260f860020a02846002815181106107df57fe5b906020010190600160f860020a031916908160001a9053508160f860020a028460038151811061080b57fe5b906020010190600160f860020a031916908160001a905350929392505050565b600454600090600160a060020a038381169116141561084f575060f860020a610890565b600554600160a060020a038381169116141561088c57507f0200000000000000000000000000000000000000000000000000000000000000610890565b5060005b919050565b6000806000806108a788888888610ba3565b92506002548851116108bc57600093506109dc565b876000815181106108c957fe5b016020015160f860020a900460f860020a029150876001815181106108ea57fe5b016020015160f860020a908190048102049050600160f860020a03198216151561092657600354600160a060020a0384811691161493506109dc565b600454600160a060020a038481169116148015906109525750600554600160a060020a03848116911614155b1561096057600093506109dc565b600454600160a060020a03848116911614801561097e575060018116155b806109a35750600554600160a060020a0384811691161480156109a357506001808216145b156109b157600093506109dc565b60f860020a600160f860020a0319831614156109d7576109d088610660565b93506109dc565b600093505b505050949350505050565b6109ef610d0c565b600080600085600081518110610a0157fe5b016020015160f860020a9081900481029350600160f860020a031984161415610aac5785600281518110610a3157fe5b016020015160f860020a900460f860020a0260f860020a9004915085600381518110610a5957fe5b016020015160f860020a900460f860020a0260f860020a90049050610a7d8561082b565b86838360005402600254010181518110610a9357fe5b906020010190600160f860020a031916908160001a9053505b5093949350505050565b600454600160a060020a0316158015610ad85750600554600160a060020a0316155b1515610ae357600080fd5b33600160a060020a031681600160a060020a031614151515610b0457600080fd5b600160a060020a0381161515610b1957600080fd5b60048054600160a060020a0333811673ffffffffffffffffffffffffffffffffffffffff19928316179283905560058054858316931692909217918290557fd1af9b1127c37da1f8a3d2e4e7d92bcbfa7bfa2ccf9ce1d29530052e366cc9019281169116604051600160a060020a039283168152911660208201526040908101905180910390a150565b6000610bad610d0c565b60006040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a313300000000815250915081876040518083805190602001908083835b60208310610c185780518252601f199092019160209182019101610bf9565b6001836020036101000a038019825116818451161790925250505091909101905082805190602001908083835b60208310610c645780518252601f199092019160209182019101610c45565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051809103902090506001818787876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610cf857600080fd5b505060206040510351979650505050505050565b602060405190810160405260008152905600a165627a7a72305820484bee41fc5181557bd97290d5852ae8ddf499ed8382be07495278f3cbf45f710029`

// DeployTicTacToe deploys a new Ethereum contract, binding an instance of TicTacToe to it.
func DeployTicTacToe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TicTacToe, error) {
	parsed, err := abi.JSON(strings.NewReader(TicTacToeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TicTacToeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TicTacToe{TicTacToeCaller: TicTacToeCaller{contract: contract}, TicTacToeTransactor: TicTacToeTransactor{contract: contract}}, nil
}

// TicTacToe is an auto generated Go binding around an Ethereum contract.
type TicTacToe struct {
	TicTacToeCaller     // Read-only binding to the contract
	TicTacToeTransactor // Write-only binding to the contract
}

// TicTacToeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TicTacToeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicTacToeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TicTacToeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicTacToeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TicTacToeSession struct {
	Contract     *TicTacToe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TicTacToeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TicTacToeCallerSession struct {
	Contract *TicTacToeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TicTacToeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TicTacToeTransactorSession struct {
	Contract     *TicTacToeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TicTacToeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TicTacToeRaw struct {
	Contract *TicTacToe // Generic contract binding to access the raw methods on
}

// TicTacToeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TicTacToeCallerRaw struct {
	Contract *TicTacToeCaller // Generic read-only contract binding to access the raw methods on
}

// TicTacToeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TicTacToeTransactorRaw struct {
	Contract *TicTacToeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTicTacToe creates a new instance of TicTacToe, bound to a specific deployed contract.
func NewTicTacToe(address common.Address, backend bind.ContractBackend) (*TicTacToe, error) {
	contract, err := bindTicTacToe(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TicTacToe{TicTacToeCaller: TicTacToeCaller{contract: contract}, TicTacToeTransactor: TicTacToeTransactor{contract: contract}}, nil
}

// NewTicTacToeCaller creates a new read-only instance of TicTacToe, bound to a specific deployed contract.
func NewTicTacToeCaller(address common.Address, caller bind.ContractCaller) (*TicTacToeCaller, error) {
	contract, err := bindTicTacToe(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TicTacToeCaller{contract: contract}, nil
}

// NewTicTacToeTransactor creates a new write-only instance of TicTacToe, bound to a specific deployed contract.
func NewTicTacToeTransactor(address common.Address, transactor bind.ContractTransactor) (*TicTacToeTransactor, error) {
	contract, err := bindTicTacToe(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TicTacToeTransactor{contract: contract}, nil
}

// bindTicTacToe binds a generic wrapper to an already deployed contract.
func bindTicTacToe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TicTacToeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicTacToe *TicTacToeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TicTacToe.Contract.TicTacToeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicTacToe *TicTacToeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicTacToe.Contract.TicTacToeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicTacToe *TicTacToeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicTacToe.Contract.TicTacToeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicTacToe *TicTacToeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TicTacToe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicTacToe *TicTacToeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicTacToe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicTacToe *TicTacToeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicTacToe.Contract.contract.Transact(opts, method, params...)
}

// Apply is a free data retrieval call binding the contract method 0xf39e9e75.
//
// Solidity: function Apply(state bytes, player address) constant returns(bytes)
func (_TicTacToe *TicTacToeCaller) Apply(opts *bind.CallOpts, state []byte, player common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "Apply", state, player)
	return *ret0, err
}

// Apply is a free data retrieval call binding the contract method 0xf39e9e75.
//
// Solidity: function Apply(state bytes, player address) constant returns(bytes)
func (_TicTacToe *TicTacToeSession) Apply(state []byte, player common.Address) ([]byte, error) {
	return _TicTacToe.Contract.Apply(&_TicTacToe.CallOpts, state, player)
}

// Apply is a free data retrieval call binding the contract method 0xf39e9e75.
//
// Solidity: function Apply(state bytes, player address) constant returns(bytes)
func (_TicTacToe *TicTacToeCallerSession) Apply(state []byte, player common.Address) ([]byte, error) {
	return _TicTacToe.Contract.Apply(&_TicTacToe.CallOpts, state, player)
}

// CheckAndApply is a free data retrieval call binding the contract method 0x38596d97.
//
// Solidity: function CheckAndApply(state bytes, v uint8, r bytes32, s bytes32) constant returns(bytes)
func (_TicTacToe *TicTacToeCaller) CheckAndApply(opts *bind.CallOpts, state []byte, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "CheckAndApply", state, v, r, s)
	return *ret0, err
}

// CheckAndApply is a free data retrieval call binding the contract method 0x38596d97.
//
// Solidity: function CheckAndApply(state bytes, v uint8, r bytes32, s bytes32) constant returns(bytes)
func (_TicTacToe *TicTacToeSession) CheckAndApply(state []byte, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return _TicTacToe.Contract.CheckAndApply(&_TicTacToe.CallOpts, state, v, r, s)
}

// CheckAndApply is a free data retrieval call binding the contract method 0x38596d97.
//
// Solidity: function CheckAndApply(state bytes, v uint8, r bytes32, s bytes32) constant returns(bytes)
func (_TicTacToe *TicTacToeCallerSession) CheckAndApply(state []byte, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return _TicTacToe.Contract.CheckAndApply(&_TicTacToe.CallOpts, state, v, r, s)
}

// CreateBoard is a free data retrieval call binding the contract method 0x1625059b.
//
// Solidity: function CreateBoard() constant returns(bytes)
func (_TicTacToe *TicTacToeCaller) CreateBoard(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "CreateBoard")
	return *ret0, err
}

// CreateBoard is a free data retrieval call binding the contract method 0x1625059b.
//
// Solidity: function CreateBoard() constant returns(bytes)
func (_TicTacToe *TicTacToeSession) CreateBoard() ([]byte, error) {
	return _TicTacToe.Contract.CreateBoard(&_TicTacToe.CallOpts)
}

// CreateBoard is a free data retrieval call binding the contract method 0x1625059b.
//
// Solidity: function CreateBoard() constant returns(bytes)
func (_TicTacToe *TicTacToeCallerSession) CreateBoard() ([]byte, error) {
	return _TicTacToe.Contract.CreateBoard(&_TicTacToe.CallOpts)
}

// GetPlayerMark is a free data retrieval call binding the contract method 0xdf6a9805.
//
// Solidity: function GetPlayerMark(player address) constant returns(bytes1)
func (_TicTacToe *TicTacToeCaller) GetPlayerMark(opts *bind.CallOpts, player common.Address) ([1]byte, error) {
	var (
		ret0 = new([1]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "GetPlayerMark", player)
	return *ret0, err
}

// GetPlayerMark is a free data retrieval call binding the contract method 0xdf6a9805.
//
// Solidity: function GetPlayerMark(player address) constant returns(bytes1)
func (_TicTacToe *TicTacToeSession) GetPlayerMark(player common.Address) ([1]byte, error) {
	return _TicTacToe.Contract.GetPlayerMark(&_TicTacToe.CallOpts, player)
}

// GetPlayerMark is a free data retrieval call binding the contract method 0xdf6a9805.
//
// Solidity: function GetPlayerMark(player address) constant returns(bytes1)
func (_TicTacToe *TicTacToeCallerSession) GetPlayerMark(player common.Address) ([1]byte, error) {
	return _TicTacToe.Contract.GetPlayerMark(&_TicTacToe.CallOpts, player)
}

// Judge is a free data retrieval call binding the contract method 0x85893cb8.
//
// Solidity: function Judge() constant returns(address)
func (_TicTacToe *TicTacToeCaller) Judge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "Judge")
	return *ret0, err
}

// Judge is a free data retrieval call binding the contract method 0x85893cb8.
//
// Solidity: function Judge() constant returns(address)
func (_TicTacToe *TicTacToeSession) Judge() (common.Address, error) {
	return _TicTacToe.Contract.Judge(&_TicTacToe.CallOpts)
}

// Judge is a free data retrieval call binding the contract method 0x85893cb8.
//
// Solidity: function Judge() constant returns(address)
func (_TicTacToe *TicTacToeCallerSession) Judge() (common.Address, error) {
	return _TicTacToe.Contract.Judge(&_TicTacToe.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0x1de21b76.
//
// Solidity: function P1() constant returns(address)
func (_TicTacToe *TicTacToeCaller) P1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "P1")
	return *ret0, err
}

// P1 is a free data retrieval call binding the contract method 0x1de21b76.
//
// Solidity: function P1() constant returns(address)
func (_TicTacToe *TicTacToeSession) P1() (common.Address, error) {
	return _TicTacToe.Contract.P1(&_TicTacToe.CallOpts)
}

// P1 is a free data retrieval call binding the contract method 0x1de21b76.
//
// Solidity: function P1() constant returns(address)
func (_TicTacToe *TicTacToeCallerSession) P1() (common.Address, error) {
	return _TicTacToe.Contract.P1(&_TicTacToe.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x30332aea.
//
// Solidity: function P2() constant returns(address)
func (_TicTacToe *TicTacToeCaller) P2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "P2")
	return *ret0, err
}

// P2 is a free data retrieval call binding the contract method 0x30332aea.
//
// Solidity: function P2() constant returns(address)
func (_TicTacToe *TicTacToeSession) P2() (common.Address, error) {
	return _TicTacToe.Contract.P2(&_TicTacToe.CallOpts)
}

// P2 is a free data retrieval call binding the contract method 0x30332aea.
//
// Solidity: function P2() constant returns(address)
func (_TicTacToe *TicTacToeCallerSession) P2() (common.Address, error) {
	return _TicTacToe.Contract.P2(&_TicTacToe.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xf2db7779.
//
// Solidity: function Verify(state bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) Verify(opts *bind.CallOpts, state []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "Verify", state, v, r, s)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0xf2db7779.
//
// Solidity: function Verify(state bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_TicTacToe *TicTacToeSession) Verify(state []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _TicTacToe.Contract.Verify(&_TicTacToe.CallOpts, state, v, r, s)
}

// Verify is a free data retrieval call binding the contract method 0xf2db7779.
//
// Solidity: function Verify(state bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) Verify(state []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _TicTacToe.Contract.Verify(&_TicTacToe.CallOpts, state, v, r, s)
}

// Can_put is a free data retrieval call binding the contract method 0x355019db.
//
// Solidity: function can_put(state bytes) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) Can_put(opts *bind.CallOpts, state []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "can_put", state)
	return *ret0, err
}

// Can_put is a free data retrieval call binding the contract method 0x355019db.
//
// Solidity: function can_put(state bytes) constant returns(bool)
func (_TicTacToe *TicTacToeSession) Can_put(state []byte) (bool, error) {
	return _TicTacToe.Contract.Can_put(&_TicTacToe.CallOpts, state)
}

// Can_put is a free data retrieval call binding the contract method 0x355019db.
//
// Solidity: function can_put(state bytes) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) Can_put(state []byte) (bool, error) {
	return _TicTacToe.Contract.Can_put(&_TicTacToe.CallOpts, state)
}

// Op_put is a free data retrieval call binding the contract method 0x8739da95.
//
// Solidity: function op_put(state bytes, x uint256, y uint256) constant returns(bytes)
func (_TicTacToe *TicTacToeCaller) Op_put(opts *bind.CallOpts, state []byte, x *big.Int, y *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "op_put", state, x, y)
	return *ret0, err
}

// Op_put is a free data retrieval call binding the contract method 0x8739da95.
//
// Solidity: function op_put(state bytes, x uint256, y uint256) constant returns(bytes)
func (_TicTacToe *TicTacToeSession) Op_put(state []byte, x *big.Int, y *big.Int) ([]byte, error) {
	return _TicTacToe.Contract.Op_put(&_TicTacToe.CallOpts, state, x, y)
}

// Op_put is a free data retrieval call binding the contract method 0x8739da95.
//
// Solidity: function op_put(state bytes, x uint256, y uint256) constant returns(bytes)
func (_TicTacToe *TicTacToeCallerSession) Op_put(state []byte, x *big.Int, y *big.Int) ([]byte, error) {
	return _TicTacToe.Contract.Op_put(&_TicTacToe.CallOpts, state, x, y)
}

// X_max is a free data retrieval call binding the contract method 0x712d215c.
//
// Solidity: function x_max() constant returns(uint256)
func (_TicTacToe *TicTacToeCaller) X_max(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "x_max")
	return *ret0, err
}

// X_max is a free data retrieval call binding the contract method 0x712d215c.
//
// Solidity: function x_max() constant returns(uint256)
func (_TicTacToe *TicTacToeSession) X_max() (*big.Int, error) {
	return _TicTacToe.Contract.X_max(&_TicTacToe.CallOpts)
}

// X_max is a free data retrieval call binding the contract method 0x712d215c.
//
// Solidity: function x_max() constant returns(uint256)
func (_TicTacToe *TicTacToeCallerSession) X_max() (*big.Int, error) {
	return _TicTacToe.Contract.X_max(&_TicTacToe.CallOpts)
}

// Y_max is a free data retrieval call binding the contract method 0x3aa805bc.
//
// Solidity: function y_max() constant returns(uint256)
func (_TicTacToe *TicTacToeCaller) Y_max(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "y_max")
	return *ret0, err
}

// Y_max is a free data retrieval call binding the contract method 0x3aa805bc.
//
// Solidity: function y_max() constant returns(uint256)
func (_TicTacToe *TicTacToeSession) Y_max() (*big.Int, error) {
	return _TicTacToe.Contract.Y_max(&_TicTacToe.CallOpts)
}

// Y_max is a free data retrieval call binding the contract method 0x3aa805bc.
//
// Solidity: function y_max() constant returns(uint256)
func (_TicTacToe *TicTacToeCallerSession) Y_max() (*big.Int, error) {
	return _TicTacToe.Contract.Y_max(&_TicTacToe.CallOpts)
}

// SetPlayer is a paid mutator transaction binding the contract method 0xf6bdd7f1.
//
// Solidity: function SetPlayer(other address) returns()
func (_TicTacToe *TicTacToeTransactor) SetPlayer(opts *bind.TransactOpts, other common.Address) (*types.Transaction, error) {
	return _TicTacToe.contract.Transact(opts, "SetPlayer", other)
}

// SetPlayer is a paid mutator transaction binding the contract method 0xf6bdd7f1.
//
// Solidity: function SetPlayer(other address) returns()
func (_TicTacToe *TicTacToeSession) SetPlayer(other common.Address) (*types.Transaction, error) {
	return _TicTacToe.Contract.SetPlayer(&_TicTacToe.TransactOpts, other)
}

// SetPlayer is a paid mutator transaction binding the contract method 0xf6bdd7f1.
//
// Solidity: function SetPlayer(other address) returns()
func (_TicTacToe *TicTacToeTransactorSession) SetPlayer(other common.Address) (*types.Transaction, error) {
	return _TicTacToe.Contract.SetPlayer(&_TicTacToe.TransactOpts, other)
}
