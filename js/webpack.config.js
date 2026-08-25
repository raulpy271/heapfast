import path from "node:path";
import { fileURLToPath } from "node:url";

// In Node.js versions prior to native support for import.meta.dirname,
// derive __dirname from import.meta.url.
// (Node 20.11+ supports import.meta.dirname and import.meta.filename.)
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default {
  mode: "development",
  entry: "./heapfast.js",
  experiments: {
    outputModule: true,
    asyncWebAssembly: true,
    sourceImport: true,
  },
  output: {
    filename: "main.js",
    path: path.resolve(__dirname, "dist"),
    module: true,
    library: {
      type: "module",
    },
  },
};
