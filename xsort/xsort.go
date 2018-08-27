// Package xsort
package xsort

import "sort"

// UintSlice attaches the methods of Interface to []uint, sorting in increasing order.
type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// UInts sort a in increasing order.
func UInts(a []uint) { sort.Sort(UintSlice(a)) }

// Uint32Slice attaches the methods of Interface to []uint32, sorting in increasing order.
type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// UInt32s sort a in increasing order.
func UInt32s(a []uint32) { sort.Sort(Uint32Slice(a)) }

// Uint64Slice attaches the methods of Interface to []uint64, sorting in increasing order.
type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// UInt64s sort a in increasing order.
func UInt64s(a []uint64) { sort.Sort(Uint64Slice(a)) }
