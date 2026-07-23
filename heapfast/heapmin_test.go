package heapfast

import "testing"

func TestAddThreeRecordsHeapMin(t *testing.T) {
	h := HeapMin[uint]{make([]Record[uint], 5), 0}
	result := []uint{1, 3, 2, 0, 0}
	h.AddItem(Record[uint]{3})
	h.AddItem(Record[uint]{2})
	h.AddItem(Record[uint]{1})
	if h.Length != 3 {
		t.Error(h)
	}
	if len(h.Records) != 5 {
		t.Error(h)
	}
	for j := range 5 {
		if h.Records[j][0] != result[j] {
			t.Error(h)
		}
	}
}

func TestAddSixRecordsHeapMin(t *testing.T) {
	h := HeapMin[uint]{make([]Record[uint], 5), 0}
	result := []uint{5, 7, 10, 9, 8, 11}
	h.AddItem(Record[uint]{9})
	h.AddItem(Record[uint]{8})
	h.AddItem(Record[uint]{10})
	h.AddItem(Record[uint]{5})
	h.AddItem(Record[uint]{7})
	h.AddItem(Record[uint]{11})
	if h.Length != 6 {
		t.Error(h)
	}
	if len(h.Records) != 6 {
		t.Error(h)
	}
	for j := range 6 {
		if h.Records[j][0] != result[j] {
			t.Error(h)
		}
	}
}

func TestPopThreeHeapMin(t *testing.T) {
	h := BuildMinHeap([]Record[uint]{{9}, {8}, {10}, {5}, {7}, {11}})
	result := []uint{9, 11, 10}
	if h.PopItem()[0] != 5 {
		t.Error(h)
	}
	if h.PopItem()[0] != 7 {
		t.Error(h)
	}
	if h.PopItem()[0] != 8 {
		t.Error(h)
	}
	if h.Length != 3 {
		t.Error(h)
	}
	if len(h.Records) != 6 {
		t.Error(h)
	}
	for i := range 3 {
		if h.Records[i][0] != result[i] {
			t.Error(h)
		}
	}
}

func TestSortMin(t *testing.T) {
	arr := []Record[uint]{{4}, {14}, {7}, {2}, {8}, {1}}
	sorted := []Record[uint]{{14}, {8}, {7}, {4}, {2}, {1}}
	heap := BuildMinHeap(arr)
	result := heap.Sort()
	for i, r := range result {
		if r != sorted[i] {
			t.Error(sorted, result)
		}
	}
}
