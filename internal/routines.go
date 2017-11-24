package internal

import (
	"context"
	ethtk "github.com/Magicking/gethitihteg"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strings"
)

var contractKey key = 4242
var contractAddr common.Address

func InitContract(ctx context.Context, auth *bind.TransactOpts) (common.Address, interface{}) {
	rawContract := common.FromHex(TicTacToeBin)
	client := CCFromContext(ctx)
	contractABI, err := abi.JSON(strings.NewReader(TicTacToeABI))
	addr, tx, c, err := ethtk.DeployContract(auth, contractABI, rawContract, client)
	if err != nil {
		log.Fatalf("InitContract: %v", err)
	}
	log.Printf("%v\nDeployed contract at %s", tx, addr.String())
	return addr, c
}

func Init(ctx context.Context, contract_address, private_key string) {
	nc := CCFromContext(ctx)
	evt_mgr, err := NewEventManager(TicTacToeABI)
	if err != nil {
		log.Fatalf("Could not create event manager")
	}
	key, err := crypto.HexToECDSA(private_key)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	_key = key
	auth := bind.NewKeyedTransactor(key)
	ctct_addr := common.HexToAddress(contract_address)
	if contract_address == "" {
		ctct_addr, _ = InitContract(ctx, auth)
	}
	ctct, err := NewTicTacToe(ctct_addr, nc)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	setContextValue(ctx, contractKey, ctct)
	contractAddr = ctct_addr
	err = DeleteAllStates(ctx)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	nc.SubscribeToEvents(ctx, ctct_addr, evt_mgr)
}
