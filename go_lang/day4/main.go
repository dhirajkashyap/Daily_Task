package main

import (
	"fmt"
)

type Person struct {
	first string
	age   int
}

func main() {
	//fmt.Println("Day4")
	x := Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Sum = %d, ", x)
	y := TestFoo(4)
	fmt.Printf("Foo return %d, ", y)
	z, s := TestBar(5)
	fmt.Printf("Bar returned %d, and %s.\n", z, s)
	fmt.Printf("Sum Foo1 = %d, Bar1 = %d\n", Foo1(x1...), Bar1(x1))
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
	p1 := Person{
		first: "James",
		age:   33,
	}
	p1.Speak()
}

func (p Person) Speak() {
	fmt.Println("My first name is", p.first, "and my age is", p.age)
}
func Sum(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}

func Foo1(ii ...int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
func Bar1(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
func TestFoo(i int) int {
	foo := i * i
	return foo
}

func TestBar(i int) (int, string) {
	bar := i * i
	var bar1 string
	bar1 = "True"
	return bar, bar1
}
