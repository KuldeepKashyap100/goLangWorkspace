package info

import "fmt"

const mainTitle = "BMI Calculator"
const seperator = "------------------"
const WeightPropmt = "Please enter your weight (KG): "
const HeightPrompt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(seperator)
}