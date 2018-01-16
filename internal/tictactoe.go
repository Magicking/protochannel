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
const TicTacToeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"getWinner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeTo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkY\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CreateBoard\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"P2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"can_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"CheckAndApply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"InfoSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeNonce\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FinalizeTimeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x_max\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Judge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"op_put\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkDiag\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"m\",\"type\":\"bytes1\"}],\"name\":\"checkX\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"checkLines\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"GetPlayerMark\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"newOperation\",\"type\":\"bytes1[3]\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"Challenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"state\",\"type\":\"bytes\"},{\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Apply\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"other\",\"type\":\"address\"}],\"name\":\"SetPlayer\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player1\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"player2\",\"type\":\"address\"}],\"name\":\"NewGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"Challenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Settle\",\"type\":\"event\"}]"

// TicTacToeBin is the compiled bytecode used for deploying new contracts.
const TicTacToeBin = `0x6060604052341561000f57600080fd5b60028054600160a060020a03338116600160a060020a0319928316179092556003600181905560008181558154831673533a245f03a1a46cacb933a3beef752fd8ff45c317918290556004805484167348d24715fe9f7daa90286141f9b17e184b5a148b17908190556006919091556007805460ff19169055600580549093169092557fd1af9b1127c37da1f8a3d2e4e7d92bcbfa7bfa2ccf9ce1d29530052e366cc901929081169116604051600160a060020a039283168152911660208201526040908101905180910390a1611acb806100eb6000396000f30060606040526004361061013d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d194c1811461014257806308f2db8c146101b757806313db0e82146101e65780631625059b1461025c5780631de21b76146102e657806330332aea146102f9578063355019db1461030c57806338596d971461035d5780633aa805bc146103be5780633eb1581d146103e357806341da8d2f14610444578063527c575e14610474578063712d215c1461048757806385893cb81461049a5780638739da95146104ad578063875b1b8a146105055780638c8c2a3214610562578063916c9d1d146105c4578063c5454d1114610620578063df6a980514610635578063eb60c1b914610654578063f2db777914610789578063f39e9e75146107ea578063f6bdd7f114610846575b600080fd5b341561014d57600080fd5b61019360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061085a95505050505050565b604051600160a060020a039092168252151560208201526040908101905180910390f35b34156101c257600080fd5b6101ca610939565b604051600160a060020a03909116815260200160405180910390f35b34156101f157600080fd5b61024860046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505084359460200135600160f860020a031916935061094892505050565b604051901515815260200160405180910390f35b341561026757600080fd5b61026f610a67565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102ab578082015183820152602001610293565b50505050905090810190601f1680156102d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102f157600080fd5b6101ca610ca3565b341561030457600080fd5b6101ca610cb2565b341561031757600080fd5b61024860046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610cc195505050505050565b341561036857600080fd5b61026f60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250610d4d915050565b34156103c957600080fd5b6103d1610d8a565b60405190815260200160405180910390f35b34156103ee57600080fd5b6101ca60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250610d90915050565b341561044f57600080fd5b610457610d9e565b604051600160f860020a0319909116815260200160405180910390f35b341561047f57600080fd5b6103d1610daa565b341561049257600080fd5b6103d1610db0565b34156104a557600080fd5b6101ca610db6565b34156104b857600080fd5b61026f60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050843594602001359350610dc592505050565b341561051057600080fd5b61024860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160f860020a0319169250610eab915050565b341561056d57600080fd5b61024860046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505084359460200135600160f860020a031916935061108592505050565b34156105cf57600080fd5b61024860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a03169250611169915050565b341561062b57600080fd5b6106336111ed565b005b341561064057600080fd5b610457600160a060020a03600435166112b9565b341561065f57600080fd5b61063360046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190806060019060038060200260405190810160405291908282606080828437509395946020808201955090358601808201945035925082915081810201604051908101604052809392919081815260200183836020028082843782019150505050505091908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061132395505050505050565b341561079457600080fd5b61024860046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505060ff8535169460208101359450604001359250611564915050565b34156107f557600080fd5b61026f60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506116c8915050565b610633600160a060020a03600435166117ad565b60035460009081908190610878908590600160a060020a0316611169565b1561089457600354600160a060020a0316925060019150610933565b6004546108ab908590600160a060020a0316611169565b156108c757600454600160a060020a0316925060019150610933565b836001815181106108d457fe5b016020015160f860020a90819004810204905060ff811615156108fd5760009250829150610933565b600260ff82160660ff1660001461091f57600454600160a060020a031661092c565b600354600160a060020a03165b6001925092505b50915091565b600554600160a060020a031681565b60008383600054026004016001018151811061096057fe5b016020015160f860020a900460f860020a02600160f860020a0319168484600054026004018151811061098f57fe5b016020015160f860020a900460f860020a02600160f860020a031916148015610a195750838360005402600401600201815181106109c957fe5b016020015160f860020a900460f860020a02600160f860020a031916848460005402600401600101815181106109fb57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b8015610a5f575083836000540260040160020181518110610a3657fe5b016020015160f860020a900460f860020a02600160f860020a03191682600160f860020a031916145b949350505050565b610a6f611a8d565b610a77611a8d565b600d604051805910610a865750595b818152601f19601f830116810160200160405290509050600060f860020a0281600081518110610ab257fe5b906020010190600160f860020a031916908160001a905350600081600181518110610ad957fe5b906020010190600160f860020a031916908160001a905350600081600281518110610b0057fe5b906020010190600160f860020a031916908160001a905350600081600381518110610b2757fe5b906020010190600160f860020a031916908160001a905350600081600481518110610b4e57fe5b906020010190600160f860020a031916908160001a905350600081600581518110610b7557fe5b906020010190600160f860020a031916908160001a905350600081600681518110610b9c57fe5b906020010190600160f860020a031916908160001a905350600081600781518110610bc357fe5b906020010190600160f860020a031916908160001a905350600081600881518110610bea57fe5b906020010190600160f860020a031916908160001a905350600081600981518110610c1157fe5b906020010190600160f860020a031916908160001a905350600081600a81518110610c3857fe5b906020010190600160f860020a031916908160001a905350600081600b81518110610c5f57fe5b906020010190600160f860020a031916908160001a905350600081600c81518110610c8657fe5b906020010190600160f860020a031916908160001a905350919050565b600354600160a060020a031681565b600454600160a060020a031681565b600080600083600281518110610cd357fe5b016020015160f860020a900460f860020a0260f860020a9004915083600381518110610cfb57fe5b016020015160f860020a900460f860020a0260f860020a90049050838282600054026004010181518110610d2b57fe5b016020015160f860020a9081900402600160f860020a03191615949350505050565b610d55611a8d565b610d6185858585611564565b1515610d6c57600080fd5b610d8185610d7c8787878761189a565b6116c8565b95945050505050565b60015481565b6000610d818585858561189a565b60075460f860020a0281565b60065481565b60005481565b600254600160a060020a031681565b610dcd611a8d565b60f860020a84600081518110610ddf57fe5b906020010190600160f860020a031916908160001a90535083600181518110610e0457fe5b016020015160f860020a900460f860020a0260f860020a900460010160f860020a0284600181518110610e3357fe5b906020010190600160f860020a031916908160001a9053508260f860020a0284600281518110610e5f57fe5b906020010190600160f860020a031916908160001a9053508160f860020a0284600381518110610e8b57fe5b906020010190600160f860020a031916908160001a905350929392505050565b60008054839060050181518110610ebe57fe5b016020015160f860020a9081900402600160f860020a03191683600481518110610ee457fe5b016020015160f860020a900460f860020a02600160f860020a031916148015610f7057508260005460020260040160020181518110610f1f57fe5b016020015160f860020a900460f860020a02600160f860020a0319168360005460010260040160010181518110610f5257fe5b016020015160f860020a900460f860020a02600160f860020a031916145b8015610fcd5750600054839060050181518110610f8957fe5b016020015160f860020a9081900402600160f860020a03191683600681518110610faf57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b801561103c57508260005460020260040160000181518110610feb57fe5b016020015160f860020a900460f860020a02600160f860020a031916836000546001026004016001018151811061101e57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b801561107e575060005483906005018151811061105557fe5b016020015160f860020a900460f860020a02600160f860020a03191682600160f860020a031916145b9392505050565b60008054849084016004018151811061109a57fe5b016020015160f860020a9081900402600160f860020a0319168460048501815181106110c257fe5b016020015160f860020a900460f860020a02600160f860020a03191614801561114c5750838360005460020260040101815181106110fc57fe5b016020015160f860020a900460f860020a02600160f860020a0319168484600054600102600401018151811061112e57fe5b016020015160f860020a900460f860020a02600160f860020a031916145b8015610a5f575083836000546002026004010181518110610a3657fe5b600080611175836112b9565b905061118384600083610948565b80611195575061119584600183610948565b806111a757506111a784600283610948565b806111b957506111b984600083611085565b806111cb57506111cb84600183611085565b806111dd57506111dd84600283611085565b80610a5f5750610a5f8482610eab565b60035433600160a060020a0390811691161480611218575060045433600160a060020a039081169116145b80611231575060025433600160a060020a039081169116145b151561123c57600080fd5b600554600160a060020a0316151561125357600080fd5b60065442901061126257600080fd5b6005547f746242f95e6909f1742b2d3bb6cdd61074d700d2dbbd0173ca59d0731bc3846f90600160a060020a0316604051600160a060020a03909116815260200160405180910390a1600554600160a060020a0316ff5b600354600090600160a060020a03838116911614156112dd575060f860020a61131e565b600454600160a060020a038381169116141561131a57507f020000000000000000000000000000000000000000000000000000000000000061131e565b5060005b919050565b600080600080611331611a8d565b60035433600160a060020a039081169116148061135c575060045433600160a060020a039081169116145b80611375575060025433600160a060020a039081169116145b151561138057600080fd5b6113ce8a8960008151811061139157fe5b90602001906020020151896000815181106113a857fe5b90602001906020020151896000815181106113bf57fe5b9060200190602002015161189a565b600354909550339450600160a060020a0380871691161480156113fe5750600454600160a060020a038581169116145b8061142e5750600354600160a060020a03858116911614801561142e5750600454600160a060020a038681169116145b151561143957600080fd5b8960018151811061144657fe5b0160200151600754600160f860020a031960f860020a928390048302811692909102161061147357600080fd5b61147d8a8a611a03565b600454909150611497908290600160a060020a03166116c8565b90506114a28161085a565b90935091508115156114b357600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038516179055896001815181106114e857fe5b01602001516007805460ff191660f860020a92839004830292909204919091179055601e420160068190557f7a180232c19fd38c83e493856a42688c477bae7e82039103b2c1ea6a6162e529908490604051600160a060020a03909216825260208201526040908101905180910390a150505050505050505050565b6000806000806115768888888861189a565b9250600488511161158a57600093506116bd565b8760008151811061159757fe5b016020015160f860020a900460f860020a0260f860020a900460018111156115bb57fe5b9150876001815181106115ca57fe5b016020015160f860020a90819004810204905060008260018111156115eb57fe5b141561160957600254600160a060020a0384811691161493506116bd565b600354600160a060020a038481169116148015906116355750600454600160a060020a03848116911614155b1561164357600093506116bd565b600354600160a060020a038481169116148015611661575060018116155b806116865750600454600160a060020a03848116911614801561168657506001808216145b1561169457600093506116bd565b60018260018111156116a257fe5b14156116b8576116b188610cc1565b93506116bd565b600093505b505050949350505050565b6116d0611a8d565b6000806000856000815181106116e257fe5b016020015160f860020a900460f860020a0260f860020a9004600181111561170657fe5b9250600183600181111561171657fe5b14156117a3578560028151811061172957fe5b016020015160f860020a900460f860020a0260f860020a900491508560038151811061175157fe5b016020015160f860020a900460f860020a0260f860020a90049050611775856112b9565b86838360005402600401018151811061178a57fe5b906020010190600160f860020a031916908160001a9053505b5093949350505050565b600354600160a060020a03161580156117cf5750600454600160a060020a0316155b15156117da57600080fd5b33600160a060020a031681600160a060020a0316141515156117fb57600080fd5b600160a060020a038116151561181057600080fd5b60038054600160a060020a0333811673ffffffffffffffffffffffffffffffffffffffff19928316179283905560048054858316931692909217918290557fd1af9b1127c37da1f8a3d2e4e7d92bcbfa7bfa2ccf9ce1d29530052e366cc9019281169116604051600160a060020a039283168152911660208201526040908101905180910390a150565b60006118a4611a8d565b60006040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a313300000000815250915081876040518083805190602001908083835b6020831061190f5780518252601f1990920191602091820191016118f0565b6001836020036101000a038019825116818451161790925250505091909101905082805190602001908083835b6020831061195b5780518252601f19909201916020918201910161193c565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051809103902090506001818787876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156119ef57600080fd5b505060206040510351979650505050505050565b611a0b611a8d565b815183600081518110611a1a57fe5b906020010190600160f860020a031916908160001a905350602082015183600281518110611a4457fe5b906020010190600160f860020a031916908160001a905350604082015183600381518110611a6e57fe5b906020010190600160f860020a031916908160001a9053509192915050565b602060405190810160405260008152905600a165627a7a723058201aed430507e532db2f7744872860ae96862b2305a601e04e3135d55b57f2f9580029`

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
