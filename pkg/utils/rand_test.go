package utils

import "testing"

func TestRandInt(t *testing.T) {
	number := RandInt(1, 5)
	if number < 1 || number > 5 {
		t.Fatalf("number should be between 1 and 5, got: '%d'", number)
	}
}

func TestRandIntReverse(t *testing.T) {
	number := RandInt(5, 1)
	if number < 1 || number > 5 {
		t.Fatalf("number should be between 1 and 5, got: '%d'", number)
	}
}
