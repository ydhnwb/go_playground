package arr

import (
	"strconv"
)

// Array function just creating an instance of products
func Array() []string {
	var products []string
	for i := 1; i <= 10; i++ {
		products = append(products, "Products "+strconv.Itoa(i))
	}
	// fmt.Printf("len of products is %d, capaicty is %d %v\n", len(products), cap(products), products)
	return products
}
