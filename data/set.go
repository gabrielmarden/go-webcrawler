package data

import "sort"

type Set struct {
	items map[string]bool
}

func NewSet() Set {
	return Set{
		items: make(map[string]bool),
	}
}

func (s *Set) Add(val string) {
	s.items[val] = true
}

func (s *Set) Delete(val string) {
	delete(s.items, val)
}

func (s *Set) Contains(val string) bool {
	return s.items[val]
}

func (s *Set) GetAll() []string {
	items := make([]string, len(s.items))
	i := 0
	for k := range s.items {
		items[i] = k
		i++
	}
	sort.Strings(items)
	return items
}

func (s *Set) Length() int {
	return len(s.items)
}
