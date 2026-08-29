/**
 * Type declarations for the `heapfast` package.
 *
 * These describe the public API exposed by `heapfast.js` (bundled into
 * `dist/main.js` by webpack, which preserves the same named/default
 * exports), backed by a TinyGo-compiled WASM module.
 */

/**
 * Minimal structural type for the `Go` class defined by Go's WASM glue
 * runtime (`wasm_exec.js`). No official types are published for it, so
 * only the members `startWasmModule` relies on are declared here.
 */
export interface Go {
  importObject: WebAssembly.Imports;
  run(instance: WebAssembly.Instance): Promise<void>;
}

/** Sort in ascending order. */
export const ASC: 0;
/** Sort in descending order. */
export const DESC: 1;

/** Sort order accepted by `heapsort` and `PriorityQueue`: `ASC` or `DESC`. */
export type OrderBy = typeof ASC | typeof DESC;

/**
 * Instantiates and starts the WASM module, if not already started.
 * Pass an existing `Go` instance to reuse it; otherwise one is created.
 * Must resolve before `heapsort`/`PriorityQueue` are used, unless you
 * let them lazily start the module for you.
 */
export function startWasmModule(go?: Go): Promise<Go>;

/**
 * Sorts a numeric typed array or plain number array in place.
 * Returns the number of sorted elements.
 */
export function heapsort(
  array: BigInt64Array | Float64Array | Int32Array | Float32Array | Uint32Array | number[],
  orderby?: OrderBy,
): number;
/**
 * Sorts an array of objects in place by a key, given either a property
 * name or a function extracting a numeric key from each element.
 * Returns the number of sorted elements.
 */
export function heapsort<T>(
  array: T[],
  orderby: OrderBy,
  key: keyof T | ((item: T) => number),
): number;

/**
 * A priority queue backed by a WASM binary min/max heap. Keys are
 * numbers; values may be any JS value and are kept on the JS side.
 */
export class PriorityQueue<V = unknown> {
  constructor(orderby?: OrderBy);
  add(key: number, value: V): void;
  pop(): [number, V];
}

declare const heapfast: {
  startWasmModule: typeof startWasmModule;
  heapsort: typeof heapsort;
  ASC: typeof ASC;
  DESC: typeof DESC;
  PriorityQueue: typeof PriorityQueue;
};

export default heapfast;
