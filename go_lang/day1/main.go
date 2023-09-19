package main

import (
	"fmt"
	"github.com/dhirajkashyap/Puppy"
)

func main() {
	var x int = 23
	fmt.Printf("value of x is %v, type is %T\n", x, x)

	s1 := puppy.Bark()
	fmt.Printf("Hello says %s \n", s1)
	s2 := puppy.Barks()
	fmt.Printf("Hello twice says %s\n", s2)

}
