package main

import (
	//"fmt" // Println function 
	"net/http" // HTTP status codes
	"github.com/gin-gonic/gin" // Gin framework 
	//"errors" // Error handling 
)

type book struct {
	ID 			string	`json:"id"`
	Title 		string 	`json:"title"`
	Author 		string 	`json:"author"`
	Quantity 	int 	`json:"quantity"`
}	// Book struct 

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 5},
	{ID: "3", Title: "The Starfield", Author: "Bathezda", Quantity: 10},
}	// Array of books 

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // 200 OK + JSON body of books array 
}	// getBooks function 

func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}	// 400 Bad Request if the JSON body cannot be bound to the newBook struct 

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)	// 201 Created + JSON body of newBook
} // createBook function 

func main() {
	router := gin.Default() // Create a router 
	router.GET("/books", getBooks) 
	router.Run("localhost:8080") // Run the server 
}