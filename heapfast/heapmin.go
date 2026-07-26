package heapfast

import (
	"cmp"
	"math"
)

type HeapMin[K cmp.Ordered, V any] heapt[K, V]

func (h *HeapMin[K, V]) heapify(i uint) {
	var l, r, lowest uint
	for i < h.Length {
		l = left(i)
		r = right(i)
		if l < h.Length && h.Records[i].Key > h.Records[l].Key {
			lowest = l
		} else {
			lowest = i
		}
		if r < h.Length && h.Records[lowest].Key > h.Records[r].Key {
			lowest = r
		}
		if i == lowest {
			break
		} else {
			h.Records[i], h.Records[lowest] = h.Records[lowest], h.Records[i]
			i = lowest
		}
	}
}

func BuildMinHeap[K cmp.Ordered, V any](records []Record[K, V]) *HeapMin[K, V] {
	heap := &HeapMin[K, V]{records, uint(len(records))}
	i := int(math.Floor((float64(len(records)) / 2) - 1))
	for ; i >= 0; i-- {
		heap.heapify(uint(i))
	}
	return heap
}

func (heap *HeapMin[K, V]) Sort() []Record[K, V] {
	for i := heap.Length - 1; i > 0; i-- {
		heap.Records[i], heap.Records[0] = heap.Records[0], heap.Records[i]
		heap.Length--
		heap.heapify(0)
	}
	return heap.Records
}

func (h *HeapMin[K, V]) AddItem(record Record[K, V]) {
	if int(h.Length) == len(h.Records) {
		h.Records = append(h.Records, record)
	} else {
		h.Records[h.Length] = record
	}
	i := h.Length
	h.Length++
	for i > 0 && h.Records[parent(i)].Key > h.Records[i].Key {
		h.Records[i], h.Records[parent(i)] = h.Records[parent(i)], h.Records[i]
		i = parent(i)
	}
}

func (h HeapMin[K, V]) Top() Record[K, V] {
	return h.Records[0]
}

func (h *HeapMin[K, V]) PopItem() Record[K, V] {
	i := h.Records[0]
	h.Records[0] = h.Records[h.Length-1]
	h.Length--
	h.heapify(0)
	return i
}
