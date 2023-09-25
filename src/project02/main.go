package main

import (
	//"fmt" // Println function
	"errors"   // Error handling
	"net/http" // HTTP status codes

	"github.com/gin-gonic/gin" // Gin framework
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
} // Book struct

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 5},
	{ID: "3", Title: "The Starfield", Author: "Bathezda", Quantity: 10},
} // Array of books

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // 200 OK + JSON body of books array
} // getBooks function

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id query parameter."})
		return
	}

	book, error := getBookById(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book) // 200 OK + JSON body of books array
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing id query parameter."})
		return
	}

	book, error := getBookById(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book is out of stock."})
		return
	} // if book is out of stock return 404 not found error message

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book) // 200 OK + JSON body of books array
}

func bookById(c *gin.Context) {
	id := c.Param("id") // Get the id from the URL
	book, error := getBookById(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		} // if
	}
	return nil, errors.New("Book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "you have error in Binding JSON in createBook."})
		return
	} // 400 Bad Request if the JSON body cannot be bound to the newBook struct

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) // 201 Created + JSON body of newBook
} // createBook function

func main() {
	router := gin.Default()                 // Create a router
	router.GET("/books", getBooks)          // GET request to /books endpoint
	router.POST("/books", createBook)       // POST request to /books endpoint
	router.GET("/books/:id", bookById)      // POST request to /books endpoint
	router.PATCH("/checkout", checkoutBook) // POST request to /books endpoint
	router.PATCH("/return", returnBook)     // POST request to /books endpoint
	router.Run("localhost:8080")            // Run the server
} // main function
