package gsa_test

import (
	"testing"

	"github.com/alkemir/eGSA/gsa"
	"github.com/alkemir/eGSA/gsa/mockmap"
)

var testStrings = []string{
	"Llanfairpwllgwyngyllgogerychwyrndrobwllllantysiliogogogoch",
	"banana",
	"anana",
	"ana",
	"zanahoria",
	"carrot",
}

var searchStrings = []struct {
	str    string
	result []gsa.ResultIndex
}{
	{"na", []gsa.ResultIndex{1, 2, 3, 4}},
}

func testImpl(t *testing.T, g gsa.GeneralizedSuffixArray) {

	for i, ts := range testStrings {
		g.Add(ts, gsa.ResultIndex(i))
	}

	s := g.Build()

	for _, ts := range searchStrings {
		rr := s.Search(ts.str)
		for _, expected := range ts.result {
			found := false
			for _, res := range rr {
				if res == expected {
					found = true
				}
			}

			if !found {
				t.Fatalf("%v not found in result for Search(%s): %v", expected, ts.str, rr)
			}
		}
	}
}

func TestMockmap(t *testing.T) {
	g := mockmap.New()
	testImpl(t, g)
}
