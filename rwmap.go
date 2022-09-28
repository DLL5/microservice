package main

import (
	"sync"
	"unsafe"
)

//RWMap 自定义map
type RWMap[key, value any] struct {
	tables map[uint]*Table[key, value]
}

//Table 自定义RWMap的hash值类型
type Table[Key, Value any] struct {
	mu    *sync.RWMutex
	lines map[any]Value
}

//hash hash算法实现
func hash[Key any](key Key) uint {
	gc := &key
	start := uintptr(unsafe.Pointer(gc))
	offset := unsafe.Sizeof(start)
	sizeOfByte := unsafe.Sizeof(byte(0))

	hashSum := uint(0)
	for ptr := start; ptr < start+offset; ptr += sizeOfByte {
		b := *(*byte)(unsafe.Pointer(ptr))

		hashSum += uint(b)
		hashSum = uint(b) + (hashSum << 6) + (hashSum << 16) - hashSum
	}
	return hashSum
}
