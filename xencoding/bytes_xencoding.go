// Package xencoding
package xencoding

import (
	"bytes"
	"encoding/binary"

	"github.com/mnhkahn/gogogo/logger"
)

func Int162Bytes(k int16) []byte {
	keyBuf := make([]byte, binary.MaxVarintLen16)
	binary.PutVarint(keyBuf, int64(k))
	return keyBuf
}

func Uint2Bytes(k uint32) []byte {
	keyBuf := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(keyBuf, uint64(k))
	return keyBuf
}

func Uints2Bytes(k []uint32) []byte {
	buf := new(bytes.Buffer)
	for _, kk := range k {
		bb := Uint2Bytes(kk)
		_, err := buf.Write(bb)
		if err != nil {
			logger.Warn(err)
		}
	}
	return buf.Bytes()
}

func Bytes2Int16(b []byte) int16 {
	if len(b) != binary.MaxVarintLen16 {
		return 0
	}

	k, _ := binary.Varint(b)
	return int16(k)
}

func Bytes2Uint(b []byte) uint32 {
	if len(b) != binary.MaxVarintLen32 {
		return 0
	}

	k, _ := binary.Uvarint(b)
	return uint32(k)
}

func Bytes2Uints(b []byte) []uint32 {
	l := len(b) / binary.MaxVarintLen32
	res := make([]uint32, l)

	j := 0
	for i := 0; i < len(b); i += binary.MaxVarintLen32 {
		if len(b) >= i+binary.MaxVarintLen32 {
			k, _ := binary.Uvarint(b[i : i+binary.MaxVarintLen32])
			res[j] = uint32(k)
			j++
		}
	}

	return res
}

func Uint642Bytes(k uint64) []byte {
	keyBuf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(keyBuf, uint64(k))
	return keyBuf
}

func Uint64s2Bytes(k []uint64) []byte {
	buf := new(bytes.Buffer)
	for _, kk := range k {
		bb := Uint642Bytes(kk)
		_, err := buf.Write(bb)
		if err != nil {
			logger.Warn(err)
		}
	}
	return buf.Bytes()
}

func Bytes2Uint64s(b []byte) []uint64 {
	l := len(b) / binary.MaxVarintLen64
	res := make([]uint64, l)

	j := 0
	for i := 0; i < len(b); i += binary.MaxVarintLen64 {
		k, _ := binary.Uvarint(b[i : i+binary.MaxVarintLen64])
		res[j] = k
		j++
	}

	return res
}

func Bytes2Uint64(b []byte) uint64 {
	if len(b) != binary.MaxVarintLen64 {
		return 0
	}
	k, _ := binary.Uvarint(b)
	return k
}
