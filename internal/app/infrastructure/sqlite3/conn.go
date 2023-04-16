package sqlite3

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"cqrs-practise/internal/cfg"
)

type Client = gorm.DB

func Connection() *Client {
	db, err := gorm.Open(sqlite.Open(cfg.Cfg.SQLite.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Error on connecting to sqlite db", err)
	}
	return db
}
