package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	userAge, err := getUserAge()

	if err != nil {
		fmt.Println(err)
		return
	}

	if userAge >= 30 && userAge < 50 {
		fmt.Println("You are eligible for our senior jobs")
	} else if userAge >= 18 {
		fmt.Println("Welcome to the club!")
	} else {
		fmt.Println("sorry, you are not eligible")
	}

	switch userAge {
	case 30:
		fmt.Println("First Option")
	case 40:
		fmt.Println("Second Option")
	default:
		fmt.Println("default option")
	}
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your age:")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	if err != nil {
		err = errors.New("invalid user age")
	}

	return int(userAge), err
}