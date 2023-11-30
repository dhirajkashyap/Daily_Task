package main

import (
	"fmt"
	"testing"
)

func testAdd(testing.T) {
	total := Add(5, 5)
	if total != 10 {
		fmt.Printf("Result Failed: Got %d, expected 10", total)
	}
}
