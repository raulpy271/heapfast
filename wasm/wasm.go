//go:build js && wasm

package main

import "github.com/raulpy271/heapfast/heapfast"
import "syscall/js"

func mean(this js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].Length() == 0 {
		return js.ValueOf(0)
	}
	arr := args[0]
	n := arr.Length()
	sum := 0.0
	for i := range n {
		sum += arr.Index(i).Float()
	}
	return js.ValueOf(sum / float64(n))
}

func NewHeap(this js.Value, args []js.Value) any {
	h := heapfast.Heap{}
	return js.ValueOf(h)
}

func main() {
	heaps := make([]heapfast.Heap, 0, 20)
	js.Global().Set("NewHeap", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.Heap{}
		heaps = append(heaps, h)
		pos := len(heaps) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		heapref["add"] = js.FuncOf(func(this js.Value, args []js.Value) any {
			v := this.Get("pos").Int()
			h := &heaps[v]
			h.AddItem(heapfast.Record{uint64(args[0].Int()), uint64(args[1].Int())})
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
			for i := range l {
				h.AddItem(heapfast.Record{uint64(keys.Index(i).Int()), uint64(values.Index(i).Int())})
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
	js.Global().Set("Heapsort", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Array")) {
			panic("The first parameter should be an array")
		}
		//js.CopyBytesToGo()
		return nil
	}))
	select {}
}
