package controllers

import (
	"ecommece/models"
	"html/template"
	"net/http"
)

func ProductIndex(response http.ResponseWriter, request *http.Request) {
	// fmt.Println("ok")
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	// fmt.Println(products)
	data := map[string]interface{}{
		"products": products,
	}

	tmp, _ := template.ParseFiles("views/index.html")
	tmp.Execute(response, data)
}
