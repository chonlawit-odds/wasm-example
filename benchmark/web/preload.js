const go = new Go();

// WebAssembly.instantiateStreaming(fetch("go.wasm"), go.importObject).then(
//   (result) => {
//     go.run(result.instance);
//   }
// );

WebAssembly.instantiateStreaming(fetch("tiny_go.wasm"), go.importObject).then(
  (result) => {
    wasm = result.instance;
    go.run(wasm);
  }
);
