//go:build js && wasm

package main

import "syscall/js"

func mean(this js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].Length() == 0 {
		return js.ValueOf(0)
	}
	arr := args[0]
	n := arr.Length()
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += arr.Index(i).Float()
	}
	return js.ValueOf(sum / float64(n))
}

func NewHeap(this js.Value, args []js.Value) any {
	h := Heap{}
	return js.ValueOf(h)
}

func main() {
	heaps := make([]Heap, 0, 20)
	js.Global().Set("NewHeap", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := Heap{}
		heaps = append(heaps, h)
		pos := len(heaps) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		heapref["add"] = js.FuncOf(func(this js.Value, args []js.Value) any {
			v := this.Get("pos").Int()
			h := &heaps[v]
			h.AddItem(Record{uint(args[0].Int()), uint(args[1].Int())})
			return nil
		})
		heapref["addBulk"] = js.FuncOf(func(this js.Value, args []js.Value) any {
			if !args[0].InstanceOf(js.Global().Get("Array")) || !args[1].InstanceOf(js.Global().Get("Array")) {
				panic("O primeiro parametro deve ser um array")
			}
			v := this.Get("pos").Int()
			h := &heaps[v]
			keys := args[0]
			values := args[1]
			l := keys.Length()
			for i := 0; i < l; i++ {
				h.AddItem(Record{uint(keys.Index(i).Int()), uint(values.Index(i).Int())})
			}
			return nil
		})
		heapref["pop"] = js.FuncOf(func(this js.Value, args []js.Value) any {
			v := this.Get("pos").Int()
			h := &heaps[v]
			i := h.PopItem()
			return []any{i[0], i[1]}
		})
		return heapref
	}))
	select {}
}
