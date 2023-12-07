package main

import "fmt"

var sum int

func main() {
	sum = 0
	func() {
		for i := 0; i < 5; i++ {
			sum += i
		}
	}()
	fmt.Println("Sum is", sum)
	x()
	fmt.Println("Now Sum is", sum)
	f := outer()
	fmt.Println("Outer is bakwas", f())

}

var x = func() {
	for i := 5; i < 10; i++ {
		sum += i
	}
}

func outer() func() int {
	return func() int {
		return 42
	}
}
