package internal

import (
	"context"
	"fmt"
	blktk "github.com/Magicking/gethitihteg"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type key int

var ctxValKey key = 0
var dbKey key = 1
var ethRpcKey key = 2
var schedulerKey key = 3
var blkKey key = 4

type Time time.Duration

func (t Time) Hours() int64 {
	return int64(time.Duration(t) / time.Hour)
}

func (t Time) Minutes() int64 {
	return int64((time.Duration(t) % time.Hour) / time.Minute)
}

func (t Time) Seconds() int64 {
	return int64((time.Duration(t) % time.Minute) / time.Second)
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hours(), t.Minutes(), t.Seconds())
}

type ctxValues struct {
	m map[key]interface{}
}

func (v ctxValues) Set(k key, val interface{}) {
	v.m[k] = val
}

func (v ctxValues) Get(k key) interface{} {
	val, ok := v.m[k]
	if !ok {
		log.Fatalf("Could not find key: %v", k)
	}
	return val
}

func NewCtxValues() *ctxValues {
	mm := make(map[key]interface{})
	cv := &ctxValues{
		m: mm,
	}
	return cv
}

func getContextValue(ctx context.Context, k key) interface{} {
	v, ok := ctx.Value(ctxValKey).(*ctxValues)
	if !ok {
		log.Fatalf("Could not obtain map context values")
	}
	return v.Get(k)
}

func setContextValue(ctx context.Context, k key, val interface{}) {
	v, ok := ctx.Value(ctxValKey).(*ctxValues)
	if !ok {
		log.Fatalf("Could not obtain map context values")
	}
	v.Set(k, val)
}

func InitContext(ctx context.Context) context.Context {
	values := NewCtxValues()
	return context.WithValue(ctx, ctxValKey, values)
}

func NewDBToContext(ctx context.Context, dbDsn string) {
	db, err := InitDatabase(dbDsn)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	setContextValue(ctx, dbKey, db)
}

func DBFromContext(ctx context.Context) *gorm.DB {
	key := dbKey
	ret, ok := getContextValue(ctx, key).(*gorm.DB)
	if !ok {
		log.Fatalf("Could not cast context with key %d", key)
	}
	return ret
}

func NewCCToContext(ctx context.Context, wsURI string, retry int) {
	cc, err := blktk.NewNodeConnector(wsURI, retry)
	if err != nil {
		log.Fatalf("Could not initialize client context: %v", err)
	}
	setContextValue(ctx, ethRpcKey, cc)
}

func CCFromContext(ctx context.Context) *blktk.NodeConnector {
	key := ethRpcKey
	ret, ok := getContextValue(ctx, key).(*blktk.NodeConnector)
	if !ok {
		log.Fatalf("Could not cast context with key %d", key)
	}
	return ret
}

func NewSchedulerToContext(ctx context.Context, tick time.Duration) {
	c := NewScheduler(ctx, tick)
	setContextValue(ctx, schedulerKey, c)
}

func SchedulerChanFromContext(ctx context.Context) chan callback {
	key := schedulerKey
	ret, ok := getContextValue(ctx, key).(chan callback)
	if !ok {
		log.Fatalf("Could not cast context with key %d", key)
	}
	return ret
}

//func NewBLKToContext(ctx context.Context, wsURI, privateKey string, entry int) {
//	blk, err := blktk.NewBlockchainContext(wsURI, privateKey, entry)
//	if err != nil {
//		log.Fatalf("Could not initialize blockchain context: %v", err)
//	}
//	setContextValue(ctx, blkKey, blk)
//}
//
//func BLKFromContext(ctx context.Context) *blktk.BlockchainContext {
//	key := blkKey
//	ret, ok := getContextValue(ctx, key).(*blktk.BlockchainContext)
//	if !ok {
//		log.Fatalf("Could not cast context with key %d", key)
//	}
//	return ret
//}
