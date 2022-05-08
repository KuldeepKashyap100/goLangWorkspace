package main

import "fmt"

type Person struct {
	name string
	age  int
}

type PersonData map[string]Person

type customNumber int

func (number customNumber) pow(power int) customNumber {
	result := number

	for i := 1; i < power; i++ {
		result *= number
	}
	return result
}

func main() {
	hobbies := []string{"Cooking", "Dancing"}
	fmt.Println(len(hobbies), cap(hobbies))

	hobbies = make([]string, 2, 10)
	fmt.Println(len(hobbies), cap(hobbies))

	var people map[string]Person = map[string]Person{
		"P1": {"Max", 32},
	}

	var shorterPeople PersonData = PersonData{
		"P1": {"kuldeep", 27},
	}

	fmt.Println(people, shorterPeople)

	var dummyNumber customNumber = 5
	// or  dummyNumber := customNumber(5)
	fmt.Println(dummyNumber.pow(3))
}
