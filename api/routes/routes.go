package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thejunghare/todo/api/controllers"
)

func Routes() {
	route := gin.Default()

	route.Use(cors.Default())

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add your frontend URL here
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	route.LoadHTMLGlob("templates/*")

	route.GET("/", controllers.GetAllTodos)

	// create todo
	route.POST("/create", controllers.CreateTodo)

	route.Run()
}
