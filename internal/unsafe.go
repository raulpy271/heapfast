package internal

import (
	"github.com/raulpy271/heapfast/heapfast"
	"unsafe"
)

func CastRecords(bytes []byte) []heapfast.Record {
	var rec heapfast.Record
	var bytes_ptr *byte = unsafe.SliceData(bytes)
	var records_ptr *heapfast.Record = (*heapfast.Record)(unsafe.Pointer(bytes_ptr))
	return unsafe.Slice(records_ptr, len(bytes)/int(unsafe.Sizeof(rec)))
}
