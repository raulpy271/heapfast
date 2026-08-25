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

func addMethodsZero[K heapfast.Number, T heapfast.Heap[K, heapfast.Zero]](heaps []T, heapref map[string]any) {
	heapref["add"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		h.AddItem(heapfast.Record[K, heapfast.Zero]{heapfast.Zero{}, K(args[0].Float())})
		return nil
	})
	heapref["addBulk"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		if !args[0].InstanceOf(js.Global().Get("Uint8Array")) {
			panic("The first parameter should be an array")
		}
		v := this.Get("pos").Int()
		h := heaps[v]
		keys := args[0]
		l := keys.Length()
		for _ = range l {
			h.AddItem(heapfast.Record[K, heapfast.Zero]{heapfast.Zero{}, K(args[0].Float())})
		}
		return nil
	})
	heapref["pop"] = js.FuncOf(func(this js.Value, args []js.Value) any {
		v := this.Get("pos").Int()
		h := heaps[v]
		i := h.PopItem()
		return i.Key
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
	heapsInt := make([]heapfast.Heap[int64, int64], 0, 20)
	heapsFloat := make([]heapfast.Heap[float64, int64], 0, 20)
	heapsIntZero := make([]heapfast.Heap[int64, heapfast.Zero], 0, 20)
	heapsFloatZero := make([]heapfast.Heap[float64, heapfast.Zero], 0, 20)
	js.Global().Set("NewHeapMaxIntKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[int64, int64]{}
		heapsInt = append(heapsInt, &h)
		pos := len(heapsInt) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsInt, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinIntKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[int64, int64]{}
		heapsInt = append(heapsInt, &h)
		pos := len(heapsInt) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsInt, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMaxFloatKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[float64, int64]{}
		heapsFloat = append(heapsFloat, &h)
		pos := len(heapsFloat) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsFloat, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinFloatKV", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[float64, int64]{}
		heapsFloat = append(heapsFloat, &h)
		pos := len(heapsFloat) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethods(heapsFloat, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMaxIntK", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[int64, heapfast.Zero]{}
		heapsIntZero = append(heapsIntZero, &h)
		pos := len(heapsIntZero) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethodsZero(heapsIntZero, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinIntK", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[int64, heapfast.Zero]{}
		heapsIntZero = append(heapsIntZero, &h)
		pos := len(heapsIntZero) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethodsZero(heapsIntZero, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMaxFloatK", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMax[float64, heapfast.Zero]{}
		heapsFloatZero = append(heapsFloatZero, &h)
		pos := len(heapsFloatZero) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethodsZero(heapsFloatZero, heapref)
		return heapref
	}))
	js.Global().Set("NewHeapMinFloatK", js.FuncOf(func(this js.Value, args []js.Value) any {
		h := heapfast.HeapMin[float64, heapfast.Zero]{}
		heapsFloatZero = append(heapsFloatZero, &h)
		pos := len(heapsFloatZero) - 1
		heapref := make(map[string]any)
		heapref["pos"] = pos
		addMethodsZero(heapsFloatZero, heapref)
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
