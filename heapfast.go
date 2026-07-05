package main

type Item struct {
	key float64;
	value any
}

type Heap struct {
	items []Item
}

func (h *Heap) AddItem(i Item) {
	h.items = append(h.items, i)
}

func (h *Heap) PopItem() Item {
	i := h.items[0]
	h.items = h.items[1:]
	return i
}
