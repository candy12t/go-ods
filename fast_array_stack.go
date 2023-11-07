package ods

type fastArrayStack struct {
	len int
	cap int
	buf []V
}

var _ List = (*fastArrayStack)(nil)
var _ Stack = (*fastArrayStack)(nil)

func NewFastArrayStack() *fastArrayStack {
	return &fastArrayStack{}
}

func (a *fastArrayStack) Size() int {
	return a.len
}

// O(1)
func (a *fastArrayStack) Get(i int) V {
	return a.buf[i]
}

// O(1)
func (a *fastArrayStack) Set(i int, v V) V {
	_v := a.buf[i]
	a.buf[i] = v
	return _v
}

// O(n-i)
func (a *fastArrayStack) Add(i int, v V) {
	if a.isFull() {
		a.resize()
	}

	// for j := a.len; j > i; j-- {
	// 	a.buf[j] = a.buf[j-1]
	// }
	copy(a.buf[i+1:a.len+1], a.buf[i:a.len])
	a.buf[i] = v
	a.len++
}

// O(n-i)
func (a *fastArrayStack) Remove(i int) V {
	v := a.buf[i]
	// for j := i; j < a.len-1; j++ {
	// 	a.buf[j] = a.buf[j+1]
	// }
	copy(a.buf[i:a.len-1], a.buf[i+1:a.len])
	a.len--

	if a.isSparse() {
		a.resize()
	}

	return v
}

// O(1)
func (a *fastArrayStack) Push(v V) {
	a.Add(a.Size(), v)
}

// O(1)
func (a *fastArrayStack) Pop() V {
	return a.Remove(a.Size() - 1)
}

func (a *fastArrayStack) resize() {
	newCap := max(2*a.len, 1)
	newBuf := make([]V, newCap)
	// for i := 0; i < a.len; i++ {
	// 	newBuf[i] = a.buf[i]
	// }
	copy(newBuf, a.buf)
	a.buf = newBuf
	a.cap = newCap
}

func (a *fastArrayStack) isFull() bool {
	return a.len == a.cap
}

func (a *fastArrayStack) isSparse() bool {
	return a.cap >= 3*a.len
}
