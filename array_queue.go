package ods

// a[Idx], a[Idx+1], ... ,a[Idx+len-1]
// 0 <= i mod cap <= cap-1
// a[Idx%cap], a[(Idx+1)%cap], ... , a[(Idx+len-1)%cap]

type arrayQueue struct {
	len     int
	cap     int
	headIdx int
	buf     []V
}

var _ Queue = (*arrayQueue)(nil)

func NewArrayQeuue() *arrayQueue {
	return &arrayQueue{}
}

func (a *arrayQueue) Enqueue(v V) bool {
	return a.Add(v)
}

func (a *arrayQueue) Dequeue() V {
	return a.Remove()
}

// O(1)
func (a *arrayQueue) Add(v V) bool {
	if a.isFull() {
		a.resize()
	}

	a.setBuf(a.len, v)
	a.len++

	return true
}

// O(1)
func (a *arrayQueue) Remove() V {
	v := a.buf[a.headIdx]

	a.headIdx = (a.headIdx + 1) % a.cap
	a.len--

	if a.isSparse() {
		a.resize()
	}

	return v
}

func (a *arrayQueue) Size() int {
	return a.len
}

func (a *arrayQueue) resize() {
	newCap := max(2*a.len, 1)
	newBuf := make([]V, newCap)

	for i := 0; i < a.len; i++ {
		newBuf[i] = a.getBuf(i)
	}

	a.buf = newBuf
	a.cap = newCap
	a.headIdx = 0
}

func (a *arrayQueue) getIdx(idx int) int {
	return (a.headIdx + idx) % a.cap
}

func (a *arrayQueue) getBuf(idx int) V {
	return a.buf[a.getIdx(idx)]
}

func (a *arrayQueue) setBuf(idx int, v V) {
	a.buf[a.getIdx(idx)] = v
}

func (a *arrayQueue) isFull() bool {
	return a.len == a.cap
}

func (a *arrayQueue) isSparse() bool {
	return a.cap >= 3*a.len
}
