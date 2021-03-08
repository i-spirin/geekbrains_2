package set

import "sync"

// // Set is an interface for structure stores unique integers
// type Set interface {
// 	Add()
// 	Has()
// 	Remove()
// }

// MuSet is structure stores unique integers. This is a thread-safe stucture uses Mutex
type MuSet struct {
	mm map[int]struct{}
	mu sync.Mutex
}

// NewMuSet returns a new MuSet
func NewMuSet() MuSet {
	return MuSet{
		mm: make(map[int]struct{}),
		mu: sync.Mutex{},
	}
}

// Add an integer to MuSet
func (s *MuSet) Add(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mm[i] = struct{}{}
}

// Has is a function to check existing for integer in the MuSet
func (s *MuSet) Has(i int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, existing := s.mm[i]
	return existing
}

// Remove an element from Muset
func (s *MuSet) Remove(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.mm, i)
}

// RWMuSet is structure stores unique integers. This is a thread-safe stucture uses RWMutex
type RWMuSet struct {
	mm map[int]struct{}
	mu sync.RWMutex
}

// NewRWMuSet return new RWMuSet object
func NewRWMuSet() RWMuSet {
	return RWMuSet{
		mm: make(map[int]struct{}),
		mu: sync.RWMutex{},
	}
}

// Add an integer to MuSet
func (s *RWMuSet) Add(i int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.mm[i] = struct{}{}
}

// Has is a function to check existing for integer in the MuSet
func (s *RWMuSet) Has(i int) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, existing := s.mm[i]
	return existing
}

// Remove an element from Muset
func (s *RWMuSet) Remove(i int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	delete(s.mm, i)
}
