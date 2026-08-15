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
		h.AddItem(heapfast.Record[K, V]{V(args[1].Int()), K(args[0].Float())})
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
			h.AddItem(heapfast.Record[K, V]{V(values.Index(i).Int()), K(keys.Index(i).Float())})
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

func jsValueToRecords[K heapfast.SizedNumber, V heapfast.SizedNumber | heapfast.Zero](v js.Value) ([]byte, []heapfast.Record[K, V]) {
	if !v.InstanceOf(js.Global().Get("Uint8Array")) {
		panic("The first parameter should be an array")
	}
	dst := make([]byte, v.Length())
	js.CopyBytesToGo(dst, v)
	return dst, internal.CastRecordsFromBytes[K, V](dst)
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
	js.Global().Set("HeapsortAscIntKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[int64, int64](args[0])
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortAscIntK", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[int64, heapfast.Zero](args[0])
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescIntKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[int64, int64](args[0])
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescIntK", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[int64, heapfast.Zero](args[0])
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortAscFloatKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[float64, int64](args[0])
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortAscFloatK", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[float64, heapfast.Zero](args[0])
		heap := heapfast.BuildMaxHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescFloatKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[float64, int64](args[0])
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	js.Global().Set("HeapsortDescFloatK", js.FuncOf(func(this js.Value, args []js.Value) any {
		dst, records := jsValueToRecords[float64, heapfast.Zero](args[0])
		heap := heapfast.BuildMinHeap(records)
		heap.Sort()
		dst = internal.CastRecordsToBytes(records)
		js.CopyBytesToJS(args[0], dst)
		return len(records)
	}))
	select {}
}
