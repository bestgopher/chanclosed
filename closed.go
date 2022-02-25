package chanclosed

import (
	"unsafe"
)

type innerChan struct {
	_      uint
	_      uint
	_      unsafe.Pointer
	_      uint16
	closed uint32
}

const offset = unsafe.Offsetof(innerChan{}.closed)

func Closed[T any](ch chan T) bool {
	return *(*uint32)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&ch)) + offset)) == 1
}
