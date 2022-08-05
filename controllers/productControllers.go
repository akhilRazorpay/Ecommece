package controllers

import (
	"ecommece/database"
	"ecommece/models"
	"fmt"
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
func Delete(w http.ResponseWriter, r *http.Request) {

	db, err := database.GetDB()
	if err != nil {
		return
	}
	product := r.URL.Query().Get("id")
	// fmt.Println(product)

	delForm, err := db.Prepare("DELETE FROM product WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(product)
	// log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Insert(response http.ResponseWriter, request *http.Request) {
	data := ""
	tmp, _ := template.ParseFiles("views/insert.html")
	tmp.Execute(response, data)
}

func Inserted(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	db, err := database.GetDB()
	if err != nil {
		return
	}
	if r.Method == "POST" {

		name := r.FormValue("name")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		image := r.FormValue("image")
		description := r.FormValue("description")
		fmt.Println(name, price, quantity, image, description)
		// fmt.Println(name, city)
		// insForm, err := db.Prepare("INSERT INTO product(name, city) VALUES(?,?)")
		insForm, err := db.Prepare("INSERT INTO product (name, price, quantity, image, description) VALUES(?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("oks1")
		insForm.Exec(name, price, quantity, image, description)
		fmt.Println("oks2")
	}
	fmt.Println("oks")
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
