package ods

import (
	"testing"
)

func TestArrayQueue(t *testing.T) {
	a := NewArrayQeuue()
	tests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range tests {
		a.Enqueue(v)
	}
	if a.Size() != len(tests) {
		t.Errorf("a.Size() returned %d, want %d\n", a.Size(), len(tests))
	}

	for _, v := range tests {
		got := a.Dequeue()
		if got != v {
			t.Errorf("a.Dequeue() returned %d, want %d\n", got, v)
		}
	}
	if a.Size() != 0 {
		t.Errorf("a.Size() returned %d, want %d\n", a.Size(), 0)
	}
}
