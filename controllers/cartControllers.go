package cartController

import (
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("views/cart.html")
	tmp.Execute(response, nil)
}
