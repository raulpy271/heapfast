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

function isFixedSizeArray(array) {
  return array instanceof BigInt64Array;
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
  let wasmArray;
  let type;
  let data = {};
  if (array instanceof BigInt64Array) {
    wasmArray = array;
    type = TYPE_INT;
  } else if (array instanceof Float64Array) {
    wasmArray = array;
    type = TYPE_FLOAT;
  }
  if (array instanceof Int32Array) {
    wasmArray = new Float64Array(array.length);
    for (let i = 0; i < array.length; i++) {
      wasmArray[i] = array[i];
    }
    type = TYPE_FLOAT;
    let len = heapsortWasm(
      new Uint8Array(wasmArray.buffer),
      orderby,
      type,
      false,
    );
    for (let i = 0; i < len; i++) {
      array[i] = wasmArray[i];
    }
    return len;
  }
  let len = heapsortWasm(
    new Uint8Array(wasmArray.buffer),
    orderby,
    type,
    false,
  );
  return len;
}

export default {
  startWasmModule,
  heapsort,
  ASC,
  DESC,
};
