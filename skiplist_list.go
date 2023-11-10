package ods

// skiplistList is implemenation of List interface using skiplist
type skiplistList struct {
}

var _ List = (*skiplistList)(nil)

func (s *skiplistList) Size() int {
	return 0
}

func (s *skiplistList) Get(idx int) V {
	return 0
}

func (s *skiplistList) Set(idx int, v V) V {
	return 0
}

func (s *skiplistList) Add(idx int, v V) {
}

func (s *skiplistList) Remove(idx int) V {
	return nil
}
