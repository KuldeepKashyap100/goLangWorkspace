package main

import "fmt"

func main() {
	/* slice functionality */
	// var productNames [4]string = [4]string{"A Book"}
	// productNames[2] = "A Carpet"

	// prices := [4]float64{1.0, 2.99, 3.99, 4.0}

	// featuredPrices := prices[1:]
	// featuredPrices[0] = 199.99
	// highlightedPrices := featuredPrices[:1]
	// fmt.Println(prices, productNames, featuredPrices)
	// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// highlightedPrices = highlightedPrices[:3]
	// fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))

	/* dynamic list */
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	updatedPrices := append(prices, 5.99, 6.99, 7.99, 8.99)
	fmt.Println(updatedPrices, prices)

	discountPrices := []float64{101.99, 80.99, 100.10}
	prices = append(prices, discountPrices...) 
	fmt.Println(prices)
}
