package internal

import (
	"github.com/raulpy271/heapfast/heapfast"
	"unsafe"
)

func CastRecordsFromBytes(bytes []byte) []heapfast.Record {
	var rec heapfast.Record
	var bytes_ptr *byte = unsafe.SliceData(bytes)
	var records_ptr *heapfast.Record = (*heapfast.Record)(unsafe.Pointer(bytes_ptr))
	return unsafe.Slice(records_ptr, len(bytes)/int(unsafe.Sizeof(rec)))
}

func CastRecordsToBytes(rec []heapfast.Record) []byte {
	var rec_ptr *heapfast.Record = unsafe.SliceData(rec)
	var bytes_ptr *byte = (*byte)(unsafe.Pointer(rec_ptr))
	return unsafe.Slice(bytes_ptr, len(rec)*int(unsafe.Sizeof(rec[0])))
}
