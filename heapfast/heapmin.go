package heapfast

import (
	"cmp"
	"math"
)

type HeapMin[T cmp.Ordered] Heap[T]

func (h *HeapMin[T]) heapify(i uint) {
	var l, r, lowest uint
	for i < h.Length {
		l = left(i)
		r = right(i)
		if l < h.Length && h.Records[i][0] > h.Records[l][0] {
			lowest = l
		} else {
			lowest = i
		}
		if r < h.Length && h.Records[lowest][0] > h.Records[r][0] {
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

func BuildMinHeap[T cmp.Ordered](records []Record[T]) *HeapMin[T] {
	heap := &HeapMin[T]{records, uint(len(records))}
	i := int(math.Floor((float64(len(records)) / 2) - 1))
	for ; i >= 0; i-- {
		heap.heapify(uint(i))
	}
	return heap
}

func SortMin[T cmp.Ordered](records []Record[T]) []Record[T] {
	heap := BuildMinHeap(records)
	for i := heap.Length - 1; i > 0; i-- {
		heap.Records[i], heap.Records[0] = heap.Records[0], heap.Records[i]
		heap.Length--
		heap.heapify(0)
	}
	return heap.Records
}

func (h *HeapMin[T]) AddItem(record Record[T]) {
	if int(h.Length) == len(h.Records) {
		h.Records = append(h.Records, record)
	} else {
		h.Records[h.Length] = record
	}
	i := h.Length
	h.Length++
	for i > 0 && h.Records[parent(i)][0] > h.Records[i][0] {
		h.Records[i], h.Records[parent(i)] = h.Records[parent(i)], h.Records[i]
		i = parent(i)
	}
}

func (h HeapMin[T]) Top() Record[T] {
	return h.Records[0]
}

func (h *HeapMin[T]) PopItem() Record[T] {
	i := h.Records[0]
	h.Records[0] = h.Records[h.Length-1]
	h.Length--
	h.heapify(0)
	return i
}
