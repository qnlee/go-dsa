package set

type Set struct {
	exists map[string]bool
}

func (s *Set) NewSet() *Set {
	return &Set{make(map[string]bool)}
}

func (s *Set) Add(elem string) {
	if _, ok := s.exists[elem]; !ok {
		s.exists[elem] = true
	}
}

func (s *Set) Remove(elem string) {
	if _, ok := s.exists[elem]; ok {
		s.exists[elem] = false
	}
}
