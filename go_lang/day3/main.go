package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
	favIC string
}

func main() {
	m := make(map[string][]string)
	m[`Bond_James`] = []string{`Shaken, Not stirred`, `martinis`, `fast cars`, `Raspberry-Coconut`}
	m[`Penny_Money`] = []string{`intelligence`, `literature`, `computer science`, `Mango-lime`}
	m[`No_Dr`] = []string{`Cats`, `icecreams`, `Mini Me`, `Jalapeno-Orange`}

	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Println(k)
		for l, v2 := range v {
			fmt.Printf("%v\t%v\n", l, v2)
		}
	}
	fmt.Println("Now deleting record")
	delete(m, "No_Dr")
	for k, v := range m {
		fmt.Println(k)
		for l, v2 := range v {
			fmt.Printf("%v\t%v\n", l, v2)
		}
	}
	fmt.Println("Adding new record now!")
	m[`No_Dr_1`] = []string{`Dogs`, `Hot Dogs`, `Mini You`, `Pineapple-Ghost-Pepper`}
	for k, v := range m {
		fmt.Println(k)
		for l, v2 := range v {
			fmt.Printf("%v\t%v\n", l, v2)
		}
	}

	p1 := struct {
		first string
		last  string
		age   int
		favIC string
	}{
		first: "James",
		last:  "Bond",
		age:   32,
		favIC: "Raspberry-Coconut",
	}
	p2 := person{
		first: "Money",
		last:  "Penny",
		age:   25,
		favIC: "Mango-Lime",
	}
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

	p2 = p1

	fmt.Println("After initialization")
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)
}
