package main

import (
	"fmt"
	"github.com/dhirajkashyap/Puppy"
)

func main() {
	var x int = 23
	fmt.Printf("value of x is %v, type is %T\n", x, x)

	s1 := puppy.Bark()
	fmt.Printf("Hello says Puppy %s \n", s1)
	s2 := puppy.Barks()
	fmt.Printf("Hello twice says Puppy %s\n", s2)
	s3 := puppy.Barkz()
	fmt.Printf("Hello says Puppy, I got a new Barkz %s\n", s3)
	s4 := puppy.BigBark()
	fmt.Printf("Hello says Puppy - When am big, I gonna say %s\n", s4)
}
