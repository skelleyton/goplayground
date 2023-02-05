package main

import (
	"fmt"
	"log"
	"net/http"

	"testapi/repository"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error string `json:"error"`
}

func getProducts(context *gin.Context) {
	id := context.Query("id")
	products := repository.Find(repository.Filter{ ID: id })

	context.JSON(http.StatusOK, products)
}

func getProduct(context *gin.Context) {
	id := context.Param("id")
	product := repository.Find(repository.Filter{ ID: id })
	if (product == nil) {
		error := Error{Error: fmt.Sprintf( "Product %s not found", id)}
		context.JSON(http.StatusNotFound, error)
	}

	context.JSON(http.StatusOK, product)
}

func postProduct(context *gin.Context) {
	var newProduct repository.Product

	if err := context.BindJSON(&newProduct); err != nil {
		log.Println(err)
		return
	}

	result := repository.Insert(newProduct)

	context.JSON(http.StatusCreated, result)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:id", getProduct)
	router.POST("/album", postProduct)

	router.Run("localhost:8080")
}
