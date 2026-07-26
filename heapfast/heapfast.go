package heapfast

import (
	"cmp"
	"math"

	"golang.org/x/exp/constraints"
)

type Record[K cmp.Ordered, V any] struct {
	Key K
	Value V
}

type heapt[K cmp.Ordered, V any] struct {
	Records []Record[K, V]
	Length  uint
}

type SizedNumber interface {
	uint64 | int64 | int32 | uint32 | float32 | float64
}

type Number interface {
	constraints.Integer | constraints.Float
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

type Heap[T cmp.Ordered, V any] interface {
	heapify(uint)
	Sort() []Record[T, V]
	Top() Record[T, V]
	AddItem(Record[T, V])
	PopItem() Record[T, V]
}
