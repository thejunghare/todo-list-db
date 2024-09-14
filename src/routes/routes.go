package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/todo/src/controllers"
)

func Routes() {
	route := gin.Default()
	route.LoadHTMLGlob("templates/*")

	route.GET("/todo", controllers.GetAllTodos)
	// show create form
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// create todo
	route.POST("/todos/create", controllers.CreateTodo)

	route.Run()
}
