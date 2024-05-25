package main

import "fmt"

func main() {
	var name string

	fmt.Println("What is your name? ")
	fmt.Scanln(&name)
	output := fmt.Sprintf("Hello, %s, nice to meet you!", name)

	fmt.Println(output)

}
