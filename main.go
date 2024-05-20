package main

import (
	// "errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []product{
	{ID: "1", Name: "Monitor", Price: 100},
	{ID: "2", Name: "Mouse", Price: 50},
	{ID: "3", Name: "Keyboard", Price: 50},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Shop API with Go!")
}

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/products", getBooks)
	router.Run("localhost:8000")
}
