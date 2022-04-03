package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tframesqui/go-event-driven/catalog/pkg/models"
	"github.com/tframesqui/go-event-driven/catalog/pkg/storage"
)

func main() {
	//elasticSearch configuration
	storage.SetupStorage()

	configRoutes()
}

func configRoutes() {
	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		productList(c)
	})

	r.POST("/product", func(c *gin.Context) {
		addProduct(c)
	})

	r.Run()
}

func productList(gin *gin.Context) {
	var products = storage.GetProducts()
	gin.JSON(http.StatusOK, products)
}

func addProduct(gin *gin.Context) {
	decoder := json.NewDecoder(gin.Request.Body)
	var product models.Product
	err := decoder.Decode(&product)

	if err != nil {
		panic(err)
	}

	storage.AddProduct(product)
}
