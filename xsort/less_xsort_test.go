// Package xsort
package xsort

import "testing"

func TestSortFunc(t *testing.T) {
	SortFunc(LessInt, SwapInt, LenInt)
}

var a = []uint{3, 2, 1}

func SwapInt(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func LenInt() int {
	return len(a)
}

func LessInt(i, j int) bool {
	return false
}
