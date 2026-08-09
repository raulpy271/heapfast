package internal

import (
	"unsafe"

	"github.com/raulpy271/heapfast/heapfast"
)

func CastRecordsFromBytes[K heapfast.SizedNumber, V heapfast.SizedNumber | heapfast.Zero](bytes []byte) []heapfast.Record[K, V] {
	var rec heapfast.Record[K, V]
	var bytes_ptr *byte = unsafe.SliceData(bytes)
	var records_ptr *heapfast.Record[K, V] = (*heapfast.Record[K, V])(unsafe.Pointer(bytes_ptr))
	return unsafe.Slice(records_ptr, len(bytes)/int(unsafe.Sizeof(rec)))
}

func CastRecordsToBytes[K heapfast.SizedNumber, V heapfast.SizedNumber](rec []heapfast.Record[K, V]) []byte {
	var rec_ptr *heapfast.Record[K, V] = unsafe.SliceData(rec)
	var bytes_ptr *byte = (*byte)(unsafe.Pointer(rec_ptr))
	return unsafe.Slice(bytes_ptr, len(rec)*int(unsafe.Sizeof(rec[0])))
}
