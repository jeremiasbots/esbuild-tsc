# ESBuild-TSC Tool

This tool is used to transform from TypeScript to JavaScript and run the code with an engine (node ​​| deno | bun), for configuration, create your esbuild-tsc.json file with the following properties (the values ​​are an example and are free except for the engine which has to be node, bun or deno):

```json
{
  "file": "index.ts",
  "tsconfig": "./tsconfig.json",
  "engine": "node"
}
```

To use the tool type `esbuild-tsc tar` (do all the alias steps first, the binaries for each architecture and system are in the bin/ folder), see the example of test/

Created by jeremiasbots/devep
