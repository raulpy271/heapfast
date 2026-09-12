// Node.js wasm loader. See wasm-loader.browser.js for why this reads the
// wasm asset via `new URL("./main.wasm", import.meta.url)` instead of
// webpack's `import source ...` WebAssembly experiment.
import { readFile } from "node:fs/promises";
import { fileURLToPath } from "node:url";

export async function loadWasmModule() {
  const url = new URL("./main.wasm", import.meta.url);
  const buffer = await readFile(fileURLToPath(url));
  return WebAssembly.compile(buffer);
}
