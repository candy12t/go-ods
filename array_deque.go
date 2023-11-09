package ods

type arrayDeque struct {
	len     int
	cap     int
	headIdx int
	buf     []V
}

var _ List = (*arrayDeque)(nil)
var _ Deque = (*arrayDeque)(nil)

func NewArrayDeque() *arrayDeque {
	return &arrayDeque{}
}

func (a *arrayDeque) Size() int {
	return a.len
}

// O(1)
func (a *arrayDeque) Get(idx int) V {
	return a.getBuf(idx)
}

// O(1)
func (a *arrayDeque) Set(idx int, v V) V {
	_v := a.getBuf(idx)
	a.setBuf(idx, v)
	return _v
}

// O(1 + min(i, n-i))
func (a *arrayDeque) Add(idx int, v V) {
	if a.isFull() {
		a.resize()
	}

	if idx < a.len/2 { // a.buf[0], ..., a.buf[idx-1]を左にずらす
		a.headIdx = a.modHeadIdx()
		for i := 0; i < idx; i++ {
			a.setBuf(i, a.getBuf(i+1))
		}
	} else { // a.buf[idx], ..., a.buf[a.len-1]を右にずらす
		for i := a.len; i > idx; i-- {
			a.setBuf(i, a.getBuf(i-1))
		}
	}

	a.setBuf(idx, v)
	a.len++
}

// O(1 + min(i, n-i))
func (a *arrayDeque) Remove(idx int) V {
	v := a.getBuf(idx)
	a.setBuf(idx, nil)

	if idx < a.len/2 { // a.buf[0], ..., a.buf[idx-1]を右にずらす
		for i := idx; i > 0; i-- {
			a.setBuf(i, a.getBuf(i-1))
		}
		a.headIdx = (a.headIdx + 1) % a.cap
	} else { // a.buf[idx+1], ..., a.buf[a.len-1]を左にずらす
		for i := idx; i < a.len-1; i++ {
			a.setBuf(i, a.getBuf(i+1))
		}
	}

	a.len--
	if a.isSparse() {
		a.resize()
	}

	return v
}

func (a *arrayDeque) AddFirst(v V) {
	a.Add(0, v)
}

func (a *arrayDeque) AddLast(v V) {
	a.Add(a.Size(), v)
}

func (a *arrayDeque) RemoveFirst() V {
	return a.Remove(0)
}

func (a *arrayDeque) RemoveLast() V {
	return a.Remove(a.Size() - 1)
}

func (a *arrayDeque) resize() {
	newCap := max(2*a.len, 1)
	newBuf := make([]V, newCap)

	for i := 0; i < a.len; i++ {
		newBuf[i] = a.getBuf(i)
	}

	a.buf = newBuf
	a.cap = newCap
	a.headIdx = 0
}

func (a *arrayDeque) getIdx(idx int) int {
	return (a.headIdx + idx) % a.cap
}

func (a *arrayDeque) getBuf(idx int) V {
	return a.buf[a.getIdx(idx)]
}

func (a *arrayDeque) setBuf(idx int, v V) {
	a.buf[a.getIdx(idx)] = v
}

func (a *arrayDeque) isFull() bool {
	return a.len == a.cap
}

func (a *arrayDeque) isSparse() bool {
	return a.cap >= 3*a.len
}

func (a *arrayDeque) modHeadIdx() int {
	if a.headIdx == 0 {
		return a.cap - 1
	} else {
		return a.headIdx - 1
	}
}
