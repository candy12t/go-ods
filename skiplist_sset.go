package ods

// skiplistSSet is implemenation of SSet interface using skiplist
type skiplistSSet struct {
}

var _ SSet = (*skiplistSSet)(nil)

func (s *skiplistSSet) Size() int {
	return 0
}

func (s *skiplistSSet) Add(idx int, v V) {
}

func (s *skiplistSSet) Remove(idx int) V {
	return nil
}

func (s *skiplistSSet) Find(v V) V {
	return nil
}
