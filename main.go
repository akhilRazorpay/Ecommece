package main

import (
	productController "ecommece/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/", productController.Index)
	http.HandleFunc("/products", productController.Index)
	http.HandleFunc("/products/index", productController.Index)

	http.ListenAndServe(":3000", nil)
}
