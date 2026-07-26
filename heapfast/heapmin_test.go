package heapfast

import "testing"

func TestAddThreeRecordsHeapMin(t *testing.T) {
	h := HeapMin[uint, uint]{make([]Record[uint, uint], 5), 0}
	result := []uint{1, 3, 2, 0, 0}
	h.AddItem(Record[uint, uint]{3, 0})
	h.AddItem(Record[uint, uint]{2, 0})
	h.AddItem(Record[uint, uint]{1, 0})
	if h.Length != 3 {
		t.Error(h)
	}
	if len(h.Records) != 5 {
		t.Error(h)
	}
	for j := range 5 {
		if h.Records[j].Key != result[j] {
			t.Error(h)
		}
	}
}

func TestAddSixRecordsHeapMin(t *testing.T) {
	h := HeapMin[uint, uint]{make([]Record[uint, uint], 5), 0}
	result := []uint{5, 7, 10, 9, 8, 11}
	h.AddItem(Record[uint, uint]{9, 0})
	h.AddItem(Record[uint, uint]{8, 0})
	h.AddItem(Record[uint, uint]{10, 0})
	h.AddItem(Record[uint, uint]{5, 0})
	h.AddItem(Record[uint, uint]{7, 0})
	h.AddItem(Record[uint, uint]{11, 0})
	if h.Length != 6 {
		t.Error(h)
	}
	if len(h.Records) != 6 {
		t.Error(h)
	}
	for j := range 6 {
		if h.Records[j].Key != result[j] {
			t.Error(h)
		}
	}
}

func TestPopThreeHeapMin(t *testing.T) {
	h := BuildMinHeap([]Record[uint, uint]{{9, 0}, {8, 0}, {10, 0}, {5, 0}, {7, 0}, {11, 0}})
	result := []uint{9, 11, 10}
	if h.PopItem().Key != 5 {
		t.Error(h)
	}
	if h.PopItem().Key != 7 {
		t.Error(h)
	}
	if h.PopItem().Key != 8 {
		t.Error(h)
	}
	if h.Length != 3 {
		t.Error(h)
	}
	if len(h.Records) != 6 {
		t.Error(h)
	}
	for i := range 3 {
		if h.Records[i].Key != result[i] {
			t.Error(h)
		}
	}
}

func TestSortMin(t *testing.T) {
	arr := []Record[uint, uint]{{4, 0}, {14, 0}, {7, 0}, {2, 0}, {8, 0}, {1, 0}}
	sorted := []Record[uint, uint]{{14, 0}, {8, 0}, {7, 0}, {4, 0}, {2, 0}, {1, 0}}
	heap := BuildMinHeap(arr)
	result := heap.Sort()
	for i, r := range result {
		if r != sorted[i] {
			t.Error(sorted, result)
		}
	}
}
