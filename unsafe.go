package main

import (
	"unsafe"
)

func CastRecords(bytes []byte) []Record {
	var rec Record
	var bytes_ptr *byte = unsafe.SliceData(bytes)
	var records_ptr *Record = (*Record)(unsafe.Pointer(bytes_ptr))
	return unsafe.Slice(records_ptr, len(bytes)/int(unsafe.Sizeof(rec)))
}
