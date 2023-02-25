package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
		return
	}

	if selectedChoice == "1" {
		calculateSumUpToNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumManually()
	} else {
		calculateListSum()
	}
}

func calculateSumUpToNumber() {
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	sum := 0
	for i := 1; i <= chosenNumber; i++ {
		sum += i
	}
	fmt.Printf("Result: %v", sum)
}

func calculateFactorial() {
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	factorial := 1
	for i := chosenNumber; i >= 1; i-- {
		factorial *= i
	}
	fmt.Printf("Result: %v", factorial)
}

func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0
	for isEnteringNumbers {
		enteredNum, err := getInputNumber()
		sum += enteredNum
		isEnteringNumbers = err == nil
	}
	fmt.Printf("Result: %v", sum)
}

func calculateListSum() {
	fmt.Print("Please enter your number: ")
	inputNumberList, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, " ")

	sum := 0

	// for i := 0; i < len(inputNumbers); i++ {
	// 	currentInputNum, err := strconv.ParseInt(inputNumbers[i], 0, 64)
	// 	if err != nil {
	// 		break
	// 	}
	// 	sum += int(currentInputNum)
	// }
	for _, value := range inputNumbers {
		currentInputNum, _ := strconv.ParseInt(value, 0, 64)
		if err != nil {
			continue
		}
		sum += int(currentInputNum)
	}
	fmt.Printf("Result: %v", sum)
}

func getInputNumber() (int, error) {
	fmt.Print("Please enter your number: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)
	inputNumber, err := strconv.ParseInt(userInput, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(inputNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Print("Please make your choice: ")
	userInput, err := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	if err != nil {
		return "", err
	}

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}
