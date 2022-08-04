package main

import (
	"ecommece/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/", controllers.ProductIndex)
	http.HandleFunc("/products", controllers.ProductIndex)
	http.HandleFunc("/products/index", controllers.ProductIndex)

	http.HandleFunc("/cart", controllers.CartIndex)
	http.HandleFunc("/cart/index", controllers.CartIndex)
	http.HandleFunc("/cart/buy", controllers.Buy)
	http.ListenAndServe(":3000", nil)
}
