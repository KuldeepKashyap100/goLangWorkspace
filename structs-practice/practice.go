package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) printData() {
	fmt.Printf("Id: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: %.2f\n", prod.price)
}

func (prod *Product) storeData() {
	content := fmt.Sprintf("Id: %v\nTitle: %v\nDescription: %v\nPrice: %.2f\n", prod.id, prod.title, prod.description, prod.price)
	file, _ := os.Create(prod.id+".txt")
	file.WriteString(content)
}
func NewProduct(productId string, title string, description string, price float64) *Product {
	return &Product{productId, title, description, price}
}

func main() {
	createdProduct := getProduct()
	createdProduct.printData()
	createdProduct.storeData()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data")
	fmt.Println("----------------------------")

	reader := bufio.NewReader(os.Stdout)

	idInput := readUserInput(reader, "Product ID:")

	titleInput := readUserInput(reader, "Title:")

	descriptionInput := readUserInput(reader, "Description:")

	priceInput := readUserInput(reader, "Price:")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
