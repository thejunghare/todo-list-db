package main

import (
	"fmt"

	"github.com/thejunghare/todo-crud-db/src/config"
	"github.com/thejunghare/todo-crud-db/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	fmt.Println("Todo with db")

	defer config.DisconnectDB(db)

	// run all routes
	routes.Routes()
}
