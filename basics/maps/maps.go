package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://google.com",
	}
	fmt.Println(websites, websites["google"])
}
