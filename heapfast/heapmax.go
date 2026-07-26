package heapfast

import (
	"cmp"
	"math"
)

type HeapMax[K cmp.Ordered, V any] heapt[K, V]

func (h *HeapMax[K, V]) heapify(i uint) {
	var l, r, largest uint
	for i < h.Length {
		l = left(i)
		r = right(i)
		if l < h.Length && h.Records[i].Key < h.Records[l].Key {
			largest = l
		} else {
			largest = i
		}
		if r < h.Length && h.Records[largest].Key < h.Records[r].Key {
			largest = r
		}
		if i == largest {
			break
		} else {
			h.Records[i], h.Records[largest] = h.Records[largest], h.Records[i]
			i = largest
		}
	}
}

func BuildMaxHeap[K cmp.Ordered, V any](records []Record[K, V]) *HeapMax[K, V] {
	heap := &HeapMax[K, V]{records, uint(len(records))}
	i := int(math.Floor((float64(len(records)) / 2) - 1))
	for ; i >= 0; i-- {
		heap.heapify(uint(i))
	}
	return heap
}

func (heap *HeapMax[K, V]) Sort() []Record[K, V] {
	for i := heap.Length - 1; i > 0; i-- {
		heap.Records[i], heap.Records[0] = heap.Records[0], heap.Records[i]
		heap.Length--
		heap.heapify(0)
	}
	return heap.Records
}

func (h *HeapMax[K, V]) AddItem(record Record[K, V]) {
	if int(h.Length) == len(h.Records) {
		h.Records = append(h.Records, record)
	} else {
		h.Records[h.Length] = record
	}
	i := h.Length
	h.Length++
	for i > 0 && h.Records[parent(i)].Key < h.Records[i].Key {
		h.Records[i], h.Records[parent(i)] = h.Records[parent(i)], h.Records[i]
		i = parent(i)
	}
}

func (h HeapMax[K, V]) Top() Record[K, V] {
	return h.Records[0]
}

func (h *HeapMax[K, V]) PopItem() Record[K, V] {
	i := h.Records[0]
	h.Records[0] = h.Records[h.Length-1]
	h.Length--
	h.heapify(0)
	return i
}
