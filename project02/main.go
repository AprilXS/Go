package main

import (
	//"fmt" // Println function
	"fmt"
	"net/http" // HTTP status codes

	"github.com/gin-gonic/gin" // Gin framework
	//"errors" // Error handling
)

// type book struct {
// 	ID       string `json:"id"`
// 	Title    string `json:"title"`
// 	Author   string `json:"author"`
// 	Quantity int    `json:"quantity"`
// } // Book struct

type book struct {
	ID       string `json:"id"`
}

// var books = []book{
// 	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
// 	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 5},
// 	{ID: "3", Title: "The Starfield", Author: "Bathezda", Quantity: 10},
// } // Array of books

var books = []book{
	{ID: "1"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // 200 OK + JSON body of books array
} // getBooks function

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "you have error."})
		fmt.Println(err.Error())
		return
	} // 400 Bad Request if the JSON body cannot be bound to the newBook struct

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) // 201 Created + JSON body of newBook
} // createBook function

// func createBook(c *gin.Context) {
// 	var newBook book

// 	if err := c.BindJSON(&newBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Validate the new book data (you can add more validation rules)
// 	if newBook.Title == "" || newBook.Author == "" || newBook.Quantity <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
// 		return
// 	}

// 	// Assuming 'books' is a slice of book
// 	books = append(books, newBook)
// 	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": newBook})
// }


func createboook(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "A new book has been created."})
}

func main() {
	router := gin.Default()            // Create a router
	router.GET("/books", getBooks)     // GET request to /books endpoint
	router.POST("/books", createBook)  // POST request to /books endpoint
	router.POST("/boook", createboook) // POST request to /boook endpoint
	router.Run("localhost:8080")       // Run the server
}
