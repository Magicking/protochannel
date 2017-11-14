package internal

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

var contractKey key = 4242

func Init(ctx context.Context, contract_address, private_key string) {
	nc := CCFromContext(ctx)
	evt_mgr, err := NewEventManager(WhitelistABI)
	if err != nil {
		log.Fatalf("Could not create event manager")
	}
	ctct_addr := common.HexToAddress(contract_address)
	nc.SubscribeToEvents(ctx, ctct_addr, evt_mgr)
	//key, err := crypto.HexToECDSA(private_key)
	//if err != nil {
	//	log.Fatalf("Init: %v", err)
	//}
	//auth := bind.NewKeyedTransactor(key)
	// Caller only for getter otherwise use above code
	ctct, err := NewWhitelistCaller(ctct_addr, nc)
	if err != nil {
		log.Fatalf("Init: %v", err)
	}
	setContextValue(ctx, contractKey, ctct)
}
