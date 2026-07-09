package main

import (
	"math"
)

type Item struct {
	key   float64
	value any
}

type Heap struct {
	items  [2]Item
	length uint
}

func left(i uint) uint {
	// i * 2 + 1
	return (i << 1) + 1
}

func right(i uint) uint {
	// i * 2 + 2
	return (i << 1) + 2
}

func parent(i uint) uint {
	return uint(math.Floor((float64(i) - 1) / 2))
}

func (h *Heap) heapify(i uint) {
	var l, r, largest uint
	for i < h.length {
		l = left(i)
		r = right(i)
		if l < h.length && h.items[i].key < h.items[l].key {
			largest = l
		} else {
			largest = i
		}
		if r < h.length && h.items[largest].key < h.items[r].key {
			largest = r
		}
		if i == largest {
			break
		} else {
			h.items[i], h.items[largest] = h.items[largest], h.items[i]
			i = largest
		}
	}

}

func (h *Heap) AddItem(i Item) {
	h.items = append(h.items, i)
}

func (h *Heap) PopItem() Item {
	i := h.items[0]
	h.items = h.items[1:]
	return i
}
