package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type todo struct {
	ID int `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{
		ID: 1,
		Item: "Learn Go",
		Completed: false,
	},
	{
		ID: 2,
		Item: "Learn React",
		Completed: false,
	},
	{
		ID: 3,
		Item: "Learn Vue",
		Completed: false,
	},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "you have errorrrrr."})
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	// router := gin.Default()
	// router.GET("/todos", getTodos)
	// router.Run("localhost:9090")

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}