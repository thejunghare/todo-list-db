package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/todo/api/config"
	"github.com/thejunghare/todo/api/models"
	"gorm.io/gorm"
)

var db *gorm.DB = config.Connectdb()

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	result := db.Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

/* func CreateTodo(c *gin.Context) {
	name := c.Request.FormValue("name")

	if name == "" {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "fields required",
		})
		return
	}

	todo := models.Todo{}
	todo.Name = name

	result := db.Create(&todo)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Error while quering to database",
		})
		return
	}

	c.HTML(http.StatusCreated, "todo_success.html", gin.H{
		"ID":   todo.ID,
		"Name": todo.Name,
	})
} */

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// Print the received todo
	// fmt.Println("Received Todo:", todo.Name)

	result := db.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while querying the database",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":   todo.ID,
		"Name": todo.Name,
	})
}
