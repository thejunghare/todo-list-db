package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/todo/src/config"
	"github.com/thejunghare/todo/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB = config.Connectdb()

type todoRequest struct {
	Name string `json:"name"`
}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	// query to find todo datas
	err := db.Find(&todos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting the todos",
		})
		return
	}

	// Creating http response
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

func CreateTodo(c *gin.Context) {
	name := c.Request.FormValue("name")

	// Marshal the request
	if name == "" {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "fields required",
		})
		return
	}

	// Matching todo models struct with todo request struct
	todo := models.Todo{}
	todo.Name = name

	// Querying to database
	result := db.Create(&todo)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Error while quering to database",
		})
		return
	}

	c.HTML(http.StatusCreated, "todo_success.html.html", gin.H{
		"ID":   todo.ID,
		"Name": todo.Name,
	})
}
