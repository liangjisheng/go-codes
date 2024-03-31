package main

type Set map[string]struct{}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Remove(key string) {
	delete(s, key)
}

func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

func NewSet() Set {
	s := make(Set)
	return s
}
