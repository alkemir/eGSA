package gsa

// ResultIndex is used as the return type for Search
type ResultIndex int32

// GeneralizedSuffixArray supports operations for a GSA
type GeneralizedSuffixArray interface {
	Searcher
	Builder
}

// Searcher can be searched for a substring, returning a list of strings
// where the substring is found.
type Searcher interface {
	Search(string) []ResultIndex
}

// Builder allows the creation of GSAs from a set of strings that are added to the builder.
type Builder interface {
	Add(string, ResultIndex)
	Build() Searcher
}
