package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m[`Bond_James`] = []string{`Shaken, Not stirred`, `martinis`, `fast cars`}
	m[`Penny_Money`] = []string{`intelligence`, `literature`, `computer science`}
	m[`No_Dr`] = []string{`Cats`, `icecreams`, `Mini Me`}

	fmt.Printf("%#v\n", m)

}
