package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Add Failed: Got %d, expected 10", total)
	} else {
		t.Logf("Add Passed: Got %d, expected 10", total)
	}
	total1 := Add(6, 5)
	if total1 != 10 {
		t.Errorf("Add Failed: Got %d, expected 10", total1)
	} else {
		t.Logf("Add Passed: Got %d, expected 10", total1)
	}
}
func TestSubtract(t *testing.T) {
	total := Subtract(6, 15)
	if total != 9 {
		t.Errorf("Subtract Failed: Got %d, expected 9", total)
	} else {
		t.Logf("Subtract Passed: Got %d, expected 9", total)
	}
	total1 := Subtract(8, 15)
	if total1 != 1 {
		t.Errorf("Subtract Failed: Got %d, expected 1", total1)
	} else {
		t.Logf("Subtract Passed: Got %d, expected 1", total1)
	}

}
