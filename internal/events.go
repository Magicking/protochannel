package internal

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventManager struct {
	abi    abi.ABI
	events map[common.Hash]abi.Event
}

func NewEventManager(abiJSON string) (*EventManager, error) {
	_abi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	events := make(map[common.Hash]abi.Event)
	for _, e := range _abi.Events {
		id := e.Id()
		events[id] = e
		log.Printf("Registering %s with hash: %s", e.Name, id.Hex())
	}
	return &EventManager{abi: _abi, events: events}, nil
}

// FireEvent handler the log and prograte event to subsystems
func (e *EventManager) FireEvent(ctx context.Context, evt *types.Log) {
	// TODO See https://github.com/ethereum/go-ethereum/pull/15832 for a nicer way to handle event
	if len(evt.Topics) == 0 {
		log.Printf("Unknown event")
		return
	}
	id := evt.Topics[0]
	_evt, found := e.events[id]
	if !found {
		log.Println("unknown event")
		return
	}
	name := _evt.Name
	switch name {
	case "NewGame":
		type newGameEvt struct {
			Player1 common.Address
			Player2 common.Address
		}
		var nge newGameEvt
		err := e.abi.Unpack(&nge, name, evt.Data)
		if err != nil {
			log.Printf("FireEvent, unpack", err)
		}
		log.Printf("New game with %s and %s\n", nge.Player1.String(), nge.Player2.String())
		break
	case "Challenge":
		type challengeEvt struct {
			Player  common.Address
			Timeout *big.Int
		}
		var ce challengeEvt
		err := e.abi.Unpack(&ce, name, evt.Data)
		if err != nil {
			log.Printf("FireEvent, unpack", err)
		}
		log.Printf("Challenged game by %s, timeout at %+v\n", ce.Player.String(), ce.Timeout)
		break
	default:
		log.Printf("Unregistered event %s", name)
	}
}
