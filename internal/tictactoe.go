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
const TicTacToeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"getWinner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkY\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CreateBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"can_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"CheckAndApply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"InfoSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Judge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"op_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkDiag\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkX\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"checkLines\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"newOperation\",\"type\":\"bytes1[3]\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"ChallengeDebug\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"GetPlayerMark\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"newOperation\",\"type\":\"bytes1[3]\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"Challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Apply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"other\",\"type\":\"address\"}],\"name\":\"SetPlayer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"player2\",\"type\":\"address\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"Challenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Settle\",\"type\":\"event\"}]"

// TicTacToeBin is the compiled bytecode used for deploying new contracts.
const TicTacToeBin = `0x6060604052341561000f57600080fd5b60028054600160a060020a033316600160a060020a0319918216179091556003600181905560008181558154831673533a245f03a1a46cacb933a3beef752fd8ff45c3179091556004805483167348d24715fe9f7daa90286141f9b17e184b5a148b1790556006556007805460ff19169055600580549091169055611d5e806100996000396000f3006060604052600436106101485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d194c1811461014d57806308f2db8c146101c257806313db0e82146101f15780631625059b146102675780631de21b76146102f157806330332aea14610304578063355019db1461031757806338596d97146103685780633aa805bc146103c95780633eb1581d146103ee57806341da8d2f1461044f578063527c575e1461047f578063712d215c1461049257806385893cb8146104a55780638739da95146104b8578063875b1b8a146105105780638c8c2a321461056d578063916c9d1d146105cf578063939040af1461062b578063c5454d1114610760578063df6a980514610775578063eb60c1b914610794578063f2db7779146108c9578063f39e9e751461092a578063f6bdd7f114610986575b600080fd5b341561015857600080fd5b61019e60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061099a95505050505050565b604051600160a060020a039092168252151560208201526040908101905180910390f35b34156101cd57600080fd5b6101d5610a79565b604051600160a060020a03909116815260200160405180910390f35b34156101fc57600080fd5b61025360046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505084359460200135600160f860020a0319169350610a8892505050565b604051901515815260200160405180910390f35b341561027257600080fd5b61027a610ba7565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102b657808201518382015260200161029e565b50505050905090810190601f1680156102e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102fc57600080fd5b6101d5610de3565b341561030f57600080fd5b6101d5610df2565b341561032257600080fd5b61025360046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610e0195505050505050565b341561037357600080fd5b61027a60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250610e8d915050565b34156103d457600080fd5b6103dc610eca565b60405190815260200160405180910390f35b34156103f957600080fd5b6101d560046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250610ed0915050565b341561045a57600080fd5b610462610ede565b604051600160f860020a0319909116815260200160405180910390f35b341561048a57600080fd5b6103dc610eea565b341561049d57600080fd5b6103dc610ef0565b34156104b057600080fd5b6101d5610ef6565b34156104c357600080fd5b61027a60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050843594602001359350610f0592505050565b341561051b57600080fd5b61025360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160f860020a0319169250610feb915050565b341561057857600080fd5b61025360046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505084359460200135600160f860020a03191693506111c592505050565b34156105da57600080fd5b61025360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506112a9915050565b341561063657600080fd5b6103dc60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190806060019060038060200260405190810160405291908282606080828437509395946020808201955090358601808201945035925082915081810201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061132d95505050505050565b341561076b57600080fd5b61077361148f565b005b341561078057600080fd5b610462600160a060020a036004351661155b565b341561079f57600080fd5b61077360046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437820191505050505050919080606001906003806020026040519081016040529190828260608082843750939594602080820195509035860180820194503592508291508181020160405190810160405280939291908181526020018383602002808284378201915050505050509190803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509496506115c595505050505050565b34156108d457600080fd5b61025360046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff85351694602081013594506040013592506117f7915050565b341561093557600080fd5b61027a60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a0316925061195b915050565b610773600160a060020a0360043516611a40565b600354600090819081906109b8908590600160a060020a03166112a9565b156109d457600354600160a060020a0316925060019150610a73565b6004546109eb908590600160a060020a03166112a9565b15610a0757600454600160a060020a0316925060019150610a73565b83600181518110610a1457fe5b016020015160f860020a90819004810204905060ff81161515610a3d5760009250829150610a73565b600260ff82160660ff16600014610a5f57600454600160a060020a0316610a6c565b600354600160a060020a03165b6001925092505b50915091565b600554600160a060020a031681565b600083836000540260040160010181518110610aa057fe5b016020015160f860020a900460f860020a02600160f860020a03191684846000540260040181518110610acf57fe5b016020015160f860020a900460f860020a02600160f860020a031916148015610b59575083836000540260040160020181518110610b0957fe5b016020015160f860020a900460f860020a02600160f860020a03191684846000540260040160010181518110610b3b57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b8015610b9f575083836000540260040160020181518110610b7657fe5b016020015160f860020a900460f860020a02600160f860020a03191682600160f860020a031916145b949350505050565b610baf611d20565b610bb7611d20565b600d604051805910610bc65750595b818152601f19601f830116810160200160405290509050600060f860020a0281600081518110610bf257fe5b906020010190600160f860020a031916908160001a905350600081600181518110610c1957fe5b906020010190600160f860020a031916908160001a905350600081600281518110610c4057fe5b906020010190600160f860020a031916908160001a905350600081600381518110610c6757fe5b906020010190600160f860020a031916908160001a905350600081600481518110610c8e57fe5b906020010190600160f860020a031916908160001a905350600081600581518110610cb557fe5b906020010190600160f860020a031916908160001a905350600081600681518110610cdc57fe5b906020010190600160f860020a031916908160001a905350600081600781518110610d0357fe5b906020010190600160f860020a031916908160001a905350600081600881518110610d2a57fe5b906020010190600160f860020a031916908160001a905350600081600981518110610d5157fe5b906020010190600160f860020a031916908160001a905350600081600a81518110610d7857fe5b906020010190600160f860020a031916908160001a905350600081600b81518110610d9f57fe5b906020010190600160f860020a031916908160001a905350600081600c81518110610dc657fe5b906020010190600160f860020a031916908160001a905350919050565b600354600160a060020a031681565b600454600160a060020a031681565b600080600083600281518110610e1357fe5b016020015160f860020a900460f860020a0260f860020a9004915083600381518110610e3b57fe5b016020015160f860020a900460f860020a0260f860020a90049050838282600054026004010181518110610e6b57fe5b016020015160f860020a9081900402600160f860020a03191615949350505050565b610e95611d20565b610ea1858585856117f7565b1515610eac57600080fd5b610ec185610ebc87878787611b2d565b61195b565b95945050505050565b60015481565b6000610ec185858585611b2d565b60075460f860020a0281565b60065481565b60005481565b600254600160a060020a031681565b610f0d611d20565b60f860020a84600081518110610f1f57fe5b906020010190600160f860020a031916908160001a90535083600181518110610f4457fe5b016020015160f860020a900460f860020a0260f860020a900460010160f860020a0284600181518110610f7357fe5b906020010190600160f860020a031916908160001a9053508260f860020a0284600281518110610f9f57fe5b906020010190600160f860020a031916908160001a9053508160f860020a0284600381518110610fcb57fe5b906020010190600160f860020a031916908160001a905350929392505050565b60008054839060050181518110610ffe57fe5b016020015160f860020a9081900402600160f860020a0319168360048151811061102457fe5b016020015160f860020a900460f860020a02600160f860020a0319161480156110b05750826000546002026004016002018151811061105f57fe5b016020015160f860020a900460f860020a02600160f860020a031916836000546001026004016001018151811061109257fe5b016020015160f860020a900460f860020a02600160f860020a031916145b801561110d57506000548390600501815181106110c957fe5b016020015160f860020a9081900402600160f860020a031916836006815181106110ef57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b801561117c5750826000546002026004016000018151811061112b57fe5b016020015160f860020a900460f860020a02600160f860020a031916836000546001026004016001018151811061115e57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b80156111be575060005483906005018151811061119557fe5b016020015160f860020a900460f860020a02600160f860020a03191682600160f860020a031916145b9392505050565b6000805484908401600401815181106111da57fe5b016020015160f860020a9081900402600160f860020a03191684600485018151811061120257fe5b016020015160f860020a900460f860020a02600160f860020a03191614801561128c57508383600054600202600401018151811061123c57fe5b016020015160f860020a900460f860020a02600160f860020a0319168484600054600102600401018151811061126e57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b8015610b9f575083836000546002026004010181518110610b7657fe5b6000806112b58361155b565b90506112c384600083610a88565b806112d557506112d584600183610a88565b806112e757506112e784600283610a88565b806112f957506112f9846000836111c5565b8061130b575061130b846001836111c5565b8061131d575061131d846002836111c5565b80610b9f5750610b9f8482610feb565b600080600080600061133d611d20565b61138b8b8a60008151811061134e57fe5b906020019060200201518a60008151811061136557fe5b906020019060200201518a60008151811061137c57fe5b90602001906020020151611b2d565b600354909550339450600160a060020a0380871691161480156113bb5750600454600160a060020a038581169116145b806113eb5750600354600160a060020a0385811691161480156113eb5750600454600160a060020a038681169116145b15156113fa5760019550611481565b8a60018151811061140757fe5b0160200151600754600160f860020a031960f860020a92839004830281169290910216106114385760029550611481565b6114428b8b611c96565b60045490915061145c908290600160a060020a031661195b565b90506114678161099a565b909350915081151561147c5760039550611481565b600495505b505050505095945050505050565b60035433600160a060020a03908116911614806114ba575060045433600160a060020a039081169116145b806114d3575060025433600160a060020a039081169116145b15156114de57600080fd5b600554600160a060020a031615156114f557600080fd5b60065442901061150457600080fd5b6005547f746242f95e6909f1742b2d3bb6cdd61074d700d2dbbd0173ca59d0731bc3846f90600160a060020a0316604051600160a060020a03909116815260200160405180910390a1600554600160a060020a0316ff5b600354600090600160a060020a038381169116141561157f575060f860020a6115c0565b600454600160a060020a03838116911614156115bc57507f02000000000000000000000000000000000000000000000000000000000000006115c0565b5060005b919050565b6000806000806115d3611d20565b60035433600160a060020a03908116911614806115fe575060045433600160a060020a039081169116145b80611617575060025433600160a060020a039081169116145b151561162257600080fd5b6116618a8960008151811061163357fe5b906020019060200201518960008151811061164a57fe5b906020019060200201518960008151811061137c57fe5b600354909550339450600160a060020a0380871691161480156116915750600454600160a060020a038581169116145b806116c15750600354600160a060020a0385811691161480156116c15750600454600160a060020a038681169116145b15156116cc57600080fd5b896001815181106116d957fe5b0160200151600754600160f860020a031960f860020a928390048302811692909102161061170657600080fd5b6117108a8a611c96565b60045490915061172a908290600160a060020a031661195b565b90506117358161099a565b909350915081151561174657600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790558960018151811061177b57fe5b01602001516007805460ff191660f860020a92839004830292909204919091179055601e420160068190557f7a180232c19fd38c83e493856a42688c477bae7e82039103b2c1ea6a6162e529908490604051600160a060020a03909216825260208201526040908101905180910390a150505050505050505050565b60008060008061180988888888611b2d565b9250600488511161181d5760009350611950565b8760008151811061182a57fe5b016020015160f860020a900460f860020a0260f860020a9004600181111561184e57fe5b91508760018151811061185d57fe5b016020015160f860020a908190048102049050600082600181111561187e57fe5b141561189c57600254600160a060020a038481169116149350611950565b600354600160a060020a038481169116148015906118c85750600454600160a060020a03848116911614155b156118d65760009350611950565b600354600160a060020a0384811691161480156118f4575060018116155b806119195750600454600160a060020a03848116911614801561191957506001808216145b156119275760009350611950565b600182600181111561193557fe5b141561194b5761194488610e01565b9350611950565b600093505b505050949350505050565b611963611d20565b60008060008560008151811061197557fe5b016020015160f860020a900460f860020a0260f860020a9004600181111561199957fe5b925060018360018111156119a957fe5b1415611a3657856002815181106119bc57fe5b016020015160f860020a900460f860020a0260f860020a90049150856003815181106119e457fe5b016020015160f860020a900460f860020a0260f860020a90049050611a088561155b565b868383600054026004010181518110611a1d57fe5b906020010190600160f860020a031916908160001a9053505b5093949350505050565b600354600160a060020a0316158015611a625750600454600160a060020a0316155b1515611a6d57600080fd5b33600160a060020a031681600160a060020a031614151515611a8e57600080fd5b600160a060020a0381161515611aa357600080fd5b60038054600160a060020a0333811673ffffffffffffffffffffffffffffffffffffffff19928316179283905560048054858316931692909217918290557fd1af9b1127c37da1f8a3d2e4e7d92bcbfa7bfa2ccf9ce1d29530052e366cc9019281169116604051600160a060020a039283168152911660208201526040908101905180910390a150565b6000611b37611d20565b60006040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a313300000000815250915081876040518083805190602001908083835b60208310611ba25780518252601f199092019160209182019101611b83565b6001836020036101000a038019825116818451161790925250505091909101905082805190602001908083835b60208310611bee5780518252601f199092019160209182019101611bcf565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051809103902090506001818787876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515611c8257600080fd5b505060206040510351979650505050505050565b611c9e611d20565b815183600081518110611cad57fe5b906020010190600160f860020a031916908160001a905350602082015183600281518110611cd757fe5b906020010190600160f860020a031916908160001a905350604082015183600381518110611d0157fe5b906020010190600160f860020a031916908160001a9053509192915050565b602060405190810160405260008152905600a165627a7a72305820cb842a26072c67a3cf98e7b0f9604694be4758f626caddc3584f6d36db6b05c00029`

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

// ChallengeDebug is a free data retrieval call binding the contract method 0x939040af.
//
// Solidity: function ChallengeDebug(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) constant returns(uint256)
func (_TicTacToe *TicTacToeCaller) ChallengeDebug(opts *bind.CallOpts, state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "ChallengeDebug", state, newOperation, v, r, s)
	return *ret0, err
}

// ChallengeDebug is a free data retrieval call binding the contract method 0x939040af.
//
// Solidity: function ChallengeDebug(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) constant returns(uint256)
func (_TicTacToe *TicTacToeSession) ChallengeDebug(state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*big.Int, error) {
	return _TicTacToe.Contract.ChallengeDebug(&_TicTacToe.CallOpts, state, newOperation, v, r, s)
}

// ChallengeDebug is a free data retrieval call binding the contract method 0x939040af.
//
// Solidity: function ChallengeDebug(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) constant returns(uint256)
func (_TicTacToe *TicTacToeCallerSession) ChallengeDebug(state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*big.Int, error) {
	return _TicTacToe.Contract.ChallengeDebug(&_TicTacToe.CallOpts, state, newOperation, v, r, s)
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

// FinalizeNonce is a free data retrieval call binding the contract method 0x41da8d2f.
//
// Solidity: function FinalizeNonce() constant returns(bytes1)
func (_TicTacToe *TicTacToeCaller) FinalizeNonce(opts *bind.CallOpts) ([1]byte, error) {
	var (
		ret0 = new([1]byte)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "FinalizeNonce")
	return *ret0, err
}

// FinalizeNonce is a free data retrieval call binding the contract method 0x41da8d2f.
//
// Solidity: function FinalizeNonce() constant returns(bytes1)
func (_TicTacToe *TicTacToeSession) FinalizeNonce() ([1]byte, error) {
	return _TicTacToe.Contract.FinalizeNonce(&_TicTacToe.CallOpts)
}

// FinalizeNonce is a free data retrieval call binding the contract method 0x41da8d2f.
//
// Solidity: function FinalizeNonce() constant returns(bytes1)
func (_TicTacToe *TicTacToeCallerSession) FinalizeNonce() ([1]byte, error) {
	return _TicTacToe.Contract.FinalizeNonce(&_TicTacToe.CallOpts)
}

// FinalizeTimeout is a free data retrieval call binding the contract method 0x527c575e.
//
// Solidity: function FinalizeTimeout() constant returns(uint256)
func (_TicTacToe *TicTacToeCaller) FinalizeTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "FinalizeTimeout")
	return *ret0, err
}

// FinalizeTimeout is a free data retrieval call binding the contract method 0x527c575e.
//
// Solidity: function FinalizeTimeout() constant returns(uint256)
func (_TicTacToe *TicTacToeSession) FinalizeTimeout() (*big.Int, error) {
	return _TicTacToe.Contract.FinalizeTimeout(&_TicTacToe.CallOpts)
}

// FinalizeTimeout is a free data retrieval call binding the contract method 0x527c575e.
//
// Solidity: function FinalizeTimeout() constant returns(uint256)
func (_TicTacToe *TicTacToeCallerSession) FinalizeTimeout() (*big.Int, error) {
	return _TicTacToe.Contract.FinalizeTimeout(&_TicTacToe.CallOpts)
}

// FinalizeTo is a free data retrieval call binding the contract method 0x08f2db8c.
//
// Solidity: function FinalizeTo() constant returns(address)
func (_TicTacToe *TicTacToeCaller) FinalizeTo(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "FinalizeTo")
	return *ret0, err
}

// FinalizeTo is a free data retrieval call binding the contract method 0x08f2db8c.
//
// Solidity: function FinalizeTo() constant returns(address)
func (_TicTacToe *TicTacToeSession) FinalizeTo() (common.Address, error) {
	return _TicTacToe.Contract.FinalizeTo(&_TicTacToe.CallOpts)
}

// FinalizeTo is a free data retrieval call binding the contract method 0x08f2db8c.
//
// Solidity: function FinalizeTo() constant returns(address)
func (_TicTacToe *TicTacToeCallerSession) FinalizeTo() (common.Address, error) {
	return _TicTacToe.Contract.FinalizeTo(&_TicTacToe.CallOpts)
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

// InfoSigner is a free data retrieval call binding the contract method 0x3eb1581d.
//
// Solidity: function InfoSigner(message bytes, v uint8, r bytes32, s bytes32) constant returns(address)
func (_TicTacToe *TicTacToeCaller) InfoSigner(opts *bind.CallOpts, message []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "InfoSigner", message, v, r, s)
	return *ret0, err
}

// InfoSigner is a free data retrieval call binding the contract method 0x3eb1581d.
//
// Solidity: function InfoSigner(message bytes, v uint8, r bytes32, s bytes32) constant returns(address)
func (_TicTacToe *TicTacToeSession) InfoSigner(message []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _TicTacToe.Contract.InfoSigner(&_TicTacToe.CallOpts, message, v, r, s)
}

// InfoSigner is a free data retrieval call binding the contract method 0x3eb1581d.
//
// Solidity: function InfoSigner(message bytes, v uint8, r bytes32, s bytes32) constant returns(address)
func (_TicTacToe *TicTacToeCallerSession) InfoSigner(message []byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _TicTacToe.Contract.InfoSigner(&_TicTacToe.CallOpts, message, v, r, s)
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

// CheckDiag is a free data retrieval call binding the contract method 0x875b1b8a.
//
// Solidity: function checkDiag(state bytes, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) CheckDiag(opts *bind.CallOpts, state []byte, m [1]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "checkDiag", state, m)
	return *ret0, err
}

// CheckDiag is a free data retrieval call binding the contract method 0x875b1b8a.
//
// Solidity: function checkDiag(state bytes, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeSession) CheckDiag(state []byte, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckDiag(&_TicTacToe.CallOpts, state, m)
}

// CheckDiag is a free data retrieval call binding the contract method 0x875b1b8a.
//
// Solidity: function checkDiag(state bytes, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) CheckDiag(state []byte, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckDiag(&_TicTacToe.CallOpts, state, m)
}

// CheckLines is a free data retrieval call binding the contract method 0x916c9d1d.
//
// Solidity: function checkLines(state bytes, player address) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) CheckLines(opts *bind.CallOpts, state []byte, player common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "checkLines", state, player)
	return *ret0, err
}

// CheckLines is a free data retrieval call binding the contract method 0x916c9d1d.
//
// Solidity: function checkLines(state bytes, player address) constant returns(bool)
func (_TicTacToe *TicTacToeSession) CheckLines(state []byte, player common.Address) (bool, error) {
	return _TicTacToe.Contract.CheckLines(&_TicTacToe.CallOpts, state, player)
}

// CheckLines is a free data retrieval call binding the contract method 0x916c9d1d.
//
// Solidity: function checkLines(state bytes, player address) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) CheckLines(state []byte, player common.Address) (bool, error) {
	return _TicTacToe.Contract.CheckLines(&_TicTacToe.CallOpts, state, player)
}

// CheckX is a free data retrieval call binding the contract method 0x8c8c2a32.
//
// Solidity: function checkX(state bytes, x uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) CheckX(opts *bind.CallOpts, state []byte, x *big.Int, m [1]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "checkX", state, x, m)
	return *ret0, err
}

// CheckX is a free data retrieval call binding the contract method 0x8c8c2a32.
//
// Solidity: function checkX(state bytes, x uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeSession) CheckX(state []byte, x *big.Int, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckX(&_TicTacToe.CallOpts, state, x, m)
}

// CheckX is a free data retrieval call binding the contract method 0x8c8c2a32.
//
// Solidity: function checkX(state bytes, x uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) CheckX(state []byte, x *big.Int, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckX(&_TicTacToe.CallOpts, state, x, m)
}

// CheckY is a free data retrieval call binding the contract method 0x13db0e82.
//
// Solidity: function checkY(state bytes, y uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCaller) CheckY(opts *bind.CallOpts, state []byte, y *big.Int, m [1]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicTacToe.contract.Call(opts, out, "checkY", state, y, m)
	return *ret0, err
}

// CheckY is a free data retrieval call binding the contract method 0x13db0e82.
//
// Solidity: function checkY(state bytes, y uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeSession) CheckY(state []byte, y *big.Int, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckY(&_TicTacToe.CallOpts, state, y, m)
}

// CheckY is a free data retrieval call binding the contract method 0x13db0e82.
//
// Solidity: function checkY(state bytes, y uint256, m bytes1) constant returns(bool)
func (_TicTacToe *TicTacToeCallerSession) CheckY(state []byte, y *big.Int, m [1]byte) (bool, error) {
	return _TicTacToe.Contract.CheckY(&_TicTacToe.CallOpts, state, y, m)
}

// GetWinner is a free data retrieval call binding the contract method 0x05d194c1.
//
// Solidity: function getWinner(state bytes) constant returns(address, bool)
func (_TicTacToe *TicTacToeCaller) GetWinner(opts *bind.CallOpts, state []byte) (common.Address, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _TicTacToe.contract.Call(opts, out, "getWinner", state)
	return *ret0, *ret1, err
}

// GetWinner is a free data retrieval call binding the contract method 0x05d194c1.
//
// Solidity: function getWinner(state bytes) constant returns(address, bool)
func (_TicTacToe *TicTacToeSession) GetWinner(state []byte) (common.Address, bool, error) {
	return _TicTacToe.Contract.GetWinner(&_TicTacToe.CallOpts, state)
}

// GetWinner is a free data retrieval call binding the contract method 0x05d194c1.
//
// Solidity: function getWinner(state bytes) constant returns(address, bool)
func (_TicTacToe *TicTacToeCallerSession) GetWinner(state []byte) (common.Address, bool, error) {
	return _TicTacToe.Contract.GetWinner(&_TicTacToe.CallOpts, state)
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

// Challenge is a paid mutator transaction binding the contract method 0xeb60c1b9.
//
// Solidity: function Challenge(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) returns()
func (_TicTacToe *TicTacToeTransactor) Challenge(opts *bind.TransactOpts, state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _TicTacToe.contract.Transact(opts, "Challenge", state, newOperation, v, r, s)
}

// Challenge is a paid mutator transaction binding the contract method 0xeb60c1b9.
//
// Solidity: function Challenge(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) returns()
func (_TicTacToe *TicTacToeSession) Challenge(state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _TicTacToe.Contract.Challenge(&_TicTacToe.TransactOpts, state, newOperation, v, r, s)
}

// Challenge is a paid mutator transaction binding the contract method 0xeb60c1b9.
//
// Solidity: function Challenge(state bytes, newOperation bytes1[3], v uint8[], r bytes32[], s bytes32[]) returns()
func (_TicTacToe *TicTacToeTransactorSession) Challenge(state []byte, newOperation [3][1]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _TicTacToe.Contract.Challenge(&_TicTacToe.TransactOpts, state, newOperation, v, r, s)
}

// Finalize is a paid mutator transaction binding the contract method 0xc5454d11.
//
// Solidity: function Finalize() returns()
func (_TicTacToe *TicTacToeTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicTacToe.contract.Transact(opts, "Finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0xc5454d11.
//
// Solidity: function Finalize() returns()
func (_TicTacToe *TicTacToeSession) Finalize() (*types.Transaction, error) {
	return _TicTacToe.Contract.Finalize(&_TicTacToe.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0xc5454d11.
//
// Solidity: function Finalize() returns()
func (_TicTacToe *TicTacToeTransactorSession) Finalize() (*types.Transaction, error) {
	return _TicTacToe.Contract.Finalize(&_TicTacToe.TransactOpts)
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
