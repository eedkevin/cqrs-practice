package sqlite3

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"cqrs-practise/internal/app/infrastructure/sqlite3/model"
	"cqrs-practise/internal/cfg"
)

type Client = gorm.DB

func Connection(cfg *cfg.Config) *Client {
	log.Println("Connecting sqlite3 database...")
	db, err := gorm.Open(sqlite.Open(cfg.SQLite.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Error on connecting to sqlite db", err)
	}
	log.Println("sqlite3 database connected")
	autoMigrate(db)
	return db
}

func autoMigrate(db *Client) {
	log.Println("Migrating database models to sqlite3...")
	db.AutoMigrate(&model.Event{}, &model.EventLog{})
	log.Println("Database models migrated")
}
