package main

import "testing"

func TestAddItem(t *testing.T) {
	i := Record{0, 0}
	h := Heap{}
	h.AddItem(i)
}

func TestPopItem(t *testing.T) {
	i := Record{0, 0}
	h := Heap{}
	h.AddItem(i)
	r := h.PopItem()
	if i[0] != r[0] {
		t.Error(i, r)
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
	arr := []Record{{10, 0}, {12, 0}, {22, 0}}
	h := Heap{arr, 1}
	h.heapify(0)
	if h.items[0][0] != 10 {
		t.Error(h)
	}
}

func TestHeapifyThreeNode(t *testing.T) {
	arr := []Record{{10, 0}, {12, 0}, {22, 0}}
	h := Heap{arr, 3}
	h.heapify(0)
	if h.items[0][0] != 22 {
		t.Error(h)
	}
	if h.items[1][0] != 12 {
		t.Error(h)
	}
	if h.items[2][0] != 10 {
		t.Error(h)
	}
}
