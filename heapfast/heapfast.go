package heapfast

import (
	"cmp"
	"math"
)

type Record[T cmp.Ordered] [2]T

type Heap[T cmp.Ordered] struct {
	Records []Record[T]
	Length  uint
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

func (h *Heap[T]) heapify(i uint) {
	var l, r, largest uint
	for i < h.Length {
		l = left(i)
		r = right(i)
		if l < h.Length && h.Records[i][0] < h.Records[l][0] {
			largest = l
		} else {
			largest = i
		}
		if r < h.Length && h.Records[largest][0] < h.Records[r][0] {
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

func BuildMaxHeap[T cmp.Ordered](records []Record[T]) *Heap[T] {
	heap := &Heap[T]{records, uint(len(records))}
	i := int(math.Floor((float64(len(records)) / 2) - 1))
	for ; i >= 0; i-- {
		heap.heapify(uint(i))
	}
	return heap
}

func SortMax[T cmp.Ordered](records []Record[T]) []Record[T] {
	heap := BuildMaxHeap(records)
	return heap.Records
}

func (h *Heap[T]) AddItem(record Record[T]) {
	if int(h.Length) == len(h.Records) {
		h.Records = append(h.Records, record)
	} else {
		h.Records[h.Length] = record
	}
	i := h.Length
	h.Length++
	for i > 0 && h.Records[parent(i)][0] < h.Records[i][0] {
		h.Records[i], h.Records[parent(i)] = h.Records[parent(i)], h.Records[i]
		i = parent(i)
	}
}

func (h Heap[T]) Max() Record[T] {
	return h.Records[0]
}

func (h *Heap[T]) PopItem() Record[T] {
	i := h.Records[0]
	h.Records[0] = h.Records[h.Length-1]
	h.Length--
	h.heapify(0)
	return i
}
