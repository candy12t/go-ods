package ods

import (
	"testing"
)

func TestDualArrayDeque(t *testing.T) {
	a := NewDualArrayDeque()
	tests := []int{1, 2, 3, 4, 5}
	for _, v := range tests {
		a.AddFirst(v)
	}

	for i := len(tests) - 1; i >= 0; i-- {
		got := a.RemoveFirst()
		if got != tests[i] {
			t.Errorf("a.RemoveFirst() returned %v, want %v\n", got, tests[i])
		}
	}
}
