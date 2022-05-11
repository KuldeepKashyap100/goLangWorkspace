package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, func(number int) int {
		return number * 3
	})
	fmt.Println(doubled, tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	updatedNumbers := []int{}

	for _, number := range *numbers {
		updatedNumbers = append(updatedNumbers, transform(number))
	}
	return updatedNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// closure functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}


