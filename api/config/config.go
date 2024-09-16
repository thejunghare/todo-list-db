package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connectdb() *gorm.DB {
	dsn := "host=localhost user=postgres password=1010 dbname=todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to mysql database")
	} else {
		fmt.Println("Connected to db")
	}

	return db
}

func Disconnectdb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection to database")
	}

	dbSQL.Close()
}
