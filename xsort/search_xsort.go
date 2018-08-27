// Package expand the sort package for some common cases.
package xsort

import "sort"

// SearchUInts searches for x in a sorted slice of uint32s and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
//
func SearchUInts(a []uint32, x uint32) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// SearchUInts searches for x in a sorted slice of uint32s and returns the index
// as specified by Search. The return value is -1 if x is not present.
// The slice must be sorted in ascending order.
//
func SearchUIntsExists(a []uint32, x uint32) int {
	pos := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if pos >= len(a) || a[pos] != x {
		return -1
	}
	return pos
}
