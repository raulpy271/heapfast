import fs from "fs";
import path from "path";
import "./wasm_exec.js";

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

// console.log("Go.run result:", result);
// console.log("Go runtime started");
// h = NewHeapMax();
// h.add(10, 11)
// h.add(20, 21)
// console.log(h.pop())
