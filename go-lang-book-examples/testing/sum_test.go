package main 

import "testing"

func TestSum(t *testing.T) {
	var result int
	n1 := 10
	n2 := 10
	expected := 15

	result = sum(n1, n2)

	if result != expected {
		t.Error("Expect 15, got", result)
	}
}