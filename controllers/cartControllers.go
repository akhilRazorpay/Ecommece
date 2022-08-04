package controllers

import (
	"ecommece/entities"
	"ecommece/models"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func CartIndex(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	strCart := session.Values["cart"].(string)
	var cart []entities.Item
	json.Unmarshal([]byte(strCart), &cart)
	data := map[string]interface{}{
		"cart": cart,
	}
	tmp, _ := template.New("cart.html").Funcs(template.FuncMap{"total": func(item entities.Item) float64 {
		return item.Product.Price * float64(item.Quantity)
	},
	}).ParseFiles("views/cart.html")
	tmp.Execute(response, data)
}

func Buy(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var productModel models.ProductModel

	product, _ := productModel.Find(id)

	session, _ := store.Get(request, "mysession")
	cart := session.Values["cart"]
	if cart == nil {
		var cart []entities.Item
		cart = append(cart, entities.Item{
			Product:  product,
			Quantity: 1,
		})
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
		session.Save(request, response)
	} else {
	}
	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}
