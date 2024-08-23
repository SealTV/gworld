package main

import "testing"

func Test_Sum(t *testing.T) {
	a := 1
	b := 2

	if a+b != sum(a, b) {
		t.Error("Sum failed")
	}
}
