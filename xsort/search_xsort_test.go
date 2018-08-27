// Package xsort
package xsort

import "testing"

func TestSearchUIntsExists(t *testing.T) {
	for i, testcase := range []struct {
		ints []uint32
		x    uint32
		pos  int
	}{
		{[]uint32{0, 1, 2}, 0, 0},
		{[]uint32{0, 1, 2}, 1, 1},
		{[]uint32{0, 1, 2}, 2, 2},
		{[]uint32{0, 1, 2}, 3, -1},
		{[]uint32{0, 10, 200}, 3, -1},
	} {
		temp := SearchUIntsExists(testcase.ints, testcase.x)
		if temp != testcase.pos {
			t.Error(i, testcase, temp)
		}
	}
}

func TestSearchUInts(t *testing.T) {
	for i, testcase := range []struct {
		ints []uint32
		x    uint32
		pos  int
	}{
		{[]uint32{0, 1, 2}, 0, 0},
		{[]uint32{0, 1, 2}, 1, 1},
		{[]uint32{0, 1, 2}, 2, 2},
		{[]uint32{0, 1, 2}, 3, 3},
		{[]uint32{0, 10, 200}, 3, 1},
	} {
		temp := SearchUInts(testcase.ints, testcase.x)
		if temp != testcase.pos {
			t.Error(i, testcase, temp)
		}
	}
}
