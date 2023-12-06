package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Result Failed: Got %d, expected 10", total)
	} else {
		t.Logf("Result Passed: Got %d, expected 10", total)
	}
	total1 := Add(6, 5)
	if total1 != 10 {
		t.Errorf("Result Failed: Got %d, expected 10", total1)
	} else {
		t.Logf("Result Passed: Got %d, expected 10", total1)
	}

}
