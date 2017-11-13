package internal

import (
	"context"
	"log"
	"time"
)

type callback func(context.Context) error

func NewScheduler(ctx context.Context, tick time.Duration) chan callback {
	var callbacks []callback
	var C chan callback

	C = make(chan callback)
	t := time.NewTicker(tick)
	go func() {
		for {
			select {
			case f := <-C:
				log.Printf("New callback registered")
				callbacks = append(callbacks, f)
			case <-t.C:
				for i, cb := range callbacks {
					if err := cb(ctx); err != nil {
						log.Printf("Callback[%d]: error: %v", i, err)
					}
				}
				//TODO death pill
			}
		}
	}()

	return C
}
