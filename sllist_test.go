package ods

import (
	"testing"
)

func TestSLList_PushAndPop(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5}
	l := NewSLList()

	for _, v := range tests {
		l.Push(v)
	}
	if got := l.Size(); got != len(tests) {
		t.Errorf("l.Size() returned %d, want %d\n", got, len(tests))
	}

	for i := len(tests) - 1; i > 0; i-- {
		if got := l.Pop(); got != tests[i] {
			t.Errorf("l.Pop() returned %v, want %v\n", got, tests[i])
		}
	}
}

func TestSLList_EnqueueAndDequeue(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5}
	l := NewSLList()

	for _, v := range tests {
		l.Enqueue(v)
	}
	if got := l.Size(); got != len(tests) {
		t.Errorf("l.Size() returned %d, want %d\n", got, len(tests))
	}

	for _, v := range tests {
		if got := l.Dequeue(); got != v {
			t.Errorf("l.Pop() returned %v, want %v\n", got, v)
		}
	}
}
