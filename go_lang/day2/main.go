package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is during initialization.")
}
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

	switch {
	case x <= 100:
		fmt.Println("Case between 0 and 100")
		break
	case x <= 200:
		fmt.Println("Case Between 101 and 200")
	default:
		fmt.Println("Case Between 201 and 500, or maybe more")
	}
	x = rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x = %d, y = %d\n", x, y)

	if (x < 4) && (y < 4) {
		fmt.Println("x and y are less than 4")
	}
	if (x > 6) && (y > 6) {
		fmt.Println("x and y are greater than 6")
	}
	if (x >= 4) && (x < 6) {
		fmt.Println("x is greater than equal to 4 and less than 6")
	}

	if y != 5 {
		fmt.Println("y aint 5 yall")
	}

	switch {
	case (x < 4) && (y < 4):
		fmt.Println("Case x & y are less than 4")
		break
	case (x > 6) && (y > 6):
		fmt.Println(" Case x & y are greater then 6")
		break
	case (x >= 4 && x < 6):
		fmt.Println("Case x greater than equal to 4 & less than 6")
		break
	case (y != 5):
		fmt.Println("Case y aint 5 yall")
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("\t%d", i)
		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}

}
