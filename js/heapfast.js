import fs from "fs";
import path from "path";
import "./wasm_exec.js";

const TYPE_INT = 0;
const TYPE_FLOAT = 1;
export const ASC = 0;
export const DESC = 1;

function createInstance(imports) {
  const p = path.join(path.resolve(), "main.wasm");
  const bytes = fs.readFileSync(p);
  const wasmModule = new WebAssembly.Module(bytes);
  return new WebAssembly.Instance(wasmModule, imports);
}

export async function startWasmModule(go) {
  if (!go) {
    go = new Go();
  }
  const instance = createInstance(go.importObject);
  const exitCode = await go.run(instance);
  if (exitCode) {
    throw new Error(`Error when running the main wasm function: $exitCode`);
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

export function heapsort(array, orderby) {
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
  throw new TypeError(
    "heapsort: unsupported array type, expected BigInt64Array, Float64Array, Int32Array, Float32Array, Uint32Array or Array",
  );
}

export default {
  startWasmModule,
  heapsort,
  ASC,
  DESC,
};
