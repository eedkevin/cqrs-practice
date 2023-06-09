package mysql

import (
	"fmt"
	"log"

	"cqrs-practise/internal/cfg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client = gorm.DB

func Connection(cfg cfg.MySQLConfig) *Client {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Schema)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error on connecting to mysql db", err)
	}
	return db
}
