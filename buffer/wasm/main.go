//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

var buffer []byte

func setBuffer(_ js.Value, args []js.Value) any {
	jsArray := args[0] // Receive Uint8Array from JavaScript
	length := jsArray.Get("length").Int()
	buffer = make([]byte, length)

	// Copy data from Uint8Array to slice in Go
	js.CopyBytesToGo(buffer, jsArray)

	return nil
}

func processBuffer(_ js.Value, _ []js.Value) any {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = buffer[i] / 2 // Divide half value
	}

	return nil
}

func getBuffer(_ js.Value, _ []js.Value) any {
	jsArray := js.Global().Get("Uint8Array").New(len(buffer))
	js.CopyBytesToJS(jsArray, buffer)
	return jsArray
}

func main() {
	js.Global().Set("setBuffer", js.FuncOf(setBuffer))
	js.Global().Set("processBuffer", js.FuncOf(processBuffer))
	js.Global().Set("getBuffer", js.FuncOf(getBuffer))
	select {}
}
