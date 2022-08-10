package models_test

import (
	"ecommece/entities"
	"ecommece/models"
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

// func IntMin(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
func TestFind(t *testing.T) {
	var tests = []struct {
		a    int64
		want entities.Product
	}{
		{1, entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}},
		// this will fail
		{0, entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}},
		{3, entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}},
		{7, entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}},
		// {-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			var productModel models.ProductModel

			product, _ := productModel.Find(tt.a)
			// fmt.Println(product)
			if product != tt.want {
				t.Errorf("got%+v, want%+v ", product, tt.want)
			}
		})
	}
}
func TestFindAll(t *testing.T) {
	var tests = []struct {
		a    string
		want []entities.Product
	}{
		{"", []entities.Product{{Id: 5, Name: "Laptop 3", Price: 60000, Quantity: 11, Description: "https://www.cnet.com/a/img/resize/4e82f3a17554a5aff8089194237de5a3acfce3b4/2022/04/27/b796792b-5b34-4405-83eb-efc66371ee06/samsung-galaxy-book-2-pro-360-01.jpg?auto=webp&fit=crop&height=630&width=1200", Image: "LAPTOP3 DESCRIPTION"},
			{Id: 6, Name: "Mobile 1", Price: 12000, Quantity: 20, Description: "https://cdn-www.mediatek.com/page/Mobile-2_2021-10-20-155734_vspa.png", Image: "Mobile1 DESCRIPTION"}, {Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"},
			{Id: 2, Name: "Tv 2", Price: 25000, Quantity: 6, Description: "https://images.philips.com/is/image/PhilipsConsumer/58PUT6604_94-IMS-en_IN?wid=840&hei=720&$jpglarge$", Image: "TV2 DESCRIPTION"},
			{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, {Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"},
			{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"},
			{Id: 18, Name: "tv2", Price: 12000, Quantity: 22, Description: "https://images.indianexpress.com/2021/10/Redmi-TV-review-1.jpg", Image: "tv2 description"}}},
		// this will fail

	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.a)
		t.Run(testname, func(t *testing.T) {
			var productModel models.ProductModel
			var products []entities.Product

			products, _ = productModel.FindAll()
			// fmt.Println(products)
			if !slices.Equal(products, tt.want) {
				t.Errorf("got%+v, want%+v ", products, tt.want)
			}
		})
	}
}
