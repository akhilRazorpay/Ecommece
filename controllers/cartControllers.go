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
		"cart":  cart,
		"total": total(cart),
	}

	tmp, _ := template.New("cart.html").Funcs(template.FuncMap{"total": func(item entities.Item) float64 {
		// fmt.Println(item.Product.Price)
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

	// fmt.Println(product)

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
	} else {
		strCart := session.Values["cart"].(string)
		var cart []entities.Item
		json.Unmarshal([]byte(strCart), &cart)
		// fmt.Println(cart)
		index := exists(id, cart)
		if index == -1 {
			cart = append(cart, entities.Item{
				Product:  product,
				Quantity: 1,
			})
		} else {
			cart[index].Quantity++
		}
		bytesCart, _ := json.Marshal(cart)
		session.Values["cart"] = string(bytesCart)
	}
	session.Save(request, response)

	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}

func Remove(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	session, _ := store.Get(request, "mysession")
	strCart := session.Values["cart"].(string)

	var cart []entities.Item
	json.Unmarshal([]byte(strCart), &cart)

	index := exists(id, cart)
	cart = remove(cart, index)

	bytesCart, _ := json.Marshal(cart)

	session.Values["cart"] = string(bytesCart)
	session.Save(request, response)

	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}

func exists(id int64, cart []entities.Item) int {
	// fmt.Println(cart, id)
	for i := 0; i < len(cart); i++ {
		if cart[i].Product.Id == id {
			return i
		}

	}
	return -1
}

func total(cart []entities.Item) float64 {
	// fmt.Println("ok")
	// fmt.Println(cart)

	var s float64 = 0
	for _, item := range cart {

		s += item.Product.Price * float64(item.Quantity)
	}
	// fmt.Println(s)
	return s
}

func remove(cart []entities.Item, index int) []entities.Item {
	// fmt.Println(cart)

	cart = append(cart[:index], cart[index+1:]...)
	return cart
}
