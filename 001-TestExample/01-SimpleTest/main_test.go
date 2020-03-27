package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(5, 6, 7)
	if x != 18 {
		t.Error("Expected", 18, "Got", x)
	}
}
