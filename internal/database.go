package internal

import (
	"fmt"
	"log"
	"time"

	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type State struct {
	gorm.Model
	State      string
	Signatures pq.StringArray `gorm:"type:varchar(255)[]"`
}

func InsertState(ctx context.Context, st *State, sigs []string) error {
	db := DBFromContext(ctx)

	st.Signatures = pq.StringArray(sigs)
	if err := db.Create(st).Error; err != nil {
		return err
	}

	return nil
}

func debugGetAllDB(ctx context.Context) {
	db := DBFromContext(ctx)
	var st []State
	cursor := db.Find(&st)
	if cursor.RecordNotFound() {
		log.Println("RecordNotFound: No states found in database")
		return
	}
	log.Println(st)
}

func LastState(ctx context.Context) (*State, error) {
	db := DBFromContext(ctx)
	var st State
	cursor := db.Last(&st)
	if cursor.RecordNotFound() {
		return nil, nil
	}
	if cursor.Error != nil {
		return nil, fmt.Errorf("Error getting last state: %v", cursor.Error)
	}
	return &st, nil
}

func DeleteAllStates(ctx context.Context) error {
	db := DBFromContext(ctx)
	cursor := db.Where(State{}).Delete(State{})
	if cursor.Error != nil {
		return fmt.Errorf("Error deleting all: %v", cursor.Error)
	}
	return nil
}

/*
func DeleteState(ctx context.Context, id uint) error {
	db := DBFromContext(ctx)
	cursor := db.Where(State{ID: id}).Delete(State{})
	if cursor.Error != nil {
		return fmt.Errorf("Error deleting for identifier (%v): %v", id, cursor.Error)
	}
	return nil
}*/

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
	if err = db.AutoMigrate(&State{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
