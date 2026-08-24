import heapfast from "../dist/main.js";

export async function mochaGlobalSetup() {
  const go = new Go();
  heapfast.startWasmModule(go);
}
