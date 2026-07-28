//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/raulpy271/heapfast/heapfast"
	"github.com/raulpy271/heapfast/internal"
)

func addMethods[K heapfast.Number, V heapfast.Integer, T heapfast.Heap[K, V]](heaps []T, heapref map[string]any) {
	heapref["add"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		h.AddItem(heapfast.Record[K, V]{K(args[0].Float()), V(args[1].Int())})
		return nil
	})
	heapref["addBulk"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		v := this.Get("pos").Int()
		h := heaps[v]
		keys := args[0]
		values := args[1]
		l := keys.Length()
		for i := range l {
			h.AddItem(heapfast.Record[K, V]{K(keys.Index(i).Float()), V(values.Index(i).Int())})
		}
		return nil
	})
	heapref["pop"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		i := h.PopItem()
		return []any{i.Key, i.Value}
	})
}

func main() {
	heapsint := make([]heapfast.Heap[int64, int64], 0, 20)
	heapsfloat := make([]heapfast.Heap[float64, int64], 0, 20)
	js.Global().Set("NewHeapMaxInt", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[int64, int64]{}
		heapsint = append(heapsint, &h)
		pos := len(heapsint) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsint, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinInt", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[int64, int64]{}
		heapsint = append(heapsint, &h)
		pos := len(heapsint) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsint, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMaxFloat", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[float64, int64]{}
		heapsfloat = append(heapsfloat, &h)
		pos := len(heapsfloat) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsfloat, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinFloat", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[float64, int64]{}
		heapsfloat = append(heapsfloat, &h)
		pos := len(heapsfloat) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsfloat, heapref)
		return heapref
	}))
	js.Global().Set("HeapsortAscInt", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[int64, int64](dst)
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescInt", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[int64, int64](dst)
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortAscFloat", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[float64, int64](dst)
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescInt", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[float64, int64](dst)
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	select {}
}
