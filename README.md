# ESBuild-TSC Tool

This tool is used to transform from TypeScript to JavaScript and run the code with an engine (node ​​| deno | bun), for configuration, create your esbuild-tsc.json file with the following properties (the values ​​are an example and are free except for the engine which has to be node, bun or deno):

```json
{
  "$schema": "https://raw.githubusercontent.com/jeremiasbots/esbuild-tsc/refs/heads/main/esbuild-tsc-schema.json",
  "file": "index.ts",
  "engine": "node",
  "tsconfig": "./tsconfig.json"
}
```

To use the tool type `esbuild-tsc tar` (do all the alias steps first, the binaries for each architecture and system are in the bin/ folder), see the example of cmd/test

To create an example project use `esbuild-tsc create <name>`

Created by jeremiasbots/devep
