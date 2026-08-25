import source wasmModule from "./main.wasm";
import "./wasm_exec.js";

const TYPE_INT = 0;
const TYPE_FLOAT = 1;
export const ASC = 0;
export const DESC = 1;

var wasmInstance;

export async function startWasmModule(go) {
  if (!go) {
    go = new Go();
  }
  if (!wasmInstance) {
    wasmInstance = new WebAssembly.Instance(wasmModule, go.importObject);
    const exitCode = await go.run(wasmInstance);
    if (exitCode) {
      throw new Error(`Error when running the main wasm function: $exitCode`);
    }
  }
  return go;
}

function heapsortWasm(buffer, orderby, type, withValues) {
  let funcname = "Heapsort";
  if (orderby === DESC) {
    funcname += "Desc";
  } else {
    funcname += "Asc";
  }
  if (type === TYPE_FLOAT) {
    funcname += "Float";
  } else {
    funcname += "Int";
  }
  if (withValues) {
    funcname += "KV";
  } else {
    funcname += "K";
  }
  return globalThis[funcname](buffer);
}

function priorityQueueWasm(orderby, type, withValues) {
  let funcname = "NewHeap";
  if (orderby === DESC) {
    funcname += "Max";
  } else {
    funcname += "Min";
  }
  if (type === TYPE_FLOAT) {
    funcname += "Float";
  } else {
    funcname += "Int";
  }
  if (withValues) {
    funcname += "KV";
  } else {
    funcname += "K";
  }
  return globalThis[funcname]();
}

export function heapsort(array, orderby, key) {
  if (!wasmInstance) {
    startWasmModule();
  }
  if (array instanceof BigInt64Array) {
    return heapsortWasm(new Uint8Array(array.buffer), orderby, TYPE_INT, false);
  }
  if (array instanceof Float64Array) {
    return heapsortWasm(
      new Uint8Array(array.buffer),
      orderby,
      TYPE_FLOAT,
      false,
    );
  }
  if (
    array instanceof Int32Array ||
    array instanceof Float32Array ||
    array instanceof Uint32Array ||
    (Array.isArray(array) && typeof array[0] === "number")
  ) {
    const wasmArray = new Float64Array(array.length);
    for (let i = 0; i < array.length; i++) {
      wasmArray[i] = array[i];
    }
    const len = heapsortWasm(
      new Uint8Array(wasmArray.buffer),
      orderby,
      TYPE_FLOAT,
      false,
    );
    for (let i = 0; i < len; i++) {
      array[i] = wasmArray[i];
    }
    return len;
  }
  if (Array.isArray(array) && key) {
    let keyfunc;
    if (typeof key === "function") {
      keyfunc = key;
    } else if (typeof key === "string") {
      keyfunc = (obj) => obj[key];
    } else {
      throw new TypeError("unsupported key type");
    }
    const wasmArray = new Float64Array(array.length * 2);
    for (let i = 0; i < array.length; i++) {
      wasmArray[i * 2] = i;
      wasmArray[i * 2 + 1] = keyfunc(array[i]);
    }
    const len = heapsortWasm(
      new Uint8Array(wasmArray.buffer),
      orderby,
      TYPE_FLOAT,
      true,
    );
    const original = array.slice(0, len);
    for (let i = 0; i < len; i++) {
      array[i] = original[wasmArray[i * 2]];
    }
    return len;
  }

  throw new TypeError(
    "heapsort: unsupported array type, expected BigInt64Array, Float64Array, Int32Array, Float32Array, Uint32Array or Array",
  );
}

export class PriorityQueue {
  constructor(orderby) {
    if (orderby === DESC) {
      this.orderby = DESC;
    } else {
      this.orderby = ASC;
    }
    this.values = [];
    this.freeHead = undefined;
    this.queue = priorityQueueWasm(this.orderby, TYPE_FLOAT, true);
    this.valuesFree = 0;
  }

  add(key, value) {
    let id;
    if (this.freeHead && this.valuesFree > 0) {
        id = this.freeHead;
        this.freeHead = this.values[this.freeHead];
        this.values[id] = value;
        this.valuesFree--;
    } else {
        id = this.values.push(value) - 1;
    }
    this.queue.add(key, id);
  }

  pop() {
    const [key, id] = this.queue.pop();
    const value = this.values[id];
    this.values[id] = this.freeHead;
    this.freeHead = id;
    this.valuesFree++;
    return [key, value];
  }
}

export default {
  startWasmModule,
  heapsort,
  ASC,
  DESC,
  PriorityQueue,
};
