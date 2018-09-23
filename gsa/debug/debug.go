package main

import (
	"fmt"

	"github.com/alkemir/eGSA/gsa"
	"github.com/alkemir/eGSA/gsa/plcp"
)

var testStrings = []string{
	"acaaacatat",
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

func main() {
	testImpl(plcp.New())
}

func testImpl(g gsa.Builder) {

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
				fmt.Printf("%v not found in result for Search(%s): %v\n", expected, ts.str, rr)
			}
		}
	}
}
