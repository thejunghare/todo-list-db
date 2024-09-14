package main

import (
	"fmt"

	"github.com/thejunghare/todo/src/config"
	"github.com/thejunghare/todo/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.Connectdb()
)

func main() {
	fmt.Println("Todo with db")

	defer config.Disconnectdb(db)

	routes.Routes()
}
