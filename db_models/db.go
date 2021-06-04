package db_models

import (
	"github.com/go-pg/pg"
	"sync"
)

var db *pg.DB

var dbLock sync.Mutex

func GetDB() *pg.DB {
	if db == nil {
		dbLock.Lock()
		if db != nil {
			return db
		}
		defer dbLock.Unlock()
		db = pg.Connect(&pg.Options{
			User:     "postgres",
			Password: "postgres",
			Addr:     "localhost:5432",
			Database: "postgres",
		})
	}
	return db
}
