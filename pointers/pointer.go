package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age)

	myAge := &age

	*myAge = 66

	fmt.Println(age, myAge, *myAge)

	newAge := doubleAge(myAge)

	fmt.Println(newAge)
}

func doubleAge(age *int) int {
	return *age * 2
}
