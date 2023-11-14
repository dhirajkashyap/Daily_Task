package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`Bond_James`] = []string{`Shaken, Not stirred`, `martinis`, `fast cars`}
	m[`Penny_Money`] = []string{`intelligence`, `literature`, `computer science`}
	m[`No_Dr`] = []string{`Cats`, `icecreams`, `Mini Me`}

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
	m[`No_Dr_1`] = []string{`Dogs`, `Hot Dogs`, `Mini You`}
	for k, v := range m {
		fmt.Println(k)
		for l, v2 := range v {
			fmt.Printf("%v\t%v\n", l, v2)
		}
	}
}
