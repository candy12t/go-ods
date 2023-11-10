package ods

type sLListNode struct {
	v    V
	next *sLListNode
}

func newSLListNode(v V) *sLListNode {
	return &sLListNode{
		v:    v,
		next: nil,
	}
}

// sLList is singly-linked list(単方向連結リスト)
type sLList struct {
	head *sLListNode
	tail *sLListNode
	len  int
}

var _ Stack = (*sLList)(nil)
var _ Queue = (*sLList)(nil)

func NewSLList() *sLList {
	return &sLList{}
}

func (l *sLList) Size() int {
	return l.len
}

// O(1)
func (l *sLList) Push(v V) {
	node := newSLListNode(v)
	node.next = l.head
	l.head = node

	if l.len == 0 {
		l.tail = node
	}
	l.len++
}

// O(1)
func (l *sLList) Pop() V {
	return l.RemoveFirst()
}

// O(1)
func (l *sLList) Enqueue(v V) bool {
	node := newSLListNode(v)
	if l.len == 0 {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.len++
	return true
}

// O(1)
func (l *sLList) Dequeue() V {
	return l.RemoveFirst()
}

func (l *sLList) RemoveFirst() V {
	if l.len == 0 {
		return nil
	}

	v := l.head.v
	l.head = l.head.next

	if l.len == 1 {
		l.tail = nil
	}
	l.len--

	return v
}
