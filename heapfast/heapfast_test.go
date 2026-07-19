package heapfast

import "testing"

func TestAddItemEmptyHeap(t *testing.T) {
	i := Record[uint]{10, 20}
	h := Heap[uint]{}
	h.AddItem(i)
	if h.Length != 1 {
		t.Error(h)
	}
	if len(h.Records) != 1 {
		t.Error(h)
	}
	if h.Records[0] != i {
		t.Error(h)
	}
}

func TestAddItemEmptyHeapWithTrash(t *testing.T) {
	i := Record[uint]{10, 20}
	h := Heap[uint]{make([]Record[uint], 5), 0}
	h.AddItem(i)
	if h.Length != 1 {
		t.Error(h)
	}
	if len(h.Records) != 5 {
		t.Error(h)
	}
	if h.Records[0] != i {
		t.Error(h)
	}
	for j := 1; j < 5; j++ {
		if h.Records[j][0] != 0 {
			t.Error(h)
		}
	}
}

func TestAddThreeRecords(t *testing.T) {
	h := Heap[uint]{make([]Record[uint], 5), 0}
	result := []uint{3, 1, 2, 0, 0}
	h.AddItem(Record[uint]{1})
	h.AddItem(Record[uint]{2})
	h.AddItem(Record[uint]{3})
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

func TestAddSixRecords(t *testing.T) {
	h := Heap[uint]{make([]Record[uint], 5), 0}
	result := []uint{11, 8, 10, 7, 5, 9}
	h.AddItem(Record[uint]{7})
	h.AddItem(Record[uint]{9})
	h.AddItem(Record[uint]{10})
	h.AddItem(Record[uint]{8})
	h.AddItem(Record[uint]{5})
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

func TestPopItem(t *testing.T) {
	i := Record[uint]{0, 0}
	h := Heap[uint]{}
	h.AddItem(i)
	r := h.PopItem()
	if i[0] != r[0] {
		t.Error(i, r)
	}
	if h.Length != 0 {
		t.Error(h)
	}
	if len(h.Records) != 1 {
		t.Error(h)
	}
}

func TestPopThree(t *testing.T) {
	h := BuildMaxHeap([]Record[uint]{{7}, {9}, {10}, {8}, {5}, {11}})
	result := []uint{8, 5, 7}
	if h.PopItem()[0] != 11 {
		t.Error(h)
	}
	if h.PopItem()[0] != 10 {
		t.Error(h)
	}
	if h.PopItem()[0] != 9 {
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

func TestLeft(t *testing.T) {
	for i := range uint(10) {
		if left(i) != (i*2 + 1) {
			t.Error(left(i), i*2+1)
		}
	}
}

func TestRight(t *testing.T) {
	for i := range uint(10) {
		if right(i) != (i*2 + 2) {
			t.Error(right(i), i*2+2)
		}
	}
}

func TestParent(t *testing.T) {
	m := map[uint]uint{1: 0, 2: 0, 3: 1, 4: 1, 5: 2, 6: 2, 7: 3, 8: 3, 9: 4}
	for i, v := range m {
		if parent(i) != v {
			t.Error(parent(i), v)
		}
	}
}

func TestHeapifyOneNode(t *testing.T) {
	arr := []Record[uint]{{10}, {12}, {22}}
	h := Heap[uint]{arr, 1}
	h.heapify(0)
	if h.Records[0][0] != 10 {
		t.Error(h)
	}
}

func TestHeapifyThreeNode(t *testing.T) {
	arr := []Record[uint]{{10}, {12}, {22}}
	h := Heap[uint]{arr, 3}
	h.heapify(0)
	if h.Records[0][0] != 22 {
		t.Error(h)
	}
	if h.Records[1][0] != 12 {
		t.Error(h)
	}
	if h.Records[2][0] != 10 {
		t.Error(h)
	}
}

func TestHeapifySixNode(t *testing.T) {
	arr := []Record[uint]{{4}, {14}, {7}, {2}, {8}, {1}}
	result := []uint{14, 8, 7, 2, 4, 1}
	h := Heap[uint]{arr, 6}
	h.heapify(0)
	for i := range 6 {
		if h.Records[i][0] != result[i] {
			t.Error(h)
		}
	}
}

func TestBuildMaxHeap(t *testing.T) {
	arr := []Record[uint]{{4}, {1}, {3}, {2}, {16}, {9}, {10}, {14}, {8}, {7}}
	result := []uint{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	h := BuildMaxHeap(arr)
	for i := range len(result) {
		if h.Records[i][0] != result[i] {
			t.Error(h)
		}
	}
}
