package ods

// chainedHashTable is hash table using chaining
type chainedHashTable struct {
	len   int
	table []arrayStack
}

func (ht *chainedHashTable) Add(v V) bool {
	if ht.Find(v) != nil {
		return false
	}

	if ht.isFull() {
		ht.resize()
	}

	ht.table[ht.hash(v)].Push(v)
	ht.len++

	return true
}

func (ht *chainedHashTable) Remove(v V) V {
	stack := ht.table[ht.hash(v)]
	for i := 0; i < stack.Size(); i++ {
		if got := stack.Get(i); v == got {
			stack.Remove(i)
			ht.len--
			return got
		}
	}
	return nil
}

func (ht *chainedHashTable) Find(v V) V {
	stack := ht.table[ht.hash(v)]
	for i := 0; i < stack.Size(); i++ {
		if got := stack.Get(i); v == got {
			return got
		}
	}
	return nil
}

func (ht *chainedHashTable) hash(v V) uint64 {
	// ((z * v) mod 2^w) div2^(w-d)
	// z: {1, 3, 5, ..., 2^w-1}からランダムに選択
	// w: 整数のビット数
	// z, d := 1, 1
	// return (z * uint64(v)) >> (w - d)
	return 0
}

func (ht *chainedHashTable) isFull() bool {
	return ht.len == len(ht.table)
}

func (ht *chainedHashTable) resize() {
}
