package models

import (
	"ecommece/database"
	"ecommece/entities"
	"fmt"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	// fmt.Println("okjhv")
	db, err := database.GetDB()
	// fmt.Println(db)

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		// rows := db.Find(&product)
		// fmt.Println(rows)
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Image)
				products = append(products, product)
			}
			// fmt.Println(products)
			return products, nil
		}
	}
}
func (*ProductModel) Find(id int64) (entities.Product, error) {
	// fmt.Println("okjhv")
	db, err := database.GetDB()
	// fmt.Println(db)

	if err != nil {
		return entities.Product{}, err
	} else {
		rows, err2 := db.Query("select * from product where id = ?", id)
		// rows := db.Find(&product)
		// fmt.Println(rows)
		if err2 != nil {
			return entities.Product{}, err2
		} else {

			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.Image)

			}
			fmt.Println(product)
			return product, nil
		}
	}
}
