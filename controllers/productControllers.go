package controllers

import (
	"ecommece/database"
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
