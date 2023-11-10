package ods

type dLListNode struct {
	v    V
	prev *dLListNode
	next *dLListNode
}

func newDLListNode(v V) *dLListNode {
	return &dLListNode{v: v}
}

// dLList is doubly-linked list(双方向連結リスト)
type dLList struct {
	dummy *dLListNode
	len   int
}

var _ Stack = (*dLList)(nil)
var _ Queue = (*dLList)(nil)
var _ List = (*dLList)(nil)

func NewDLList() *dLList {
	dummy := new(dLListNode)
	dummy.prev = dummy
	dummy.next = dummy

	return &dLList{
		dummy: dummy,
		len:   0,
	}
}

func (l *dLList) Size() int {
	return l.len
}

// O(1+min(i, n-i))
func (l *dLList) getNode(idx int) *dLListNode {
	var node *dLListNode

	if idx < l.len/2 {
		node = l.dummy.next
		for i := 0; i < idx; i++ {
			node = node.next
		}
	} else {
		node = l.dummy
		for i := l.len; idx < i; i-- {
			node = node.prev
		}
	}

	return node
}

// O(1+min(i, n-i))
func (l *dLList) Get(idx int) V {
	return l.getNode(idx).v
}

// O(1+min(i, n-i))
func (l *dLList) Set(idx int, v V) V {
	node := l.getNode(idx)
	_v := node.v
	node.v = v
	return _v
}

// O(1+min(i, n-i))
func (l *dLList) Add(idx int, v V) {
	l.addBefore(l.getNode(idx), v)
}

func (l *dLList) addBefore(node *dLListNode, v V) {
	newNode := newDLListNode(v)
	newNode.prev = node.prev
	newNode.next = node
	newNode.prev.next = newNode
	newNode.next.prev = newNode
	l.len++
}

// O(1+min(i, n-i))
func (l *dLList) Remove(idx int) V {
	node := l.getNode(idx)
	node.prev.next = node.next
	node.next.prev = node.prev
	l.len--
	return node.v
}

func (l *dLList) Enqueue(v V) bool {
	l.Add(l.Size(), v)
	return true
}

func (l *dLList) Dequeue() V {
	return l.Remove(0)
}

func (l *dLList) Push(v V) {
	l.Add(0, v)
}

func (l *dLList) Pop() V {
	return l.Remove(0)
}
