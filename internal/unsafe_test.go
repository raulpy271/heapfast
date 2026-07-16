package internal

import (
	"encoding/binary"
	"testing"
	"unsafe"
	"github.com/raulpy271/heapfast/heapfast"
)

func TestCastRecords(t *testing.T) {
	recs := [3]heapfast.Record{{10, 20}, {11, 21}, {12, 22}}
	raw := make([]byte, 0, int(unsafe.Sizeof(recs[0]))*len(recs))
	raw, _ = binary.Append(raw, binary.NativeEndian, recs)
	result := CastRecords(raw)
	for i, rec := range recs {
		if rec != result[i] {
			t.Error(recs, result)
		}
	}
}
