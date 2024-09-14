package controllers

// var db *gorm.DB = config.ConnectDB()

// struct for request
/* type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
} */

// struct for response
/* type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
} */

// Create todo
/* func CreateTodo(c *gin.Context) {
	var data todoRequest

	// Marshal the request
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while marshalling the data",
		})
		return
	}

	// Matching todo models struct with todo request struct
	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	// Querying to database
	result := db.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while quering to database",
		})
		return
	}

	// Match result to create response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	c.JSON(http.StatusCreated, response)
} */

// Get all todos
/* func GetAllTodos(c *gin.Context) {
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
} */

// Update todos
/* func UpdateTodo(c *gin.Context) {
	var data todoRequest

	// Defining request parameter to get todo id
	reqParamID := c.Param("idTodo")
	idTodo := cast.ToUint(reqParamID)

	// Marshal the request body json to request body struct
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while Marshalling the request data",
		})
		return
	}

	// Initiate models todo
	todo := models.Todo{}

	// Querying find todo data by todo id from request parameter
	todoByID := db.Where("id = ?", idTodo).First(&todo)
	if todoByID.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Todo not found",
		})
		return
	}

	// Matching todo request with todo models
	todo.Name = data.Name
	todo.Description = data.Description

	// update the todo data
	result := db.Save(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while updating the data",
		})
		return
	}

	// Matching result todo to response struct
	var response todoRequest
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response
	c.JSON(http.StatusCreated, response)
} */
