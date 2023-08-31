package main

import "fmt"

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const (
	_  = iota
	c4 = (1 << (iota * 10))
	c5
	c6
	c7
)

func main() {
	var i int
	fmt.Printf("i = %d\n", i)
	i = 40
	fmt.Println("Hello Gophers!")
	const name, age = "Dhiraj", 45
	fmt.Printf("Name %s is %d years old i %d\n", name, age, i)
	fmt.Println(c1, c2, c3)
	fmt.Println(c4, c5, c6, c7)
	fmt.Printf("%b %b %b %b\n", c4, c5, c6, c7)
}
