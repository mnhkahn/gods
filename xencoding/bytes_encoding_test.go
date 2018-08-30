// Package xencoding
package xencoding

import (
	"math"
	"reflect"
	"runtime/debug"
	"testing"

	"github.com/mnhkahn/gogogo/logger"
)

func TestUint64AndByte(t *testing.T) {
	a := []uint64{1, 2, math.MaxUint64}
	temp := Uint64s2Bytes(a)
	c := Bytes2Uint64s(temp)
	if !reflect.DeepEqual(a, c) {
		t.Error(a, temp, c)
	}
}

func TestDebugStack(t *testing.T) {
	a := []uint32{1, 2}
	temp := Uints2Bytes(a)
	logger.Warn(len(a), cap(a), len(temp), cap(temp))
	Stack(temp)
	StackInt(temp)
	StackInt32Slice(temp)
}

func Stack(aa []byte) {
	debug.PrintStack()
}

func StackInt(aa []byte) int {
	debug.PrintStack()
	return 1
}

func StackInt32Slice(aa []byte) (int, []uint32) {
	debug.PrintStack()
	return 1, Bytes2Uints(aa)
}
