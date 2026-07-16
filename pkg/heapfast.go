package heapfast

import (
	"math"
)

type Record [2]uint64

type Heap struct {
	items  []Record
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
		if l < h.length && h.items[i][0] < h.items[l][0] {
			largest = l
		} else {
			largest = i
		}
		if r < h.length && h.items[largest][0] < h.items[r][0] {
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

func BuildMaxHeap(records []Record) *Heap {
	heap := &Heap{records, uint(len(records))}
	i := int(math.Floor((float64(len(records)) / 2) - 1))
	for ; i >= 0; i-- {
		heap.heapify(uint(i))
	}
	return heap
}

func SortMax(records []Record) []Record {
	heap := BuildMaxHeap(records)
	return heap.items
}

func (h *Heap) AddItem(record Record) {
	if int(h.length) == len(h.items) {
		h.items = append(h.items, record)
	} else {
		h.items[h.length] = record
	}
	i := h.length
	h.length++
	for i > 0 && h.items[parent(i)][0] < h.items[i][0] {
		h.items[i], h.items[parent(i)] = h.items[parent(i)], h.items[i]
		i = parent(i)
	}
}

func (h Heap) Max() Record {
	return h.items[0]
}

func (h *Heap) PopItem() Record {
	i := h.items[0]
	h.items[0] = h.items[h.length-1]
	h.length--
	h.heapify(0)
	return i
}
