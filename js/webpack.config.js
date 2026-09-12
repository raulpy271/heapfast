import path from "node:path";
import { fileURLToPath } from "node:url";

// In Node.js versions prior to native support for import.meta.dirname,
// derive __dirname from import.meta.url.
// (Node 20.11+ supports import.meta.dirname and import.meta.filename.)
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const outputPath = path.resolve(__dirname, "dist");

const experiments = {
  outputModule: true,
};

// heapfast.js imports the wasm loader from "./wasm-loader.js", a module that
// doesn't exist on disk. The Node configs below alias it to a loader variant
// bundled the ordinary way (their output is only ever consumed directly by
// Node, never re-bundled by Vite/Rollup, so webpack's own asset pipeline is
// fine there). The browser config, by contrast, treats it as *external* and
// left completely unbundled — see the comment on browserConfig for why.
const nodeLoaderAlias = {
  "./wasm-loader.js": path.resolve(__dirname, "wasm-loader.node.js"),
};

// Treat main.wasm as a plain binary asset: webpack copies it into the output
// directory (content-hashed) and rewrites `new URL("./main.wasm", ...)` to
// point at the emitted file. Fine for the Node builds; see browserConfig for
// why this is *not* used there.
const wasmAssetModule = {
  rules: [{ test: /\.wasm$/, type: "asset/resource" }],
};

// Browser/ESM build: for bundlers targeting the browser (e.g. a ReactJS app
// built with Vite or webpack). Using target "web" alone (instead of the
// combined ["web", "node"] target) keeps the emitted code to the
// fetch()-based path only, so downstream browser bundles never contain
// references to Node core modules like "fs"/"url".
//
// Unlike the Node builds, wasm-loader.js is marked *external* here and never
// touched by webpack at all: webpack's own asset handling for
// `new URL("./main.wasm", import.meta.url)` rewrites that call into
// `new URL(r(732), r.b)` — a runtime module-registry lookup, not a source
// literal. Vite/Rollup only recognize the `new URL(...)` asset-copy
// convention when the first argument is a literal string, so a
// webpack-processed reference is exactly as opaque to them as the old
// `import source ...` WebAssembly experiment was (see git history/upstream
// bug report). Leaving wasm-loader.browser.js as a real, unbundled file
// (copied verbatim into dist/ by the "build" script — see package.json)
// means downstream bundlers read its actual, untouched
// `new URL("./main.wasm", import.meta.url)` literal and handle it exactly
// as they would in a project's own source.
const browserConfig = {
  mode: "production",
  devtool: "source-map",
  entry: "./heapfast.js",
  target: ["web"],
  experiments,
  externalsType: "module",
  externals: {
    "./wasm-loader.js": "./wasm-loader.js",
  },
  output: {
    filename: "main.browser.js",
    path: outputPath,
    module: true,
    library: { type: "module" },
  },
};

// Node ESM build: for `import` in Node.js. Keeps the historical "main.js"
// filename so the existing test suite (which imports "../dist/main.js")
// keeps working unmodified.
const nodeEsmConfig = {
  mode: "production",
  devtool: "source-map",
  entry: "./heapfast.js",
  target: ["node"],
  experiments,
  module: wasmAssetModule,
  resolve: { alias: nodeLoaderAlias },
  output: {
    filename: "main.js",
    path: outputPath,
    module: true,
    library: { type: "module" },
  },
};

// Node CJS build: for `require()` in Node.js.
const nodeCjsConfig = {
  mode: "production",
  devtool: "source-map",
  entry: "./heapfast.js",
  target: ["node"],
  experiments: {
    ...experiments,
    outputModule: false,
  },
  module: wasmAssetModule,
  resolve: { alias: nodeLoaderAlias },
  output: {
    filename: "main.cjs",
    path: outputPath,
    module: false,
    library: { type: "commonjs2" },
  },
};

export default [browserConfig, nodeEsmConfig, nodeCjsConfig];
