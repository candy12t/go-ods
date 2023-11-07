package ods

type arrayStack struct {
	len int
	cap int
	buf []V
}

var _ List = (*arrayStack)(nil)
var _ Stack = (*arrayStack)(nil)

func NewArrayStack() *arrayStack {
	return &arrayStack{}
}

func (a *arrayStack) Size() int {
	return a.len
}

// O(1)
func (a *arrayStack) Get(i int) V {
	return a.buf[i]
}

// O(1)
func (a *arrayStack) Set(i int, v V) V {
	_v := a.buf[i]
	a.buf[i] = v
	return _v
}

// O(n-i)
func (a *arrayStack) Add(i int, v V) {
	if a.isFull() {
		a.resize()
	}

	for j := a.len; j > i; j-- {
		a.buf[j] = a.buf[j-1]
	}
	a.buf[i] = v
	a.len++
}

// O(n-i)
func (a *arrayStack) Remove(i int) V {
	v := a.buf[i]
	for j := i; j < a.len-1; j++ {
		a.buf[j] = a.buf[j+1]
	}
	a.len--

	if a.isSparse() {
		a.resize()
	}

	return v
}

// O(1)
func (a *arrayStack) Push(v V) {
	a.Add(a.Size(), v)
}

// O(1)
func (a *arrayStack) Pop() V {
	return a.Remove(a.Size() - 1)
}

func (a *arrayStack) resize() {
	newCap := max(2*a.len, 1)
	newBuf := make([]V, newCap)
	for i := 0; i < a.len; i++ {
		newBuf[i] = a.buf[i]
	}
	a.buf = newBuf
	a.cap = newCap
}

func (a *arrayStack) isFull() bool {
	return a.len == a.cap
}

func (a *arrayStack) isSparse() bool {
	return a.cap >= 3*a.len
}
