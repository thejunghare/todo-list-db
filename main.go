package main

import (
	"fmt"

	"github.com/thejunghare/todo/api/config"
	"github.com/thejunghare/todo/api/routes"
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
