package main

import "fmt"

func main() {
	fmt.Println("Hello Gophers!")
	const name, age = "Dhiraj", 45
	fmt.Printf("Name %s is %d years old\n", name, age)
	fmt.Println(`
		This is raw string output - 😊. - Trying with emojis
		`)
}
