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
	for i := 0; i < 120; {
		fmt.Printf("\t%d", i)
		i++
	}

	xs := []string{"Aakhri Rasta", "Bees Saal Baad", "Chandni", "Deewar", "Ek Kahani"}
	fmt.Println(xs)
	fmt.Printf("Type %T\n", xs)
	for _, v := range xs {
		fmt.Printf("%v\t", v)
	}
	fmt.Printf("\n")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])
	fmt.Println(xs[4])
	xs = append(xs, "Farishtey", "Golmal", "Hero Number 1", "Insanyat")
	for i := 0; i < len(xs); i++ {
		fmt.Printf("All Movie[%d] = %s \n", i, xs[i])
	}
	zs := xs[4:7]
	for i := 0; i < len(xs[4:7]); i++ {
		fmt.Printf("New X Movie[%d] = %s \n", i, xs[4:7][i])
		fmt.Printf("New Z Movie[%d] = %s \n", i, zs[i])
	}

	fmt.Println("Newer X Movie = %s", xs[4:])
	fmt.Println("Newest X Movie = %s", xs)

	ys := xs[0:3]
	for i := 0; i < len(ys); i++ {
		fmt.Printf("Original Y Movie[%d] = %s \n", i, ys[i])
	}

	as := append(xs[:4], xs[5:]...)
	for i := 0; i < len(as); i++ {
		fmt.Printf("New deleting slice Movie[%d] = %s \n", i, as[i])
	}

	ii := make([]int, 0, 10)
	fmt.Println(ii)
	fmt.Println(len(ii))
	fmt.Println(cap(ii))

	ii = append(ii, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(ii)
	fmt.Println(len(ii))
	fmt.Println(cap(ii))

	ii = append(ii, 10, 11, 12, 13, 14)
	fmt.Println(ii)
	fmt.Println(len(ii))
	fmt.Println(cap(ii))

	jb := []string{"James", "Bond", "Martini", "BMW"}
	mp := []string{"Money", "Penny", "Mojito", "Audi"}
	bond007 := [][]string{jb, mp}
	fmt.Println(jb)
	fmt.Println(mp)
	fmt.Println(bond007)

	a := []int{0, 1, 2, 3, 4, 5, 6}
	//b := a
	b := make([]int, 7)
	copy(b, a)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	b[0] = 8
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
