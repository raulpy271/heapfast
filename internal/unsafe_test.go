package internal

import (
	"encoding/binary"
	"github.com/raulpy271/heapfast/heapfast"
	"testing"
	"unsafe"
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
