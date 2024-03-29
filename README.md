## CUE Playground

The CUE Playground is a [TypeScript](https://www.typescriptlang.org/) application that is backed by a [WASM-compiled
Go](https://github.com/golang/go/wiki/WebAssembly) application. [React](https://reactjs.org/) is used to render the UI,
with [Bootstrap](https://getbootstrap.com/) providing the styling.

For now we use a simple three-pane UI:

```
+------------------------------------------------------------------+
|                                                                  |
|  Header                                                          |
|                                                                  |
+---------------------------------+--------------------------------+
|                                 |                                |
|  Input                          |    Output                      |
|                                 |                                |
+---------------------------------+--------------------------------+
```

The output is the JSON-marshalled result of the CUE input.

### Developing locally

To develop the application locally, within the `play` directory at the repo root:

```bash
# Running dist.sh outside of the netlify environment works in a "dev"
# mode which ultimately results in running npm run serve
./dist.sh
```

### Details

* The TypeScript single-page application entry point is `src/index.tsx`
* The entire application runs via a [Webpack](https://webpack.js.org/) pipeline
* ...

### Requirements for local development

* [NodeJS](https://nodejs.org/) `>= v12.14.1`
* [Go](https://golang.org/dl/) (stable version)

### TODO

* UI/UX
  * Support `fmt` and `trim` dropdown button option for the input pane contents
  * Support `txtar` input that then gets used as an overlay
  * Extend input dropdown to support different types of input (CUE, JSON, Yaml, Go)
  * Extend output dropdown modes
* Development improvements
  * Integrate automatic recompilation of `main.wasm` into the webpack watch pipeline
  * Ensure, via GitHub Action checks, that `.go` and `.tsx` files are formatted
