package main

import (
	"ecommece/controllers"
	"net/http"
)

func main() {

	http.HandleFunc("/", controllers.ProductIndex)
	http.HandleFunc("/products", controllers.ProductIndex)
	http.HandleFunc("/products/index", controllers.ProductIndex)
	http.HandleFunc("/products/delete", controllers.Delete)

	http.HandleFunc("/products/insert", controllers.Insert)
	http.HandleFunc("/products/inserted", controllers.Inserted)

	http.HandleFunc("/cart", controllers.CartIndex)
	http.HandleFunc("/cart/index", controllers.CartIndex)
	http.HandleFunc("/cart/buy", controllers.Buy)
	http.HandleFunc("/cart/remove", controllers.Remove)

	http.ListenAndServe(":3000", nil)
}
