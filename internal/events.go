package internal

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strings"
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
	if len(evt.Topics) == 0 {
		log.Printf("Unknown event")
		return
	}
	id := evt.Topics[0]
	_evt, found := e.events[id]
	if found {
		log.Printf("Event %s fired !", _evt.Name)
		log.Printf("Topic 0: %v", id.Hex())
		log.Printf("Data: %v", evt.Data)
	}
}
