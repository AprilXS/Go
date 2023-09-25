package main

import (
	"errors"
	"net/http"
	"strconv"

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

func getTodoById(id string) (*todo, error){
	for i, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			return &todos[i], nil
		} 
	} // end of for loop 
	return nil, errors.New("todo not found") // return error if not found
} // end of getTodoById

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"}) // return error if not found
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"}) // return error if not found
		return
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo) // get todo by id
	router.PATCH("/todos/:id", toggleTodoStatus) // toggle todo status
	router.Run("localhost:9090")
}