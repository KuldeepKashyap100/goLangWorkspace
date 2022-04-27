package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kuldeep/bmiCalculator/info"
)



func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Seperator)
	fmt.Print(info.WeightPropmt)
	weightInput, _ := reader.ReadString('\n');
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n');

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 32)
	height, _ := strconv.ParseFloat(heightInput, 32)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)

}