package heapfast

import (
	"cmp"
	"math"
)

type Record[T cmp.Ordered] [2]T

type heapt[T cmp.Ordered] struct {
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

type Heap[T cmp.Ordered] interface {
	heapify(uint)
	Sort() []Record[T]
	Top() Record[T]
	AddItem(Record[T])
	PopItem() Record[T]
}
