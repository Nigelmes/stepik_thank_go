package main

// реализуйте быстрое множество
type IntSet struct {
	elems map[int]struct{}
}

func MakeIntSet() IntSet {
	return IntSet{make(map[int]struct{})}
}

func (s IntSet) Contains(elem int) bool {
	if _, ok := s.elems[elem]; ok {
		return true
	}
	return false
}

func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s.elems[elem] = struct{}{}
	return true
}
