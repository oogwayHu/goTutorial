// Instead of testing with one piece of data
// This takes multiple input data and expected answer
// This gives a more comprehensive test

package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 6, 10}, 18},
		test{[]int{52, 40}, 92},
		test{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "But calculated", x)
		}
	}
}
