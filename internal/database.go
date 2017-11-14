package internal

import (
	"fmt"
	"log"
	"time"

	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Msg struct {
	Addr string
	Data string
}

func InsertMsg(ctx context.Context, msg *Msg) error {
	db := DBFromContext(ctx)
	if err := db.Create(msg).Error; err != nil {
		return err
	}

	return nil
}

func DeleteMsg(ctx context.Context, addr string) error {
	db := DBFromContext(ctx)
	cursor := db.Where(Msg{Addr: addr}).Delete(Msg{})
	if cursor.Error != nil {
		return fmt.Errorf("Error deleting for URL (%v): %v", addr, cursor.Error)
	}
	return nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Msg{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
