// Package xsort
package xsort

import "sort"

// LessSwap ...
type LessSwap struct {
	less func(i, j int) bool
	swap func(i, j int)
	len  func() int
}

// Less ...
func (ls *LessSwap) Less(i, j int) bool {
	return ls.less(i, j)
}

// Swap ...
func (ls *LessSwap) Swap(i, j int) {
	ls.swap(i, j)
}

// Len ...
func (ls *LessSwap) Len() int {
	return ls.len()
}

// SortFunc ...
func SortFunc(less func(i, j int) bool, swap func(i, j int), len func() int) {
	sort.Sort(&LessSwap{less, swap, len})
}

// LessUint ...
type LessUint struct {
	data []uint32
	less func(i, j int) bool
}

// Less ...
func (ls *LessUint) Less(i, j int) bool {
	return ls.less(i, j)
}

// Swap ...
func (ls *LessUint) Swap(i, j int) {
	ls.data[i], ls.data[j] = ls.data[j], ls.data[i]
}

// Len ...
func (ls *LessUint) Len() int {
	return len(ls.data)
}

// SortUintLess ...
func SortUintLess(data []uint32, less func(i, j int) bool) {
	sort.Sort(&LessUint{data, less})
}
