package ods

import (
	"testing"
)

func TestFastArrayStack(t *testing.T) {
	a := NewFastArrayStack()

	tests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range tests {
		a.Push(v)
	}
	if a.Size() != len(tests) {
		t.Errorf("a.Size() returned %d, want %d", a.Size(), len(tests))
	}

	for i := len(tests) - 1; i >= 0; i-- {
		got := a.Pop()
		if want := tests[i]; got != want {
			t.Errorf("a.Pop() returned %d, want %d", got, want)
		}
	}
	if a.Size() != 0 {
		t.Fatalf("a.Size() returned %d, want %d", a.Size(), 0)
	}
}
