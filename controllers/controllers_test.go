package controllers_test

import (
	"ecommece/controllers"
	"ecommece/entities"

	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestExists(t *testing.T) {
	var tests = []struct {
		a    int64           // input
		b    int             // output
		want []entities.Item // input
	}{
		{0, -1, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},
		{1, 0, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},
		{01, -1, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}}, // {1, 0, 0},

	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {

			ans := controllers.Exists(tt.a, tt.want)
			if ans != tt.b {
				t.Errorf("got %d, want %d", ans, tt.b)
			}
		})
	}
}

func TestTotal(t *testing.T) {
	var tests = []struct {
		a    float64         // output
		want []entities.Item // input
	}{
		{313300, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},
		{292300, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},
		{01, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}}, // {1, 0, 0},
		{0, []entities.Item{}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%g", tt.a)
		t.Run(testname, func(t *testing.T) {

			ans := controllers.Total(tt.want)
			if ans != tt.a {
				t.Errorf("got %g, want %g", ans, tt.a)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		a    int             // input
		b    []entities.Item // output
		want []entities.Item // input
	}{
		{1, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}},
			[]entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
				{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
				{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
				{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},

		{0, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
			{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}},
			[]entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
				{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
				{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
				{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},

		{3, []entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
			{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
			{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1}},
			[]entities.Item{{Product: entities.Product{Id: 1, Name: "Tv 1", Price: 20000, Quantity: 4, Description: "https://images.samsung.com/is/image/samsung/in-full-hd-tv-te50fa-ua43te50fakxxl-frontblack-231881877?$650_519_PNG$", Image: "TV1 DESCRIPTION"}, Quantity: 4},
				{Product: entities.Product{Id: 4, Name: "Laptop 2", Price: 54100, Quantity: 6, Description: "https://cdn.thewirecutter.com/wp-content/uploads/2020/04/laptops-lowres-2x1-.jpg?auto=webp&quality=75&crop=2:1&width=1024", Image: "LAPTOP2 DESCRIPTION"}, Quantity: 3},
				{Product: entities.Product{Id: 7, Name: "Mobile 2", Price: 21000, Quantity: 25, Description: "https://img.tatacliq.com/images/i8/437Wx649H/MP000000013198982_437Wx649H_202205211422291.jpeg", Image: "Mobile2 DESCRIPTION"}, Quantity: 1},
				{Product: entities.Product{Id: 3, Name: "Laptop 1", Price: 50000, Quantity: 11, Description: "https://m.media-amazon.com/images/I/718ETwvLVOL._SL1500_.jpg", Image: "LAPTOP1 DESCRIPTION"}, Quantity: 1}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {

			ans := controllers.RemoveCart(tt.want, tt.a)
			if !slices.Equal(ans, tt.b) {
				t.Errorf("got %+v, want %+v", ans, tt.b)
			}
		})
	}
}
