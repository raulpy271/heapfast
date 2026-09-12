// Browser wasm loader.
//
// Deliberately uses `new URL("./main.wasm", import.meta.url)` instead of
// webpack's `import source ... from "./main.wasm"` (WebAssembly source-phase
// import). That experiment relies on a webpack-only runtime helper
// (`r.p`/`r.vs`) to resolve and fetch the emitted `.module.wasm` chunk, which
// only webpack's own bundler machinery knows how to rewrite at build time.
// Rollup/esbuild/Vite don't recognize it, so the wasm file was silently
// dropped from non-webpack production builds (see project memory /
// upstream report). `new URL(..., import.meta.url)` is the bundler-agnostic
// asset-reference convention: webpack 5, Vite and Rollup all statically
// detect this exact pattern, copy the referenced file into the build
// output, and rewrite the URL to point at it.
export async function loadWasmModule() {
  const url = new URL("./main.wasm", import.meta.url);
  const response = await fetch(url);
  if (typeof WebAssembly.compileStreaming === "function") {
    try {
      return await WebAssembly.compileStreaming(response.clone());
    } catch {
      // Fall through to the ArrayBuffer path below. Some static file
      // servers don't set `Content-Type: application/wasm`, which makes
      // `compileStreaming` reject even though the bytes are fine.
    }
  }
  const buffer = await response.arrayBuffer();
  return WebAssembly.compile(buffer);
}
