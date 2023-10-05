package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("Random value is %d\n", x)

	if x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}
}
