package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The resulst must be 5")
	}
}

func TestSub(t *testing.T) {

	result := sub(2, 2)

	if result != 0 {
		t.Error("The resulst must be 0")
	}
}

func TestMult(t *testing.T) {

	result := mult(2, 2)

	if result != 4 {
		t.Error("The resulst must be 4")
	}
}
