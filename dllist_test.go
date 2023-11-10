package ods

import (
	"testing"
)

func TestDLList_NewDLList(t *testing.T) {
	l := NewDLList()
	wantSize := 0

	if l.Size() != wantSize {
		t.Errorf("l.Size() returned %d, want %d\n", l.Size(), wantSize)
	}

	if l.dummy != l.dummy.prev {
		t.Errorf("l.dummy.prev is not dummy(%v), got is %v\n", l.dummy, l.dummy.prev)
	}

	if l.dummy != l.dummy.next {
		t.Errorf("l.dummy.next is not dummy(%v), got is %v\n", l.dummy, l.dummy.next)
	}
}

func TestDLList_Get(t *testing.T) {
	wantValue := 1

	l := NewDLList()
	node := newDLListNode(wantValue)
	node.next = l.dummy
	node.prev = l.dummy
	l.dummy.next = node
	l.dummy.prev = node
	l.len++

	if got := l.Get(0); got != wantValue {
		t.Errorf("l.Get(0) returned %v, want %v\n", got, wantValue)
	}
}

func TestDLList_Add(t *testing.T) {
	l := NewDLList()
	wantSize := 1
	wantValue := 1
	l.Add(0, wantValue)

	if l.Size() != wantSize {
		t.Errorf("l.Size() returned %d, want %d\n", l.Size(), wantSize)
	}

	if got := l.dummy.next.v; got != wantValue {
		t.Errorf("l.dummy.next.v returned %v, want %v\n", got, wantValue)
	}

	if got := l.dummy.prev.v; got != wantValue {
		t.Errorf("l.dummy.prev.v returned %v, want %v\n", got, wantValue)
	}
}

func TestDLList_PushAndPop(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5}
	l := NewDLList()

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

func TestDLList_EnqueueAndDequeue(t *testing.T) {
	tests := []int{1, 2, 3, 4, 5}
	l := NewDLList()

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
