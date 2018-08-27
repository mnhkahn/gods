// Package xsort
package xsort

import (
	"sort"
	"testing"
)

var uints = [...]uint{74, 59, 238, 784, 9845, 959, 905, 0, 0, 42, 7586, 5467984, 7586}
var uint32s = [...]uint32{74, 59, 238, 784, 9845, 959, 905, 0, 0, 42, 7586, 5467984, 7586}
var uint64s = [...]uint64{74, 59, 238, 784, 9845, 959, 905, 0, 0, 42, 7586, 5467984, 7586}

func TestSortUintSlice(t *testing.T) {
	data := uints
	a := UintSlice(data[0:])
	sort.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", uints)
		t.Errorf("   got %v", data)
	}
}

func TestSortUint32Slice(t *testing.T) {
	data := uint32s
	a := Uint32Slice(data[0:])
	sort.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", uints)
		t.Errorf("   got %v", data)
	}
}

func TestSortUint64Slice(t *testing.T) {
	data := uint64s
	a := Uint64Slice(data[0:])
	sort.Sort(a)
	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", uints)
		t.Errorf("   got %v", data)
	}
}
