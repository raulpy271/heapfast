import heapfast from "../heapfast.js";

export function mochaGlobalSetup() {
  const go = new Go();
  heapfast.startWasmModule(go);
}
