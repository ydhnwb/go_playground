package arr

import (
	"fmt"
	"strconv"
)

func arr() {
	var products []string
	fmt.Printf("%v\n", products)
	for i := 1; i <= 10; i++ {
		products = append(products, "Products "+strconv.Itoa(i))
	}
	fmt.Printf("len of products is %d, capaicty is %d %v\n", len(products), cap(products), products)

}
