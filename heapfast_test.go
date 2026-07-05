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
	if len(i.value) != len(r.value) {
		t.Error(i, r)
	}

}
