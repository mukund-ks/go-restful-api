package main

import (
	"errors"
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

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func getRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Simulating Shop API with Go!"})
}

func createProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusOK, newProduct)
}

func getProductById(id string) (*product, error) {
	for i, p := range products {
		if p.ID == id {
			return &products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}

func main() {
	router := gin.Default()
	router.GET("/", getRoot)
	router.GET("/products", getProducts)
	router.GET("/products/:id", productById)
	router.POST("/products", createProduct)
	router.Run("localhost:8000")
}
