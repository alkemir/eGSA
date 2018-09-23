// Package brute provides an implementation of the methods of a GSA
// with very bad performance by sorting all suffixes using a golangs sort.
package brute

import (
	"fmt"
	"sort"
	"sync"

	"github.com/alkemir/eGSA/gsa"
)

type store struct {
	sa []*suffixArray
	l  sync.Mutex
}

type suffixArray struct {
	a   []int
	str string
	idx gsa.ResultIndex
}

// New returns a new GSA
func New() gsa.Builder {
	return &store{}
}

func findInArray(pat string, a *suffixArray) bool {
	fmt.Println(a.str)
	left := 0
	right := len(a.str) - 1

	for left <= right {
		mid := left + (right-left-1)/2

		l := a.a[mid] + len(pat)
		if l > len(a.str) {
			l = len(a.str)
		}
		suf := a.str[a.a[mid]:l]
		fmt.Println(suf, left, right)

		if pat < suf {
			right = mid - 1
		} else if pat > suf {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

func (s *store) Search(str string) []gsa.ResultIndex {
	s.l.Lock()
	res := make([]gsa.ResultIndex, 0)
	for _, a := range s.sa {
		if findInArray(str, a) {
			res = append(res, a.idx)
		}
	}

	s.l.Unlock()
	return res
}

func (s *store) Add(str string, idx gsa.ResultIndex) {
	suffixes := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		suffixes = append(suffixes, str[i:])
	}

	sort.Sort(Lex(suffixes))

	SA := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		SA[i] = len(str) - len(suffixes[i])
	}

	fmt.Println(SA)

	s.l.Lock()
	s.sa = append(s.sa, &suffixArray{a: SA, idx: idx, str: str})
	s.l.Unlock()
}

func (s *store) Build() gsa.Searcher {
	return s
}

type Lex []string

func (s Lex) Len() int      { return len(s) }
func (s Lex) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Lex) Less(i, j int) bool {
	compLen := len(s[i])
	if compLen > len(s[j]) {
		compLen = len(s[j])
	}

	if s[i][:compLen] < s[j][:compLen] {
		return true
	} else if s[i][:compLen] > s[j][:compLen] {
		return false
	}
	return len(s[i]) < len(s[j])
}
