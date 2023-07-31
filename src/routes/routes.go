package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thejunghare/todo-crud-db/src/controllers"
)

func Routes() {
	route := gin.Default()

	route.GET("/todo", controllers.GetAllTodos)

	route.Run()
}
