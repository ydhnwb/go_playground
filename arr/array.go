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

// RemoveWithKeepOrder is removing an element while maintains the order
func RemoveWithKeepOrder(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// Remove the element of an array
func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
