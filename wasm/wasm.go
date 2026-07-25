//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/raulpy271/heapfast/heapfast"
	"github.com/raulpy271/heapfast/internal"
)

func addMethods[R uint32 | int32 | uint64 | int64, T heapfast.Heap[R]](heaps []T, heapref map[string]any) {
	heapref["add"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		h.AddItem(heapfast.Record[R]{R(args[0].Int()), R(args[1].Int())})
		return nil
	})
	heapref["addBulk"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Array")) || !args[1].InstanceOf(js.Global().Get("Array")) {
			panic("O primeiro parametro deve ser um array")
		}
		v := this.Get("pos").Int()
		h := heaps[v]
		keys := args[0]
		values := args[1]
		l := keys.Length()
		for i := range l {
			h.AddItem(heapfast.Record[R]{R(keys.Index(i).Int()), R(values.Index(i).Int())})
		}
		return nil
	})
	heapref["pop"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		i := h.PopItem()
		return []any{i[0], i[1]}
	})

}

func main() {
	heaps32 := make([]heapfast.Heap[int32], 0, 20)
	heaps64 := make([]heapfast.Heap[int64], 0, 20)
	js.Global().Set("NewHeapMax", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[int32]{}
		heaps32 = append(heaps32, &h)
		pos := len(heaps32) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods[int32, heapfast.Heap[int32]](heaps32, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMin", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[int32]{}
		heaps32 = append(heaps32, &h)
		pos := len(heaps32) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods[int32, heapfast.Heap[int32]](heaps32, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMax64", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[int64]{}
		heaps64 = append(heaps64, &h)
		pos := len(heaps64) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods[int64, heapfast.Heap[int64]](heaps64, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMin64", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[int64]{}
		heaps64 = append(heaps64, &h)
		pos := len(heaps64) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods[int64, heapfast.Heap[int64]](heaps64, heapref)
		return heapref
	}))
	js.Global().Set("HeapsortAsc", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[int32](dst)
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDesc", js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		dst := make([]byte, args[0].Length())
		js.CopyBytesToGo(dst, args[0])
		records := internal.CastRecordsFromBytes[int32](dst)
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	select {}
}
