package rds

import (
	"go-todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(cnf config.Rds) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cnf.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
