package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kuldeep/bmiCalculator/info"

)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPropmt)
	height = getUserInput(info.HeightPrompt)
	return
}


func getUserInput(promptText string) float64 {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	enteredValue, _ := strconv.ParseFloat(userInput, 32)
	return enteredValue
}