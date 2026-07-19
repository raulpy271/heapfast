package internal

import (
	"unsafe"

	"github.com/raulpy271/heapfast/heapfast"
)

func CastRecordsFromBytes[T uint64 | int64 | int32 | uint32 | float32 | float64](bytes []byte) []heapfast.Record[T] {
	var rec heapfast.Record[T]
	var bytes_ptr *byte = unsafe.SliceData(bytes)
	var records_ptr *heapfast.Record[T] = (*heapfast.Record[T])(unsafe.Pointer(bytes_ptr))
	return unsafe.Slice(records_ptr, len(bytes)/int(unsafe.Sizeof(rec)))
}

func CastRecordsToBytes[T uint64 | int64 | int32 | uint32 | float32 | float64](rec []heapfast.Record[T]) []byte {
	var rec_ptr *heapfast.Record[T] = unsafe.SliceData(rec)
	var bytes_ptr *byte = (*byte)(unsafe.Pointer(rec_ptr))
	return unsafe.Slice(bytes_ptr, len(rec)*int(unsafe.Sizeof(rec[0])))
}
