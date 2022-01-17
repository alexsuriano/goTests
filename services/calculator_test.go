package services

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 2) != 4 {
		t.Error("2 + 2 must be 4")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Error("2 * 3 must be 6")
	}
}
