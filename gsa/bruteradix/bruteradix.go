// Package bruteradix is the same as brute but uses radix sort for sorting
//the suffixes.

package bruteradix

import (
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

type bucket struct {
	first *bucketElement
	tail  *bucketElement
	len   int
}

type bucketElement struct {
	next *bucketElement
	str  string
}

// New returns a new GSA
func New() gsa.Builder {
	return &store{}
}

func findInArray(pat string, a *suffixArray) bool {
	left := 0
	right := len(a.str) - 1

	for left <= right {
		mid := left + (right-left-1)/2

		l := a.a[mid] + len(pat)
		if l > len(a.str) {
			l = len(a.str)
		}
		suf := a.str[a.a[mid]:l]

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
	suffixes := make([]bucketElement, 0, len(str))
	for i := 0; i < len(str); i++ {
		suffixes = append(suffixes, bucketElement{str: str[i:]})
	}

	result := radixSort0(suffixes)

	SA := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		SA[i] = len(str) - len(result[i].str)
	}

	s.l.Lock()
	s.sa = append(s.sa, &suffixArray{a: SA, idx: idx, str: str})
	s.l.Unlock()
}

func (s *store) Build() gsa.Searcher {
	return s
}

func radixSort0(ss []bucketElement) []bucketElement {
	firstBucket := &bucket{first: &ss[0], tail: &ss[0], len: len(ss)}

	for i := 1; i < len(ss); i++ {
		firstBucket.tail.next = &ss[i]
		firstBucket.tail = &ss[i]
	}

	result := make([]bucketElement, len(ss))
	radixSort(result, firstBucket, 0, 0)
	return result
}

func radixSort(result []bucketElement, b *bucket, k int, bucketIndex int) {
	// fill subbuckets
	subbuckets := make([]*bucket, 256)
	for e := b.first; e != nil; e = e.next {
		n := &bucketElement{str: e.str}
		bucketPos := 0
		if k < len(e.str) {
			bucketPos = int([]byte(e.str)[k])
		}

		if subbuckets[bucketPos] == nil {
			subbuckets[bucketPos] = &bucket{first: n, tail: n, len: 1}
		} else {
			subbuckets[bucketPos].tail.next = n
			subbuckets[bucketPos].tail = n
			subbuckets[bucketPos].len++
		}
	}

	// sort subbuckets
	currentIndex := bucketIndex
	for i := 0; i < len(subbuckets); i++ {
		if subbuckets[i] != nil {

			if subbuckets[i].len == 1 {
				result[currentIndex] = *subbuckets[i].first
				currentIndex++
				continue
			}

			radixSort(result, subbuckets[i], k+1, currentIndex)
			currentIndex += subbuckets[i].len
		}
	}
}
