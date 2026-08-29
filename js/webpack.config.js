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
  asyncWebAssembly: true,
  sourceImport: true,
};

// Browser/ESM build: for bundlers targeting the browser (e.g. a ReactJS app
// built with webpack). Using target "web" alone (instead of the combined
// ["web", "node"] target) keeps webpack's generated wasm-chunk loader to the
// fetch()-based path only, so downstream browser bundles never contain
// references to Node core modules like "fs"/"url".
const browserConfig = {
  mode: "production",
  devtool: "source-map",
  entry: "./heapfast.js",
  target: ["web"],
  experiments,
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
  output: {
    filename: "main.cjs",
    path: outputPath,
    module: false,
    library: { type: "commonjs2" },
  },
};

export default [browserConfig, nodeEsmConfig, nodeCjsConfig];
