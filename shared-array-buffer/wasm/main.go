//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func grayscale(_ js.Value, args []js.Value) any {
	ptr := args[0] // SharedArrayBuffer

	// Get buffer size
	length := ptr.Get("byteLength").Int()

	// Access to WebAssembly memory
	memory := js.Global().Get("Uint8Array").New(ptr)

	// Loop RGBA
	for index := 0; index < length; index += 4 {
		red := memory.Index(index).Int()
		green := memory.Index(index + 1).Int()
		blue := memory.Index(index + 2).Int()

		// Filter grayscale
		gray := uint8((0.299 * float32(red)) +
			(0.587 * float32(green)) +
			(0.114 * float32(blue)))

		memory.SetIndex(index, gray)       // Red
		memory.SetIndex((index + 1), gray) // Green
		memory.SetIndex((index + 2), gray) // Blue
		// Alpha (index + 3) not changed
	}

	return nil
}

func main() {
	js.Global().Set("grayscale", js.FuncOf(grayscale))

	select {}
}
