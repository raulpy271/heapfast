package internal

import (
	"encoding/binary"
	"testing"
	"unsafe"

	"github.com/raulpy271/heapfast/heapfast"
)

func TestCastRecordsFromBytes(t *testing.T) {
	recs := [3]heapfast.Record[uint64, int64]{{10, 20}, {11, 21}, {12, 22}}
	raw := make([]byte, 0, int(unsafe.Sizeof(recs[0]))*len(recs))
	raw, _ = binary.Append(raw, binary.NativeEndian, recs)
	result := CastRecordsFromBytes[uint64, int64](raw)
	for i, rec := range recs {
		if rec != result[i] {
			t.Error(recs, result)
		}
	}
}

func TestCastRecordsToBytes(t *testing.T) {
	recs := [3]heapfast.Record[uint64, int64]{{10, 20}, {11, 21}, {12, 22}}
	raw := make([]byte, 0, int(unsafe.Sizeof(recs[0]))*len(recs))
	raw, _ = binary.Append(raw, binary.NativeEndian, recs)
	result := CastRecordsToBytes(recs[:])
	for i, b := range result {
		if b != raw[i] {
			t.Error(recs, result)
		}
	}
}

func TestCastRecordsFromBytesWithZeroValue(t *testing.T) {
	recs := [3]heapfast.Record[uint64, heapfast.Zero]{{struct{}{}, 20}, {struct{}{}, 21}, {struct{}{}, 22}}
	if int(unsafe.Sizeof(recs)) != 3 * 8 {
		t.Error("Records with empty value should have the size of a 64 integer")
	}
	raw := make([]byte, 0, int(unsafe.Sizeof(recs[0]))*len(recs))
	raw, _ = binary.Append(raw, binary.NativeEndian, recs)
	result := CastRecordsFromBytes[uint64, heapfast.Zero](raw)
	for i, rec := range recs {
		if rec != result[i] {
			t.Error(recs, result)
		}
	}
}
