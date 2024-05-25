package main

import "fmt"

func main() {
	var String string

	fmt.Println("What is the input string: ")
	fmt.Scanln(&String)

	lenght := len(String)

	resultado := fmt.Sprintf("%s has %d characters", String, lenght)

	fmt.Println(resultado)

}
