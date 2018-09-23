// Is safe for concurrent use but does not support duplicates. Does not support the empty string.
package external

import (
	"sync"

	"github.com/alkemir/eGSA/gsa"
)

const sep byte = 1
const end byte = 0

type store struct {
	path   string
	cat    []byte
	catLen int
	head   *linkedString
	tail   *linkedString
	l      sync.Mutex
}

type linkedString struct {
	str  string
	next *linkedString
}

func Create(path string) gsa.Builder {
	return &store{path: path}
}

func Load(path string) gsa.Searcher {
	return &store{path: path}
}

func (s *store) Search(str string) []gsa.ResultIndex {
	s.l.Lock()
	s.l.Unlock()
	return res
}

func (s *store) Add(str string, idx gsa.ResultIndex) {
	s.l.Lock()
	if s.head == nil {
		s.head = &linkedString{str: str}
		s.tail = s.head
	} else {
		s.tail.next = &linkedString{str: str}
		s.tail = s.tail.next
	}
	s.catLen += len(str) + 1
	s.l.Unlock()
}

func (s *store) Build() gsa.Searcher {
	s.l.Lock()

	s.cat = make([]byte, 0, s.catLen+1)
	for p := s.head; p != nil; p = p.next {
		s.cat = append(s.cat, []byte(p.str)...)
		s.cat = append(s.cat, sep)
	}
	s.cat = append(s.cat, end)

	// Compute SA, LCP
	// Save it to disk.
	// merge
	s.l.Unlock()
	return s
}
