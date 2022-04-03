package storage

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/tframesqui/go-event-driven/catalog/pkg/models"
)

var products = []models.Product{
	{Name: "TESTE"},
}

func GetProducts() []models.Product {	
	return products
}

func AddProduct(product models.Product) {
	put, err := es.Index("product", esutil.NewJSONReader(&product))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(put)
}
