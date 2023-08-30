package main

import "fmt"

func main() {
	var i int
	fmt.Printf("i = %d\n", i)
	i = 40
	fmt.Println("Hello Gophers!")
	const name, age = "Dhiraj", 45
	fmt.Printf("Name %s is %d years old i %d\n", name, age, i)
}
