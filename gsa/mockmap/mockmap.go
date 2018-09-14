// Package mockmap provides an implementation of the methods of a GSA
// with very bad performance by using a map as a string store and searching
// through all of them. It handles duplicates and is safe for concurrent use.
package mockmap

import (
	"strings"
	"sync"

	"github.com/alkemir/eGSA/gsa"
)

type store struct {
	m map[string][]gsa.ResultIndex
	l sync.Mutex
}

// New returns a new GSA mockup based on a map
func New() gsa.GeneralizedSuffixArray {
	return &store{m: make(map[string][]gsa.ResultIndex)}
}

func (s *store) Search(str string) []gsa.ResultIndex {
	s.l.Lock()
	res := make([]gsa.ResultIndex, 0)
	for k, v := range s.m {
		if strings.Contains(k, str) {
			res = append(res, v...)
		}
	}

	s.l.Unlock()
	return res
}

func (s *store) Add(str string, idx gsa.ResultIndex) {
	s.l.Lock()
	r := s.m[str]
	r = append(r, idx)
	s.m[str] = r
	s.l.Unlock()
}

func (s *store) Build() gsa.Searcher {
	return s
}
