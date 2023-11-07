package ods

type V any

// Queue is FIFO (first-in-first-out) queue
type Queue interface {
	Enqueue(V) bool
	Dequeue() V
}

// Stack is LIFO (last-in-first-out) queue
type Stack interface {
	Push(V)
	Pop() V
}

type Deque interface {
	AddFirst(V)
	AddLast(V)
	RemoveFirst() V
	RemoveLast() V
}

type List interface {
	Size() int
	Get(int) V
	Set(int, V) V
	Add(int, V)
	Remove(int) V
}

// USet is unordered Set
type USet interface {
	Size() int
	Add(int, V)
	Remove(int) V
	Find(V) V
}

// SSet is sorted Set
type SSet interface {
	Size() int
	Add(int, V)
	Remove(int) V
	Find(V) V
}
