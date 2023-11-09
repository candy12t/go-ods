package ods

type dualArrayDeque struct {
	front arrayStack
	back  arrayStack
}

var _ List = (*dualArrayDeque)(nil)
var _ Deque = (*dualArrayDeque)(nil)

func NewDualArrayDeque() *dualArrayDeque {
	return &dualArrayDeque{}
}

func (a *dualArrayDeque) Size() int {
	return a.front.Size() + a.back.Size()
}

// O(1)
func (a *dualArrayDeque) Get(idx int) V {
	if idx < a.front.Size() {
		return a.front.Get(a.getFrontIdx(idx))
	} else {
		return a.back.Get(a.getBackIdx(idx))
	}
}

// O(1)
func (a *dualArrayDeque) Set(idx int, v V) V {
	if idx < a.front.Size() {
		return a.front.Set(a.getFrontIdx(idx), v)
	} else {
		return a.back.Set(a.getBackIdx(idx), v)
	}
}

// O(1 + min(i, n-i))
func (a *dualArrayDeque) Add(idx int, v V) {
	if idx < a.front.Size() {
		a.front.Add(a.getFrontIdx(idx)+1, v)
	} else {
		a.back.Add(a.getBackIdx(idx), v)
	}
	a.balance()
}

// O(1 + min(i, n-i))
func (a *dualArrayDeque) Remove(idx int) V {
	var v V
	if idx < a.front.Size() {
		v = a.front.Remove(a.getFrontIdx(idx))
	} else {
		v = a.back.Remove(a.getBackIdx(idx))
	}
	a.balance()
	return v
}

func (a *dualArrayDeque) AddFirst(v V) {
	a.Add(0, v)
}

func (a *dualArrayDeque) AddLast(v V) {
	a.Add(a.Size(), v)
}

func (a *dualArrayDeque) RemoveFirst() V {
	return a.Remove(0)
}

func (a *dualArrayDeque) RemoveLast() V {
	return a.Remove(a.Size() - 1)
}

func (a *dualArrayDeque) getFrontIdx(idx int) int {
	return a.front.Size() - idx - 1
}

func (a *dualArrayDeque) getBackIdx(idx int) int {
	return idx - a.front.Size()
}

func (a *dualArrayDeque) balance() {
	if 3*a.front.Size() < a.back.Size() || 3*a.back.Size() < a.front.Size() {
		len := a.Size()
		newFrontLen := len / 2
		newFrontCap := max(2*newFrontLen, 1)
		newFrontBuf := make([]V, newFrontCap)
		for i := 0; i < newFrontLen; i++ {
			newFrontBuf[newFrontLen-i-1] = a.Get(i)
		}

		newBackLen := len - newFrontLen
		newBackCap := max(2*newBackLen, 1)
		newBackBuf := make([]V, newBackCap)
		for i := 0; i < newBackLen; i++ {
			newBackBuf[i] = a.Get(newFrontLen + i)
		}

		a.front.buf = newFrontBuf
		a.front.len = newFrontLen
		a.front.cap = newFrontCap
		a.back.buf = newBackBuf
		a.back.len = newBackLen
		a.back.cap = newBackCap
	}
}
