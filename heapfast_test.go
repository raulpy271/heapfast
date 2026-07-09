package main

import "testing"

func TestAddItem(t *testing.T) {
	i := Item{0, nil}
	h := Heap{}
	h.AddItem(i)
}

func TestPopItem(t *testing.T) {
	i := Item{0, nil}
	h := Heap{}
	h.AddItem(i)
	r := h.PopItem()
	if i.key != r.key {
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
	arr := []Item{{10, nil}, {12, nil}, {22, nil}}
	h := Heap{arr, 1}
	h.heapify(0)
	if h.items[0].key != 10 {
		t.Error(h)
	}
}

func TestHeapifyThreeNode(t *testing.T) {
	arr := []Item{{10, nil}, {12, nil}, {22, nil}}
	h := Heap{arr, 3}
	h.heapify(0)
	if h.items[0].key != 22 {
		t.Error(h)
	}
	if h.items[1].key != 12 {
		t.Error(h)
	}
	if h.items[2].key != 10 {
		t.Error(h)
	}
}
